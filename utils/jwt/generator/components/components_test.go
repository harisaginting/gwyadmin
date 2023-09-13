package components

import (
	"fmt"
	"testing"
)

func TestObtainPID(t *testing.T) {

	result := ObtainPID()
	fmt.Println(result)
	if result == 0 {
		t.Errorf("No Process ID obtained.")
	}
}

func TestTimeSinceOrganizationEpoch(t *testing.T) {

	result := TimeSinceOrganizationEpoch(1609459201000)
	fmt.Println(result)
	if result <= 1000000 {
		t.Errorf("Improper Epoch Time calculated.")
	}
}

func TestRandStringCreatesRandomString(t *testing.T) {
	result := RandString()
	fmt.Println(result)
	if result == "" {
		t.Errorf("No random string constructed or returned.")
	}
}
