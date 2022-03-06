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
	"lpb/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-gl/mathgl/mgl64"
)

//----------------------- Constansts ----------------------- //

//----------------------- Variables ----------------------- //

var (
	// Dimensions of sphere
	R float64

	// Total progress:
	progress float64
	// Increment of progress per waypoint
	increment float64

	// Queue with waiting jobs:
	jobqueue chan mgl64.Vec3
	// Waypoints of profile:
	waypoints []mgl64.Vec2
	index     int

	// Stop Receiver:
	stop_receiver chan bool
)

//----------------------- Functions ----------------------- //

/*
Receive new programs,update queue on finish
*/
func StartReceiveQueue(p *tea.Program, c <-chan mgl64.Vec3) {
	for {
		select {
		case <-stop_receiver:
			return
		case <-c:
			// Send tui a handshake signal, the program was received:
			p.Send(tui.HandshakeMsg{})
			// Inform tui on all program states:
			// When program is #1 send in progress
			// Else sende pending
		default:
		}
	}
}

/* Receive Progress-Information from tracking module */
func PushProgress(progress float64) {

}

/*
* Initializes size of checkbox and resets logic to known state
 */
func Initialize(radius float64, queue chan mgl64.Vec3) error {
	R = radius
	jobqueue = queue
	return nil
}

/*
* Load new waypoints and reset progress:
 */
func StartNewSeries(series []mgl64.Vec2, len int) error {

	waypoints = series
	progress = 0
	index = 0
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

	// Check if there are remaining waypoints
	if index < len(waypoints) {
		// Check if trackpoint is reached
		if IsTargetReached(R, vec, waypoints[index]) {
			progress = progress + increment
			index++
			return progress, nil
		} else {
			return progress, errors.New("Waypoint not reached!")
		}
	} else {
		return progress, nil
	}

}
