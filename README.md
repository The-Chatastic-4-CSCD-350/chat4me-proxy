# Chat 4 Me request proxy server
This is a server used for act as a proxy for requests from the [Chat 4 Me app](https://github.com/The-Chatastic-4-CSCD-350/Chat-4-Me) to forward completion requests to the OpenAI servers and return the results

# Configuration
The chat4me-proxy server is configured by editing (or creating) config.json. Currently it is expected to reside in the same directory as the current working directory (so if the executable is run with `./chat4me-proxy`, it will (try to) read `./config.json`.
## Example config.json
```JSON
{
  "apiKey": "youropenaiapikeyhere",
  "organizationID": "youropenaiorganizationidhere",
  "logDir": "./log",
  "verbose": true
}
```

# Instructions for development
1. Install [Vagrant](https://www.vagrantup.com/) and [VirtualBox](https://www.virtualbox.org/)
2. Run `vagrant up`. It wil set up the virtual machine and install everything necessary for it.
3. To access/interact with the virtual machine from the command line, run `vagrant ssh`
4. To access the server from a browser, go to https://192.168.56.4/. You will likely get a warning in your browser about the server using a self-signed certificate. You can ignore this warning for development.
5. To delete the VM, run `vagrant destroy --force`
