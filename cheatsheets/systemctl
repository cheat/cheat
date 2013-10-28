# List all loaded/active units
systemctl list-units

# Check the status of a service
systemctl status foo.service

# Start a service
systemctl start foo.service

# Restart a service
systemctl restart foo.service

# Stop a service
systemctl stop foo.service

# Reload a service's configuration
systemctl reload foo.service

# Enable a service to startup on boot
systemctl enable foo.service

# Disable a service to startup on boot
systemctl disable foo.service

# List the dependencies of a service
# when no service name is specified, lists the dependencies of default.target
systemctl list-dependencies foo.service 

# List currently loaded targets
systemctl list-units --type=target

# Change current target
systemctl isolate foo.target

# Change default target
systemctl enable foo.target
