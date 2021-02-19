.PHONY: test
test:
	go test ./options

.PHONY: docker
docker:
	docker build -t davidoram/kratos-selfservice-ui-go:latest .

.PHONY: run
run:
	go run . --kratos-public-url http://127.0.0.1:4433/ --kratos-browser-url http://127.0.0.1:4433/ --kratos-admin-url http://127.0.0.1:4433/ --base-url / --port 4455

.PHONY: quickstart-standalone-up quickstart-standalone-down
quickstart-standalone-up:
	# docker pull oryd/kratos:latest-sqlite
	# docker pull davidoram/kratos-selfservice-ui-go:latest
	docker-compose -f quickstart.yml -f quickstart-standalone.yml up --build --force-recreate

quickstart-standalone-down:
	docker-compose -f quickstart.yml -f quickstart-standalone.yml down

.PHONY: cypress
cypress:
	cd cypress-tests && npm run cypress:open
