package main

import network "gitstuff/AT/internal"

func main() {
	err := network.Start()
	if err != nil {
		//TODO Log the error
	}
}
