heroku_deploy: heroku_container_push heroku_container_release

heroku_container_push:
	heroku container:push web -a william-molsbee

heroku_container_release:
	heroku container:release web -a william-molsbee

docker_deploy: docker_build docker_run

docker_build:
	docker build -t blog .

docker_run:
	docker run --name blog -d blog

go_run:
	docker run --name postgres -e POSTGRES_USER=blogger -e POSTGRES_PASSWORD=password -e POSTGRES_DB=blog -p "5432:5432" -d postgres:12
	yarn --cwd ./frontend/ build && go run main.go

go_clean:
	docker stop postgres
	docker rm postgres