package covidsummaryapi

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "inter.lmwn.kongphop/api/openapi-spec"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	CovidSummary_Registry(router)
	return router
}

func CovidSummary_Registry(route *gin.Engine) { // if not working try *gin.Engine
	route.GET("/covid/summary", GenerateSummary)

	docs.SwaggerInfo.BasePath = "/"
	route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
