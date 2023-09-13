package generator

import (
	"regexp"

	"github.com/harisaginting/gwyn/utils/jwt/generator/components"
)

// GenerateIdentifier is the function that takes all of the components and generates the
func GenerateIdentifier() string {

	var code string
	valid := false
	rgx, _ := regexp.Compile("^[0-9a-zA-Z_]{6}$")
	for {
		result := components.RandString()
		if match := rgx.MatchString(result); match {
			code = result
			valid = true
		}
		if valid {
			break
		}
	}
	return code
}
