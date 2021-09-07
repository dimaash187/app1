# GoLANG App1 BACKEND

This project is a backend server to calculate factorial. Makefile provided. The project is run on the port 8081.
It exposes 3 endpoints:

    1. /info [ GET request ]
    2. /factorial/{n} [ GET request with a numeric n argument ]
    3. /hit-counter [ PUT request ], increments a server wide hit counter.

## Build

Run `make build` to build the project. Main executable will be generated in the project root. 
Also a docker image [ app1 ] can be built with `make build-docker`.

Then run the docker image: `docker run -p 8081:8081 app1`

## Running unit tests

Run `make test` to execute the unit tests.

