package main

import (
	ginprometheus "github.com/zsais/go-gin-prometheus"
	covidsummaryapi "inter.lmwn.kongphop/internal/app/covid-summary-api"
)

// @title           LMWN Assesment Covid Summary API
// @version         1.0
// @description     This is a covid summary API server.

// @contact.name   Kuranasaki
// @contact.email  kongphop.c@kuranasaki.work

// @license.name  NIGGY License
// @license.url   https://github.com/Admin-OR-1-1/2110336-SE2-Crafty/blob/main/LICENSE

// @host      localhost:8080
// @BasePath  /

func main() {
	r := covidsummaryapi.SetupRouter()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
	r.Run()
}
