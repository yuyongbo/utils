package utils

import (
	"fmt"
	"testing"
)

func TestStart(t *testing.T) {
	//Start(logger)
	CurrentTest(logger, 10000, 10000)
	t.Log()
}
func logger() error {
	fmt.Sprintf("test")
	return nil
}
