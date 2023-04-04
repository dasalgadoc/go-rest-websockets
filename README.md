<h1 align="center">
  🚀 🐹 Go Rest & Websockets examples 
</h1>

<p align="center">
    <a href="#"><img src="https://img.shields.io/badge/technology-go-blue.svg" alt="Go"/></a>
</p>

This repository serves as an example of a completely dockerized Rest API, CRUD with a Database, and Websocket notifying events.
It was created to investigate some techniques in Go. (with [Gorilla](https://github.com/gorilla#gorilla-toolkit) 😢)

__Keywords:__
[JWT](./resources/notes/JWTNOTES.md), Postgres, Rest API, Environment Variables, UUID, Middlewares, [Specification Pattern](./resources/notes/SPECIFICATION_NOTES.md)

## 🧲 Environment Setup

### 🛠️ Needed tools

1. Go 1.20.2 or higher
2. Docker and Docker compose (I use Docker version 23.01.1 and docker-compose v2.17.0)
   __Note:__ This repository includes a way to run it using Docker without having to install Go locally. It can move slowly for exploration and quickly for changes while still being useful. 

### 🏃🏻 Application execution

1. Make sure to download all Needed tools
2. Clone the repository
```
  git clone https://github.com/dasalgadoc/go-rest-websockets.git
```
3. There are two options for moving further at this point: semi-dockerization or full-dockerization

#### 🐬 Semi-dockerization

This approach uses Docker to manage a Postgres database and localhost to compile and run the Go source code.

- Build up go project

```bash
  go mod download
  go get .
```

- Starts the database docker

```bash
docker-compose up --build -d postgres
```

- Uncomment environment variable in __.env__ file

```makefile
DATABASE_URL=postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable
```

- Run the API
```bash
  go run main.go
```

- Test the Ping Endpoint
```bash
curl --request GET \
  --url http://localhost:8081/ping
```

#### 🐳 Full-dockerization

In this method, a Postgres database is managed by Docker, and [Golang Alpine](https://hub.docker.com/_/golang) code is copied, built, and sent as another docker image to execute.

- Make sure that the __DATABASE_URL__ environment variable is commented __.env__ file
```makefile
# DATABASE_URL=postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable
```

- Start the docker compose
```bash
docker-compose up --build
```

- Test the Ping Endpoint
```bash
curl --request GET \
  --url http://localhost:8085/ping
```

### 🤳🏻 Util commands

You can run queries on Postgres docker using:
```bash
docker exec -it go-rest-websockets-postgres-1 psql -U postgres -d postgres -c "<QUERY>"
```

You can remove all the docker compose volumes using:
```bash
docker-compose down -v
```

## 🧳 Project use

If everything went ok, you can test the ping endpoint.
Remember, if you use the Semi-dockerization the selected port its 8081 (.env PORT variable)
Full-dockerization, will forward the 8081 port to 8085 (docker-compose).

### Ping Endpoint
Basic Heathcheck
```bash
curl --request GET \
  --url http://localhost:<PORT>/ping
```

### SignUp Endpoint
```bash
curl --request POST \
  --url http://localhost:<PORT>/signup \
  --header 'Content-Type: application/json' \
  --data '{
	"email": "example@example.com",
	"password": "password"
}'
```

### Login
With a proper credentials this endpoint will generate a JWT token, useful for the next request
```bash
curl --request POST \
  --url http://localhost:<PORT>/login \
  --header 'Content-Type: application/json' \
  --data '{
	"email": "example@example.com",
	"password": "password"
}'
```

### Get logged user 
```bash
curl --request GET \
  --url http://localhost:<PORT>/me \
  --header 'x-auth-token: <TOKEN>'
```

### Create Post
This endpoint generates the POST_ID, you can query the database for more.
```bash
curl --request POST \
  --url http://localhost:<PORT>/post \
  --header 'Content-Type: application/json' \
  --header 'x-auth-token: <TOKEN>' \
  --data '{
	"post_content": "My custom content"
}'
```

### Get Post
```bash
curl --request GET \
  --url http://localhost:<PORT>/post/<POST_ID> \
  --header 'x-auth-token: <TOKEN>'
```

### Put Post
```bash
curl --request PUT \
  --url http://localhost:<PORT>/post/<POST_ID> \
  --header 'Content-Type: application/json' \
  --header 'x-auth-token: <TOKEN>' \
  --data '{
	"post_content": "New content"
}'
```

### Delete post
```bash
curl --request DELETE \
  --url http://localhost:<PORT>/post/<POST_ID> \
  --header 'x-auth-token: <TOKEN>'
```

## 📕 Library considerations:

This repository explores Rest API and WebSockets using these libraries.

```bash
go get github.com/gorilla/mux 
go get github.com/gorilla/websocket
```

Despite, I already used Gin in [Solid Example](https://github.com/dasalgadoc/solid-example-go), I decided to use the Gorilla alternatives without knowing these libraries
was marked as archived at the end of 2022.

I won't change this repository to use others libraries, but I will link some alternatives:

- ⭐️ [Gin](https://gin-gonic.com/)
- [Echo](https://echo.labstack.com/)
- [Chi](https://go-chi.io/#/) 
- [Goji](https://goji.io/)

For websockets

- [Fasthttp WebSocket](https://github.com/fasthttp/websocket)
- [Nhooyr WebSocket](https://github.com/nhooyr/websocket)
- [WS](https://github.com/gobwas/ws)
- [Centrifugo](https://github.com/centrifugal/centrifugo)
