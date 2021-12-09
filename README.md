# (g)o (w)eb (s)erver

A simple web server in golang used for hosting [chenbrian.ca](https://chenbrian.ca)
As it is a static site, this server only has to serve `GET` requests and deal with simple routing/error handling.

Configuration is done in the `CONFIG.yml` file. 
By convention this will serve the `index.html` found at the path requested by the URL.
An example of the configuration file is provided in the `EXAMPLE_CONFIG.yml` file.

Probably best to put this behind a nginx reverse proxy.
Usage is simple:
    `go run ./main.go`




