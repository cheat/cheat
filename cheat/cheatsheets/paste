# Concat columns from files
paste file1 file2 ...

# List the files in the current directory in three columns:
ls | paste - - -

# Combine pairs of lines from a file into single lines:
paste -s -d '\t\n' myfile

# Number the lines in a file, similar to nl(1):
sed = myfile | paste -s -d '\t\n' - -

# Create a colon-separated list of directories named bin,
# suitable for use in the PATH environment variable:
find / -name bin -type d | paste -s -d : -