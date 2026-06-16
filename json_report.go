package main

import (
	"encoding/json"
	"maps"
	"os"
	"slices"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	keys := slices.Sorted(maps.Keys(pages))

	report := make([]PageData, 0, len(pages))

	for _, key := range keys {
		report = append(report, pages[key])
	}

	data, err := json.MarshalIndent(report, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
