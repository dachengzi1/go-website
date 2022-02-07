package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Persion struct {
	Name string
	Age  int
}

func TestJsonMarshalToString(t *testing.T) {
	p := Persion{"张三", 1}
	str, err := JsonMarshalToString(p)
	assert.Nil(t, err)
	assert.Equal(t, str, "{\"Name\":\"张三\",\"Age\":1}")
}
