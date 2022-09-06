package message

import "encoding/json"

type CommandMessage struct {
	Payload interface{} `json:"payload"`
	Command string      `json:"command"`
}

func ParseCommandMessage(in []byte) (*CommandMessage, error) {
	var msg CommandMessage

	err := json.Unmarshal(in, &msg)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}
