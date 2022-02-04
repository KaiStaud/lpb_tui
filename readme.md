## Hardware
Any embedded linux device which meets the following requirements can be used:
- 2 CAN Controllers (integrated or external)
- 2 I2C Bus
- 5 GPIO Pins

Although LPBs Firmware is tested and developed on Beaglebords Beaglebone Black, any device
which fullfills the listed requirements can be used as a coordinator node.

## Go Dependencies
LPBs functionality depends partially on external provided community modules.
Before installing lpb on your device fetch the following dependencies with go-get:

### Hardware Dependencies
The controller node uses two CAN Networks to communicate with its slaves.
For isolated testing, all Frames will be redirected to internally to channels ("internal loopback"), but the following
package still needs to be included:

- Socket Can (https://github.com/brutella/can)
- gpio (https://github.com/google/periph)
- gosensors(https://github.com/ssimunic/gosensors)
- gopsutil(https://github.com/shirou/gopsutil)
For linear algebra the Inverse Kinematics are calculated with help of the followin package:
- MathGL (https://github.com/go-gl/mathgl)

### Data Handling
Data is stored in an SQLite Database, accessing it is done via GORM.
A individual yaml configuration file is used for generating reports and log files.
Corresponding entries are read and processed with viper.

- Viper (https://github.com/spf13/viper)
- GORM (https://gorm.io/gorm)
- SQLite(https://gorm.io/driver/sqlite)

### User Interface

LPB uses Charms TUI Libraries for rendering an user interface;
currently the following components are used:
- Bubbletea (https://github.com/charmbracelet/bubbletea)
- Bubbles (https://github.com/charmbracelet/bubbles)
- Lipgloss (github.com/charmbracelet/lipgloss)

### Tools
- Zap (https://github.com/uber-go/zap)