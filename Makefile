.PHONY: test
test:
	go test ./options

.PHONY: docker
docker:
	docker build -t davidoram/kratos-selfservice-ui-go:latest .

.PHONY: run
run:
	go run . --kratos-public-url http://127.0.0.1:4433/ --kratos-browser-url http://127.0.0.1:4433/ --kratos-admin-url http://127.0.0.1:4434/ --base-url / --port 4455

.PHONY: quickstart
quickstart:
	docker-compose -f quickstart.yml -f quickstart-standalone.yml up --build --force-recreate

.PHONY: cypress
cypress:
	cd cypress-tests && npm run cypress:open --browser firefox

.PHONY: cypress-headless
cypress-headless:
	cd cypress-tests && npm run cypress:run --headless --spec "cypress/integration/*.js"
