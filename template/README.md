# Golang WebApp Guide

## Template

Here we cleaned up the code and move the templating functionality into a separate package. This will allow us to isolate that functionality.

At this point we also start to make use of our first external package, which is [Goji](github.com/zenazn/goji/web). Goji is a web microframework and provides us with flexible routing and extensible middleware.

### Usage

    go run server.go
