#!/bin/bash
mkdir /srv/curso-go-web
cd /srv/curso-go-web
chmod +x curso-go-web-linux-binary
cp curso-go-web.service /lib/systemd/system
chmod 755 /lib/systemd/system/curso-go-web.service
systemctl enable curso-go-web.service
systemctl start curso-go-web.service
systemctl stop curso-go-web.service
systemctl status curso-go-web.service
apt install nginx
add-apt-repository ppa:certbot/certbot
certbot --nginx -d vpsdfddf.publiccloud.com.br
ufw app list
ufw allow 'Nginx Full'

ls -l /etc/nginx/sites-available
vim /etc/nginx/sites-available/default

# Conteudo para o arquivo /etc/nginx/sites-available/default
server {
	listen 80;
	return 301 https://$host$request_uri;
}

server {
	listen 443 ssl;
	listen [::]:443 ssl;

	server_name vpsdddfdf.publiccloud.com.br;

	ssl_certificate /etc/letsencrypt/libs/vpsdddfdf.publiccloud.com.br/fullchain.pem
	ssl_certificate_key /etc/letsencrypt/libs/vpsdddfdf.publiccloud.com.br/prvtkey.pem
	include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
	ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

	location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header X-Forwarded-Protocol $scheme;
        proxy_set_header X-Forwarded-Host $http_host;
	}
}
