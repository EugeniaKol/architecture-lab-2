package main

import (
	"flag"
	"io"
	"io/ioutil"
	"os"
	"strings"

	lab2 "github.com/roman-mazur/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	// TODO: Add other flags support for input and output configuration.
	inputFile  = flag.String("f", "", "Input file")
	outputFile = flag.String("o", "", "Output file")
)

func main() {
	flag.Parse()

	var src io.Reader
	var dest io.Writer

	if *inputExpression != "" {
		src = strings.NewReader(*inputExpression)
	} else if *inputFile != "" {
		data, err := ioutil.ReadFile(*inputFile)
		if err != nil {
			_, _ = os.Stderr.WriteString(err.Error())
			return
		}
		src = strings.NewReader(string(data))
	} else {
		_, _ = os.Stderr.WriteString("error: no expression provided")
		return
	}

	if *outputFile != "" {
		if file, err := os.Create(*outputFile); err == nil {
			dest = file
		} else {
			_, _ = os.Stderr.WriteString(err.Error())
			return
		}
	} else {
		dest = os.Stdout
	}

	handler := lab2.ComputeHandler{
		Input:  src,
		Output: dest,
	}

	if err := handler.Compute(); err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		return
	}
	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

	// handler := &lab2.ComputeHandler{
	// 	Input:  src,
	// 	Output: dest,
	// }

	// err := handler.Compute()
	// res, _ := lab2.ToInfix()
	// fmt.Println(res)
}
