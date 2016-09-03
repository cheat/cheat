# Patch one file
patch version1 < version.patch

# Reverse a patch
patch -R version1 < version.patch

# Patch all files in a directory, adding any missing new files
# -p strips leading slashes
$ cd dir
$ patch -p1 -i ../big.patch

# Patch files in a directory, with one level (/) offset
patch -p1 -r version1/ < version.patch
