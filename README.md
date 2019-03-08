# Standard Auth API (written in Golang) <!-- omit in toc -->

- [The project](#the-project)
- [Init the project](#init-the-project)
- [Documentation](#documentation)
- [Dependances](#dependances)

## The project
This project is a simple auth API written in Go. It's been thought to follow the best code practices and easily maintainable and expandable.

It includes a handling for login, registering, recovering and getting information on an account. It creates and uses a token.


It has been inspired by [https://medium.com/@adigunhammedolalekan/build-and-deploy-a-secure-rest-api-with-go-postgresql-jwt-and-gorm-6fadf3da505b](the article of Adigun Hammed Olalekan). Go check his work.

## Init the project
In order for the project to run, you need to set up your environment. To do so, please create a file named `my.env` at the root of the project, following the same pattern as the provided `.env` file.  
Then you can `go build` the app.

Or you can just execute the `install.sh` file, which does exactly the same.

## Documentation
Complete documentation can be found in the `doc` directory. It has been generated using [apiDoc](https://apidocjs.com). An online version is available on [my own server]().

## Dependances
This project uses some dependances :
- [github.com/gorilla/mux](https://github.com/gorilla/mux)
- [github.com/joho/godotenv](https://github.com/joho/godotenv)

Go check their work !
