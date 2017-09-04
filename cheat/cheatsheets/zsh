# A plain old glob
print -l *.txt
print -l **/*.txt

# Show text files that end in a number from 1 to 10
print -l **/*<1-10>.txt

# Show text files that start with the letter a
print -l **/[a]*.txt

# Show text files that start with either ab or bc
print -l **/(ab|bc)*.txt

# Show text files that don't start with a lower or uppercase c
print -l **/[^cC]*.txt

# Show only directories
print -l **/*(/)

# Show only regular files
print -l **/*(.)

# Show empty files
print -l **/*(L0)

# Show files greater than 3 KB
print -l **/*(Lk+3)

# Show files modified in the last hour
print -l **/*(mh-1)

# Sort files from most to least recently modified and show the last 3
print -l **/*(om[1,3])

# `.` show files, `Lm-2` smaller than 2MB, `mh-1` modified in last hour,
# `om` sort by modification date, `[1,3]` only first 3 files
print -l **/*(.Lm-2mh-1om[1,3])

# Show every directory that contain directory `.git`
print -l **/*(e:'[[ -d $REPLY/.git ]]':)

# Return the file name (t stands for tail)
print -l *.txt(:t)

# Return the file name without the extension (r stands for remove_extension)
print -l *.txt(:t:r)

# Return the extension
print -l *.txt(:e)

# Return the parent folder of the file (h stands for head)
print -l *.txt(:h)

# Return the parent folder of the parent
print -l *.txt(:h:h)

# Return the parent folder of the first file
print -l *.txt([1]:h)

# Parameter expansion
files=(*.txt)          # store a glob in a variable
print -l $files
print -l $files(:h)    # this is the syntax we saw before
print -l ${files:h}
print -l ${files(:h)}  # don't mix the two, or you'll get an error
print -l ${files:u}    # the :u modifier makes the text uppercase

# :s modifier
variable="path/aaabcd"
echo ${variable:s/bc/BC/}    # path/aaaBCd
echo ${variable:s_bc_BC_}    # path/aaaBCd
echo ${variable:s/\//./}     # path.aaabcd (escaping the slash \/)
echo ${variable:s_/_._}      # path.aaabcd (slightly more readable)
echo ${variable:s/a/A/}      # pAth/aaabcd (only first A is replaced)
echo ${variable:gs/a/A/}     # pAth/AAAbcd (all A is replaced)

# Split the file name at each underscore
echo ${(s._.)file}

# Join expansion flag, opposite of the split flag.
array=(a b c d)
echo ${(j.-.)array} # a-b-c-d
