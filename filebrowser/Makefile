NAME=creator-suite-frontend
VERSION=1.0.0

.PHONY: build start push

build:build-version tag-latest

build-version:
	docker build -t ${NAME}:${VERSION}  .

tag-latest:
	docker tag ${NAME}:${VERSION} ${NAME}:latest

start:
	docker run -it --rm ${NAME}:${VERSION} /bin/bash

push:build-version tag-latest
	docker push ${NAME}:${VERSION}; docker push ${NAME}:latest

prepare:
	@npm install

test:prepare
	@npm run dev

clean:
	@rm -rf dist
