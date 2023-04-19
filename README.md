# Chat 4 Me request proxy server
This is a server used for act as a proxy for requests from the Chat 4 Me app to forward completion requests to the OpenAI servers and return the results

# Instructions for development
1. Install [Vagrant](https://www.vagrantup.com/) and [VirtualBox](https://www.virtualbox.org/)
2. Run `vagrant up`. It wil set up the virtual machine and install everything necessary for it.
3. To access/interact with the virtual machine from the command line, run `vagrant ssh`
4. To access the server from a browser, go to https://192.168.56.4/. You will likely get a warning in your browser about the server using a self-signed certificate. You can ignore this warning for development.
5. To delete the VM, run `vagrant destroy --force`
