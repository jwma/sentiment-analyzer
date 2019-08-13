# Sentiment analyzer

[中文](README.zh_cn.md "中文") | [English](README.md "English")

原本就计划开发一个关于 Kubernetes 的项目，只是没想好怎么去切入这个大课题，
不久前受到[一篇文章](https://medium.com/free-code-camp/learn-kubernetes-in-under-3-hours-a-detailed-guide-to-orchestrating-containers-114ff420e882)的启发，所以我也将会开发一个类似的项目。

一个简单的基于微服务架构的中文情绪分析应用，它会被部署到 Kubernetes 进行服务。

技术栈：
- Python
    - Flask
    - SnowNLP
- Golang
    - Gin
- Vue.js
- Docker
- Kubernetes

## 使用 docker-compose 启动应用
使用 `docker-compose up -d --build` 启动应用。

在浏览器访问 `http://127.0.0.1:3000`。

使用 `docker-compose down -v` 停止应用。