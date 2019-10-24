[![Build Status](https://travis-ci.org/XiovV/vidy.svg?branch=master)](https://travis-ci.org/XiovV/vidy)

# Vidy

Vidy is a streaming service where users can upload any media they want and stream it from our servers.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
* Go 1.13.3

### Development Setup

Steps to get the development code running on your system:

```
git clone https://github.com/XiovV/vidy.git
cd vidy/
```

To run the register or login service run this in their directories:
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

## Release History
* 0.4.0
  * ADD: Browsing and streaming media
* 0.3.0
  * CHANGE: Dashboard UI
* 0.2.0
  * CHANGE: Login and Registration UI
* 0.1.1
  * FIX: Bug where a user was able to create an account with an email address that already existed
* 0.1.0
  * ADD: Working registration and login system

## Built With
* [Go](https://golang.org/) - Go Language
* [Gorilla Mux](https://github.com/gorilla/mux) - Go Multiplexer
* [Bootstrap](https://getbootstrap.com/) - CSS Framework
* [Parason](https://colorlib.com/preview/theme/parason/index.html) - Landing Page Template
* [Login Form 10 by Colorlib](https://colorlib.com/wp/template/login-form-v10/) - Login and Registration Template
