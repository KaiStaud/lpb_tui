/*
*  package for synchronising operational modes of robot.
*  The operational mode can be changed by the following packages:
*	- CLI
*	- Rotary Encoder Interface
*	- Error during Runtime (IK-Error)
*	- Teaching during Idle-Time (IK-Teaching)
 */

package modes

import "errors"

/* Available modes to robot */

type mode int32

const (
	BOOTING  = iota
	IDLE     = iota
	RUNNING  = iota
	TEACHING = iota
	ERROR    = iota
)

var master_mode mode = BOOTING

/* Change mode: */
func SetMode(new_mode mode) (mode, error) {

	/* Last action must be finished before changing to new mode*/
	if master_mode == RUNNING {
		return master_mode, errors.New("Action needs to be finished")
	} else {
		return new_mode, nil
	}
}
