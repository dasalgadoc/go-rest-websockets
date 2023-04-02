# go-rest-websockets
Project to explore rest and websockets

Without docker compose
Build Image
```bash
docker build . -t go-rest-websockets-postgres-container
```
docker run -p 54321:5432 go-rest-websockets-postgres-container

```bash
docker-compose up
```

```bash
docker exec -it go-rest-websockets-postgres-1 psql -U postgres -d postgres -c "SELECT * FROM users"
```