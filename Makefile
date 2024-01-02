remote := "root@${LINODE_IP}"



### docker build stuff
pb-docker: clean-pb-docker build-pb-docker run-pb-docker

clean-pb-docker:
	-docker container stop pocketbase
	-docker container rm pocketbase

build-pb-docker:
	docker build --target pocketbase -t pocketbase -f Dockerfile-go .

run-pb-docker:
	docker run -v ./pb_data:/app/pb_data -d -p 8090:8090 --name pocketbase pocketbase


dispatcher-docker: build-dispatcher-docker run-dispatcher-docker

build-dispatcher-docker:
	docker build --target dispatcher -t dispatcher -f Dockerfile-go .

run-dispatcher-docker:
	docker run -d -p 8080:8080 --name dispatcher dispatcher


### local build stuff
build-local: build-dispatcher build-pb

build-pb:
	go build -o bin/pb ./cmd/pocketbase

build-dispatcher:
	go build -o bin/dispatcher ./cmd/dispatcher

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
