package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONMapDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Apple Mac Book Pro","price":24000000,"image_url":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["price"])
	fmt.Println(result["image_url"])
}

func TestJSONMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P00001",
		"name":  "Apple Mac Book Pro",
		"price": 24000000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
