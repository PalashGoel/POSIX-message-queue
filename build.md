## PreRequisites

linux env
gcc
Go latest version (1.20) with cgo enabled 
Docker
minikube

Command to set local docker as docker for Minikube
eval $(minikube docker-env)  — bashrc file

command to enable cgo
export CGO_ENABLED="1"

## Build Binary

Run the command ***make build*** inside the POSIX-message-queue folder,
then there will be 2 binary files created **sender** and **receiver**

## Build image

Run the command ***make image*** inside the POSIX-message-queue folder,
then run the command ***docker images*** to see the latest created images with image names as **posix-mq-sender and posix-mq-receiver**

## Build load

Run the command ***make load*** inside the POSIX-message-queue folder,will load the images in the minikube.
then run the command ***minikube image ls --format=table*** to see all the images in the minikube

## Build unload

Run the command ***make unload*** inside the POSIX-message-queue folder,will unload the images in the minikube.
then run the command ***minikube image ls --format=table*** to see all the images in the minikube
