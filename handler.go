package lab2

import (
	//"fmt"
	"io"
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
	//f := ToInfix("33")
	//TODO: Implement.
	// expr, err := ioutil.ReadAll(ch.Input)

	// if err != nil {
	// 	return err
	// }
	// if res, err := ToInfix(expr); err != nil {
	// 	return err
	// } else {
	// 	buf := []byte(fmt.Sprintf("%s", res))
	// 	if _, e := ch.Output.Write(buf); e != nil {
	// 		return e
	// 	}
	// }

	return nil
}
