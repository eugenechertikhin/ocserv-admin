package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"ocserv-admin/internal/model"
)

func UpdateData() ([]model.Base, []model.Online, error) {
	// load etc/ocpasswd
	f, err := os.Open(*PASSWD_FILE)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	base := []model.Base{}
	online := []model.Online{}

	scanF := bufio.NewScanner(f)
	for scanF.Scan() {
		line := scanF.Text()
		fields := strings.Split(line, ":")

		r := model.Base{
			User:     fields[0],
			Group:    fields[1],
			Password: fields[2],
		}
		base = append(base, r)
	}

	if err := scanF.Err(); err != nil {
		return nil, nil, err
	}

	// Exec SHOW_USERS
	cmd := exec.Command("occtl", "show", "users")

	on, err := cmd.Output()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to execute command: %v", err)
	}

	scanOn := bufio.NewScanner(bytes.NewReader(on))

	for scanOn.Scan() {
		line := scanOn.Text()
		fields := strings.Fields(line)

		if id, err := strconv.Atoi(fields[0]); err == nil {
			dur, _ := DurationConv(fields[6])
			r := model.Online{
				Id:     id,
				User:   fields[1],
				ExtIp:  fields[3],
				IntIp:  fields[4],
				Device: fields[5],
				Since:  dur,
			}
			online = append(online, r)
		}
	}

	if err := scanF.Err(); err != nil {
		return nil, nil, err
	}

	return base, online, nil
}
