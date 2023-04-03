# go-rest-websockets
Project to explore rest and websockets

Keywords:
JWT, postgres, API, Environment variables, UUID, Middlewares

Without docker compose
Build Image
```bash
docker build . -t go-rest-websockets-postgres-container
docker run -p 54321:5432 go-rest-websockets-postgres-container
```

```bash
docker-compose up
docker-compose down -v
```

```bash
docker exec -it go-rest-websockets-postgres-1 psql -U postgres -d postgres -c "SELECT * FROM users"
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
