package main

import (
	"fmt"
	"os"

	"github.com/runtakun/kinesis-connector-go/kcl"
)

type exampleProcessor struct {
}

func (p *exampleProcessor) Initialize(shardID string) error {
	fmt.Fprintf(os.Stderr, "processor initialized. shardID: %s\n", shardID)
	return nil
}
func (p *exampleProcessor) ProcessRecords(records []*kcl.Record, cp *kcl.CheckPointer) error {
	fmt.Fprintf(os.Stderr, "process records\n")
	return nil
}

func (p *exampleProcessor) Shutdown(reason string, cp *kcl.CheckPointer) error {
	fmt.Fprintf(os.Stderr, "shutdown. resion: %s", reason)
	return nil
}

type exampleLogger struct {
}

func (l *exampleLogger) Printf(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", v...)
}

func init() {
	kcl.SetLogger(&exampleLogger{})
}

func main() {
	kcl.RunCLI(&exampleProcessor{})
}
