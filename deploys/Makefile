account:
	kubectl apply -f ./account_service/deployment.yaml

account-delete:
	kubectl delete -f ./account_service/deployment.yaml

account-restart:
	make account-delete
	make account

email:
	kubectl apply -f ./email_service/secret.yaml
	kubectl apply -f ./email_service/deployment.yaml

email-delete:
	kubectl delete -f ./email_service/deployment.yaml
	kubectl delete -f ./email_service/secret.yaml

email-restart:
	make email-delete
	make email

postgres:
	kubectl apply -f ./postgres/volumes.yaml
	kubectl apply -f ./postgres/secret.yaml
	kubectl apply -f ./postgres/deployment.yaml

postgres-delete:
	kubectl delete -f ./postgres/secret.yaml
	kubectl delete -f ./postgres/deployment.yaml
	kubectl delete -f ./postgres/volumes.yaml

postgres-restart:
	make postgres-delete
	make postgres

.PHONY: postgres
