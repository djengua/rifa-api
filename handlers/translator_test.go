package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/djengua/rifa-api/handlers"
)

func TestTranslateAPI(t *testing.T) {
	// Arrange
	tt := []struct {
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Endpoint:            "/hello",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "hello",
		},
		{
			Endpoint:            "/hello?language=german",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "german",
			ExpectedTranslation: "hallo",
		},
	}
	handler := http.HandlerFunc(handlers.TranslateHandler)

	// Act
	for _, test := range tt {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", test.Endpoint, nil)

		handler.ServeHTTP(rr, req)
		// Assert
		if rr.Code != http.StatusOK {
			t.Errorf(`expected status %d but received %d`, test.StatusCode, rr.Code)
		}

		var resp handlers.Resp
		// Decodes the body of the response into a struct to be asserted
		json.Unmarshal(rr.Body.Bytes(), &resp)

		if resp.Language != test.ExpectedLanguage {
			t.Errorf(`expected language "%s" but received %s`, test.ExpectedLanguage, resp.Language)
		}

		if resp.Translation != test.ExpectedTranslation {
			t.Errorf(`expected language "%s" but received %s`, test.ExpectedTranslation, resp.Translation)
		}

	}

}