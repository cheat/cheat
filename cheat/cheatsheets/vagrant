# Initate Vagrant
mkdir vag-vm; cd vag-vm
vagrant init

# Add a box to vagrant repo
vagrant box add hashicorp/precise32

# Add a box  Vagrant file
config.vm.box = "hashicorp/precise32"

# Add vm to public network as host
config.vm.network "public_network"

# Add provision script to vagrant file
config.vm.provision :shell, path: "provision.sh"

# Start vm 
vagrant up

# Connect to started instance
vagrant ssh

# Shutdown vm
vagrant halt

# Hibernate vm
vagrant suspend

# Set vm to initial state by cleaning all data
vagrant destroy

# Restart vm with new provision script
vagran reload --provision
