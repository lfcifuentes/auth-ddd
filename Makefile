include .env

docker_up:
	docker-compose up --build -d

docker_down:
	docker-compose down

docs_generate:
	@echo "Generating docs...";
	swag init --parseDependency true;
	@echo "Docs generated successfully";

docs_format:
	@echo "Formatting code...";
	swag fmt
	@echo "Code formatted successfully";
