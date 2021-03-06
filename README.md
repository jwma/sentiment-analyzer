# Sentiment analyzer

[中文](README.zh_cn.md "中文") | [English](README.md "English")

I originally planned to develop a project about Kubernetes, 
but I didn't think about how to cut into this big topic. 
I was inspired by [this article](https://medium.com/free-code-camp/learn-kubernetes-in-under-3-hours-a-detailed-guide-to-orchestrating-containers-114ff420e882) 
recently, so I will also develop a similar project.

A simple Chinese service analysis application based on micro service architecture, 
which will be deployed to Kubernetes for service.

Technology stack:
- Python
    - Flask
    - SnowNLP
- Golang
    - Gin
- Vue.js
- Docker
- Kubernetes

## Play with docker-compose
Start project with this command `docker-compose up -d --build`.

Open your web browser and access `http://127.0.0.1:3000`.

Stop project with this command `docker-compose down -v`.

## Screenshot
![Screenshot](screenshot.gif?raw=true "screenshot")
