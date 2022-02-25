package Proxy

import (
	"fmt"
	"time"
)

type Youtube interface {
	ListVideo() []string
}

type YoutubeImpl struct {
}

func (y YoutubeImpl) ListVideo() []string {
	return []string{
		"https://youtube/1",
		"https://youtube/2",
	}
}

type CacheYoutube struct {
	YoutubeService Youtube
	cache          []string
}

func (y *CacheYoutube) Run() {
	for {
		select {
		case <-time.After(5 * time.Second):
			y.cache = []string{}
		}
	}
}

func (y *CacheYoutube) ListVideo() []string {
	if len(y.cache) == 0 {
		y.cache = y.YoutubeService.ListVideo()
		fmt.Println("Set Cache")
	}
	return y.cache
}
