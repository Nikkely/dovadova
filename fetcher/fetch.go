package fetcher

import (
	"encoding/json"
	"os"
)

func FetchCardList(outputPath string, headless bool) (err error) {
	var list []string
	if list, err = GetCardList(headless); err != nil {
		return err
	}

	raw, err := json.MarshalIndent(list, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(outputPath, raw, 0664)
}
