[![Build Status](https://travis-ci.org/XiovV/vidy.svg?branch=master)](https://travis-ci.org/XiovV/vidy)

# Vidy

Vidy is a streaming service where users can upload any media they want and stream it from our servers.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Running

Steps to get the development code running on your system:

```
git clone https://github.com/XiovV/vidy.git
```
```
cd vidy/
```

To run the register and login service run this:
```
go run main.go database_handlers.go handlers.go helpers.go jwt.go models.go
```
To run the dashboard service run this:
```
go run main.go handlers.go helpers.go
```

## Running the tests

```
go test -v ./...
```


## Built With

* [Gorilla Mux](https://github.com/gorilla/mux) - Go Multiplexer
* [Bootstrap](https://getbootstrap.com/) - CSS Framework
