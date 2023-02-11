package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/joshuaswickirl/proglog/internal/server"
)

func TestHandleProduce(t *testing.T) {
	t.Run("returns correct body", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/",
			strings.NewReader(`{"record": {"value": "TGV0J3MgR28gIzEK"}}`))
		res := httptest.NewRecorder()

		s := server.NewHTTPServer()
		s.HandleProduce(res, req)

		expectedResBody := `{"offset":0}` + "\n"
		if res.Body.String() != expectedResBody {
			t.Errorf("got %q, want %q", res.Body.String(), expectedResBody)
		}
	})

	t.Run("offset increments", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/",
			strings.NewReader(`{"record": {"value": "TGV0J3MgR28gIzIK"}}`))
		res := httptest.NewRecorder()
		req2, _ := http.NewRequest(http.MethodPost, "/",
			strings.NewReader(`{"record": {"value": "TGV0J3MgR28gIzMK"}}`))
		res2 := httptest.NewRecorder()

		s := server.NewHTTPServer()
		s.HandleProduce(res, req)
		s.HandleProduce(res2, req2)

		offset := parseResponseBody(t, res2)

		expectedOffset := 1
		if offset != expectedOffset {
			t.Errorf("got %d, want %d", offset, expectedOffset)
		}
	})
}

type responseBody struct {
	Offset int
}

func parseResponseBody(t *testing.T, r *httptest.ResponseRecorder) int {
	rb := &responseBody{}

	err := json.Unmarshal(r.Body.Bytes(), rb)
	if err != nil {
		t.Error(err)
	}

	return rb.Offset
}
