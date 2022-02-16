package framehandling

import (
	"errors"
	"lpb/storage"

	"gorm.io/gorm"
)

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
	SiliconID       int
	PositionInChain int
	FrameID         int
	CoordinateX     int32
	CoordinateY     int32
	CoordinateZ     int32
	RotationX       int32
	RotationY       int32
	RotationZ       int32
}

var (
	arms map[int]DataFrame
)

// Initialize the timeout and load frame data
func Init(db *gorm.DB) error {
	arms = make(map[int]DataFrame)
	// Get all registered arms from DB
	constr, err := storage.GetArms(db)
	if err != nil {
		return err
	} else {
		// Copy Data and initialize Coordinates / Routation to zero (will be transmitted seperatly)
		for _, v := range constr {
			arms[v.SiliconID] = DataFrame{v.SiliconID, v.NumberInChain, 0, 0, 0, 0, 0, 0, 0}
			// Compute the correct FrameID

		}
		return nil
	}

}

// Accept/Rejects CAN-Frame
// Frames properties must match to be accepted:
// Correct Frame-ID
// Signature
// Number-in-chain
func ProcessFrame(can_frame DataFrame) error {

	// Is FrameID registered?
	if (arms[can_frame.SiliconID] == DataFrame{}) {
		return errors.New("Couldn't match Silicon ID")
	} else if arms[can_frame.SiliconID].PositionInChain == 0 {
		return errors.New("Zero not allowed as NumberInChain!")
	} else if arms[can_frame.SiliconID].SiliconID != can_frame.SiliconID {
		return errors.New("Silicon IDs mismatched!")
	} else if arms[can_frame.SiliconID].PositionInChain != can_frame.PositionInChain {
		return errors.New("Position in Chain inconsistent")
	} else {
		// All Criteria matched
		return nil
	}
}

func DeleteFrameHistory() {

}

/* Simulation run in goroutine.  For running a simulation, the Field 'simulaition' must be set to 'yes' in config.yaml*/

// Send Dataframes every 10ms to simulate Changes in Position
func SimulateTeachingFrames() {

}

// Calculate a linear movement between Setpoints
func SimulateMotionFrames() {

}

// Stop sending Frames for 20ms
func SimulateFrameTimeout() {

}
