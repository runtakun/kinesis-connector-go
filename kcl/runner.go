package kcl

import "io"

func RunCLI(rp RecordProcessor) error {

	ih := NewCLIHandler()
	cp := &CheckPointer{ih, false}
	mh := &messageHandler{ih, rp, cp}

	return mh.doAction()
}

func RunFile(rp RecordProcessor, in io.Reader, out io.Writer, err io.Writer) error {

	ih := NewIOHandler(in, out, err)
	cp := &CheckPointer{ih, false}
	mh := &messageHandler{ih, rp, cp}

	return mh.doAction()
}
