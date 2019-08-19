# Infrastructure

[中文](README.zh_cn.md "中文") | [English](README.md "English")

用来存放 Kubernetes 的相关文件，如 service、deployment 等配置文件。

## 创建 Deployment
#### sa-logic
`kubectl apply -f sa-logic-deployment.yaml` 

#### sa-webapp
`kubectl apply -f infrastructure/sa-webapp-deployment.yaml`

#### sa-frontend
`kubectl apply -f infrastructure/sa-frontend-deployment.yaml`

## 创建 Service
#### sa-logic
`kubectl apply -f sa-logic-service.yaml` 

#### sa-webapp
`kubectl apply -f infrastructure/sa-webapp-service.yaml`

#### sa-frontend
`kubectl apply -f infrastructure/sa-frontend-service.yaml`

## 部署到 Minikube 集群
创建一个 minikube 集群: `minikube start --memory='4000mb' --cpus=4 --image-mirror-country='cn'`

获取 sa-frontend 的访问地址: `minikube service sa-frontend-service`