/*
* Unit Tests for tracking package
 */

package tracking

import (
	"errors"
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

func TestCalculateMovementProgress(t *testing.T) {

	testseries := []mgl64.Vec2{{1, 2}, {3, 4}, {5, 6}, {7, 8}}

	StartNewSeries(testseries, 4)

	testCases := []struct {
		description string
		vec         mgl64.Vec2
		progress    float64
		err         error
	}{
		{"25 Percent", mgl64.Vec2{1, 2}, 25, nil},
		{"50 Percent", mgl64.Vec2{3, 4}, 50, nil},
		{"Vector not found in slice", mgl64.Vec2{4, 4}, 50, errors.New("Waypoint not reached!")},
		{"Correct vector after error", mgl64.Vec2{5, 6}, 75, nil},
		{"All done", mgl64.Vec2{7, 8}, 100, nil},
		{"Additional vector after finish", mgl64.Vec2{0, 4}, 100, nil},
	}

	Initialize(1)

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {

			progress, err := CalculateMovementProgress(testCase.vec)
			if err != testCase.err {
				t.Logf("Returned incorrect error: Expected %v, Got %v", testCase.err, err)
			}
			if progress == testCase.progress {
				t.Logf("Progress correctly calculated: Expected %f, Got %f", testCase.progress, progress)
			} else {
				t.Logf("Progress INcorrectly calculated: Expected %f, Got %f", testCase.progress, progress)

			}
		})
	}
}
