package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFatal(t *testing.T) {
	defer func() {
		assert.NotNil(t, recover())
	}()
	Fatal(fmt.Errorf("error"))
}
