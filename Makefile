REGISTRY = localhost:5000
VERSION = 0.0.1

IMG = $(REGISTRY)/from-helm:$(VERSION)

.PHONY: build push

build:
	docker build -t $(IMG) .

push:
	docker push $(IMG)