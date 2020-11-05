package lab2

import (
	"bufio"
	"io"
)

type Stack []string

type Exp struct {
	Input io.Reader
}

func IsOperand(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') {
			return false
		}
	}
	return true
}

func IsOperator(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "^" || s == "/" {
		return true
	}
	return false
}

func (exp *Exp) ToInfix(output chan string) {
	scanner := bufio.NewScanner(exp.Input)
	scanner.Split(bufio.ScanWords)
	var stack Stack

	for scanner.Scan() {
		item := scanner.Text()
		if IsOperand(item) {
			stack = append(stack, item)

		} else if IsOperator(item) {
			if len(stack) < 2 {
				output <- "incorrect number of operands"
				close(output)
				return
			}
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op2 := stack[len(stack)-1]

			stack = stack[:len(stack)-1]
			newItem := "(" + op2 + " " + item + " " + op1 + ")"
			stack = append(stack, newItem)

		} else {
			output <- "forbidden symbol at: "
			output <- item
			close(output)
			return
		}
	}
	if len(stack) == 1 {
		output <- stack[0]
	} else if len(stack) == 0 {
		output <- "empty input"
	}
	close(output)
}
