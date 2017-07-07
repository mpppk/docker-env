# docker-env
docker-env is the CLI tool for generate Dockerfile/docker-compose.yml with host environment variables.

## Generate Dockerfile

```bash
$ docker-env -q "proxy" node:latest
FROM node:latest
ENV HTTP_PROXY http://proxy.example.com
ENV HTTPS_PROXY http://proxy.example.com
ENV NO_PROXY willignore.com
$ docker build -t mynode .
$ docker run mynode -it /bin/bash
# Now you are in the container!
```

## Generate docker-compose.yml

```bash
$ docker-env -q "proxy" -f compose node:latest redis:latest
services:
  nodelatest:
    environment:
      HTTP_PROXY: http://proxy.example.com
      HTTPS_PROXY: http://proxy.example.com
      NO_PROXY: willignore.com
    image: node:latest
version: "3"
  redislatest:
    environment:
      HTTP_PROXY: http://proxy.example.com
      HTTPS_PROXY: http://proxy.example.com
      NO_PROXY: willignore.com
    image: redis:latest
$ docker-compose up nodelatest /bin/bash
# Now you are in the container!
```
## Installation

### Standalone

Download from [release page](https://github.com/mpppk/docker-env/releases) and put it anywhere in your executable path.

### From source

```bash
$ go get github.com/mpppk/docker-env
```

