# Sentiment analyzer logic
Sentiment analyzer 主要逻辑，获取输入句子的情绪分析结果，是一个微服务，使用`` Python 开发。

## 使用 Docker 运行
使用 `docker build -t sa-logic:1.0.0 .` 构建 Docker 镜像。

使用 `docker run --rm -p 5000:5000 sa-logic:1.0.0` 启动应用。