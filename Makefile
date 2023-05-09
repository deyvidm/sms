remote := "root@${LINODE_IP}"

all: build run

build: 
	go build -o sms

run: 
	./sms 

host: build
	./sms serve --http="170.187.194.105:8090"

deploy:
	rsync -av -e ssh --exclude-from='.gitignore' . $(remote):/root/sms-backend


