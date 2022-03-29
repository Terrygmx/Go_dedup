package main

import (
	"testing"
)

func TestJsonConfigReader(t *testing.T) {
	conf, _ := ReadConf("config.json")
	t.Logf("%+v\n", conf)
}
