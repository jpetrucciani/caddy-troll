# caddy-troll

[![built in go](https://img.shields.io/badge/built%20in-go-%2301ADD8)](https://go.dev/)

`caddy-troll` is a caddy v2 plugin that allows you to mess with people who may be scanning your server for vulnerabilities. It allows you to send back malformed/confusing responses, spoof your server headers, redirect randomly, and even send back responses that may crash naive clients! **Use at your own risk!**

# Installation

This repo uses [nix](https://nixos.org/download.html) + [direnv](https://direnv.net/) to easily and automatically install dependencies and run caddy with this plugin enabled in an easy way. Once both nix and direnv are installed, run `direnv allow` in the root of the project to install all the required dependencies.

# Building

Use [xcaddy](https://github.com/caddyserver/xcaddy) to build, or use nix!

## xcaddy example:

```bash
xcaddy build --output ./caddy --with github.com/jpetrucciani/caddy-troll@main
```

## nix example:

caddy with caddy-troll already included:

```nix
TODO
```

build your own!

```nix
TODO
```

# How to run

There are two ways to run the project.

1. The `run` command which will rebuild the go caddy plugin when files are changed as well as run the `run-troll` command.
1. The `run-troll` command which will run Caddy in watch mode on the Caddyfile in the conf directory.

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

### `localhost:6666/only_naughty`:

Disables other hacks so only the naughty strings hack is applied
