# Infrastructure

[中文](README.zh_cn.md "中文") | [English](README.md "English")

Used to store Kubernetes related files, such as service, 
deployment and other configuration files.

## Create Deployment
#### sa-logic
`kubectl apply -f sa-logic-deployment.yaml` 

#### sa-webapp
`kubectl apply -f infrastructure/sa-webapp-deployment.yaml`

#### sa-frontend
`kubectl apply -f infrastructure/sa-frontend-deployment.yaml`

## Create Service
#### sa-logic
`kubectl apply -f sa-logic-service.yaml` 

#### sa-webapp
`kubectl apply -f infrastructure/sa-webapp-service.yaml`

#### sa-frontend
`kubectl apply -f infrastructure/sa-frontend-service.yaml`

## Play with Minikube
Create minikube cluster: `minikube start --memory='4000mb' --cpus=4 --image-mirror-country='cn'`

Get sa-frontend service access: `minikube service sa-frontend-service`