package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readCpuTemperature() (float64, error) {
	file, err := os.Open("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	temperatureStr := scanner.Text()
	temperature, err := strconv.ParseFloat(temperatureStr, 64)
	if err != nil {
		return 0, err
	}
	return temperature / 1000.0, nil
}
func readGpuTemperature() (float64, error) {
	file, err := os.Open("/sys/class/hwmon/hwmon1/temp1_input")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	temperatureStr := scanner.Text()
	temperature, err := strconv.ParseFloat(temperatureStr, 64)
	if err != nil {
		return 0, err
	}
	return temperature / 1000.0, nil

}
func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Options:")
		fmt.Println("1: CPU Temperature")
		fmt.Println("2: GPU Temperature")
		fmt.Println("3: Exit")

		fmt.Print("Enter your choice: ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		switch userInput {
		case "1":
			cpuTemperature, err := readCpuTemperature()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Failed to read CPU temperature:", err)
			} else {
				fmt.Printf("CPU Temperature: %.2f °C\n", cpuTemperature)
			}
		case "2":
			gpuTemperature, err := readGpuTemperature()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Failed to read GPU temperature:", err)
			} else {
				fmt.Printf("GPU Temperature: %.2f °C\n", gpuTemperature)
			}
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
		}

		fmt.Println("Press Enter to continue...")
		reader.ReadString('\n')
	}
}
