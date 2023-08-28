package main

import (
	"HourlyNewsGo/newsapiscrape"
	"container/list"
	"encoding/json"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const jsonCache = "./cache/news.json"

type searchStatus int

const (
	ready = iota
	searched
)

type Manager struct {
	status searchStatus
	news   newsapiscrape.News
}

func startSearch(manager *Manager) {
	if manager.status == searched {
		return
	}
	manager.status = searched
	search := newsapiscrape.Search{}
	search.SetKey("f8b7c43989b44e07af5c870fed7944ec")
	news, err := search.Search()
	if err != nil {
		os.Exit(1)
	}
	manager.news = *news
	createCacheFile(news, manager)
}

func createCacheFile(news *newsapiscrape.News, manager *Manager) error {
	jsonString, err := json.Marshal(news)
	if err != nil {
		return err
	}
	if err = os.WriteFile(jsonCache, jsonString, 0666); err != nil {
		return err
	}
	go waitTime(manager)

	return nil
}

func query(usermap map[string]uuid.UUID, adminlist list.List) gin.HandlerFunc {
	fn := func(reqcontext *gin.Context) {
		// make sure user is an admin
	}
	return fn
}

// async function
func waitTime(manager *Manager) {
	time.Sleep((60 * time.Second))
	manager.status = ready
}
