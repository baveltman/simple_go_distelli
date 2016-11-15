package simple_go_distelli

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInsertGetUser(t *testing.T) {
	obj := SimpleStruct{
		FirstName: "myName",
	}

	assert.Equal(t, obj.FirstName, "myName", "obj not created correctly")
}
