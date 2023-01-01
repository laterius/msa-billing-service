build:
	docker build -f docker/Dockerfile . -t 34234247632/billing-service:v1.6

push:
	docker push 34234247632/billing-service:v1.6

docker-start:
	cd docker && docker-compose up -d

docker-stop:
	cd docker && docker-compose down


