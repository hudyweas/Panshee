postgres:
	kubectl apply -f ./postgres/volumes.yaml
	kubectl apply -f ./postgres/secret.yaml
	kubectl apply -f ./postgres/deployment.yaml

postgres-delete:
	kubectl delete -f ./postgres/secret.yaml
	kubectl delete -f ./postgres/deployment.yaml
	kubectl delete -f ./postgres/volumes.yaml

email:
	kubectl apply -f ./email/secret.yaml
	kubectl apply -f ./email/deployment.yaml


.PHONY: postgres
