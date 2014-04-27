# List all IPv4 network files
sudo lsof -i4

# List all IPv6 network files
sudo lsof -i6

# To find listening ports:
lsof -Pnl +M -i4

# To find which program is using the port 80:
lsof -i TCP:80

# List all processes accessing a particular file/directory
lsof </path/to/file>

# List all files open for a particular user
lsof -u <username>

# List all files/network connections a given process is using
lsof -c <command-name>

# See this primer: http://www.danielmiessler.com/study/lsof/
# for a number of other useful lsof tips
