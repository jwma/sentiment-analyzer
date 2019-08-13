package main

import (
	"bytes"
	"encoding/json"
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
	AnalyseSentimentData
	Sentiments float64 `json:"sentiments"`
}

type Result struct {
	AnalyseSentimentData
	Level int `json:"level"`
}

func sendAnalyseRequest(reqData *AnalyseSentimentData) ([]byte, error) {
	reqBody, _ := json.Marshal(reqData)
	url := os.Getenv("API_HOST") + "/analyse/sentiment"
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

func calcLevel(sentiments float64) (level int) {
	if sentiments < 0.1 {
		level = 1
	} else if sentiments >= 0.1 && sentiments < 0.2 {
		level = 2
	} else if sentiments >= 0.2 && sentiments < 0.3 {
		level = 3
	} else if sentiments >= 0.3 && sentiments < 0.4 {
		level = 4
	} else if sentiments >= 0.4 && sentiments < 0.5 {
		level = 5
	} else if sentiments >= 0.5 && sentiments < 0.6 {
		level = 6
	} else if sentiments >= 0.6 && sentiments < 0.7 {
		level = 7
	} else if sentiments >= 0.7 && sentiments < 0.8 {
		level = 8
	} else {
		level = 9
	}
	return level
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
		result := Result{AnalyseSentimentData: resp.Data.AnalyseSentimentData, Level: calcLevel(resp.Data.Sentiments)}
		c.JSON(http.StatusOK, result)
	})

	err := router.Run()
	if err != nil {
		panic("启动失败...")
	}
}
