# Sentiment analyzer logic

[中文](README.zh_cn.md "中文") | [English](README.md "English")

Sentiment analyzer main logic, Is a micro service.

## Play with Docker
Build the docker image with this command `docker build -t sa-logic:1.0.0 .`

Start this project with this command `docker run --rm -p 5000:5000 sa-logic:1.0.0`.

## Analyse API
Request:
```
curl -X POST \
  http://localhost:5000/analyse/sentiment \
  -H 'Content-Type: application/json' \
  -d '{"sentence": "很酷哦"}'
```

Response:
```json
{
    "code": 0,
    "data": {
        "sentence": "很酷哦",
        "sentiments": 0.838566002771668
    }
}
```