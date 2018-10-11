package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/rybitskyi/go-chat-server/api"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infof("Start Chat Server")
	r := gin.Default()

	api.New(r)

	err := http.ListenAndServe(":8081", r)
	panic(err)
}
