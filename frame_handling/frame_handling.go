package framehandling

import (
	"errors"
	"lpb/storage"
	"lpb/tui"
	"math/rand"
	"time"

	"github.com/brutella/can"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-gl/mathgl/mgl64"
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
	arms           = make(map[int]DataFrame)
	frames_channel chan DataFrame // All simulated / CAN-Frames are routed over this channel
	id_loopback    chan int       // Loops back id from database into simulation
	// Channels to stop looping goroutines
	stop_sim chan bool
	shutdown chan bool

	NullFrameError         error = errors.New("Nulled Frame detected")
	NICNullError           error = errors.New("NumberInChain incorrect in DB!")
	NICNotTransmittedError error = errors.New("NumberInChain not transmitted!")
	NICError               error = errors.New("Couldn't match NUmberInChain!")
	FrameIDError           error = errors.New("Couldn't match Frame ID!")
)

// Initialize the timeout and load frame data
func Init(frame_channel chan DataFrame, id_sim_channel chan int) error {
	//arms = make(map[int]DataFrame)
	// Get all registered arms from DB
	id_loopback = id_sim_channel
	constr, err := storage.GetArms()
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
		return NullFrameError
	} else if arms[can_frame.SiliconID].PositionInChain == 0 {
		return NICNullError
	} else if arms[can_frame.SiliconID].PositionInChain != can_frame.PositionInChain {
		if can_frame.PositionInChain == 0 {
			return NICNotTransmittedError
		} else {
			return NICError
		}

	} else if arms[can_frame.SiliconID].FrameID != can_frame.FrameID {
		return FrameIDError
	} else {
		// All Criteria matched
		return nil
	}
}

func DeleteFrameHistory() {

}

func GeneralFrameHandler(sim_frame <-chan DataFrame, can_frame <-chan can.Frame) {
	for {
		select {
		case <-sim_frame:
			if ProcessFrame(<-sim_frame) != nil {

			} else {
				//TODO:Quit out of Run-Mode by sending error_frame, quit out of TUI with error
			}
		case <-can_frame:
		case <-shutdown: // shutdown goroutine to cleanly exit program
			return
		default:
		}
	}
}

/* Simulation run in goroutine. */

//For running a simulation, the Field 'simulation' must be set to 'yes' in simulation.yaml
//Set the number of arms under 'simulated_links'*/
func InitSimulation(count int) error {
	// Generate IDs and Positions
	for i := 0; i < count; i++ {
		if i == 2 {
			// Edit Element 2 intentionally, to contain wrong NIC
			arms[i] = DataFrame{i, 0, i * 100, 0, 0, 0, 0, 0, 0}
		} else {
			arms[i] = DataFrame{i, i, i * 100, 0, 0, 0, 0, 0, 0}
		}
	}
	return nil
}

// Switches asyncronously the used simulation
func StartSimSwitch(p *tea.Program, idle <-chan bool, movement <-chan bool) {
	for {
		select {
		case <-shutdown:
			return
		case <-idle:
			SwitchSimulation(<-idle, false)
			p.Send(tui.HandshakeMsg{})
		case <-movement:
			SwitchSimulation(<-idle, false)
			p.Send(tui.HandshakeMsg{})
		default:
		}
	}
}

// Preselect simulation:
func SwitchSimulation(Idle bool, Motion bool) {
	stop_sim <- true
	if Idle {
		go SimulateIdleFrames(len(arms), frames_channel)
	} else {
		go SimulateMotionFrames(id_loopback, frames_channel)
	}
}

//Send Dataframes every 10ms to simulate Changes in Position on channel
//Idle-Simulation Stops, when quit is sent into channel.
// Count begins at 1.
func SimulateIdleFrames(count int, tx_frame chan<- DataFrame) {
	for {
		select {
		case <-stop_sim:
			return
		default:
			// Do other stuff
			for _, v := range arms {
				v.CoordinateX = rand.Int31n(100)
				v.CoordinateY = rand.Int31n(100)
				v.CoordinateZ = rand.Int31n(100)
				v.RotationX = rand.Int31n(360)
				v.RotationY = rand.Int31n(360)
				v.RotationZ = rand.Int31n(360)
				tx_frame <- v
				time.Sleep(100 * time.Millisecond)

			}
		}
	}

}

// Calculate a linear movement between Setpoints
// Set field "time_slots" to > 1sec for correct frames, < 1sec for frame Timeouts
func SimulateMotionFrames(id <-chan int, tx_frame chan<- DataFrame) {
	for {
		select {
		case <-stop_sim:
			return
		default:
			// Get all waypoins from db by id
			waypoints := make([]mgl64.Vec2, 10)
			// Play back all waypoints
			for _, v := range waypoints {

				var temp_frame DataFrame
				temp_frame.SiliconID = int(arms[0].SiliconID)
				temp_frame.PositionInChain = arms[0].PositionInChain
				temp_frame.FrameID = int(arms[0].FrameID)
				temp_frame.CoordinateX = int32(v.X())
				temp_frame.CoordinateY = int32(v.Y())
				tx_frame <- temp_frame
				time.Sleep(1 * time.Second)
			}
		}
	}

}
