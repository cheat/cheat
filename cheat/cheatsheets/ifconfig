# Display network settings of the first ethernet adapter
ifconfig wlan0

# Display all interfaces, even if down
ifconfig -a

# Take down / up the wireless adapter
ifconfig wlan0 {up|down} 

# Set a static IP and netmask
ifconfig eth0 192.168.1.100 netmask 255.255.255.0

# You may also need to add a gateway IP
route add -net 192.168.1.0 netmask 255.255.255.0 gw 192.168.1.1
