package worker

import (
	"reflect"
	"testing"

	"github.com/jasonjoo2010/goschedule/utils"
	"github.com/stretchr/testify/assert"
)

type Demo struct {
	x, y int
}

func TestRegister(t *testing.T) {
	Register(Demo{})
	Register(&Demo{})
	assert.NotNil(t, GetType(utils.TypeName(Demo{})))
	RegisterName("a", Demo{})
	assert.Equal(t, reflect.TypeOf(Demo{}), GetType("a"))
	RegisterName("a", &Demo{})
	assert.NotEqual(t, reflect.TypeOf(Demo{}), GetType("a"))
	assert.Equal(t, reflect.TypeOf(&Demo{}), GetType("a"))
}