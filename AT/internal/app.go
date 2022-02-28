package network

import (
	"fmt"
	"os"
)

//Start starts the application
func Start() error {
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}

	itmhandler, e := NewItemHandle()

	e.GET("/item", itmhandler.GetItems)
	e.GET("/item/:id", itmhandler.GetItem)
	e.DELETE("/item/:id", itmhandler.DeleteItem)
	e.PUT("/item/:id", itmhandler.UpdateItem)
	e.POST("/item", itmhandler.CreateItem)

	//TODO below remove fmt
	itmhandler.e.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	return itmhandler.e.Start(fmt.Sprintf("localhost:%s", port))
}
