package bluetooth

import (
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var NonNumberRegex = regexp.MustCompile(`[^0-9]+`)

type UPowerInfo struct {
	Percentage uint
}

func GetUPowerDump() (string, error) {
	b, err := exec.Command("upower", "-d").Output()
	if err != nil {
		return "", err
	}

	out := strings.TrimSpace(string(b))

	return out, nil
}

func GetUPowerInfo(objectPath string) (*UPowerInfo, error) {
	info := new(UPowerInfo)

	b, err := exec.Command("upower", "-i", objectPath).Output()
	if err != nil {
		return nil, err
	}
	out := strings.TrimSpace(string(b))

	lines := strings.Split(out, "\n")

	for _, line := range lines {
		l := strings.TrimSpace(line)
		s := strings.Fields(l)

		if s[0] == "percentage:" {
			v := NonNumberRegex.ReplaceAllString(s[1], "")
			x, err := strconv.Atoi(v)
			if err != nil {
				x = 0
			}

			info.Percentage = uint(x)
		}
	}

	return info, nil
}
