# find all file name ending with .pdf and remove them
find -name *.pdf | xargs rm -rf

# if file name contains spaces you should use this instead
find -name *.pdf | xargs -I{} rm -rf '{}'

# Will show every .pdf like:
#	&toto.pdf=
#	&titi.pdf=
# -n1 => One file by one file. ( -n2 => 2 files by 2 files )

find -name *.pdf | xargs -I{} -n1 echo '&{}='
