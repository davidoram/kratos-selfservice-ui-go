# kratos-selfservice-ui-go
ORY Kratos Self-Service UI written in golang 1.16 using the labstack Echo framework

A self service UI for [Kratos](https://www.ory.sh/kratos) based on the NodeJS version but written in go 1.16.


Provides:

- Registration

# Quickstart

- Run `make quickstart-standalone-up` to run the systems in docker compose
- `open http://127.0.0.1:4455/dashboard`
- Because you are not logged in you will be taken to the login page by default


https://www.cypress.io/blog/2019/05/02/run-cypress-with-a-single-docker-command/


# Cypress tests

The following steps show you how to run individual cypress tests interactively, using the cypress UI.

- In one terminal run `make quickstart-standalone-up`
- In another  run `make cypress`, and then choose the tests to run.
