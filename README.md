# kratos-selfservice-ui-go
ORY Kratos Self-Service UI written in golang 1.16.

A self service UI for [Kratos](https://www.ory.sh/kratos) based on the NodeJS version but written in go 1.16.

The application provides the following self service UI pages:

- Registration
- Login
- Logout
- User settings
  - Update profile
  - Change password
  - Password reset

and then once logged in provides the following additional page:

- Dashboard

# Quickstart

- Start docker
- Run `make quickstart` to run the systems in docker compose
- `open http://127.0.0.1`
- Because you are not logged in you will be taken to the login page by default

# Cypress tests

The following steps show you how to run individual cypress tests interactively, using the cypress UI.

- In one terminal run `make quickstart`
- In another  run `make cypress`, and then choose the tests to run.

# Tailwind CSS

To create the initial css file:

```
nvm use
npx tailwindcss-cli@latest build -o static_src/css/tailwind.css
```

Static assets served via [HashFS](https://github.com/benbjohnson/hashfs) that appends hashes to embedded static assets for aggresive HTTP caching.

# Icons

Icons from [Bootstrap](https://icons.getbootstrap.com/)

# Stimulus JS

- See [Reference](https://stimulus.hotwire.dev/reference/controllers)
- Used to show/hide the mobile menu when clicked

# TODO
 - Add gzip middleware
 - Fix traefik warnings "level=warning msg="Could not find network named 'internal' for container '/kratos-selfservice-ui-go_kratos-selfservice-ui-go_1'! Maybe you're missing the project's prefix in the label? Defaulting to first available network." container=kratos-selfservice-ui-go-kratos-selfservice-ui-go-2fd978669efd2f580e1ac7fcb67271ea7d966fbcdc75a2498c15786af0ff702d serviceName=kratos-selfservice-ui-go-kratos-selfservice-ui-go providerName=docker"

