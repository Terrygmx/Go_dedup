package main

// 合并文件
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func mergeFile(paths []string, originalFile string) error {
	fnWriter := fmt.Sprintf("%s_dedup", originalFile)
	writerFile, err := os.OpenFile(fnWriter, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(writerFile)
	defer writerFile.Close()
	var lf byte = '\n'
	for _, path := range paths {
		cnt := 0
		fn := fmt.Sprintf("%s_dedup", path)
		f, err := os.Open(fn)
		if err != nil {
			return err
		}
		defer f.Close()
		reader := bufio.NewReader(f)
		for {
			byt, err := reader.ReadBytes(lf)
			if err == io.EOF {
				break
			}
			writer.Write(byt)
			cnt++
			if cnt >= 10000 {
				cnt = 0
				writer.Flush()
			}
		}
		writer.Flush()

	}
	return nil
}
