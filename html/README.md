# Golang WebApp Guide

## HTML

We now return an HTML response, and include more than 1 page.

### Usage

Template files are placed in the `templates` directory and for each page we create a `HandleFunc` that points to the relevant HTML template.

    go run server.go
