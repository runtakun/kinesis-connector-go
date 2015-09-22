package kcl

import (
	"io"
	"os"
)

func RunCLI(rp RecordProcessor) {

	ih := newIOHandler(os.Stdin, os.Stdout)
	cp := &CheckPointer{ih, true}
	mh := &messageHandler{ih, rp, cp}

	run(mh)
}

func RunIO(rp RecordProcessor, in io.Reader, out io.Writer) {

	ih := newIOHandler(in, out)
	cp := &CheckPointer{ih, true}
	mh := &messageHandler{ih, rp, cp}

	run(mh)
}

func run(mh *messageHandler) {
	for {
		if err := mh.doAction(); err != nil {
			logger.Printf("action process err: %s", err.Error())
			break
		}
	}
}
