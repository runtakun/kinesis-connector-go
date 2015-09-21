package kcl

import (
	"bufio"
	"bytes"
	"encoding/json"
	"os"
)

func Run(rp RecordProcessor) error {

	cp := &CheckPointer{}

	message, err := getMessage()
	if err != nil {
		return err
	}

	return doAction(rp, cp, message)
}

func getMessage() (*message, error) {
	bio := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer
	for {
		line, hasMore, err := bio.ReadLine()
		if err != nil {
			return nil, err
		}
		buf.Write(line)
		if !hasMore {
			break
		}
	}

	var message *message
	if err := json.Unmarshal(buf.Bytes(), message); err != nil {
		return nil, err
	}

	return message, nil
}
