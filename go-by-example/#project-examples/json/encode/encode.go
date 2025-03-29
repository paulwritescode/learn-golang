package encode

import (
	"encoding/json"
	"log"

	"github.com/paulwritescode/json/api/example/data"
)

func EncodeJson(Slug data.MyData) []byte {
	slug, err := json.MarshalIndent(Slug, "", "  ")
	if err != nil {
		log.Fatal("Error turning to json", err)
	}

	return slug
}
