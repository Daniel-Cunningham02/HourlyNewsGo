package main

import (
	"HourlyNewsGo/newsapiscrape"
	"container/list"
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
	usermap := make(map[string]uuid.UUID) //ip to UUID
	adminlist := list.New()
	manager := Manager{status: ready, news: newsapiscrape.News{}}
	fmt.Println("Setup\n----------")
	fmt.Println("Starting setup")
	router := Setup(&usermap, adminlist, &manager)
	fmt.Println("Setup Done\n----------")
	startSearch(&manager)
	router.Run(":8080") /*
		search := newsapiscrape.Search{}
		search.SetKey("f8b7c43989b44e07af5c870fed7944ec")
		news, err := search.Search()
		if err != nil {
			os.Exit(1)
		}*/
}

func sendNews(usermap map[string]uuid.UUID, manager *Manager) gin.HandlerFunc {
	t := time.Now()
	if t.Hour() > 7 && t.Hour() < 20 && t.Minute() == 59 {
		startSearch(manager)
	}
	fn := func(reqcontext *gin.Context) {
		key := reqcontext.Param("key")
		if _, contains := usermap[key]; contains {
			reqcontext.JSON(http.StatusOK, manager.news.Articles)
		}
	}
	return fn
}

func createapikey(usermap map[string]uuid.UUID) gin.HandlerFunc {
	fn := func(reqcontext *gin.Context) {
		target := uuid.New()
		key := reqcontext.ClientIP()
		if _, contain := usermap[key]; !contain {
			usermap[key] = target
		}
		reqcontext.String(http.StatusOK, usermap[key].String())
	}
	return gin.HandlerFunc(fn)
}

func shutdown(usermap *map[string]uuid.UUID, adminlist *list.List) gin.HandlerFunc {
	fn := func(reqcontext *gin.Context) {
		id := reqcontext.Param("key")
		var contains bool
		i := adminlist.Front()
		for i != adminlist.Back().Next() {
			if i.Value == id {
				contains = true
				break
			}
			i = i.Next()
		}
		if contains == true {
			var response string
			WriteAdminList(adminlist, &response)
			GenerateStatistics(usermap, &response) // going to implement later
			reqcontext.String(http.StatusOK, response)
			os.Exit(0)
			return
		}
		reqcontext.Status(http.StatusForbidden)
	}

	return fn
}

func createJsonFromMap(usermap map[string]uuid.UUID) []byte {
	json, err := json.Marshal(usermap)
	if err != nil {
		fmt.Println("Error when creating json")
	}
	return json
}

func Setup(usermap *map[string]uuid.UUID, adminlist *list.List, manager *Manager) *gin.Engine {
	var newCache bool
	cache := "./cache" // Set Cache Directory string for easy access
	fmt.Println("Finding Cache")
	cacheDirInfo, openError := os.Stat(cache)
	if openError != nil || cacheDirInfo == nil {
		CreateCache(cache)
		newCache = true
	}
	if newCache != true {
		readSuccess := ReadAdminList(*adminlist)
		if readSuccess != true {
			println("Failed to read admin list!")
		}
	}
	t := time.Now()
	date := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, time.Now().Location())

	os.WriteFile(filepath.Join(cache, "server.json"), []byte("{\n \"date\": \""+date.String()+"\"\n}"), 0666)
	router := gin.Default()
	router.GET("/news?=:key", sendNews(*usermap, manager))
	router.POST("/query?=:querytype", query(*usermap, *adminlist))
	router.PUT("/key", createapikey(*usermap))
	//router.GET("/quit?=:key", shutdown(&usermap))
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
	_, adminFileErr := os.Create(filepath.Join(cacheDir, "admin.json"))
	if adminFileErr != nil {
		println("Admin File Not Created")
	}
	fmt.Println("Cache created")

}
