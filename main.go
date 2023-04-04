package main

import "dasalgadoc.com/rest-websockets/api/application"

func main() {
	var app = application.BuildApplication()
	app.Broker.Start(application.BindRoutes)
}
