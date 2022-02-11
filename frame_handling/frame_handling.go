package framehandling

/*
Data Recording is used to provide data during dispatching jobs.
Also the module captures and buffers data  during teaching phase to provide
tracking data to teached programs

Data Recording captures the data of CAN-Frames, but for simulating Frames are
simulated by the Frame Simulator.

If there aren't any new frames after retriggering the timeout, the module signlizes an
Missing-Frame Error
*/

type DataFrame struct {
	ID              int
	PositionInChain int
	CoordinateX     int32
	CoordinateY     int32
	CoordinateZ     int32
	RotationX       int32
	RotationY       int32
	RotationZ       int32
}

var (
	arms map[int32]DataFrame
)

// Initialize the timeout and load frame data
func Init() {
	arms = make(map[int32]DataFrame)
	// Get all registered arms from DB

	/* 	constr, err := storage.GetConstraints()
	   	if err == nil {
	   		for _, v := range constr {
	   			var initframe DataFrame
	   			initframe.ID = v.ID
	   			initframe.PositionInChain = v.NumberInChain
	   			AddArm(initframe)
	   		}
	   	}
	   	return err */

}

// Accept/Rejects CAN-Frame
// Frames properties must match to be accepted:
// Correct Frame-ID
// Signature
// Number-in-chain
func ProcessFrame(can_frame DataFrame) {

}

func DeleteFrameHistory() {

}

/* Simulation */
func SimulateFrames() {

}

func SimulateFrameTimeout() {

}
