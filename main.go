package main

import (
	"errors"
	"flag"
	"github.com/anhgelus/gomath/interpreter"
	"github.com/anhgelus/gomath/lexer"
	"github.com/anhgelus/gomath/utils"
	"os"
	"strings"
)

var (
	input   string
	decimal bool
)

func init() {
	flag.StringVar(&input, "i", "", "input file")
	flag.BoolVar(&decimal, "d", false, "decimal output")
	flag.Parse()
}

func main() {
	b, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}
	content := strings.Split(string(b), "\n")
	lexed, err := lexer.Lex(content)
	if err != nil {
		panic(err)
	}
	//for _, l := range lexed {
	//	s := ""
	//	for _, v := range l {
	//		s += v.String() + " "
	//	}
	//	println(s[:len(s)-1])
	//}
	p, err := interpreter.Parse(lexed)
	if err != nil {
		panic(err)
	}
	//m, err := json.MarshalIndent(p, "", "  ")
	//if err != nil {
	//	panic(err)
	//}
	//println(string(m))
	for i, stmt := range p.Body {
		err = stmt.Eval(&interpreter.Options{Decimal: decimal})
		if err != nil {
			panic(errors.Join(utils.GenErrorLine(i), err))
		}
	}
}
