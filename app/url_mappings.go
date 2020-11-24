package app

import (
	"github.com/ponsonio/reto_tekcom/controllers"
)

func mapUrls() {
	router.POST("v1/multiplication/", controllers.Multiply)
}
