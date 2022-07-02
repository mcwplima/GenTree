package server

import (
	"context"
	"fmt"
	locallog "log"
	"net/http"

	"gentree/config"
	"gentree/middleware"
	"gentree/router"
)

// Start the webservice
func Start(context context.Context) {

	config, ok := config.FromContext(context)
	if !ok {
		locallog.Fatal("Unable to retrive configuration")
	}

	router := router.Routes()

	locallog.Println(fmt.Sprintf("Starting Web Service on port %d", config.Core.Port))

	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Core.Port), middleware.WrapHandler(context, router))
	if err != nil {
		locallog.Fatal(err.Error())
	}
}
