package main

func main() {
	server := ApiServer{":3000"}
	server.Run()
}