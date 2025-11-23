package main

import (
	"log/slog"
	"net/http"
	"testing"

	"github.com/oscargsdev/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	app := &application{
		logger: slog.New(slog.DiscardHandler),
	}

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	res := ts.get(t, "/ping")
	assert.Equal(t, res.status, http.StatusOK)
	assert.Equal(t, res.body, "OK")
}
