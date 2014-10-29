package structs_test

import (
	"encoding/json"
	// "fmt"
	"github.com/fatih/structs"
	// "reflect"
	// "github.com/markbates/going/nulls"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

type Author struct {
	Name        string                   `json:"name"`
	Nationality string                   `json:"nationality"`
	Books       []map[string]interface{} `json:"books"`
	//Books       []Book `json:"books"`
}

type Book struct {
	Title string `json:"title" structs:"title"`
	Year  int    `json structs:"year"`
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
	}]
}`

func Test(t *testing.T) {
	assert.Nil(t, nil)

	author := Author{}
	assert.Nil(t, json.NewDecoder(strings.NewReader(jsonData)).Decode(&author))
	json.NewEncoder(os.Stdout).Encode(author)

	for _, bookMap := range author.Books {
		book := Book{}
		assert.Nil(t, mapstructure.Decode(bookMap, &book))
		if bookMap["year"] == nil {
			assert.Equal(t, 0, book.Year)
		} else {
			assert.Equal(t, bookMap["year"], book.Year)
		}
		// json.NewEncoder(os.Stdout).Encode(bookMap)
		json.NewEncoder(os.Stdout).Encode(structs.Map(book))
		//fmt.Println(reflect.TypeOf(&book).Elem().Field(0).Tag.Get("json"))
	}
}
