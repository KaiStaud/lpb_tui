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
		//result      mgl64.Vec2
	}{
		{"correctly identified too large target vector", 10, 0, 0},
		{"correctly identified another too large target vector", 5, 5, 0},
		{"correctly identified an not supported 3D target vector", 5, 5, 5},
		{"found an correct solution to target vector", 3, 4, 0},
		{"found an solution to smal target vector", 1, 3, 0},
	}

	for _, testCase := range testCases {
		err := CalculateVectors(testCase.input_x, testCase.input_y, testCase.input_z)
		/*
			if err != testCase.expected {
				t.Logf("Test failed")
			} else {
				t.Logf("Test passed")
			}
		*/
	}
	//}
}
