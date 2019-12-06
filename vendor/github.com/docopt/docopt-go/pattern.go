package docopt

import (
	"fmt"
	"reflect"
	"strings"
)

type patternType uint

const (
	// leaf
	patternArgument patternType = 1 << iota
	patternCommand
	patternOption

	// branch
	patternRequired
	patternOptionAL
	patternOptionSSHORTCUT // Marker/placeholder for [options] shortcut.
	patternOneOrMore
	patternEither

	patternLeaf = patternArgument +
		patternCommand +
		patternOption
	patternBranch = patternRequired +
		patternOptionAL +
		patternOptionSSHORTCUT +
		patternOneOrMore +
		patternEither
	patternAll     = patternLeaf + patternBranch
	patternDefault = 0
)

func (pt patternType) String() string {
	switch pt {
	case patternArgument:
		return "argument"
	case patternCommand:
		return "command"
	case patternOption:
		return "option"
	case patternRequired:
		return "required"
	case patternOptionAL:
		return "optional"
	case patternOptionSSHORTCUT:
		return "optionsshortcut"
	case patternOneOrMore:
		return "oneormore"
	case patternEither:
		return "either"
	case patternLeaf:
		return "leaf"
	case patternBranch:
		return "branch"
	case patternAll:
		return "all"
	case patternDefault:
		return "default"
	}
	return ""
}

type pattern struct {
	t patternType

	children patternList

	name  string
	value interface{}

	short    string
	long     string
	argcount int
}

type patternList []*pattern

func newBranchPattern(t patternType, pl ...*pattern) *pattern {
	var p pattern
	p.t = t
	p.children = make(patternList, len(pl))
	copy(p.children, pl)
	return &p
}

func newRequired(pl ...*pattern) *pattern {
	return newBranchPattern(patternRequired, pl...)
}

func newEither(pl ...*pattern) *pattern {
	return newBranchPattern(patternEither, pl...)
}

func newOneOrMore(pl ...*pattern) *pattern {
	return newBranchPattern(patternOneOrMore, pl...)
}

func newOptional(pl ...*pattern) *pattern {
	return newBranchPattern(patternOptionAL, pl...)
}

func newOptionsShortcut() *pattern {
	var p pattern
	p.t = patternOptionSSHORTCUT
	return &p
}

func newLeafPattern(t patternType, name string, value interface{}) *pattern {
	// default: value=nil
	var p pattern
	p.t = t
	p.name = name
	p.value = value
	return &p
}

func newArgument(name string, value interface{}) *pattern {
	// default: value=nil
	return newLeafPattern(patternArgument, name, value)
}

func newCommand(name string, value interface{}) *pattern {
	// default: value=false
	var p pattern
	p.t = patternCommand
	p.name = name
	p.value = value
	return &p
}

func newOption(short, long string, argcount int, value interface{}) *pattern {
	// default: "", "", 0, false
	var p pattern
	p.t = patternOption
	p.short = short
	p.long = long
	if long != "" {
		p.name = long
	} else {
		p.name = short
	}
	p.argcount = argcount
	if value == false && argcount > 0 {
		p.value = nil
	} else {
		p.value = value
	}
	return &p
}

func (p *pattern) flat(types patternType) (patternList, error) {
	if p.t&patternLeaf != 0 {
		if types == patternDefault {
			types = patternAll
		}
		if p.t&types != 0 {
			return patternList{p}, nil
		}
		return patternList{}, nil
	}

	if p.t&patternBranch != 0 {
		if p.t&types != 0 {
			return patternList{p}, nil
		}
		result := patternList{}
		for _, child := range p.children {
			childFlat, err := child.flat(types)
			if err != nil {
				return nil, err
			}
			result = append(result, childFlat...)
		}
		return result, nil
	}
	return nil, newError("unknown pattern type: %d, %d", p.t, types)
}

func (p *pattern) fix() error {
	err := p.fixIdentities(nil)
	if err != nil {
		return err
	}
	p.fixRepeatingArguments()
	return nil
}

