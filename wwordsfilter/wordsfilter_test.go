package wwordsfilter

import (
	"log"
	"testing"
)

func Init() {
	SetWordsFilter()
}

func TestContentFilter(t *testing.T) {
	Init()
	//fmt.Println(words)
	if ContentFilter("SB") {
		log.Fatal("SB 是敏感词")
	}
	log.Fatal("SB 不是敏感词")
}
