# Golang WebApp Guide

## Template

This is a minor change to the [HTML](https://github.com/reinbach/golang-webapp-guide/tree/master/html) application we had. The following are the major changes that we make;

   * move the rendering function, the templates and static files into a separate directory for better organization.
   * add `context` functionality, this allows us to pass data to the templates
   * and we are able to pass a list of templates to be rendered.

At this point we also start to make use of our first external package, which is [Goji](github.com/zenazn/goji/web). Goji is a web microframework and provides us with flexible routing and extensible middleware. We will probably make use of one more external package further down.

### Usage

Template files are now placed in the `template/html` directory and statis files are placed in the `template/static` directory.

    go run server.go
