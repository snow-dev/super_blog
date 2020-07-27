default:
	@echo "=============building Local API============="
	docker build -f Dockerfile -t super_blog .

up: default
	@echo "=============starting super_blog locally============="
	docker-compose up -d

logs:
	docker-compose logs -f

down:
	docker-compose down

test:
	go test -v -cover ./...

clean: down
	@echo "=============cleaning up============="
	rm -f super_blog
	docker system prune -f
	docker volume prune -f