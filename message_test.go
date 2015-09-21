package kinesis_connector_go_test

import (
	"bytes"
	"fmt"

	"code.google.com/p/go-uuid/uuid"

	. "github.com/runtakun/kinesis-connector-go/kcl"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type testProcessor struct {
}

func (p *testProcessor) Initialize(shardID string) error {
	return nil
}
func (p *testProcessor) ProcessRecords(records []*Record, cp *CheckPointer) error {
	return nil
}

func (p *testProcessor) Shutdown(reason string, cp *CheckPointer) error {
	return nil
}

type ginkgoLogger struct {
}

func (l *ginkgoLogger) Init() {
}

func (l *ginkgoLogger) Log(msg string) {
	GinkgoWriter.Write([]byte(msg))
	GinkgoWriter.Write([]byte{0xa}) //new line
}

var _ = Describe("Message", func() {

	BeforeEach(func() {
		SetLogger(&ginkgoLogger{})
	})

	Context("receive intialize message", func() {

		var stdout *bytes.Buffer
		var processor RecordProcessor

		BeforeEach(func() {
			stdout = new(bytes.Buffer)
			processor = &testProcessor{}
		})

		It("should be initialize processor", func() {
			stdin := bytes.NewBufferString(fmt.Sprintf(`{"action": "initialize", "shardId": "%s"}`, uuid.NewRandom().String()))
			Expect(RunFile(processor, stdin, stdout)).ShouldNot(HaveOccurred())
			Expect(stdout.Bytes()).Should(MatchJSON(`{"action": "status", "responseFor": "initialize"}`))
		})
	})
})
