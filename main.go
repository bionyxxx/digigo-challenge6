package main

import "Challenge6/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
