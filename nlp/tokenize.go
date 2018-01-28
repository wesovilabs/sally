package nlp

import (
	"strings"
	"text/scanner"
)

//Tokenize structure used to handle tokenization
type Tokenize struct {
	BaseNlp
}

//TokenizeOutputItem structure for items in array
type TokenizeOutputItem struct {
	Pos  scanner.Position
	Text string
}

//TokenizeOutput structure
type TokenizeOutput struct {
	Items []*TokenizeOutputItem
}

func (to *TokenizeOutput) addItem(pos scanner.Position, text string) {
	to.Items = append(to.Items, &TokenizeOutputItem{pos, text})
}

//Process function used to do tokenization
func (tk *Tokenize) Process() DataOutput {
	out := &TokenizeOutput{}
	var s scanner.Scanner
	s.Init(strings.NewReader(tk.dataInput.SrcCntent))
	var tok rune
	for tok != scanner.EOF {
		tok = s.Scan()
		out.addItem(s.Pos(), s.TokenText())
	}
	return out
}
