package search

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// Search
// search by query json
func Search[T any](client *elasticsearch.Client, body Body[T], o ...func(*esapi.SearchRequest)) (Body[T], error) {
	ret, err := client.Search(o...)
	if err != nil {
		return body, err
	}

	if ret.IsError() {
		return body, fmt.Errorf("search error: %s", ret.Status())
	}

	_b, err := io.ReadAll(ret.Body)
	if err != nil {
		return body, err
	}

	err = json.Unmarshal(_b, &body)
	if err != nil {
		return body, err
	}

	return body, nil
}
