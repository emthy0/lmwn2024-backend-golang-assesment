package covidsummaryapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"inter.lmwn.kongphop/pkg/covid"
)

type APIResponse struct {
	Province map[string]int `json:"Province"`
	Age      map[string]int `json:"Age"`
}

type APIResponseError struct {
	Message string `json:"message"`
}

// GenerateSummary
// @Description Generate Summary of Covid Patient in Thailand grouped by Province and Age
// @Tags Covid
// @Produce json
// @Success 200 {object} APIResponse "Fetch and Summarize Covid Data Success"
// @Failure 500 {object} APIResponseError "Server Error"
// @Router /covid/summary [get]
func GenerateSummary(ctx *gin.Context) {
	covid, err := covid.NewCovid()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, APIResponseError{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, APIResponse{
			Province: covid.GroupByProvince(),
			Age:      covid.GroupByAge(),
		})
	}
}
