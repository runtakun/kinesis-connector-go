package kcl

import (
	"errors"
	"fmt"
)

type checkPointResponse struct {
	Action     string `json:"action"`
	CheckPoint string `json:"checkpoint"`
}

type CheckPointer struct {
	checkPointAllowed bool
}

func (cp *CheckPointer) CheckPointAll() error {
	return cp.doCheckPoint(&checkPointResponse{})
}

func (cp *CheckPointer) CheckPointSeq(seq string) error {
	return cp.doCheckPoint(&checkPointResponse{CheckPoint: seq})
}

func (cp *CheckPointer) doCheckPoint(resp *checkPointResponse) error {
	if !cp.checkPointAllowed {
		return errors.New("checkpoint is not allowed")
	}

	message, err := getMessage()
	if err != nil {
		return err
	}

	if message.Action != "checkpoint" {
		return fmt.Errorf("invalid action: %s", message.Action)
	} else if message.Error != nil {

	}

	return nil
}
