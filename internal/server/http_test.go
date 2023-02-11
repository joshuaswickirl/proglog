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

		offset := parseOffsetFromResponseBody(t, res2)

		expectedOffset := 1
		if offset != expectedOffset {
			t.Errorf("got %d, want %d", offset, expectedOffset)
		}
	})
}

type responseBody struct {
	Offset int
}

func parseOffsetFromResponseBody(t *testing.T, r *httptest.ResponseRecorder) int {
	rb := &responseBody{}

	err := json.Unmarshal(r.Body.Bytes(), rb)
	if err != nil {
		t.Error(err)
	}

	return rb.Offset
}

func TestHandleConsume(t *testing.T) {
	t.Run("returns correct body", func(t *testing.T) {
		req0, _ := http.NewRequest(http.MethodPost, "/",
			strings.NewReader(`{"record": {"value": "TGV0J3MgR28gIzEK"}}`)) // need to mock this
		res0 := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/",
			strings.NewReader(`{"offset": 0}`))
		res := httptest.NewRecorder()

		s := server.NewHTTPServer()
		s.HandleProduce(res0, req0) // should not be needed
		s.HandleConsume(res, req)

		expectedResBody := `{"record":{"value":"TGV0J3MgR28gIzEK","offset":0}}` + "\n"
		if res.Body.String() != expectedResBody {
			t.Errorf("got %q, want %q", res.Body.String(), expectedResBody)
		}
	})

	t.Run("offset works", func(t *testing.T) {
		t.Skip()
		req, _ := http.NewRequest(http.MethodGet, "/",
			strings.NewReader(`{"offset": 1}`))
		res := httptest.NewRecorder()
		req2, _ := http.NewRequest(http.MethodGet, "/",
			strings.NewReader(`{"record":{"value":"TGV0J3MgR28gIzIK","offset":1}}`))
		res2 := httptest.NewRecorder()

		s := server.NewHTTPServer()
		s.HandleConsume(res, req)
		s.HandleConsume(res2, req2)

		offset := parseOffsetFromResponseBody(t, res2)

		expectedOffset := 1
		if offset != expectedOffset {
			t.Errorf("got %d, want %d", offset, expectedOffset)
		}
	})
}
