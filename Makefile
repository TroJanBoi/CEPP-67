all:
	echo "No default"

up-db:
	docker compose up -d postgresql

up-be:
	docker compose up -d postgresql be

up-fe:
	docker compose up -d web

up:
	docker compose up -d

down:
	docker compose down -v
