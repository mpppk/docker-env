package env

import (
	"os"
	"regexp"
	"strings"
)

type Store map[string]string

func (s Store) Filter(filter string) (Store, error) {
	filteredStore := map[string]string{}
	for k, v := range s {
		matched, err := regexp.MatchString(strings.ToUpper(filter), strings.ToUpper(k))
		if err != nil {
			return nil, err
		}

		if matched {
			filteredStore[k] = v
		}
	}
	return filteredStore, nil
}

func New() Store {
	store := map[string]string{}
	for _, keyAndValue := range os.Environ() {
		storeSlice := strings.Split(keyAndValue, "=")
		value := strings.Join(storeSlice[1:], "=")
		value = strings.Trim(value, "\n")
		value = strings.TrimSpace(value)
		key := strings.TrimSpace(storeSlice[0])

		if value == "" || strings.Contains(value, "\033[") {
			continue
		}
		store[key] = value
	}
	return store
}
