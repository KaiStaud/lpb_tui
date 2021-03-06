/*
Handles inverse kinematic calculations
*/

package inverse_kinematics

import (
	"errors"
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

//  Slice contains absolute length of vectors
var number_of_arms int
var ROM range_of_motion

/*
Initializes Inverse Kinematics constraints.
Parameters are length of a single arm ( expected equal size among all arms),
Minimum reachable z-height of IK-Chain (available after setup) and the number
of attached arms in the IK Chain.
*/
func Init(length float64, ROM_min float64, num_subvectors int) {
	// Calculate the maximum range of motion (ROM)
	ROM.max = length * float64(num_subvectors)
	// Set the minimum ROM
	ROM.min = ROM_min

}

type support_vector struct {
	x     float64
	y     float64
	z     float64
	size  float64
	angle float64
}

type range_of_motion struct {
	min float64
	max float64
}

/*
Calculate the resulting vector of each arm.
After adding all vectors together, the resulting vector
equals the passed one.
*/
func CalculateVectors(x float64, y float64, z float64) (mgl64.Vec2, error) {

	var sv support_vector
	var nv float64 // Normal vectors size

	// Support is limited to 2 arm elements:
	// The passed vector's size shouldn't be greater, then
	// the sum of all lengths stored in the slice vector_size!
	length := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

	if length < ROM.min {
		return mgl64.Vec2{x, y}, errors.New("Passed vector's size to small!")
	} else if length > ROM.max {
		return mgl64.Vec2{x, y}, errors.New("Passed vector's size to large!")
	} else {

		// Calculate the two vectors of the arms

		/* Calculate the first vector */
		// To calculate the first vector, a support vector (passed vector with half lenght) and an attached normal vector
		// is used to calculate the vector of the first arm.

		// The temporary vector can be calculated on the targets vector angle, and its half size:
		sv.angle = math.Atan(y / x)            // in radians!
		sv.size = 0.5 * y / math.Sin(sv.angle) // Half of sine!
		sv.x = sv.size * math.Cos(sv.angle)
		sv.y = sv.size * math.Sin(sv.angle)

		// Calculate the coordinates of the normal vector:
		// The normal vector is located at legth/2,
		nv = math.Sqrt(math.Pow(35, 2) - math.Pow(sv.size, 2)) // Size of normal vector

		/*
			The normal vector creates two singularities.
			The normal vector is located above the support vector,
			thus nv is multiplicated with -1.
		*/
		normal_vec := mgl64.Vec2{(-1) * sv.y * nv / sv.size, (-1) * (-1) * sv.x * nv / sv.size}

		// sv + nv = arm_vector
		// nv dot sv = 0 ( orthogonal / 90?? angle)
		support_vec := mgl64.Vec2{sv.x, sv.y}

		// The option with greater y coordinate is choosen to reduce mechanical stress
		first_arm_vec := support_vec.Add(normal_vec)

		/* Calculate the second vector: */
		// The sum of the arm vectors equals the passed destination vector.
		// The remaining vector can be calculated by substracting the passed ond with the sum of support_vec and normal_vec
		destination_vec := mgl64.Vec2{x, y}
		sum := support_vec.Add(normal_vec)
		second_arm_vec := destination_vec.Sub(sum)

		// Check if calculations are correct by adding the two found vectors together.
		// If the inverse_kinematics are correct, the sum should equal the passed destination vector
		if destination_vec.ApproxEqual(first_arm_vec.Add(second_arm_vec)) {
			return first_arm_vec.Add(second_arm_vec), nil
		} else {
			return first_arm_vec.Add(second_arm_vec), errors.New("Inverse Kinematics returned incorrect vectors!")
		}
	}
}
