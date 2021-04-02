.PHONY: test
test:
	go test ./options

# Compile the app and package in a docker image
#
.PHONY: docker docker-push
docker:
	docker build -t davidoram/kratos-selfservice-ui-go:latest .

# Push the application image to dockerhub
#
docker-push:
	docker image push davidoram/kratos-selfservice-ui-go:latest

clean:
	rm -rf static
	mkdir -p static/images static/css

build-css: static_src/css/* tailwind.config.js
	npx tailwindcss-cli@latest build ./static_src/css/tailwind.css -o ./static/css/tailwind.css

copy-images: static_src/images/*
	mkdir -p static/images
	cp -r static_src/images/ static/images/

.PHONY: fastrun gen-keys compile-docker
fastrun:
	go run . --kratos-public-url http://127.0.0.1:4433/ \
		--kratos-browser-url http://127.0.0.1:4433/ \
		--kratos-admin-url http://127.0.0.1:4434/ \
		--base-url / \
		--port 4455 \
		--cookie-store-key-pairs '6QKIvm1ZwLD+hrS6zysrs50a8gOU8O385BkVEDdlDN0= 2m/+Pva16CPu3pDs4DLfmR7q74WmI0Bv+3bxdUtHmSQ='

# Run the app standalone
run: clean build-css copy-images fastrun

# Generate keys for secure cookie management
#
gen-keys:
	go run . --gen-cookie-store-key-pair


# To get started with Tailwind
.PHONY: install-tailwind
install-tailwind:
	npm install tailwindcss tailwindcss-cli@latest @tailwindcss/forms


# Run up the application
#
.PHONY: quickstart
quickstart:
	docker-compose -f quickstart.yml \
		-f quickstart-standalone.yml \
		up \
		--build \
		--force-recreate

# Helper targets to open veruious pages on the system, after you have run 'make quickstart'
#
.PHONY: open-mail
open-mail:
	open http://localhost:8025

.PHONY: open-traefik
open-traefik:
	open http://localhost:8080

.PHONY: open-app
open-app:
	open http://localhost:4455

.PHONY: open-all
open-all: open-mail open-traefik open-app

# Build a docker image containing cypress, plus the extra node modules
# used by tests
#
.PHONY: cypress-docker cypress-docker-push
cypress-docker:
	rm -rf $(CURDIR)/cypress-tests/node_modules
	mkdir -p $(CURDIR)/cypress-tests/node_modules
	cd cypress-tests && docker build -t davidoram/kratos-go-cypress:latest .

# Push the custom cypress docker image to dockerhub
#
cypress-docker-push:
	docker image push davidoram/kratos-go-cypress:latest

# Run cypress interactively.
#
# Pre-req is X server
# On Mac install XQuartz https://www.xquartz.org/
# Enable 'Security Preferences' > 'Allow connections from network clients'
#
# I followed these instructions to get the tests running over X
# https://medium.com/@mreichelt/how-to-show-x11-windows-within-docker-on-mac-50759f4b65cb
#
.PHONY: cypress
cypress:
	xhost + 127.0.0.1
	docker run -it --network="host" -v $(CURDIR)/cypress-tests:/e2e -w /e2e -v /tmp/.X11-unix:/tmp/.X11-unix -e DISPLAY=host.docker.internal:0 --entrypoint cypress davidoram/kratos-go-cypress:latest open --project .

# Run the cypress tests in 'headless' mode, good for CI/CD.
#
# Note: Use --network="host" in your docker run command, then 127.0.0.1 in your
#       docker container will point to your docker host
#
.PHONY: cypress-headless
cypress-headless:
	docker run -it --network="host" -v $(CURDIR)/cypress-tests:/e2e -w /e2e davidoram/kratos-go-cypress:latest

# List the browsers in the cypress image
#
.PHONY: cypress-info
cypress-info:
	docker run -it -v $(CURDIR)/cypress-tests:/e2e -w /e2e --entrypoint=cypress davidoram/kratos-go-cypress:latest info

