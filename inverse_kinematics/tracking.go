/*
* Module tracks progress of active profile.
* Tracking-Points are receveived over the CAN-Bus Socket.
* The Communication module will process and route the received frames.
*
* The module tracks the movement by supervising stored tracking vectors.
* A tracking point is reached, when the received vector is pointing into
* the attached box around one of the tracking vectors
 */
package tracking

import (
	"github.com/go-gl/mathgl/mgl64"
)

var(
	// Dimensions of checkbox
	X float64
	Y float64
	Z float64	

	//
)
/*
* Initializes size of checkbox and resets logic to known state
*/
func Initialize(x float64, y float64, z float64) error  {

	X = x
	Y = y
	z = z
}
/*
* Set new target position. 
* The passed slice is used as intermittent waypoints.
* The passed waypoints are used for calculating the total progress in percent 
*/
func NewTrackingTarget(tv mgl64.Vec3)(){

	// Create six planes, which envelope the vector:

		// Calculate the necessary vectors:
		
		// Temporary vectors span planes around origin:
		tx = mgl64.Vec3{x,0,0}
		ty = mgl64.Vec3{0,y,0}
		tz = mgl64.Vec3{0,0,z}

		// Create new vectors around target vector
		xv = tv.Add(tx)
		yv = tv.Add(ty)
		zv = tv.Add(tz)

		// Span the individual planes around the target vector:
		// P = tv + r * (tv - xv) *  s * (tv - yv) -> XY Plane
		
		// P = tv + r * (tv - xv) *  s * (tv - zv) -> XZ Plane

		// P = tv + r * (tv - yv) *  s * (tv - zv) -> YZ Plane


		// Mirror the three planes around tv:



	//	

}


/*
* Process the received tracking-message
*/
func CalculateMovementProgress(vec mgl64.Vec2)(progress_percent int, err error)  {

	// Check if vector 
	return nil
}

)
/*
* 
*/