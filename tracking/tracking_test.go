/*
* Unit Tests for tracking package
 */

package tracking

import (
	"testing"

	"github.com/go-gl/mathgl/mgl64"
)

func TestIsTargetReached(t *testing.T) {

	testCases := []struct {
		description     string
		test_vector     mgl64.Vec2
		target_vector   mgl64.Vec2
		radius          float64
		expected_output bool
	}{
		{"Target equal passed vector", mgl64.Vec2{3, 4}, mgl64.Vec2{3, 4}, 1, true},
		{"Coordinates swapped", mgl64.Vec2{3, 4}, mgl64.Vec2{4, 3}, 5, false},
		{"Outside of sphere", mgl64.Vec2{3, 4}, mgl64.Vec2{6, 6}, 5, false},
		{"Origin", mgl64.Vec2{3, 4}, mgl64.Vec2{0, 0}, 5, false},
	}

	Initialize(1)

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			inSphere := IsTargetReached(testCase.radius, testCase.test_vector, testCase.target_vector)

			if testCase.expected_output != inSphere {
				t.Logf("Sphere check false: %s", testCase.description)
			} else {
				t.Logf("Sphere check correct: %s", testCase.description)
			}
		})
	}
}
