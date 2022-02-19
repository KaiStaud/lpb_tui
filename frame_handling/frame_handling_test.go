package framehandling

import (
	"errors"
	"testing"
)

func TestProcessFrame(t *testing.T) {
	InitSimulation(2)
	testCases := []struct {
		description     string
		test_frame      DataFrame
		expected_result error
	}{
		{"Correct", DataFrame{1, 1, 100, 0, 0, 0, 0, 0, 0}, nil},
		{"False Silicon ID", DataFrame{1, 1, 0, 0, 0, 0, 0, 0, 0}, SIDError},
		{"NIC is Zero", DataFrame{1, 0, 0, 0, 0, 0, 0, 0, 0}, SIDError},
		{"NullFrame", DataFrame{}, SIDError},
	}

	// Insert into table
	for _, subtest := range testCases {
		// Do test
		t.Run(subtest.description, func(t *testing.T) {
			err := ProcessFrame(subtest.test_frame)
			if errors.Is(err, subtest.expected_result) {
				t.Log("Test OK")
			} else {
				t.Errorf("Checking incorrect, expected %v, got %v", subtest.expected_result, err)
			}

		})

	}
}
