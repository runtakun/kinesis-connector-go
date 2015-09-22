package kcl

import (
	"errors"
	"fmt"
)

type CheckPointer struct {
	ih                *ioHandler
	checkPointAllowed bool
}

func (cp *CheckPointer) CheckPointAll() error {
	return cp.doCheckPoint("")
}

func (cp *CheckPointer) CheckPointSeq(seq string) error {
	return cp.doCheckPoint(seq)
}

func (cp *CheckPointer) doCheckPoint(seq string) error {
	if !cp.checkPointAllowed {
		return errors.New("checkpoint is not allowed")
	}

	if err := cp.ih.sendCheckpoint(seq); err != nil {
		return err
	}

	msg, err := cp.ih.receiveMessage()
	if err != nil {
		return err
	}

	if msg.Action != "checkpoint" {
		return fmt.Errorf("invalid action: %s", msg.Action)
	} else if msg.Error != nil {
		return errors.New(*msg.Error)
	}

	return nil
}
