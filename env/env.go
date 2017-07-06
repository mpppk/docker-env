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
		matched, err := regexp.MatchString(filter, k)
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
	for _, rawValue := range os.Environ() {
		storeSlice := strings.Split(rawValue, "=")
		store[storeSlice[0]] = strings.Join(storeSlice[1:], "=")
	}
	return store
}
