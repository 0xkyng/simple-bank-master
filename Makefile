createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple-bank

dropdb:
	docker exec -it postgres12 dropdb simple-bank

.PHONY: createdb