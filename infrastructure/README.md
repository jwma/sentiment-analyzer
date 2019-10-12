# Infrastructure

[中文](README.zh_cn.md "中文") | [English](README.md "English")

Used to store Kubernetes related files, such as service, 
deployment and other configuration files.

## Create Deployment
#### sa-logic
`kubectl apply -f deployment/sa-logic.yaml`

#### sa-webapp
`kubectl apply -f deployment/sa-webapp.yaml`

#### sa-frontend
`kubectl apply -f deployment/sa-frontend.yaml`

## Create Service
#### sa-logic
`kubectl apply -f service/sa-logic.yaml` 

#### sa-webapp
`kubectl apply -f service/sa-webapp.yaml`

#### sa-frontend
`kubectl apply -f service/sa-frontend.yaml`

## Play with Minikube
Create minikube cluster: `minikube start --memory='4000mb' --cpus=4 --image-mirror-country='cn'`

Get sa-frontend service access: `minikube service sa-frontend-service`
