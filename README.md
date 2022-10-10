# caddy-troll

[![built in go](https://img.shields.io/badge/built%20in-go-%2301ADD8)](https://go.dev/)

# Installation

This repo uses [Nix](https://nixos.org/download.html) + [Direnv](https://direnv.net/) to easily and automatically install and run the project. Once both are installed, run `direnv allow` in the root of the project to install all the required dependencies.

# How to run

There are two ways to run the project.

1. The `run-troll` command which will run Caddy in watch mode on the Caddyfile in the conf directory.
2. The `run` command which will rebuild the go caddy plugin when files are changed as well as run the `run-troll` command.

## Current Hacks

The local server runs on `localhost:6666`. Some of the hacks can be run in isolation using different routes. Here is the current list of supported routes.

### `localhost:6666/`

Responds with the string "test" to check the server is running correctly.

### `localhost:6666/random_server_header`:

Sets server headers designed to confused people by lying.

For example, we may set the Server header to "nginx" when this server is actually using Caddy.

### `localhost:6666/not_random_server_header`:

Disables the random server header hack

### `localhost:6666/only_gzip`:

Disables other hacks so only the gzip hack is applied

### `localhost:6666/only_redirect`:

Disables other hacks so only the redirect hack is applied

### `localhost:6666/only_xml`:

Disables other hacks so only the xml hack is applied
