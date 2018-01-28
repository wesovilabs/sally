package main

import (
	"fmt"

	"github.com/wesovilabs/sally/nlp"
)

func main() {
	tech := &nlp.Tokenize{}
	tech.SetLanguage(nlp.ES)
	tech.SetDataInput(&nlp.DataInput{SrcCntent: "Hello my friend!\n How are you?\n I'm fine thanks, and you?\n I'm ok too"})
	res := nlp.Run(tech)
	for _, item := range res.(*nlp.TokenizeOutput).Items {
		fmt.Println(*item)
	}
	fmt.Println("*******************")
	tech.SetDataInput(&nlp.DataInput{SrcCntent: "Hello my friend!\n How are you?\n I am fine thanks, and you?\n I am ok too"})
	res = nlp.Run(tech)
	for _, item := range res.(*nlp.TokenizeOutput).Items {
		fmt.Println(*item)
	}
}
