# Go Auth Service
Dockerized auth microservice based on Golang and Gin

## Features
* [ ] JWT
* [ ] User registration and authentication
* [ ] User profile management
* [ ] Remembering password
* [ ] Sign-in/sign-up via social network


## Requirements
- [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/)
- [Git](https://git-scm.com/)


## Installation and Usage
```shell
$ git clone git@github.com:dmitrymomot/go-auth-service.git
$ cd ./go-auth-service
```

### Runing without proxy. It will be available in your browser by reference http://localhost:8080
```shell
$ docker-compose up -d
```

### Runing with docker proxy container. It will be available in your browser by reference http://auth.go.dev
```shell
$ docker-compose -f proxy.yml up -d
```

### Quick deploy to a remote server on [DigitalOcean](https://m.do.co/c/15cdc2182d4b)
```shell
$
```


## Depends On
- [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)
- [joho/godotenv](https://github.com/joho/godotenv)
