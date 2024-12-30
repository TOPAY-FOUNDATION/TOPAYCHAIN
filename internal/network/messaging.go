package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

func SendMessage(address string, msg Message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to serialize message: %v", err)
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/message", address), "application/json", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK status: %d", resp.StatusCode)
	}

	return nil
}

func HandleMessage(data []byte) (Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return Message{}, fmt.Errorf("failed to deserialize message: %v", err)
	}
	return msg, nil
}
