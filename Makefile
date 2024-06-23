all: | unit-test smoke-test run;

setup:
	docker-compose build

clean:
	docker-compose down --volumes --remove-orphans

run: clean
	docker-compose up api

test: clean
	docker-compose run --rm -e CGO_ENABLED=0 -e GOOS=linux api test -cover -v ./...

smoke-test: setup
	docker-compose run --rm test