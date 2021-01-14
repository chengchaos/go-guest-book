package learning

import (
	"fmt"
	"github.com/chengchao/go-guest-book/learning/infra"
)

func getRetriever() infra.Retriever {
	return &infra.RetrieverImpl{}
}

func Run() {

	var retriever infra.Retriever = getRetriever()
	result := retriever.Get("https://chengchaos.github.io")
	fmt.Printf("%s\n", result)
}
