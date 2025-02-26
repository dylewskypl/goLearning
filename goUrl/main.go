package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	var database JSONDb
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	file, err := os.Open("urls.json")

	if err != nil {
		if os.IsNotExist(err) {
			database = JSONDb{}
			database.Urls = []Url{}
		}
	}

	data, err := ioutil.ReadAll(file)
	err = json.Unmarshal(data, &database)
	if err != nil {
		fmt.Println(err.Error())
	}
	js, err := json.Marshal(database)
	ioutil.WriteFile("urls.json", js, 644)
	r := gin.Default()
	r.GET("/list", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": database.Urls})
	})
	r.GET("/add", func(c *gin.Context) {
		var input struct {
			Url string `form:"url" binding:"required"`
		}

		c.ShouldBindQuery(&input)

		if containsOrginal(database.Urls, input.Url) {
			c.JSON(409, gin.H{"exists": input})
			return
		}

		toAdd := Url{}
		toAdd.OrginalUrl = input.Url

		rand.Seed(time.Now().UnixNano()) // Seed with current time
		result := make([]byte, 5)
		for i := range result {
			result[i] = charset[rand.Intn(len(charset))]
		}

		toAdd.ShortUrl = string(result)
		database.Urls = append(database.Urls, toAdd)
		js, _ := json.Marshal(database)
		ioutil.WriteFile("urls.json", js, 644)
		c.JSON(201, gin.H{"url": toAdd})
	})
	r.GET("/go", func(c *gin.Context) {
		var input struct {
			Url string `form:"target" binding:"required"`
		}

		c.ShouldBindQuery(&input)

		if containsShort(database.Urls, input.Url) {
			c.Redirect(301, getUrl(database.Urls, input.Url))
			return
		}
	})
	r.Run(":8080")
}

func containsOrginal(slice []Url, target string) bool {
	for _, item := range slice {
		if item.OrginalUrl == target {
			return true
		}
	}
	return false
}

func containsShort(slice []Url, target string) bool {
	for _, item := range slice {
		if item.ShortUrl == target {
			return true
		}
	}
	return false
}

func getUrl(slice []Url, target string) string {
	for _, item := range slice {
		if item.ShortUrl == target {
			return item.OrginalUrl
		}
	}
	return ""
}
