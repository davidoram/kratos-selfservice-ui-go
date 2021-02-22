# kratos-selfservice-ui-go
ORY Kratos Self-Service UI written in golang 1.16 using the labstack Echo framework

A self service UI for [Kratos](https://www.ory.sh/kratos) based on the NodeJS version but written in go 1.16.


Provides:

- Registration

# Quickstart

- Run `make quickstart` to run the systems in docker compose
- `open http://127.0.0.1:4455/dashboard`
- Because you are not logged in you will be taken to the login page by default
https://gist.github.com/Ocramius/d44b9c500d3fd19e863c621500adeec0

# Cypress tests

The following steps show you how to run individual cypress tests interactively, using the cypress UI.

- In one terminal run `make quickstart`
- In another  run `make cypress`, and then choose the tests to run.
