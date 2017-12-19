# Initial check-in of file (leaving file active in filesystem)
ci -u <filename>

# Check out with lock
co -l <filename>

# Check in and unlock (leaving file active in filesystem)
ci -u <filename>

# Display version x.y of a file
co -px.y <filename>

# Undo to version x.y (overwrites file active in filesystem with the specified revision)
co -rx.y <filename>

# Diff file active in filesystem and last revision
rcsdiff <filename>

# Diff versions x.y and x.z
rcsdiff -rx.y -rx.z <filename>

# View log of check-ins
rlog <filename>

# Break an RCS lock held by another person on a file
rcs -u <filename>
