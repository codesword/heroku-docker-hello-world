APP_NAME = hello-world-example
PROCESS_TYPE = web

all:	test

setup:
	heroku plugins:install heroku-container-registry
	heroku container:login

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o artifact .

image: build
	docker build -t registry.heroku.com/$(APP_NAME)/$(PROCESS_TYPE) .

push: image
	docker push registry.heroku.com/$(APP_NAME)/$(PROCESS_TYPE)

deploy: build
	heroku container:push $(PROCESS_NAME) --app $(APP_NAME)

run:
	docker run registry.heroku.com/$(APP_NAME)/$(PROCESS_TYPE)

clean:
