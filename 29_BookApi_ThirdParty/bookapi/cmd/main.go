package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jy-pandey/bookapi/db"
	"github.com/jy-pandey/bookapi/handlers"
	"github.com/jy-pandey/bookapi/jobs"
	"github.com/jy-pandey/bookapi/routes"
	"go.uber.org/fx"
)


func main() {
	fmt.Println("Third Party API integration")

	app := fx.New(
		fx.Provide(
			gin.Default,                 // gives *gin.Engine .. anyone has need can use it
			db.ConnectDB,                // gives *gorm.DB
			handlers.NewBookHandler,     // gives *BookHandler
			handlers.NewTopBooksHandler, // gives *TopBookHandler
			jobs.NewCronJob,             // giver *CronJob // pointer to CronJob
		),
		fx.Invoke(routes.RegisterRoutes, StartCronJobRunner, StartServer),
	)
	app.Run()

}
func StartCronJobRunner(cj *jobs.CronJob) {
	go cj.StartCronJob()
}
func StartServer(router *gin.Engine) {
	router.Run(":8080") // https://localhost:8080/
}
