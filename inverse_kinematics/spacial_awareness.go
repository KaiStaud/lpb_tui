/*
Initializes the spacial awareness of lpb.
This prevents from executing profiles, which can not
be correctly executed due to physical constraints.
Sets base parameters for inverse kinematics.
*/

package spacial_awareness

import "github.com/golang/geo/r3"

/* -------------------- Types ---------------------- */
type arm_element struct {
	alpha  float32 // minimal angle to predecessor arm
	beta   float32 // maximal angle to predecessor arm
	lenght int     // length in mm
}

type tcp_element struct {
	alpha  float32
	beta   float32
	delta  float32
	gamma  float32
	length int
}

/* ------------------- Variables ------------------ */
var (
	max_vector r3.PreciseVector
)

/* ------------------- Functions ------------------ */
/*
Initializes physical layout / constraints
*/

/*
* @ brief: Checks if passed vector is longer than fully extended arm
* @ param: x,y,z coordinates of vector
* @ return: vector is in range [true/false]
 */
