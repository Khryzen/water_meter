package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/mbdeguzman/water_district/models"
	"github.com/tarm/serial"
	"github.com/uadmin/uadmin"
)

const bufferSize = 128

func ReadingService() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			Readings()
		}
	}
}

func Readings() {
	uadmin.Trail(uadmin.INFO, "Serial reader service is running...")

	config := &serial.Config{
		Name:        "/dev/ttyUSB0",
		Baud:        9600,
		ReadTimeout: time.Second * 2,
	}
	port, err := serial.OpenPort(config)
	if err != nil {
		uadmin.Trail(uadmin.ERROR, "Unable to open serial port. %s", err.Error())
		return
	}
	defer port.Close()

	buffer := make([]byte, bufferSize)
	for {
		// Read data into the buffer
		n, _ := port.Read(buffer)

		// Convert read bytes to a string and remove any whitespace
		str := string(buffer[:n])
		processed := strings.Join(strings.Fields(str), "")

		// Print only non-empty data
		if len(processed) > 0 {
			uadmin.Trail(uadmin.INFO, "%s", processed)
			date := strings.Split(processed, ":")[0]
			device_serial := strings.Split(processed, ":")[1]
			start := strings.Split(processed, ":")[2]
			end := strings.Split(processed, ":")[3]
			reading := models.Reading{}
			device := models.Device{}
			client := models.Client{}

			reading.DateString = date
			reading.Beginning, _ = strconv.ParseFloat(start, 64)
			reading.Ending, _ = strconv.ParseFloat(end, 64)

			uadmin.Get(&device, "serial_number = ?", device_serial)
			uadmin.Get(&client, "device_id = ?", device.ID)
			reading.DeviceID = device.ID

			uadmin.Save(&reading)
		}

		// Clear the buffer
		for i := range buffer {
			buffer[i] = 0
		}
	}
}
