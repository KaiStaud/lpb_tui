/*
Handles inverse kinematic calculations
*/

package inverse_kinematics

import (
	"errors"
	"math"
	"github.com/golang/geo/r3"

)

//  Slice contains absolute length of vectors
var arm_size []float64
var number_of_arms int
var range_of_motion float64
/*
Initialize the slice with vector lengths.
The slice will be used to calculate the resulting
vectors for each arm
*/
func InitVectors(length int, num_subvectors int) {

	// Write value into slice:

	// Calculate the maximum range of motion (ROM)

}

type support_vector struct{
	x float64
	y float64
	z float64
	size float64
	angle float64
}

/*
Calculate the resulting vector of each arm.
After adding all vectors together, the resulting vector
equals the passed one.
*/
func CalculateVectors(x float64, y float64, z float64 ) error {

	var sv support_vector
	var nv int // Normal vectors size

	// Support is limited to 2 arm elements:
	// The passed vector's size shouldn't be greater, then
	// the sum of all lengths stored in the slice vector_size!
	var length := math.Sqrt( math.Pow(x,2)+ math.Pow(y,2));

	if length > range_of_motion{
		return errors.New("Passed vector's size to large!")
	}
	
	// Calculate the two vectors of the arms
	else{
		
		/* Calculate the first vector */
		// To calculate the first vector, a support vector (passed vector with half lenght) and an attached normal vector
		// is used to calculate the vector of the first arm.

		// The temporary vector can be calculated on the targets vector angle, and its half size:
		sv.angle = math.Atan(y/z)
		sv.size = lenght/2
		sv.x = sv.size * math.sin(sv.angle)

		// Calculate the coordinates of the normal vector:
		// The normal vector is located at legth/2,  
		nv = math.sqrt(math.Pow(arm_size[0],2) - math.Pow(sv.length,2))

		// sv + nv = arm_vector
		// nv dot sv = 0 ( orthogonal / 90° angle)
		var support_vec = r3.NewPreciseVector(sv.x,sv.y,0)
		var normal_vec = r3.NewPreciseVector(-1/sv.x * nv, 1/sv.y,0 * nv,0) // Here will a singularity occur.
																			// The option with greater y coordinate is choosen to reduce mechanical stress
		var first_arm_vec = r3.PreciseVectorFromVector(support_vec.Add(normal_vec))

		/* Calculate the second vector: */
	}
}
