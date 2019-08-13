package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

type AnalyseSentimentResponse struct {
	Code int        `json:"code"`
	Data *Sentiment `json:"data"`
}

type AnalyseSentimentData struct {
	Sentence string `json:"sentence"`
}

type Sentiment struct {
	Sentence   string  `json:"sentence"`
	Sentiments float64 `json:"sentiments"`
}

func sendAnalyseRequest(reqData *AnalyseSentimentData) ([]byte, error) {
	reqBody, _ := json.Marshal(reqData)
	url := os.Getenv("API_HOST") + "/analyse/sentiment"
	fmt.Println(url)
	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func analyseSentiment(reqData *AnalyseSentimentData) *AnalyseSentimentResponse {
	resp, _ := sendAnalyseRequest(reqData)
	rs := &AnalyseSentimentResponse{
		Data: &Sentiment{},
	}
	_ = json.Unmarshal(resp, rs)
	return rs
}

func main() {
	router := gin.Default()
	// 开启 cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	router.POST("/analyse", func(c *gin.Context) {
		reqData := &AnalyseSentimentData{}
		_ = c.BindJSON(reqData)
		if reqData.Sentence == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "缺失参数",
			})
			return
		}
		resp := analyseSentiment(reqData)
		c.JSON(http.StatusOK, resp.Data)
	})

	err := router.Run()
	if err != nil {
		panic("启动失败...")
	}
}
