/*
* Test file for Inverse Kinematics
* Test Cases:
* - To Large target vector
* - correct target vector
 */

package inverse_kinematics

import (
	"testing"
)

func TestInverseKinematics(t *testing.T) {

	testCases := []struct {
		description string
		input_x     float64
		input_y     float64
		input_z     float64
	}{
		{"45° Orientation,", 40, 40, 0},
		{"45° Orientation,Case 2,", 30, 30, 0},
		{"Unit Vector", 0, 0, 0},
		{"Vector out of ROM", 60, 40, 0},
	}

	Init(35, 2, 2)
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			vec, err := CalculateVectors(testCase.input_x, testCase.input_y, testCase.input_z)

			if err != nil {
				if err.Error() == "Passed vector's size to large!" {
					t.Log("Found incorrect vector!")
				} else if err.Error() == "Passed vector's size too small!" {
					t.Fatalf("Failed test with error %v", err)
				}
			} else {
				t.Logf("IK correct, Testcase: %s", testCase.description)
				t.Log("Summ of Efectors =", vec)
			}
		})
	}
}
