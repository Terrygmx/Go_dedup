package main

// 使用哈希表实现去重
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func dedup(fn string) error {
	f, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer f.Close()
	rowDic := make(map[string][]byte)
	var lf byte = byte(rowDe[0])
	key := keyIndex[0]
	reader := bufio.NewReader(f)
	for {
		byt, err := reader.ReadBytes(lf)
		if err == io.EOF {
			break
		}
		parseRes := Parse(byt)
		var psptType string
		serial := string(parseRes[key])
		if len(parseRes[4]) > 0 {
			psptType = string(parseRes[4])
		} else {
			psptType = " "
		}
		if _, ok := rowDic[serial]; ok {
			if IndividualPsptType[psptType] {
				rowDic[serial] = byt
			}
		} else {
			rowDic[serial] = byt
		}

	}
	fnWriter := fmt.Sprintf("%s_dedup", fn)
	writerFile, err := os.OpenFile(fnWriter, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(writerFile)
	defer writerFile.Close()
	var cnt int
	for _, row := range rowDic {
		writer.Write(row)
		cnt++
		if cnt >= 10000 {
			writer.Flush()
			cnt = 0
		}
	}
	writer.Flush()

	return nil

}
