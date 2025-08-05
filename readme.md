# README

## Prerequisites
- connect to mysql and import albums.sql
```sh
docker cp albums.sql mysql:/root/albums.sql
docker exec -it mysql mysql -u root -p
create database recordings;
use recordings;
source /root/albums.sql
```

- export ENV variable
```sh
export DB_USERNAME=root
export DB_PASSWORD=root
```

## Run
`cd data_access`
`go run .`


ref: 
(https://go.dev/doc/tutorial/web-service-gin)[https://go.dev/doc/tutorial/web-service-gin]

## How To Run
- `go run .`
- dan di another terminal
```sh
# get all album
curl http://localhost:8000/albums

# get album by id
curl http://localhost:8000/albumbs/1

# post new album
curl http://localhost:8000/albums \
--include \
--request "POST" \
--header "Content-Type: application/json" \
--data '{"id": "5", "artist": "John Doe", "title": "Everybody loves", "price": 9.67}'
```
- 
