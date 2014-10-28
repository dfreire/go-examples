package structs_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type Book struct {
	Title  string `json:"title"`
	Year   int    `json:"year"`
	Author `json:"author"`
}

type Author struct {
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
}

const testJson string = `[{
	"title": "O ano da morte de Ricardo Reis",
	"year": 1984,
	"author": {
		"name": "José Saramago",
		"nationality": "Portugal"
	}
}, {
	"title": "Ensaio sobre a cegueira",
	"year": 1995,
	"author": {
		"name": "José Saramago",
		"nationality": "Portugal"
	}
}, {
	"title": "El mal de Montano",
	"year": 2002,
	"author": {
		"name": "Enrique Vila-Matas",
		"nationality": "Spain"
	}
}, {
	"title": "Paris no se acaba nunca",
	"year": 2003,
	"author": {
		"name": "Enrique Vila-Matas",
		"nationality": "Spain"
	}
}]`

func Test(t *testing.T) {
	assert.Nil(t, nil)

	jose := Author{"José Saramago", "Portugal"}
	enrique := Author{"Enrique Vila-Matas", "Spain"}

	books := []Book{
		Book{"O ano da morte de Ricardo Reis", 1984, jose},
		Book{"Ensaio sobre a cegueira", 1995, jose},
		Book{"El mal de Montano", 2002, enrique},
		Book{"Paris no se acaba nunca", 2003, enrique},
	}
	json.NewEncoder(os.Stdout).Encode(books)
}
