package network

import (
	config "gitstuff/AT/config"
	"strings"
)

//Start starts the application
func Start() error {
	port := config.GetPortNum()

	itmhandler, e := NewItemHandle()

	e.GET("/item", itmhandler.GetItems)
	e.GET("/item/:id", itmhandler.GetItem)
	e.DELETE("/item/:id", itmhandler.DeleteItem)
	e.PUT("/item/:id", itmhandler.UpdateItem)
	e.POST("/item", itmhandler.CreateItem)

	//TODO below remove fmt
	var sb strings.Builder
	sb.WriteString("Listening on port ")
	sb.WriteString(port)

	var sb2 strings.Builder
	sb2.WriteString("localhost:")
	sb2.WriteString(port)

	itmhandler.e.Logger.Print(sb.String())
	return itmhandler.e.Start(sb2.String())
}
