# To list VMs on current tenant:
nova list

# To list VMs of all tenants (admin user only):
nova list --all-tenants

# To boot a VM on a specific host:
nova boot --nic net-id=<net_id> \
          --image <image_id> \
          --flavor <flavor> \
          --availability-zone nova:<host_name> <vm_name>

# To stop a server
nova stop <server>

# To start a server
nova start <server>

# To attach a network interface to a specific VM:
nova interface-attach --net-id <net_id> <server>
