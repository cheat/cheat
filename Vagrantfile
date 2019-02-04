# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "alpine/alpine64"

  config.vm.provider "virtualbox" do |vb|
    vb.memory = "256"
  end

  config.vm.provision "shell", inline: <<-SHELL
     sudo apk update
     sudo apk add py-pip
     su vagrant && sudo -H pip install docopt pygments termcolor flake8
     cd /vagrant && sudo python setup.py install
  SHELL
end
