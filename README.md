

---

# Temperature Monitor

This Go package provides a simple command-line utility to monitor CPU and GPU temperatures on Linux systems.

## Installation

To use this package, you need to have Go installed on your system. You can install it using the following command:

```bash
go get -u github.com/yourusername/tempermonitor
```

Replace `yourusername` with your GitHub username.

## Usage

After installing the package, you can use the `tempermonitor` command in your terminal. It provides the following options:

1. **CPU Temperature**: Displays the current CPU temperature in Celsius.
2. **GPU Temperature**: Displays the current GPU temperature in Celsius.
3. **Exit**: Exits the program.

To run the utility, simply execute the following command:

```bash
tempermonitor
```

Follow the on-screen instructions to select an option.

## Notes

- This package is designed to work on Linux systems.
- The CPU temperature is read from `/sys/class/thermal/thermal_zone0/temp`.
- The GPU temperature is read from `/sys/class/hwmon/hwmon1/temp1_input`.
- Press Enter to continue after viewing the temperature.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

