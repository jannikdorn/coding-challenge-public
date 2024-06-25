#!/bin/sh
sed -i "s|__BACKEND_URL__|${BACKEND_URL}|g" /usr/share/nginx/html/static/js/app.js
# Start Nginx
nginx -g 'daemon off;'