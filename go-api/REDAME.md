# go-api
## requirements
docker compose

## setup
### 1. create docker
```
docker compose up -d --build
```
### 2. create authers table and insert test data.
ref: ```schema.sql```
ref: ```main.go```

### 3. exec go.main
```
docker compose exec app go run main.go
```

## Appendix
### 1. sqlc doc
https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html

### 2. how to gererate db package
```
docker compose exec app sqlc generate
```
