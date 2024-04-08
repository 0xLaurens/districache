package command

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseCmd(raw []byte) (Command, error) {
	rawStr := string(raw)
	parts := strings.Split(rawStr, " ")
	invalid := fmt.Errorf("invalid command (%s)", parts)

	if len(parts) == 0 || len(parts) < 2 {
		return nil, invalid
	}
	base := BaseCMD{
		Cmd: CMD(parts[0]),
		Key: []byte(parts[1]),
	}

	switch CMD(parts[0]) {
	case CMDGet, CMDHas, CMDDelete:
		return base, nil
	case CMDSet:
		if len(parts) < 3 || len(parts) > 4 {
			return nil, invalid
		}

		cmd := SetCMD{
			BaseCMD: base,
			Value:   []byte(parts[2]),
			TTL:     0,
		}

		if len(parts) == 4 {
			ttl, err := strconv.Atoi(parts[3])
			if err != nil {
				return nil, invalid
			}
			cmd.TTL = time.Millisecond * time.Duration(ttl)
		}

		return cmd, nil
	}

	return nil, invalid
}
