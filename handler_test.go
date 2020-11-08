package lab2

import (
	"bytes"
	"strings"

	. "gopkg.in/check.v1"
)

func (s *TestSuite) TestHandler1(c *C) {
	input := "42 1 -"
	exp := strings.NewReader(input)
	buf := new(bytes.Buffer)

	handler := ComputeHandler{
		Input:  exp,
		Output: buf,
	}

	handler.Compute()

	c.Assert(buf.String(), Equals, "(42 - 1)")
}

func (s *TestSuite) TestHandler2(c *C) {
	input := "5 6+"
	exp := strings.NewReader(input)

	buf := new(bytes.Buffer)

	handler := ComputeHandler{
		Input:  exp,
		Output: buf,
	}

	err := handler.Compute()

	c.Assert(err, NotNil)
}
