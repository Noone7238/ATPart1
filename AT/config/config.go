package config

import "os"

func GetPortNum() string {
	port := os.Getenv("ITEM_PORT_NUM")

	//if it's empty
	if port == "" {
		port = "8080"
	}

	return port
}
