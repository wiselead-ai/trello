package trello

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateCard(t *testing.T) {
	// Arrange
	card := TrelloCard{
		Name:        "Test Card",
		Description: "This is a test card",
		ListID:      "testListID",
	}

	apiKey := "testApiKey"
	token := "testToken"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if !strings.Contains(r.URL.String(), apiKey) || !strings.Contains(r.URL.String(), token) {
			t.Errorf("API key or token missing in URL")
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	httpCli := server.Client()
	trelloAPI := NewTrelloAPI(httpCli, apiKey, token)
	trelloAPI.baseURL = server.URL

	// Act
	err := trelloAPI.CreateCard(context.Background(), card)

	// Assert
	require.NoError(t, err)
}
