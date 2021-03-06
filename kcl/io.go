package kcl

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
)

type ioHandler struct {
	In  *bufio.Scanner
	Out *bufio.Writer
}

func newIOHandler(stdin io.Reader, stdout io.Writer) *ioHandler {
	return &ioHandler{bufio.NewScanner(stdin), bufio.NewWriter(stdout)}
}

type response struct {
	Action     string `json:"action"`
	For        string `json:"responseFor,omitempty"`
	CheckPoint string `json:"checkpoint,omitempty"`
}

func (ih *ioHandler) sendStatus(action string) error {
	resp := response{Action: "status", For: action}
	return ih.sendResponse(resp)
}

func (ih *ioHandler) sendCheckpoint(seq string) error {
	resp := response{Action: "checkpoint", CheckPoint: seq}
	return ih.sendResponse(resp)
}

func (ih *ioHandler) sendResponse(resp response) error {
	defer ih.Out.Flush()
	return json.NewEncoder(ih.Out).Encode(&resp)
}

func (ih *ioHandler) receiveMessage() (*message, error) {
	buf, err := ih.readLine()
	if err != nil {
		return nil, err
	}

	logger.Printf("incoming message: %s", buf)

	var msg message
	if err := json.Unmarshal(buf, &msg); err != nil {
		return nil, err
	}

	return &msg, nil
}

func (ih *ioHandler) readLine() ([]byte, error) {
	if !ih.In.Scan() {
		return nil, errors.New("failed to scan input")
	}

	return ih.In.Bytes(), nil
}