func (p *pattern) fixIdentities(uniq patternList) error {
	// Make pattern-tree tips point to same object if they are equal.
	if p.t&patternBranch == 0 {
		return nil
	}
	if uniq == nil {
		pFlat, err := p.flat(patternDefault)
		if err != nil {
			return err
		}
		uniq = pFlat.unique()
	}
	for i, child := range p.children {
		if child.t&patternBranch == 0 {
			ind, err := uniq.index(child)
			if err != nil {
				return err
			}
			p.children[i] = uniq[ind]
		} else {
			err := child.fixIdentities(uniq)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *pattern) fixRepeatingArguments() {
	// Fix elements that should accumulate/increment values.
	var either []patternList

	for _, child := range p.transform().children {
		either = append(either, child.children)
	}
	for _, cas := range either {
		casMultiple := patternList{}
		for _, e := range cas {
			if cas.count(e) > 1 {
				casMultiple = append(casMultiple, e)
			}
		}
		for _, e := range casMultiple {
			if e.t == patternArgument || e.t == patternOption && e.argcount > 0 {
				switch e.value.(type) {
				case string:
					e.value = strings.Fields(e.value.(string))
				case []string:
				default:
					e.value = []string{}
				}
			}
			if e.t == patternCommand || e.t == patternOption && e.argcount == 0 {
				e.value = 0
			}
		}
	}
}

func (p *pattern) match(left *patternList, collected *patternList) (bool, *patternList, *patternList) {
	if collected == nil {
		collected = &patternList{}
	}
	if p.t&patternRequired != 0 {
		l := left
		c := collected
		for _, p := range p.children {
			var matched bool
			matched, l, c = p.match(l, c)
			if !matched {
				return false, left, collected
			}
		}
		return true, l, c
	} else if p.t&patternOptionAL != 0 || p.t&patternOptionSSHORTCUT != 0 {
		for _, p := range p.children {
			_, left, collected = p.match(left, collected)
		}
		return true, left, collected
	} else if p.t&patternOneOrMore != 0 {
		if len(p.children) != 1 {
			panic("OneOrMore.match(): assert len(p.children) == 1")
		}
		l := left
		c := collected
		var lAlt *patternList
		matched := true
		times := 0
		for matched {
			// could it be that something didn't match but changed l or c?
			matched, l, c = p.children[0].match(l, c)
			if matched {
				times++
			}
			if lAlt == l {
				break
			}
			lAlt = l
		}
		if times >= 1 {
			return true, l, c
		}
		return false, left, collected
	} else if p.t&patternEither != 0 {
		type outcomeStruct struct {
			matched   bool
			left      *patternList
			collected *patternList
			length    int
		}
		outcomes := []outcomeStruct{}
		for _, p := range p.children {
			matched, l, c := p.match(left, collected)
			outcome := outcomeStruct{matched, l, c, len(*l)}
			if matched {
				outcomes = append(outcomes, outcome)
			}
		}
		if len(outcomes) > 0 {
			minLen := outcomes[0].length
			minIndex := 0
			for i, v := range outcomes {
				if v.length < minLen {
					minIndex = i
				}
			}
			return outcomes[minIndex].matched, outcomes[minIndex].left, outcomes[minIndex].collected
		}
		return false, left, collected
	} else if p.t&patternLeaf != 0 {
		pos, match := p.singleMatch(left)
		var increment interface{}
		if match == nil {
			return false, left, collected
		}
		leftAlt := make(patternList, len((*left)[:pos]), len((*left)[:pos])+len((*left)[pos+1:]))
		copy(leftAlt, (*left)[:pos])
		leftAlt = append(leftAlt, (*left)[pos+1:]...)
		sameName := patternList{}
		for _, a := range *collected {
			if a.name == p.name {
				sameName = append(sameName, a)
			}
		}

		switch p.value.(type) {
		case int, []string:
			switch p.value.(type) {
			case int:
				increment = 1
			case []string:
				switch match.value.(type) {
				case string:
					increment = []string{match.value.(string)}
				default:
					increment = match.value
				}
			}
			if len(sameName) == 0 {
				match.value = increment
				collectedMatch := make(patternList, len(*collected), len(*collected)+1)
				copy(collectedMatch, *collected)
				collectedMatch = append(collectedMatch, match)
				return true, &leftAlt, &collectedMatch
			}
			switch sameName[0].value.(type) {
			case int:
				sameName[0].value = sameName[0].value.(int) + increment.(int)
			case []string:
				sameName[0].value = append(sameName[0].value.([]string), increment.([]string)...)
			}
			return true, &leftAlt, collected
		}
		collectedMatch := make(patternList, len(*collected), len(*collected)+1)
		copy(collectedMatch, *collected)
		collectedMatch = append(collectedMatch, match)
		return true, &leftAlt, &collectedMatch
	}
	panic("unmatched type")
}

func (p *pattern) singleMatch(left *patternList) (int, *pattern) {
	if p.t&patternArgument != 0 {
		for n, pat := range *left {
			if pat.t&patternArgument != 0 {
				return n, newArgument(p.name, pat.value)
			}
		}
		return -1, nil
	} else if p.t&patternCommand != 0 {
		for n, pat := range *left {
			if pat.t&patternArgument != 0 {
				if pat.value == p.name {
					return n, newCommand(p.name, true)
				}
				break
			}
		}
		return -1, nil
	} else if p.t&patternOption != 0 {
		for n, pat := range *left {
			if p.name == pat.name {
				return n, pat
			}
		}
		return -1, nil
	}
	panic("unmatched type")
}

func (p *pattern) String() string {
	if p.t&patternOption != 0 {
		return fmt.Sprintf("%s(%s, %s, %d, %+v)", p.t, p.short, p.long, p.argcount, p.value)
	} else if p.t&patternLeaf != 0 {
		return fmt.Sprintf("%s(%s, %+v)", p.t, p.name, p.value)
	} else if p.t&patternBranch != 0 {
		result := ""
		for i, child := range p.children {
			if i > 0 {
				result += ", "
			}
			result += child.String()
		}
		return fmt.Sprintf("%s(%s)", p.t, result)
	}
	panic("unmatched type")
}

func (p *pattern) transform() *pattern {
	/*
		Expand pattern into an (almost) equivalent one, but with single Either.

		Example: ((-a | -b) (-c | -d)) => (-a -c | -a -d | -b -c | -b -d)
		Quirks: [-a] => (-a), (-a...) => (-a -a)
	*/
	result := []patternList{}
	groups := []patternList{patternList{p}}
	parents := patternRequired +
		patternOptionAL +
		patternOptionSSHORTCUT +
		patternEither +
		patternOneOrMore
	for len(groups) > 0 {
		children := groups[0]
		groups = groups[1:]
		var child *pattern
		for _, c := range children {
			if c.t&parents != 0 {
				child = c
				break
			}
		}
		if child != nil {
			children.remove(child)
			if child.t&patternEither != 0 {
				for _, c := range child.children {
					r := patternList{}
					r = append(r, c)
					r = append(r, children...)
					groups = append(groups, r)
				}
			} else if child.t&patternOneOrMore != 0 {
				r := patternList{}
				r = append(r, child.children.double()...)
				r = append(r, children...)
				groups = append(groups, r)
			} else {
				r := patternList{}
				r = append(r, child.children...)
				r = append(r, children...)
				groups = append(groups, r)
			}
		} else {
			result = append(result, children)
		}
	}
	either := patternList{}
	for _, e := range result {
		either = append(either, newRequired(e...))
	}
	return newEither(either...)
}

func (p *pattern) eq(other *pattern) bool {
	return reflect.DeepEqual(p, other)
}

func (pl patternList) unique() patternList {
	table := make(map[string]bool)
	result := patternList{}
	for _, v := range pl {
		if !table[v.String()] {
			table[v.String()] = true
			result = append(result, v)
		}
	}
	return result
}

func (pl patternList) index(p *pattern) (int, error) {
	for i, c := range pl {
		if c.eq(p) {
			return i, nil
		}
	}
	return -1, newError("%s not in list", p)
}

func (pl patternList) count(p *pattern) int {
	count := 0
	for _, c := range pl {
		if c.eq(p) {
			count++
		}
	}
	return count
}

func (pl patternList) diff(l patternList) patternList {
	lAlt := make(patternList, len(l))
	copy(lAlt, l)
	result := make(patternList, 0, len(pl))
	for _, v := range pl {
		if v != nil {
			match := false
			for i, w := range lAlt {
				if w.eq(v) {
					match = true
					lAlt[i] = nil
					break
				}
			}
			if match == false {
				result = append(result, v)
			}
		}
	}
	return result
}

func (pl patternList) double() patternList {
	l := len(pl)
	result := make(patternList, l*2)
	copy(result, pl)
	copy(result[l:2*l], pl)
	return result
}

func (pl *patternList) remove(p *pattern) {
	(*pl) = pl.diff(patternList{p})
}

func (pl patternList) dictionary() map[string]interface{} {
	dict := make(map[string]interface{})
	for _, a := range pl {
		dict[a.name] = a.value
	}
	return dict
}
