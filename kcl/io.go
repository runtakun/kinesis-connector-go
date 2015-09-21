package kcl

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type ioHandler struct {
	In  *bufio.Scanner
	Out *bufio.Writer
	Err *bufio.Writer
}

func NewCLIHandler() *ioHandler {
	return &ioHandler{bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout), bufio.NewWriter(os.Stderr)}
}

func NewIOHandler(stdin *os.File, stdout *os.File, stderr *os.File) *ioHandler {
	return &ioHandler{bufio.NewScanner(stdin), bufio.NewWriter(stdout), bufio.NewWriter(stderr)}
}

type response struct {
	Action     string `json:"action"`
	For        string `json:"responseFor,omitifempty"`
	CheckPoint string `json:"checkpoint,omitifempty"`
}

func (ih *ioHandler) sendStatus(action string) error {
	resp := &response{Action: "status", For: action}
	return ih.sendResponse(resp)
}

func (ih *ioHandler) sendCheckpoint(seq string) error {
	resp := &response{Action: "checkpoint", CheckPoint: seq}
	return ih.sendResponse(resp)
}

func (ih *ioHandler) sendResponse(resp *response) error {
	return json.NewEncoder(ih.Out).Encode(resp)
}

func (ih *ioHandler) receiveMessage() (*message, error) {
	buf, err := ih.readLine()
	if err != nil {
		return nil, err
	}

	var msg *message
	if err := json.Unmarshal(buf, msg); err != nil {
		return nil, err
	}

	return msg, nil
}

func (ih *ioHandler) readLine() ([]byte, error) {
	if !ih.In.Scan() {
		return nil, errors.New("failed to scan input")
	}

	return ih.In.Bytes(), nil
}
