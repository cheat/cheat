# Show hit for rules with auto refresh
watch --interval 0 'iptables -nvL | grep -v "0     0"'

# Show hit for rule with auto refresh and highlight any changes since the last refresh
watch -d -n 2 iptables -nvL

# Block the port 902 and we hide this port from nmap.
iptables -A INPUT -i eth0 -p tcp --dport 902 -j REJECT --reject-with icmp-port-unreachable

# Note, --reject-with accept:
#	icmp-net-unreachable
#	icmp-host-unreachable
#	icmp-port-unreachable <- Hide a port to nmap
#	icmp-proto-unreachable
#	icmp-net-prohibited
#	icmp-host-prohibited or
#	icmp-admin-prohibited
#	tcp-reset

# Add a comment to a rule:
iptables ... -m comment --comment "This rule is here for this reason"


# To remove or insert a rule:
# 1) Show all rules
iptables -L INPUT --line-numbers
# OR iptables -nL --line-numbers

# Chain INPUT (policy ACCEPT)
#     num  target prot opt source destination
#     1    ACCEPT     udp  --  anywhere  anywhere             udp dpt:domain
#     2    ACCEPT     tcp  --  anywhere  anywhere             tcp dpt:domain
#     3    ACCEPT     udp  --  anywhere  anywhere             udp dpt:bootps
#     4    ACCEPT     tcp  --  anywhere  anywhere             tcp dpt:bootps

# 2.a) REMOVE (-D) a rule. (here an INPUT rule)
iptables -D INPUT 2

# 2.b) OR INSERT a rule.
iptables -I INPUT {LINE_NUMBER} -i eth1 -p tcp --dport 21 -s 123.123.123.123 -j ACCEPT -m comment --comment "This rule is here for this reason"
