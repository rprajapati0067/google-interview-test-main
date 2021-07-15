package app

import "github.com/rprajapati0067/golang-interview-test-main/api/controllers"

func mapUrls() {
	//router.GET("marco", polo.Polo)
	router.POST("/subscribers", controllers.CreateSubscribers)
	//router.POST("/repositories", repositories.CreateRepos)

}
