### run docker-compose
compose-up:
	docker-compose up --build -d && docker-compose logs -f
.PHONY: compose-up

### down docker-compose
compose-down:
	docker-compose down --remove-orphans
.PHONY: compose-down

test:
	go test -v --cover github.com/aeonva1ues/dockered_go/cmd/main
.PHONY: test