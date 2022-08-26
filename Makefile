build:
	docker build -f docker/Dockerfile . -t 34234247632/otus-msa-hw2:v1.0

push:
	docker push 34234247632/otus-msa-hw2:v1.0

docker-start:
	cd docker && docker-compose up -d

docker-stop:
	cd docker && docker-compose down

k8s-deploy:
	kubectl create ns otus-msa-hw2
	helm upgrade --install -n otus-msa-hw2 otus-msa-hw2 helm/chart

k8s-remove:
	kubectl delete ns otus-msa-hw2

newman:
	newman run postman/collection.json
