#!/usr/bin/env bash

set -euo pipefail
export DEBIAN_FRONTEND=noninteractive
export GO_VERSION=1.18

add-apt-repository -y ppa:longsleep/golang-backports
apt-get -y update && apt-get -y upgrade

apt-get -y install git nginx golang-${GO_VERSION}
ln -sf /usr/lib/go-${GO_VERSION}/bin/* /usr/local/bin/

rm -f /etc/nginx/sites-enabled/* /etc/nginx/sites-available/*
ln -sf /vagrant/vagrant/chat4me-server.nginx /etc/nginx/sites-enabled/
rm -rf /var/www/*
ln -sf /vagrant/html /var/www/

# VirtualBox shared folders don't play nicely with sendfile.
sed -e 's/sendfile on;/sendfile off;/' -i /etc/nginx/nginx.conf

# Make sure our shared directories are mounted before nginx starts
systemctl disable nginx
sed -i 's/WantedBy=multi-user.target/WantedBy=vagrant.mount/' /lib/systemd/system/nginx.service

# generate self-signed certificate since some browsers like Firefox and Chrome automatically do HTTPS requests
# this will likely show a warning in the browser, which you can ignore
openssl req -x509 -nodes -days 7305 -newkey rsa:2048 \
	-keyout /etc/ssl/private/nginx-selfsigned.key \
	-out /etc/ssl/certs/nginx-selfsigned.crt \
	-subj "/CN=192.168.56.4"

systemctl daemon-reload
systemctl enable nginx
systemctl restart nginx &
wait

cd /vagrant/chat4me-router
go build

echo "Done setting up chat4me-router. You can access it for development at https://192.168.56.4/"