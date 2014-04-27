# TCPDump is a packet analyzer. It allows the user to intercept and display TCP/IP
# and other packets being transmitted or received over a network. (cf Wikipedia).
# Note: 173.194.40.120 => google.com

# Intercepts all packets on eth0
tcpdump -i eth0

# Intercepts all packets from/to 173.194.40.120
tcpdump host 173.194.40.120

# Intercepts all packets on all interfaces from / to 173.194.40.120 port 80
# -nn => Disables name resolution for IP addresses and port numbers.
tcpdump -nn -i any host 173.194.40.120 and port 80

# Make a grep on tcpdump (ASCII)
# -A  => Show only ASCII in packets.
# -s0 => By default, tcpdump only captures 68 bytes.
tcpdump -i -A any host 173.194.40.120 and port 80 | grep 'User-Agent'

# With ngrep
# -d eth0 => To force eth0 (else ngrep work on all interfaces)
# -s0 => force ngrep to look at the entire packet. (Default snaplen: 65536 bytes)
ngrep 'User-Agent' host 173.194.40.120 and port 80

# Intercepts all packets on all interfaces from / to 8.8.8.8 or 173.194.40.127 on port 80
tcpdump 'host ( 8.8.8.8 or 173.194.40.127 ) and port 80' -i any

# Intercepts all packets SYN and FIN of each TCP session.
tcpdump 'tcp[tcpflags] & (tcp-syn|tcp-fin) != 0'

# To display SYN and FIN packets of each TCP session to a host that is not on our network
tcpdump 'tcp[tcpflags] & (tcp-syn|tcp-fin) != 0 and not src and dst net local_addr'

# To display all IPv4 HTTP packets that come or arrive on port 80 and that contain only data (no SYN, FIN no, no packet containing an ACK)
tcpdump 'tcp port 80 and (((ip[2:2] - ((ip[0]&0xf)<<2)) - ((tcp[12]&0xf0)>>2)) != 0)'

# Saving captured data
tcpdump -w file.cap

# Reading from capture file
tcpdump -r file.cap

# Show content in hexa
# Change -x to -xx => show extra header (ethernet).
tcpdump -x

# Show content in hexa and ASCII
# Change -X to -XX => show extra header (ethernet).
tcpdump -X

# Note on packet maching:
# Port matching:
# - portrange 22-23
# - not port 22
# - port ssh
# - dst port 22
# - src port 22
#
# Host matching:
# - dst host 8.8.8.8
# - not dst host 8.8.8.8
# - src net 67.207.148.0 mask 255.255.255.0
# - src net 67.207.148.0/24
