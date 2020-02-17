package main
 
import (
	"github.com/firdaus-git/restapi/app"
	"github.com/firdaus-git/restapi/config"
)
 
func main() {
	config := config.GetConfig()
 
	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}