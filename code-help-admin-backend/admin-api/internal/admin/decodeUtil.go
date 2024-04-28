package admin

import (
	"encoding/json"
	"net/http"
)

func decode[T any](response *http.Response, output *T) error {
	defer func() { _ = response.Body.Close() }()
	err := json.NewDecoder(response.Body).Decode(&output)
	if err != nil {
		return err
	}

	return nil
}
