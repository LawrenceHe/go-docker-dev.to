package main

func main() {
	route := initiateRoutes();

	route.Run(":8080")

}
