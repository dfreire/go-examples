package json_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJson(t *testing.T) {
	data := make(map[string]string)
	data["a"] = "A"
	data["b"] = "B"

	var encoded bytes.Buffer
	assert.Nil(t, json.NewEncoder(&encoded).Encode(data))

	var decoded map[string]string
	assert.Nil(t, json.Unmarshal(encoded.Bytes(), &decoded))

	assert.Equal(t, decoded, data)
}
