REPO?=onuryilmaz/boilerplate-go
VERSION=latest

EXPOSED_PORT= "8080"

.PHONY: all

build:
	docker build --no-cache -t $(REPO):$(VERSION) .

push:
	docker push $(REPO):$(VERSION)

run:
	docker run --rm -p $(EXPOSED_PORT):8080 $(REPO):$(VERSION) --log-level=debug --port=8080

test:
	docker build --no-cache --rm --target=tester .