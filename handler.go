package lab2

import (
	//"fmt"

	"io"
	"io/ioutil"
	"strings"
	//"io/ioutil"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	// TODO: Add necessary fields.
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	//TODO: Implement.

	expr, err := ioutil.ReadAll(ch.Input)
	exp := &Exp{Input: strings.NewReader(string(expr))}
	output := make(chan string)
	go exp.ToInfix(output)
	var res string
	for data := range output {
		res = res + data
	}

	if err != nil {
		return err
	} else {
		buf := []byte(res)

		if _, err := ch.Output.Write(buf); err != nil {
			return err
		}
	}

	return nil
}
