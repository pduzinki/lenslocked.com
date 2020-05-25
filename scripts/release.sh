#!/bin/bash

cd "/home/pduzinki/go/src/lenslocked.com"

echo "=== releasing lenslocked.com ==="
echo "  deleting the local binary if it exists..."
rm lenslocked.com
echo "  done."

echo "  deleting existing code..."
ssh root@142.93.100.106 "rm -rf /root/go/src/lenslocked.com"
echo "  code deleted successfully."

echo "  uploading code..."
rsync -avr --exclude '.git/*' --exclude 'tmp/*' \
  --exclude 'images/*' ./ \
  root@142.93.100.106:/root/go/src/lenslocked.com/

echo "  code uploaded successfully."

echo "  go getting dependencies..."
ssh root@142.93.100.106 "export GOPATH=/root/go; \
  /usr/local/go/bin/go get golang.org/x/crypto/bcrypt"
ssh root@142.93.100.106 "export GOPATH=/root/go; \
  /usr/local/go/bin/go get github.com/gorilla/mux"
ssh root@142.93.100.106 "export GOPATH=/root/go; \
  /usr/local/go/bin/go get github.com/gorilla/schema"
ssh root@142.93.100.106 "export GOPATH=/root/go; \
  /usr/local/go/bin/go get github.com/lib/pq"
ssh root@142.93.100.106 "export GOPATH=/root/go; \
  /usr/local/go/bin/go get github.com/jinzhu/gorm"
ssh root@142.93.100.106 "export GOPATH=/root/go; \
  /usr/local/go/bin/go get github.com/gorilla/csrf"

echo "  building the code on remote server..."
ssh root@142.93.100.106 'export GOPATH=/root/go; \
  cd /root/app; \
  /usr/local/go/bin/go build -o ./server \
  /root/go/src/lenslocked.com/*.go'
echo "  code built successfully."

echo "  moving assets..."
ssh root@142.93.100.106 "cd /root/app; \
  cp -R /root/go/src/lenslocked.com/assets ."
echo "  assets moved successfully."

echo "  moving views..."
ssh root@142.93.100.106 "cd /root/app; \
  cp -R /root/go/src/lenslocked.com/views ."
echo "  views moved successfully."

echo "  moving Caddyfile..."
ssh root@142.93.100.106 "cd /root/app; \
  cp /root/go/src/lenslocked.com/Caddyfile ."
echo "  caddyfile moved successfully."

echo "  restarting the server..."
ssh root@142.93.100.106 "sudo service lenslocked.com restart"
echo "  server restarted successfully."

echo "  restarting Caddy server..."
ssh root@142.93.100.106 "sudo service caddy restart"
echo "  Caddy restarted successfully."

echo "=== done releasing lenslocked.com ==="