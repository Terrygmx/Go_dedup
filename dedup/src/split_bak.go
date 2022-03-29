package main

// 按照记录主键分组，进行文件拆分
import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strings"
)

//type rowDef struct {
//	provinceCode	string
//	userId	string
//	serial	string
//	psptId	string
//	psptType	string
//}
func Parse_bak(row []byte) (out [][]byte) {
	var de byte = byte(colDe[0])
	buf := make([]byte, 0)
	for i := range row {
		if row[i] != de {
			buf = append(buf, row[i])
		} else {
			out = append(out, buf)
			buf = []byte{}
		}
	}
	return
}
func split_bak(fn string) error {
	f, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer f.Close()
	arrOfFile := make([]*os.File, splitFileCount)
	arrOfWriter := make([]*bufio.Writer, splitFileCount)
	reader := bufio.NewReader(f)
	var lf byte = byte(rowDe[0])
	key := keyIndex[0]
	for {
		byt, err := reader.ReadBytes(lf)
		if err == io.EOF {
			break
		}
		parseRes := Parse(byt)
		// 按照哈希求模的原理进行文件拆分
		hashRes := md5.Sum(parseRes[key])
		hashOfSeiral := (int(hashRes[15]) + int(hashRes[14])<<2) % splitFileCount
		if arrOfFile[hashOfSeiral] == nil {
			fnArr := strings.Split(fn, "/")
			fnTail := fnArr[len(fnArr)-1]
			fnHere := fmt.Sprintf("%s/%s_%d", outputPath, fnTail, hashOfSeiral)
			var err error
			arrOfFile[hashOfSeiral], err = os.OpenFile(fnHere, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				return err
			}
			arrOfWriter[hashOfSeiral] = bufio.NewWriter(arrOfFile[hashOfSeiral])
			defer arrOfFile[hashOfSeiral].Close()
		}
		arrOfWriter[hashOfSeiral].Write(byt)

	}
	for _, each := range arrOfWriter {
		each.Flush()
	}
	return nil

}
