package parser

import (
	"encoding/json"
	"robot-monitor/data"
)

type JsonParser struct{}

func JSON() *JsonParser {
	return &JsonParser{}
}

func (jp *JsonParser) Parse(jsonContent []byte) (robotStatus data.RobotStatus, err error) {
	err = json.Unmarshal(jsonContent, &robotStatus)
	return
}
