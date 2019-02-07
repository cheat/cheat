# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/bionic64"

  config.vm.provider "virtualbox" do |vb|
    vb.memory = "512"
  end

  config.vm.provision "shell", privileged: false, inline: <<-SHELL
     sudo apt-get update
     sudo apt-get install -y python-pip
     sudo -H pip install flake8
     pip install --user docopt pygments termcolor
     cd /vagrant/ && python setup.py install --user
     echo 'export PATH=$PATH:/home/vagrant/.local/bin' >> /home/vagrant/.bashrc
  SHELL

end
