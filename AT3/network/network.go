package network

import (
	"fmt"
	"os"
)

//Start starts the application
func Start() {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}

	//fornow
	samphandler, e := NewSampleHandle()

	e.GET("/sample", samphandler.GetSamples)
	e.GET("/sample/:id", samphandler.GetSample)
	e.DELETE("/sample/:id", samphandler.DeleteSample)
	e.PUT("/sample/:id", samphandler.UpdateSample)
	e.POST("/sample", samphandler.CreateSample)

	samphandler.e.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	samphandler.e.Logger.Fatal(samphandler.e.Start(fmt.Sprintf("localhost:%s", port)))
}
