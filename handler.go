package lab2

import (
	"io"
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

	var res string
	//_, err := ioutil.ReadAll(ch.Input)

	// if strings.Contains(string(input), " + ") == false && strings.Contains(string(input), " - ") == false && strings.Contains(string(input), " * ") == false && strings.Contains(string(input), " / ") == false && strings.Contains(string(input), " +") == false && strings.Contains(string(input), " -") == false && strings.Contains(string(input), " *") == false && strings.Contains(string(input), " /") == false {
	// 	return errors.New("Incorrect synthax")
	// }

	//exp := &Exp{Input: strings.NewReader(string(input))}
	exp := &Exp{Input: ch.Input}

	output := make(chan string)

	go exp.ToInfix(output)

	for data := range output {
		res = res + data
	}

	// if err != nil {
	// 	return err
	// } else {
	buf := []byte(res)

	if _, err := ch.Output.Write(buf); err != nil {
		return err
	}
	//}

	return nil
}
