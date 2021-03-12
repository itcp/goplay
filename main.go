package main

import (
	"fmt"
	"goplay/router"
	"net/http"
	"time"
	
)

func main()  {
	router := router.InitRouter()

	s := &http.Server{
			Addr:           fmt.Sprintf(":%d", 8000),
			Handler:        router,
			ReadTimeout:    60 * time.Second,
			WriteTimeout:   60 * time.Second,
			MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
			panic(err.Error())
	}

}