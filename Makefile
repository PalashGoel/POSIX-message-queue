build:
	env GOOS=linux GOARCH=amd64 go build -o sender sender.go
	env GOOS=linux GOARCH=amd64 go build -o receiver receiver.go

image:
	docker build -t posix-mq-sender -f Dockerfile_sender --build-arg binary=sender . 
	docker build -t posix-mq-receiver -f Dockerfile_receiver --build-arg binary=receiver .


load:
	minikube image load posix-mq-receiver:latest
	minikube image load posix-mq-sender:latest

unload:
	minikube image unload posix-mq-receiver:latest
	minikube image unload posix-mq-sender:latest

create:
	kubectl create -f pod.yaml

delete:
	kubectl delete -f pod.yaml

get:
	kubectl get po posix-mq-pod

describe:
	kubectl describe po posix-mq-pod

exec:
	kubectl exec -it posix-mq-pod -c posix-mq-sender -- sh

log:
	kubectl logs -f posix-mq-pod -c posix-mq-receiver 
