postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=G8keeper -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple-bank

dropdb:
	docker exec -it postgres12 dropdb simple-bank

.PHONY: postgres createdb dropdb