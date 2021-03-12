.PHONY: test
test:
	go test ./options

.PHONY: docker
docker:
	docker build -t davidoram/kratos-selfservice-ui-go:latest .

clean:
	rm -rf static
	mkdir -p static/images static/css

build-css: static_src/css/* tailwind.config.js
	npx tailwindcss-cli@latest build ./static_src/css/tailwind.css -o ./static/css/tailwind.css

copy-images: static_src/images/*
	mkdir -p static/images
	cp -r static_src/images/ static/images/

.PHONY: run gen-keys compile-docker
run: clean build-css copy-images
	tree static
	go run . --kratos-public-url http://127.0.0.1:4433/ \
		--kratos-browser-url http://127.0.0.1:4433/ \
		--kratos-admin-url http://127.0.0.1:4434/ \
		--base-url / \
		--port 4455 \
		--cookie-store-key-pairs '6QKIvm1ZwLD+hrS6zysrs50a8gOU8O385BkVEDdlDN0= 2m/+Pva16CPu3pDs4DLfmR7q74WmI0Bv+3bxdUtHmSQ='
gen-keys:
	go run . --gen-cookie-store-key-pair

.PHONY: quickstart
quickstart:
	docker-compose -f quickstart.yml -f quickstart-standalone.yml up --build --force-recreate

.PHONY: cypress
cypress:
	cd cypress-tests && npm run cypress:open --browser firefox

.PHONY: cypress-headless
cypress-headless:
	cd cypress-tests && npm run cypress:run --headless --spec "cypress/integration/*.js"

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

