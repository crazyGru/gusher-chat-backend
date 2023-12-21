package access

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IdProvider interface {
	Get(key string) (*Identity, error)
}

type MockApiProvider struct {
	idList map[string]Identity
}

func NewMockApiProvider(idList map[string]Identity) *MockApiProvider {
	return &MockApiProvider{idList: idList}
}

func (m *MockApiProvider) Get(key string) (*Identity, error) {
	id, exists := m.idList[key]
	if !exists {
		return nil, fmt.Errorf("failed to find user for key: %s", key)
	}
	return &id, nil
}

type LegacyApiProvider struct {
	url string
}

func NewLegacyApiProvider(url string) *LegacyApiProvider {
	return &LegacyApiProvider{url: url}
}

func (l *LegacyApiProvider) Get(key string) (*Identity, error) {
	var err error
	resp, err := http.Get(l.url + "?id=" + key)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	result := &Identity{}
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
