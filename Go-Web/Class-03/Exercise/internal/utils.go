package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func LoadProductJSON(filename string) (contentString string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}

	defer file.Close()

	reader := io.Reader(file)
	content, err := io.ReadAll(reader)

	if err != nil {
		return
	}

	contentString = string(content)

	return
}

func DumpProductJSON(data string, products StorageProduct) (StorageProduct, error) {
	stream := strings.NewReader(data)
	decoder := json.NewDecoder(stream)

	if err := decoder.Decode(&products.Products); err == io.EOF {
		fmt.Println("EOF")
	} else if err != nil {
		fmt.Println(err)
	}

	return products, nil
}

func PickleProductJSON(products StorageProduct) (jsonData []byte, err error) {
	jsonData, err = json.Marshal(products.Products)
	if err != nil {
		return
	}

	return
}

type ProductJSON interface {
	PickleProductJSON() (jsonData []byte, err error)
}

func JSON(product ProductJSON) (jsonData []byte, err error) {
	return product.PickleProductJSON()
}
