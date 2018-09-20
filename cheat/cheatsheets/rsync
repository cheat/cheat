# To copy files from remote to local, maintaining file properties and sym-links (-a), zipping for faster transfer (-z), verbose (-v).  
rsync -avz host:file1 :file1 /dest/
rsync -avz /source host:/dest

# Copy files using checksum (-c) rather than time to detect if the file has changed. (Useful for validating backups). 
rsync -avc /source/ /dest/

# Copy contents of /src/foo to destination:

# This command will create /dest/foo if it does not already exist
rsync -auv /src/foo /dest

# Explicitly copy /src/foo to /dest/foo
rsync -auv /src/foo/ /dest/foo

# Copy file from local to remote over ssh with non standard port 1234 to destination folder in remoteuser's home directory
rsync -avz -e "ssh -p1234" /source/file1 remoteuser@X.X.X.X:~/destination/
