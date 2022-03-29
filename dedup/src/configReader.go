package main

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
)

type Delimiter_conf struct {
	Row []int `json:"row"`
	Col []int `json:"col"`
}
type Filter_conf struct {
	Index int      `json:"index"`
	Value []string `json:"value"`
}
type Conf struct {
	Delimiter   *Delimiter_conf `json:"delimiter"`
	Filters     []*Filter_conf  `json:"filters"`
	DedupKey    []int           `json:"dedupKey"`
	ParallelCnt int             `json:"parallelCnt"`
}

func ReadConf(fn string) (conf *Conf, error error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	byt := []byte{}
	for {

		byt, err = reader.ReadBytes('|')

		if err == io.EOF {
			break
		}
	}
	if err = json.Unmarshal(byt, &conf); err != nil {
		return nil, err
	}
	return conf, nil

}
