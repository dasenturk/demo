.PHONY: postgres adminer migrateup migratedown

postgres:
	docker run --rm -ti -P --publish 127.0.0.1:5432:5432 -e POSTGRES_PASSWORD=secret postgres

adminer:
	docker run --rm -it --network host adminer

migrateup:
	migrate -source file://migrations \
			-database postgres://postgres:secret@localhost/postgres?sslmode=disable up

migratedown:
	migrate -source file://migrations \
			-database postgres://postgres:secret@localhost/postgres?sslmode=disable down