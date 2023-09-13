package generator

import (
	"fmt"
	"testing"
)

func TestGenerateTheIdentifier(t *testing.T) {
	result := GenerateIdentifier()
	fmt.Println(result)
	if result == "" {
		t.Errorf("Generated Identifier is blank.")
	}
}
