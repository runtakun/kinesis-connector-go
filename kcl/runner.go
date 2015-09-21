package kcl

import "io"

func RunCLI(rp RecordProcessor) error {

	ih := newCLIHandler()
	cp := &CheckPointer{ih, true}
	mh := &messageHandler{ih, rp, cp}

	return mh.doAction()
}

func RunFile(rp RecordProcessor, in io.Reader, out io.Writer, err io.Writer) error {

	ih := newIOHandler(in, out, err)
	cp := &CheckPointer{ih, true}
	mh := &messageHandler{ih, rp, cp}

	return mh.doAction()
}
