package structs_test

import (
	"encoding/json"
	"github.com/fatih/structs"
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
}

//Books       []Book `json:"books"`

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
	}]
}`

func Test(t *testing.T) {
	assert.Nil(t, nil)

	/*
		catalog := []Author{
			Author{"José Saramago", "Portugal", []Book{
				Book{"O ano da morte de Ricardo Reis", 1984},
				Book{"Ensaio sobre a cegueira", 1995},
			}},
			Author{"Enrique Vila-Matas", "Spain", []Book{
				Book{"El mal de Montano", 2002},
				Book{"Paris no se acaba nunca", 2003},
			}},
		}
	*/

	author := Author{}
	assert.Nil(t, json.NewDecoder(strings.NewReader(jsonData)).Decode(&author))
	//json.NewEncoder(os.Stdout).Encode(author)

	books := []Book{}
	assert.Nil(t, mapstructure.Decode(author.Books, &books))
	//json.NewEncoder(os.Stdout).Encode(books)

	json.NewEncoder(os.Stdout).Encode(books)
	json.NewEncoder(os.Stdout).Encode(structs.Map(author))
}
