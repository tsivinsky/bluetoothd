package bluetooth

import (
	"os/exec"
	"strings"
)

type BluetoothDevice struct {
	ID         string
	Name       string
	ObjectPath string
	Percentage uint
}

func GetConnectedDevices() ([]*BluetoothDevice, error) {
	var devices []*BluetoothDevice

	b, err := exec.Command("bluetoothctl", "devices").Output()
	if err != nil {
		return nil, err
	}
	out := string(b)

	if out == "" {
		return devices, nil
	}

	lines := strings.Split(out, "\n")

	for _, line := range lines {
		s := strings.Split(line, " ")

		if len(s) < 2 {
			continue
		}

		id := s[1]
		name := strings.Join(s[2:], " ")

		devices = append(devices, &BluetoothDevice{
			ID:   id,
			Name: name,
		})
	}

	upowerDump, err := GetUPowerDump()
	if err != nil {
		return nil, err
	}

	dumpLines := strings.Split(upowerDump, "\n")
	for i, line := range dumpLines {
		for _, device := range devices {
			if strings.Contains(line, device.Name) {
				s := strings.Split(dumpLines[i-2], " ")
				if len(s) == 2 {
					objectPath := s[1]
					device.ObjectPath = objectPath
				}
			}
		}
	}

	for _, device := range devices {
		if device.ObjectPath == "" {
			continue
		}

		info, err := GetUPowerInfo(device.ObjectPath)
		if err != nil {
			continue
		}

		device.Percentage = info.Percentage
	}

	return devices, nil
}
