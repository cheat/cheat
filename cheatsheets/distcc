# INSTALL
# ==============================================================================
# Edit /etc/default/distcc and set theses vars
# STARTDISTCC="true"
# ALLOWEDNETS="127.0.0.1 192.168.1.0/24"# Your computer and local computers
# #LISTENER="127.0.0.1"# Comment it
# ZEROCONF="true"# Auto configuration

# REMEMBER 1:
# Start/Restart your distccd servers before using one of these commands.
# service distccd start

# REMEMBER 2:
# Do not forget to install on each machine DISTCC.
# No need to install libs ! Only main host need libs !

# USAGE
# ==============================================================================

# Run make with 4 thread (a cross network) in auto configuration.
# Note: for gcc, Replace CXX by CC and g++ by gcc
ZEROCONF='+zeroconf' make -j4 CXX='distcc g++'

# Run make with 4 thread (a cross network) in static configuration (2 ip)
# Note: for gcc, Replace CXX by CC and g++ by gcc
DISTCC_HOSTS='127.0.0.1 192.168.1.69' make -j4 CXX='distcc g++'

# Show hosts aviables
ZEROCONF='+zeroconf' distcc --show-hosts
