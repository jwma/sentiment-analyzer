# Sentiment analyzer logic
Sentiment analyzer 主要逻辑，获取输入句子的情绪分析结果，是一个微服务，使用`` Python 开发。

## 使用 Docker 运行
使用 `docker build -t sa-logic:1.0.0 .` 构建 Docker 镜像。

使用 `docker run --rm -p 5000:5000 sa-logic:1.0.0` 启动应用。

## 情绪分析接口
请求:
```
curl -X POST \
  http://localhost:5000/analyse/sentiment \
  -H 'Content-Type: application/json' \
  -d '{"sentence": "很酷哦"}'
```

响应:
```json
{
    "code": 0,
    "data": {
        "sentence": "很酷哦",
        "sentiments": 0.838566002771668
    }
}
```