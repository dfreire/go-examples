package structs_test

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

type Book struct {
	Title  string
	Author string
	Year   int
}

const jsonData string = `{
	"name": "Enrique Vila-Matas",
	"nationality": "Spain",
	"books": [{
		"title": "El mal de Montano",
		"year": 2002
	}, {
		"title": "Paris no se acaba nunca"
	}, {
		"title": "Doctor Pasavento",
		"year": 2005,
		"publisher": "Anagrama"
	}, {
		"title": "Dublinesca",
		"year": null
	}]
}`

func Test(t *testing.T) {
	assert.Nil(t, nil)

	book := Book{"El mal de Montano", "E. Vila-Matas", 2002}
	assert.Equal(t, "E. Vila-Matas", book.Author)

	bookUpdated1 := Book{}
	getUpdatedStruct(book, &bookUpdated1, `{"author": "Enrique Vila-Matas"}`)
	assert.Equal(t, "Enrique Vila-Matas", bookUpdated1.Author)

	bookUpdated2 := Book{}
	getUpdatedStruct(bookUpdated1, &bookUpdated2, `{"publisher": "Anagrama"}`)
	assert.Equal(t, bookUpdated2, bookUpdated1)

	bookUpdated3 := Book{}
	getUpdatedStruct(bookUpdated1, &bookUpdated3, `{"year": null}`)
	assert.Equal(t, bookUpdated3.Title, bookUpdated1.Title)
	assert.Equal(t, bookUpdated3.Author, bookUpdated1.Author)
	assert.Equal(t, bookUpdated3.Year, 0)

	bookUpdated4 := Book{}
	getUpdatedStruct(bookUpdated1, &bookUpdated4, `{"title": null}`)
	assert.Equal(t, bookUpdated4.Title, "")
	assert.Equal(t, bookUpdated4.Author, bookUpdated1.Author)
	assert.Equal(t, bookUpdated4.Year, bookUpdated1.Year)
}

func getUpdatedStruct(orig interface{}, dest interface{}, changeJson string) {
	origMap := structs.Map(orig)
	fmt.Println("---")
	json.NewEncoder(os.Stdout).Encode(origMap)

	var changeMap map[string]interface{}
	json.Unmarshal([]byte(changeJson), &changeMap)
	json.NewEncoder(os.Stdout).Encode(changeMap)

	destMap := make(map[string]interface{})
	for key, value := range origMap {
		key := strings.ToLower(key)
		destMap[key] = value
	}
	for key, value := range changeMap {
		key := strings.ToLower(key)
		destMap[key] = value
	}

	mapstructure.Decode(destMap, dest)
	json.NewEncoder(os.Stdout).Encode(dest)
}
