package main

import (
	"github.com/toluwase1/busha-assessment/controllers"
	_ "github.com/toluwase1/busha-assessment/docs"
)
// @title        Busha Assessment, A movie server
// @version      1
// @description  Repo can be found here: https://github.com/toluwase1/busha-backend-assessment

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host       localhost:8080
// @BasePath  /
// @securityDefinitions.basic  BasicAuth
func main() {
	server := &controllers.Server{}
	server.Start()
}




