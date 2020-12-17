PROJECT = k8s-dev-cwxstat
NAME = gotls
TAG = dev


docker-build:
	docker build --no-cache -t gcr.io/$(PROJECT)/$(NAME):$(TAG) -f Dockerfile .

push:
	docker push gcr.io/$(PROJECT)/$(NAME):$(TAG)

build:
	go build -v .

run:
	docker run --rm -it -p 3000:3000  gcr.io/$(PROJECT)/$(NAME):$(TAG)
