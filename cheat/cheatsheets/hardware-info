# Display all hardware details
sudo lshw

# List currently loaded kernel modules
lsmod

# List all modules available to the system
find /lib/modules/$(uname -r) -type f -iname "*.ko"

# Load a module into kernel
modprobe modulename

# Remove a module from kernel 
modprobe -r modulename

# List devices connected via pci bus
lspci

# Debug output for pci devices (hex)
lspci -vvxxx

# Display cpu hardware stats
cat /proc/cpuinfo

# Display memory hardware stats
cat /proc/meminfo

# Output the kernel ring buffer
dmesg

# Ouput kernel messages
dmesg --kernel
