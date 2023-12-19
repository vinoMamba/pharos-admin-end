package main

import (
	"github.com/gin-contrib/cors"
	"github.com/vinoMamba.com/pharos-admin-end/config"
	"github.com/vinoMamba.com/pharos-admin-end/handler"
	"github.com/vinoMamba.com/pharos-admin-end/server"
	"github.com/vinoMamba.com/pharos-admin-end/storage"
)

func init() {
	config.LoadConfig(".")
	storage.DbConn()
}

func main() {
	r := server.SetupServer()
	r.Use(cors.Default())
	handler.HandleUser(r)
	handler.HandleUpms(r)
	r.Run(":3000")
}
