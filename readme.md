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
- Cobra (https://github.com/spf13/cobra)
- Viper (https://github.com/spf13/viper)
- Socket Can (https://github.com/brutella/can)
- MathGL (https://github.com/go-gl/mathgl)
- GORM (https://gorm.io/gorm)
- SQLite(https://gorm.io/driver/sqlite)
- Bubbletea (https://github.com/charmbracelet/bubbletea)
