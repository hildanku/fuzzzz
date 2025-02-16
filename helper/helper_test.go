package helper

import (
	"bytes"
	"log"
	"testing"
)

func TestHandleError(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	HandleError(nil)
	if buf.Len() != 0 {
		t.Errorf("Expected no log output, got %s", buf.String())
	}

	buf.Reset()

	HandleError(log.Output(2, "test error"))
	if buf.Len() == 0 {
		t.Errorf("Expected no error, got none")
	}
}
