package covidsummaryapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSummaryRoute(t *testing.T) {
	r := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/covid/summary", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	body, err := io.ReadAll(w.Body)
	if err != nil {
		t.Error("Error reading response body " + err.Error())
	}
	var c APIResponse
	if err = json.Unmarshal(body, &c); err != nil {
		t.Error("Error unmarshalling JSON " + err.Error())
	}
}
