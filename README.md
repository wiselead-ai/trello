# Trello API Client

This project provides a Go client for interacting with the Trello API, including functionality to create Trello cards.

## Installation

To install the package, use:

```sh
go get github.com/wiselead-ai/trello
```

## Usage

### Creating a Trello Card

```go
package main

import (
	"context"
	"net/http"
	"github.com/wiselead-ai/trello"
)

func main() {
	httpCli := &http.Client{}
	apiKey := "yourApiKey"
	token := "yourToken"

	trelloAPI := trello.NewTrelloAPI(httpCli, apiKey, token)

	card := trello.TrelloCard{
		Name:        "New Card",
		Description: "Card description",
		ListID:      "yourListID",
	}

	err := trelloAPI.CreateCard(context.Background(), card)
	if err != nil {
		// Handle error
	}
}
```

## Testing

To run the tests, use:

```sh
go test ./...
```
