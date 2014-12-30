# Display all interfaces with addresses
ip addr

# Take down / up the wireless adapter
ip link set dev wlan0 {up|down}

# Set a static IP and netmask
ip addr add 192.168.1.100/32 dev eth0

# Remove a IP from an interface
ip addr del 192.168.1.100/32 dev eth0

# Remove all IPs from an interface
ip address flush dev eth0

# Display all routes
ip route

# Display all routes for IPv6
ip -6 route

# Add default route via gateway IP
ip route add default via 192.168.1.1

# Add route via interface
ip route add 192.168.0.0/24 dev eth0

# Change your mac address 
ip link set dev eth0 address aa:bb:cc:dd:ee:ff

# View neighbors (using ARP and NDP) 
ip neighbor show
