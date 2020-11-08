package lab2

import (
    . "gopkg.in/check.v1"
    "strings"
    "testing"
    "fmt"
)

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s * TestSuite) TestToInfix(c *C) {
    for _, testCase := range []struct{
        dummy string
        expected string
    }{
        {
            dummy: "",
            expected: "empty input",
        },
        {
            dummy: "a - b *",
            expected: "incorrect number of operands",
        },
        {
            dummy: "C 16 - a *",
            expected: "((C - 16) * a)",
        },
        {
            dummy: "a b * c ^ d /",
            expected: "(((a * b) ^ c) / d)",
        },
        {
            dummy: "+5 73 + 20 t ^ /",
            expected: "((5 + 73) / (20 ^ t))",
        },
        {
            dummy: "5 6 d * -  T 30 * ^ c +",
            expected: "(((5 - (6 * d)) ^ (T * 30)) + c)",
        },
        {
            dummy: "a b ^ c * d - e f / g h + / +",
            expected: "((((a ^ b) * c) - d) + ((e / f) / (g + h)))",
        },

        {
            dummy: "a b# + c -",
            expected: "forbidden symbol at: b#",
        },
    }{
        exp := &Exp{Input: strings.NewReader(testCase.dummy)}

        output := make(chan string)
        go exp.ToInfix(output)

        var res string
        for data := range output{
            res = res + data
        }

        c.Assert(res, Equals, testCase.expected)
    }
}

func ExampleExp_ToInfix() {
    exp := &Exp{Input: strings.NewReader("a b c * +")}

    output := make(chan string)
    go exp.ToInfix(output)

    var res string
    for data := range output{
        res = res + data
    }

    fmt.Println(res)
    // Output: (a + (b * c))
}
