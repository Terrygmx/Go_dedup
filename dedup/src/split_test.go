package main

import (
	"testing"
)

func TestSplit(t *testing.T) {
	t.Log(split("sample.txt"))
}

func TestDedup(t *testing.T) {
	t.Log(dedup("split_output/sample.txt_99"))
}


//func TestDedupAll(t *testing.T) {
//	t.Log(dedup("split_output/sample.txt_99"))
//}