package trello

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/wiselead-ai/httpclient"
)

type TrelloCard struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	ListID      string `json:"idList"`
}

type TrelloAPI struct {
	httpCli *http.Client
	apiKey  string
	token   string
	baseURL string
}

func NewTrelloAPI(httpCli *http.Client, apiKey, token string) *TrelloAPI {
	return &TrelloAPI{
		httpCli: httpCli,
		apiKey:  apiKey,
		token:   token,
		baseURL: "https://api.trello.com/1",
	}
}

func (t *TrelloAPI) CreateCard(ctx context.Context, card TrelloCard) error {
	payload, err := json.Marshal(card)
	if err != nil {
		return fmt.Errorf("error marshaling card data: %v", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/cards?idList=%s&key=%s&token=%s", t.baseURL, card.ListID, t.apiKey, t.token),
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return fmt.Errorf("error creating POST request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := httpclient.DoWithRetry(t.httpCli, req)
	if err != nil {
		return fmt.Errorf("error performing POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}
	return nil
}
