package structs_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	// "github.com/fatih/structs"
	// "github.com/markbates/going/nulls"
	// "github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

type Author struct {
	Name        string                   `json:"name"`
	Nationality string                   `json:"nationality"`
	Books       []map[string]interface{} `json:"books"`
	// Books       []Book `json:"books"`
}

type Book struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
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

	author := Author{}
	assert.Nil(t, json.NewDecoder(strings.NewReader(jsonData)).Decode(&author))
	json.NewEncoder(os.Stdout).Encode(author)

	for _, bookMap := range author.Books {
		fmt.Println("---")
		//json.NewEncoder(os.Stdout).Encode(bookMap)
		fmt.Println(bookMap)
		//json.NewEncoder(os.Stdout).Encode(getValidMap(Book{}, bookMap))
		fmt.Println(getValidMap(Book{}, bookMap))
	}
}

func getValidMap(v interface{}, in map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	t := reflect.TypeOf(v)
	tags := []string{}
	for i := 0; i < t.NumField(); i++ {
		tags = append(tags, strings.Split(t.Field(i).Tag.Get("json"), ",")[0])
	}
	for key, value := range in {
		for _, tag := range tags {
			if strings.Contains(tag, key) {
				out[key] = value
			}
		}
	}
	return out
}
