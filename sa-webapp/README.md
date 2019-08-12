# Sentiment analyzer webapp
Sentiment analyzer API layer, create by Golang.

## Play with Docker
Build the docker image with this command `docker build -t sa-webapp:1.0.0 .`

Start this project with this command `docker run -e "API_HOST=http://127.0.0.1:5000" --rm -p 8080:8080 sa-webapp:1.0.0`.
Normally, you are not able to access the `sa-logic` service correctly, Because it has its own independent network, 
I will use docker-compose to combine all services before using Kubernetes.

## Analyse API
Request:
```
curl -X POST \
  http://127.0.0.1:8080/analyse \
  -H 'Content-Type: application/json' \
  -d '{
	"sentence": "很酷哦"
}'
```

Response:
```json
{
    "sentence": "很酷哦",
    "sentiments": 0.838566002771668
}
```