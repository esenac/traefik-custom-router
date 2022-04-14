package traefik_custom_router_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	traefikcustomrouter "github.com/esenac/treafik-custom-router"
)

func TestUserRedirect(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "", want: "http://public"},
		{input: "a-value", want: "http://private"},
	}

	for _, tc := range tests {

		cfg := traefikcustomrouter.CreateConfig()
		cfg.Public = "http://public"
		cfg.Private = "http://private"

		ctx := context.Background()
		next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

		handler, err := traefikcustomrouter.New(ctx, next, cfg, "user-redirect")
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		req, err := http.NewRequestWithContext(ctx, http.MethodHead, "http://localhost", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("Authorization", tc.input)

		handler.ServeHTTP(recorder, req)
		location, _ := recorder.Result().Location()
		assertEquals(t, location.String(), tc.want)
		assertEquals(t, recorder.Result().StatusCode, http.StatusSeeOther)
	}
}

func assertEquals[K int | string](t *testing.T, current, expected K) {
	t.Helper()

	if current != expected {
		t.Errorf("invalid value: %v, expected: %v", current, expected)
	}
}
