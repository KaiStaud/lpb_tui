/*
* Module tracks progress of active profile.
* Tracking-Points are receveived over the CAN-Bus Socket.
* The Communication module will process and route the received frames.
*
* The module tracks the movement by supervising stored tracking vectors.
* A tracking point is reached, when the received vector is pointing into
* the attached sphere around one of the tracking vectors
 */
package tracking

import (
	"errors"

	"github.com/go-gl/mathgl/mgl64"
)

var (
	// Dimensions of sphere
	R float64

	// Total progress:
	progress float64
	// Increment of progress per waypoint
	increment float64

	// Waypoints of profile:
	waypoints []mgl64.Vec2
)

/*
* Initializes size of checkbox and resets logic to known state
 */
func Initialize(radius float64) error {
	R = radius
	return nil
}

/*
* Load new waypoints and reset progress:
 */
func StartNewSeries(series []mgl64.Vec2, len int) error {

	waypoints = series
	progress = 0
	increment = float64(100 / len)
	return nil
}

/*
Checks if passed vector is located outside of checksphere.
Returns true if vector to check (vec) is inside of spheres radius, otherwise returns false

Vector to check : vec
Spheres radius : radius
Destination vector on which sphere is attached : dest
*/
func IsTargetReached(radius float64, vec mgl64.Vec2, dest mgl64.Vec2) bool {
	d := dest.Sub(vec)

	d.Len()
	if d.Len() < R {
		return true
	} else {
		return false
	}

}

/*
* Process the received tracking-message and calculate the total progress
 */
func CalculateMovementProgress(vec mgl64.Vec2) (progress_percent float64, err error) {

	// Check if trackpoint is reached
	if IsTargetReached(R, vec, waypoints[0]) {
		progress = progress + increment
		return progress, nil
	} else {
		return progress, errors.New("Waypoint not reached!")
	}
}
