# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
	config.vm.box = "ubuntu/focal64"

	config.vm.network "forwarded_port", guest: 80, host: 80
	config.vm.network "forwarded_port", guest: 443, host: 443

	config.vm.network "private_network", ip: "192.168.56.4"

	# config.vm.synced_folder "../data", "/vagrant_data"

	# Provider-specific configuration so you can fine-tune various
	# backing providers for Vagrant. These expose provider-specific options.
	# Example for VirtualBox:
	#
	# config.vm.provider "virtualbox" do |vb|
	#   # Display the VirtualBox GUI when booting the machine
	#   vb.gui = true
	#
	#   # Customize the amount of memory on the VM:
	#   vb.memory = "1024"
	# end
	#
	# View the documentation for the provider you are using for more
	# information on available options.

	config.vm.provision :shell, path: "vagrant/bootstrap.sh", env: {
		
	}
end
