#replace : with new line
echo $PATH|tr ":" "\n" #equivalent with:
echo $PATH|tr -t ":" \n 

#remove all occurance of "ab"
echo aabbcc |tr -d "ab"
#ouput: cc

#complement "aa"
echo aabbccd |tr -c "aa" 1
#output: aa11111 without new line
#tip: Complement meaning keep aa,all others are replaced with 1

#complement "ab\n"
echo aabbccd |tr -c "ab\n" 1
#output: aabb111 with new line

#Preserve all alpha(-c). ":-[:digit:] etc" will be translated to "\n". sequeeze mode.
echo $PATH|tr -cs "[:alpha:]" "\n" 

#ordered list to unordered list
echo "1. /usr/bin\n2. /bin" |tr -cs " /[:alpha:]\n" "+"
