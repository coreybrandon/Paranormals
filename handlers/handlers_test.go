package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProducts(t *testing.T) {
	t.Run("returns all products", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/products", nil)
		response := httptest.NewRecorder()

		GetProducts(response, request)

		got := response.Body.String()
		want := `[{"id":"1","name":"test","life":100}]`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
