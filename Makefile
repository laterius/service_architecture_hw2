build:
	docker build -f docker/Dockerfile . -t 34234247632/otus-msa-hw2:v1.0

push:
	docker push eznd/otus-msa-hw2:latest

docker-start:
	cd docker && docker-compose up -d

docker-stop:
	cd docker && docker-compose down

k8s-deploy:
	helm/deploy.sh

k8s-remove:
	helm/remove.sh

newman:
	newman run postman/collection.json
