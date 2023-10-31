remote := "root@${LINODE_IP}"

all: build test

build: build-server build-dispatcher 

build-server: 
	go build -o bin/sms ./web-server

build-dispatcher:
	go build -o bin/dispatcher ./dispatcher

run-server: 
	./bin/sms

run-dispatcher:
	./bin/dispatcher

test:
	echo "tests OK lol"

host: build
	./sms serve --http="170.187.194.105:8090"

deploy:
	rsync -av -e ssh --exclude-from='.gitignore' . $(remote):/root/sms-backend


