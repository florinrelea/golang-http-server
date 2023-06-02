package main

func main() {
	server := CreateNewApiServer(":3000")

	server.Run()
}
