/*
* Triggers an testing routine on controller and device level.
* If Mode is set to "internal loopback", the module will get the requested information over package "simulation"
* Otherwise the devices will perform connectivity and sensor tests and report back to the controller.
* Operational Test should be performed in regular intervals.
*If not set otherwise, viper will create an default one to enshure this.
 */

package optest

import (
	"errors"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//----------------------- Types ----------------------- //
type wire_status int
type imu_status int
type drive_consistency int
type communication_status int
type node_consistency int

type TestValues struct {
	wire_output string
	imu_error   string
	drive_error string
	comms_alive string
	nodes_ok    string
}

//----------------------- Constansts ----------------------- //

const (

	// Wire Status
	firm_contact  wire_status = iota
	disconnected  wire_status = iota
	loose_contact wire_status = iota

	// IMU Status
	internal_available      imu_status = iota
	external_available      imu_status = iota
	no_replies              imu_status = iota
	internal_with_auxiliary imu_status = iota

	// Drive consistency
	analog_feedback_inconsistent  drive_consistency = iota
	digital_feedback_inconsistent drive_consistency = iota
	feedback_consistent           drive_consistency = iota

	// Communication Status
	alive          communication_status = iota
	partialy_alive communication_status = iota
	dead           communication_status = iota
	with_errors    communication_status = iota

	// Node Consistency
	all_correct  node_consistency = iota
	ids_incorect node_consistency = iota
	hotplugged   node_consistency = iota
)

//----------------------- Variables ----------------------- //

var ()

//----------------------- Interfaces ----------------------- //

type test_replies interface {
	wired() wire_status
	imu_running() imu_status
	drive_consistent() drive_consistency
	communication_active() communication_status
	nodes_consistent() node_consistency
}

//----------------------- Functions ----------------------- //

/*
* Enable loopback test with values from config.yaml
* Set enable_loopback to false if loopback is not desired
 */
func SetConfig(path string, name string) error {

	if _, err := os.Stat(path); err == nil {

		return LoadFromYAML(path, name)
	} else if errors.Is(err, os.ErrNotExist) {
		return errors.New("No Config file found!")
	} else {
		return err
	}
}

/*
* Load passed test configuration
 */
func LoadFromYAML(path string, name string) error {
	viper.SetConfigName(name)
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Couldn't read in config file %v", err)
	}

	var test_output TestValues
	err = viper.Unmarshal(&test_output)
	if err != nil {
		log.Fatalf("Couldn't decode configs values into struct %v", err)
	}

	// Watch for Changes during Runtime:
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		err = viper.Unmarshal(&test_output)
		if err != nil {
			log.Fatalf("could not decode config after changed: %v", err)
		}
	})
	return err
}
