ss
Utility to investigate sockets

Args
-4/-6 list ipv4/ipv6 sockets
-n numeric addresses instead of hostnames
-l list listing sockets
-u/-t/-x list udp/tcp/unix sockets
-p Show process(es) that using socket

# show all listing tcp sockets including the corresponding process
ss -tlp

# show all sockets connecting to 192.168.2.1 on port 80
ss -t dst 192.168.2.1:80

# show all ssh related connection
ss -t state established '( dport = :ssh or sport = :ssh )'
