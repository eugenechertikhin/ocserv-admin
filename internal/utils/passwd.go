package utils

import (
	"bufio"
	"os"
	"strings"
	"ocserv-admin/internal/model"
)

func LoadPasswd(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		fields := strings.Split(line, ":")

		r := model.Base{
			User:     fields[0],
			Group:    fields[1],
			Password: fields[2],
		}
		model.Basic = append(model.Basic, r)
	}

	if err := scan.Err(); err != nil {
		return err
	}

	return nil
}
