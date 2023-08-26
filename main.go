package main

import (
	"HourlyNewsGo/newsapiscrape"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	usermap := make(map[string]bool) //uuid to bool
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
		}
		for i := 0; i < len(news.Articles); i++ {
			fmt.Println("Title: " + news.Articles[i].Title + "\nAuthor: " + news.Articles[i].Author + "\nContent: " + news.Articles[i].Content + "\n")
		}*/
}

func sendNews(usermap map[string]bool) gin.HandlerFunc {
	fn := func(reqcontext *gin.Context) {
		key := reqcontext.Param("key")
		if _, contains := usermap[key]; contains {
			search := newsapiscrape.Search{}
			search.SetKey("f8b7c43989b44e07af5c870fed7944ec")
			news, err := search.Search()
			if err != nil {
				os.Exit(1)
			}
			reqcontext.JSON(http.StatusOK, news.Articles)
		}
	}
	return fn
}

func query(usermap map[string]bool) gin.HandlerFunc {
	fn := func(reqcontext *gin.Context) {
		// make sure user is an admin
	}
	return fn
}

func createapikey(usermap map[string]bool) gin.HandlerFunc {
	fn := func(reqcontext *gin.Context) {
		key := uuid.New().String()
		usermap[key] = true
		reqcontext.String(http.StatusOK, key)
	}
	return gin.HandlerFunc(fn)
}

func displaykeys(usermap map[string]bool) gin.HandlerFunc {
	fn := func(reqcontext *gin.Context) {
		/*if value, contains := adminmap[idmap[reqcontext.ClientIP()]]; contains {
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
		reqcontext.Status(http.StatusForbidden)*/
	}
	return fn
}

func createJsonFromMap(usermap map[string]bool) []byte {
	json, err := json.Marshal(usermap)
	if err != nil {
		fmt.Println("Error when creating json")
	}
	return json
}

func adminify(usermap map[string]bool) gin.HandlerFunc {
	fn := func(reqcontext *gin.Context) {
		/*if _, contains := adminmap[idmap[reqcontext.ClientIP]; contains {
			adminmap[idmap[reqcontext.ClientIP()]] = true
		}
		reqcontext.Status(http.StatusAccepted)*/
	}
	return fn
}

func Setup(usermap map[string]bool) *gin.Engine {
	cacheDir := "./cache"
	fmt.Println("Finding Cache")
	cacheDirInfo, openError := os.Stat(cacheDir)
	if openError != nil || cacheDirInfo == nil {
		CreateCache(cacheDir)
	}
	os.WriteFile(filepath.Join(cacheDir, "server.json"), []byte("{\n \"date\": \"July 26th, 2023\"\n}"), 0666)
	router := gin.Default()
	router.GET("/news?=:key", sendNews(usermap))
	router.POST("/query?=:query", query(usermap))
	router.GET("/key", createapikey(usermap))
	router.GET("/quit", displaykeys(usermap))
	router.GET("/admin", adminify(usermap))
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
		os.Exit(1)
	}
	fmt.Println("Cache created")

}
