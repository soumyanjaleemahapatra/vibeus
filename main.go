package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/soumyanjaleemahapatra/vibeus/internal/api"
	"github.com/soumyanjaleemahapatra/vibeus/internal/pkg/config"
	"github.com/soumyanjaleemahapatra/vibeus/internal/pkg/store"
)

var router = gin.Default()

func main() {
	conf := config.SetUpConfiguration()
	dbStore, err := store.Init(conf)
	if err != nil {
		log.Fatalf("failed to init database store, %s", err)
	}
	defer dbStore.Conn.Close()
	a := api.New(conf, dbStore)
	a.Run()
}
