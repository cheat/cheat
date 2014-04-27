# WARNING ! netstat is deprecated. Look below.

# To view which users/processes are listening to which ports:
sudo netstat -lnptu

# To view routing table (use -n flag to disable DNS lookups):
netstat -r

# Which process is listening to port <port>
netstat -pln | grep <port> | awk '{print $NF}'

Example output: 1507/python

# Fast display of ipv4 tcp listening programs
sudo netstat -vtlnp --listening -4

# WARNING ! netstat is deprecated.
# Replace it by:
ss

# For netstat-r
ip route

# For netstat -i
ip -s link

# For netstat-g
ip maddr
