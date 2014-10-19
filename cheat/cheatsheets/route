# To display routing table IP addresses instead of host names:
route -n

# To add a default gateway:
route add default gateway 192.168.0.1

# To add the normal loopback entry, using netmask 255.0.0.0 and associated with the "lo" device (assuming this device was previously set up correctly with ifconfig(8)).
route add -net 127.0.0.0 netmask 255.0.0.0 dev lo

# To add a route to the local network 192.56.76.x via "eth0".  The word "dev" can be omitted here.
route add -net 192.56.76.0 netmask 255.255.255.0 dev eth0

# To delete the current default route, which is labeled "default" or 0.0.0.0 in the destination field of the current routing table.
route del default

# To add a default  route (which will be used if no other route matches).  All packets using this route will be gatewayed through "mango-gw". The device which will actually be used for that route depends on how we can reach "mango-gw" - the static route to "mango-gw" will have to be set up before.
route add default gw mango-gw

# To add the route to the "ipx4" host via the SLIP interface (assuming that "ipx4" is the SLIP host).
route add ipx4 sl0

# To add the net "192.57.66.x" to be gateway through the former route to the SLIP interface.
route add -net 192.57.66.0 netmask 255.255.255.0 gw ipx4

# To install a rejecting route for the private network "10.x.x.x."
route add -net 10.0.0.0 netmask 255.0.0.0 reject

# This is an obscure one documented so people know how to do it. This sets all of the class D (multicast) IP routes to go via "eth0". This is the correct normal configuration line with a multicasting kernel
route add -net 224.0.0.0 netmask 240.0.0.0 dev eth0
