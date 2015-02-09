# Actively follow log (like tail -f)
journalctl -f

# Display all errors since last boot
journalctl -b -p err

# Filter by time period
journalctl --since=2012-10-15 --until="2011-10-16 23:59:59"

# Show list of systemd units logged in journal
journalctl -F _SYSTEMD_UNIT

# Filter by specific unit
journalctl -u dbus

# Filter by executable name
journalctl /usr/bin/dbus-daemon

# Filter by PID
journalctl _PID=123

# Filter by Command, e.g., sshd
journalctl _COMM=sshd

# Filter by Command and time period
journalctl _COMM=crond --since '10:00' --until '11:00'

# List all available boots 
journalctl --list-boots

# Filter by specific User ID e.g., user id 1000 
journalctl _UID=1000
