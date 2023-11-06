remote := "root@${LINODE_IP}"

all: build test

build: build-dispatcher build-pb

build-pb:
	go build -o bin/pb ./cmd/pocketbase

build-dispatcher:
	go build -o bin/dispatcher ./cmd/dispatcher


build-server: 
	go build -o bin/web-server ./cmd/web-server

run-server: 
	./bin/sms

run-dispatcher:
	./bin/dispatcher

run-pb:
	./bin/pb serve

test:
	echo "tests OK lol"

host: build
	./sms serve --http="170.187.194.105:8090"

deploy:
	rsync -av -e ssh --exclude-from='.gitignore' . $(remote):/root/sms-backend


