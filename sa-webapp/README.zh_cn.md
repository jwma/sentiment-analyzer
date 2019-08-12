# Sentiment analyzer webapp
Sentiment analyzer API 服务，处理 `sa-frontend` 请求，调用 `sa-logic` 微服务，使用 Go 语言开发。

## 使用 Docker 运行
使用 `docker build -t sa-webapp:1.0.0 .` 构建 Docker 镜像。

使用 `docker run -e "API_HOST=http://127.0.0.1:5000" --rm -p 8080:8080 sa-webapp:1.0.0`. 启动应用。
正常来说，这样启动是无法正确访问到 `sa-logic` 服务的，因为它们存在于不同的网络当中，这里不会讲如何通过创建自定义网络来把服务串联起来，
在正式使用 Kubernetes 之前，我将会使用 docker-compose 把所有服务组合在一起进而能够提供服务。

## 情绪分析接口
请求:
```
curl -X POST \
  http://127.0.0.1:8080/analyse \
  -H 'Content-Type: application/json' \
  -d '{
	"sentence": "很酷哦"
}'
```

响应:
```json
{
    "sentence": "很酷哦",
    "sentiments": 0.838566002771668
}
```