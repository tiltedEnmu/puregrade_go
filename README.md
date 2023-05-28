# puregrade_go

This project is a server for a review aggregator (like metacritic or IGN, for example). The project aims to demonstrate my knowledge and skills in web development. The server can process http and grpc* requests and is written according to clean architecture principles.

Stack: PostgreSQL + sqlx, Docker, Gin, gRPC, jwt, testify (for tests).

## How to start

> Important! Make sure you have the Go compiler v17.0+ (For the second method) and Docker installed.

There are two launch ways. First:
1. Clone this repository: `git clone https://github.com/ZaiPeeKann/puregrade_go`
2. Go to the folder with the project
3. Set up environment variables, in this option we use `db.host: "host.docker.internal"` and leave the rest as default (It's better not to invent anything)
4. run `docker-compose up`
> What's the problem? gRPC server will not work. By default, 1 container = 1 process, but this is fixed by 2 containers or using a specific utility (so to speak, wait for the patch, gentlemen). This is where the alternative comes into play.

Way 2:
1. Clone this repository: `git clone https://github.com/ZaiPeeKann/puregrade_go`
2. Go to the folder with the project
3. Set up environment variables, in this option we use `db.host: "0.0.0.0"` and leave the rest as default
4. Here we either run `docker-compose up` and then off the container with the service and run our own locally like this: `go run ./cmd/main.go`. Or we run `docker-compose up` after cutting out all mentions of the server from the `docker-compose.yml` file.
> This allows you to feel all the functionality, but it takes longer to start (Again, we are waiting for the patch)

## Usage

With default settings, the http server port is 8000, and grpc is 9000. You can test it by sending requests, for example, via Postman

For more detailed documentation go here: https://www.notion.so/puregrade/Puregrade-docs-aa1883df47f348cfa98e2f334a2f858f?pvs=4
> Later it will be updated as well as this project

## Problems

Here I tried to collect the most likely problems with launching and using the project.
1. One of the possible startup errors is an error in which one of the ports that the project wants to listen to is already busy. In this case, see what processes are using the ports, like so: `netstat -ao`, and then kill the process.
2. You may have confused the `db.host variable in configs/example/config.yml`. To make sure you don't screw up in this matter, take a close look at the section **How to start**
3. If you notice any other error, create an issue or write to me in pm in tg: https://t.me/tiltedEnmu

*grpc** - *Only registration and authorization implemented, just to show what I can do.*