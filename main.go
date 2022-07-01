package main

import (
	"context"
	"flag"
	"gentree/config"
	"gentree/log"
	"gentree/server"
	"gentree/storage"
)

func main() {
	configfile := flag.String("c", "config/app.conf", "The configuration file")
	flag.Parse()
	cfg := config.ReadConfig(*configfile)

	store := storage.FromConfig(cfg)
	logger := log.FromConfig(cfg)

	ctx := context.Background()
	ctx = config.NewContext(ctx, cfg)
	ctx = storage.NewContext(ctx, store)
	ctx = log.NewContext(ctx, logger)

	server.Start(ctx)
}
