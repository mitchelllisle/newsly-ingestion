package main

import (
	"fmt"
	"newsly-ingestion/internal/arxiv"
)

func main() {
	arxivAPI := arxiv.InitArxivAPI("http://export.arxiv.org/api/query")

	articles := arxivAPI.GetArticlesForTerm("privacy")

	fmt.Println(articles)
}
