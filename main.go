package main

import (
	"HourlyNewsGo/newsapiscrape"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	usermap := make(map[string]string) //ip to UUID
	//adminmap := vector.New(0)
	fmt.Println("Setup\n----------")
	fmt.Println("Starting setup")
	router := Setup(usermap)
	fmt.Println("Setup Done\n----------")
	router.Run(":8080") /*
		search := newsapiscrape.Search{}
		search.SetKey("f8b7c43989b44e07af5c870fed7944ec")
		news, err := search.Search()
		if err != nil {
			os.Exit(1)
		}*/
}

func sendNews(usermap map[string]string) gin.HandlerFunc {
	fn := func(reqcontext *gin.Context) {
		key := reqcontext.Param("key")
		if _, contains := usermap[key]; contains {
			search := newsapiscrape.Search{}
			search.SetKey("f8b7c43989b44e07af5c870fed7944ec")
			news, err := search.Search()
			if err != nil {
				os.Exit(1) // just for testing right now
			}
			reqcontext.JSON(http.StatusOK, news.Articles)
		}
	}
	return fn
}

func query(usermap map[string]string) gin.HandlerFunc {
	fn := func(reqcontext *gin.Context) {
		// make sure user is an admin
	}
	return fn
}

func createapikey(usermap map[string]string) gin.HandlerFunc {
	fn := func(reqcontext *gin.Context) {
		target := uuid.New().String()
		key := reqcontext.ClientIP()
		if _, contain := usermap[key]; !contain {
			usermap[key] = target
		}
		reqcontext.String(http.StatusOK, usermap[key])
	}
	return gin.HandlerFunc(fn)
}

/*func shutdown(usermap map[string]bool) gin.HandlerFunc {
	fn := func(reqcontext *gin.Context) {
		if value, contains := adminmap[idmap[reqcontext.ClientIP()]]; contains {
			if value == true {
				file, err := os.Create(filepath.Join("./cache", "user.json"))
				if err != nil {
					fmt.Println("Error creating file")
				}
				json := createJsonFromMap(idmap)
				os.WriteFile(file.Name(), json, 0666)
				reqcontext.String(http.StatusOK, "Server quit")
				return
			}
		}
		reqcontext.Status(http.StatusForbidden)
	}
	return fn
}*/

func createJsonFromMap(usermap map[string]string) []byte {
	json, err := json.Marshal(usermap)
	if err != nil {
		fmt.Println("Error when creating json")
	}
	return json
}

func Setup(usermap map[string]string) *gin.Engine {
	cache := "./cache" // Set Cache Directory string for easy access
	fmt.Println("Finding Cache")
	cacheDirInfo, openError := os.Stat(cache)
	if openError != nil || cacheDirInfo == nil {
		CreateCache(cache)
	}
	t := time.Now()
	date := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, time.Now().Location())

	os.WriteFile(filepath.Join(cache, "server.json"), []byte("{\n \"date\": \""+date.String()+"\"\n}"), 0666)
	router := gin.Default()
	router.GET("/news?=:key", sendNews(usermap))
	router.POST("/query?=:querytype", query(usermap))
	router.GET("/key", createapikey(usermap))
	//router.GET("/quit", shutdown(usermap))
	return router
}

func CreateCache(cacheDir string) {
	fmt.Println("Creating cache")
	dirErr := os.Mkdir(cacheDir, os.ModePerm)
	if dirErr != nil {
		fmt.Println("Failed to create directory\n" + dirErr.Error())
	}
	_, err := os.Create(filepath.Join(cacheDir, "server.json"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1) // for testing
	}
	fmt.Println("Cache created")

}
