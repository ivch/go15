package printer

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

func Test_Listen(t *testing.T) {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w

	testPrinter := make(chan string)
	go Listen(testPrinter)
	testPrinter <- "test msg"

	outputChan := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outputChan <- buf.String()
	}()

	time.Sleep(1 * time.Second)

	w.Close()
	os.Stdout = old

	ar := <-outputChan
	er := fmt.Sprintf("%s%s%s", screenTitle, "test msg", gameLegend)

	if ar != er {
		t.Error("actual text doesn't equal to expected")
	}
}
