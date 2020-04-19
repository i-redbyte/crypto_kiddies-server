package rest

import (
	"io"
	"net/http"
	"net/http/httptest"
)

//Method for tests
func PerformRequest(r http.Handler, method, path string, data io.Reader, token string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, data)
	w := httptest.NewRecorder()
	if token != "" {
		req.Header.Set("Authorization", token)
	}
	r.ServeHTTP(w, req)
	return w
}
