# eventual

A simple event calendar

## First Time Setup

You'll need to install the following tools:
- [Go](https://go.dev/)
- [Node](https://nodejs.org/)
- [TypeScript](https://www.typescriptlang.org/)

Optionally, you may need to install:
- [Docker](https://www.docker.com/)
- [flyctl](https://fly.io/docs/flyctl/)

Clone the repo and install the dependencies:

```sh
go mod tidy && npm install
```

## Local Development

The client, server, and dependent services are built and run separately. To run them, open a terminal and run the next set of commands in separate tabs:

```sh
docker compose up
```

```sh
go run ./cmd/eventual
```

```sh
npm run start
```

Then navigate to http://localhost:3000 where you'll find the development server running for the client, which is configured to work with the server.

## Production Environment

Both the client and server are hosted on [fly.io](https://fly.io/).
