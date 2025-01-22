package main

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)

const (
	literal   string = "literal"
	number    string = "number"
	seperator string = "seperator"
	operator  string = "operator"
)

var (
	InvalidSeparatorErr = errors.New("invalid separator")
	operators           = []string{"+", "-", "*", "/", "^", "%", "="}
)

type Lexer struct {
	Type, Value string
}

func lex(content []string) ([][]*Lexer, error) {
	var lexers [][]*Lexer
	for i, line := range content {
		if strings.HasPrefix(line, "#") {
			continue
		}
		var lexer []*Lexer
		for _, w := range strings.Split(line, " ") {
			word, err := lexWord(w)
			if err != nil {
				return nil, errors.Join(errors.New("line "+strconv.Itoa(i+1)+" has an error"), err)
			}
			lexer = append(lexer, word...)
		}
		lexers = append(lexers, lexer)
	}
	return lexers, nil
}

func lexWord(w string) ([]*Lexer, error) {
	if isDigit(w) {
		return []*Lexer{
			{number, w},
		}, nil
	}
	if strings.Contains(w, ",") {
		if strings.Contains(w[:len(w)-1], ",") {
			return nil, errors.Join(InvalidSeparatorErr, errors.New(w+" contains an invalid comma"))
		}
		return []*Lexer{
			{literal, w[:len(w)-1]},
			{seperator, ","},
		}, nil
	}
	if isOperator(w) {
		return []*Lexer{
			{operator, w},
		}, nil
	}
	return []*Lexer{
		{literal, w},
	}, nil
}

func isDigit(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func isOperator(s string) bool {
	return slices.Contains(operators, s)
}

func (l *Lexer) String() string {
	return l.Type + "(" + l.Value + ")"
}
