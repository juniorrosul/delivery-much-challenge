# Delivery Much Tech Challenge

This is the implementation of Delivery Much Tech Challenge. The main objective of this challenge was to implement an endpoint for searching recipes with one, two or three different ingredients. The application was built using hexagonal architecture, this approach allows the isolation of the business logic from the external connection. The isolation of the components allow better test scenarios for all components.

The file structure on this project are described bellow:

```
./
├── adapters
│   ├── primary
│   │   ├── middleware
│   │   │   └── evaluateparameters.go
│   │   └── server.go
│   └── secondary
│       ├── giphy
│       │   ├── giphyapi.go
│       │   └── models.go
│       ├── http.go
│       ├── http_test.go
│       └── rpa
│           ├── recipe.go
│           ├── recipe_mock.go
│           └── recipe_test.go
├── application
│   ├── recipe
│   │   └── models.go
│   └── recipepuppy
│       └── models.go
├── bin
│   └── server
├── cmd
│   └── dmserver
│       └── main.go
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
└── README.md

12 directories, 18 files
```

## First steps

To run or build the application you need to configure dot env file.
For faster configuration you can copy the example

```
cp .env.example .env
```

On this configuration only two settings are available:

- `SERVER_PORT`: the port to run the server
- `GIPHY_API_KEY`: the api key from giphy service
	- Check [this link](https://developers.giphy.com/docs/api#quick-start-guide) for more information about the key.


## Building

To build the application you can use directly go or use the docker:

Using Go:
```
go build -o bin/server cmd/dmserver/*.go
```

Using docker:
```
docker run --rm -v "${PWD}":/app golang:latest sh -c 'cd /app; make build-go'
```

A shortcut was included on Makefile: `make build`.

