package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
)

// os.Args参数含义 1.待处理文件名 2.拆分文件放入的路径 3.拆分文件数量
var (
	pathChan           chan string
	wg                 sync.WaitGroup
	splitFileCount     int
	conf               *Conf
	keyIndex           []int
	parallelPool       int            // 设置并发数
	rowDe              []int          // 换行字符序列
	colDe              []int          // 换列字符序列
	outputPath         string         // 临时文件目录
	filters            []*Filter_conf // 过滤器
	IndividualPsptType map[string]bool
)

func initFilter() {
	IndividualPsptType = make(map[string]bool)
	for _, f := range filters {
		for _, v := range f.Value {
			IndividualPsptType[v] = true
		}
	}
}

func main() {
	if len(os.Args) >= 5 {
		var err error
		splitFileCount, err = strconv.Atoi(os.Args[3])
		if err != nil || splitFileCount <= 0 {
			return
		}
		conf, err = ReadConf(os.Args[4])

		if err != nil {
			fmt.Println("读取配置文件错误！配置文件格式请参考配置样例")
			panic(err)
		}
		parallelPool = conf.ParallelCnt
		rowDe = conf.Delimiter.Row
		colDe = conf.Delimiter.Col
		keyIndex = conf.DedupKey
		outputPath = os.Args[2]
		filters = conf.Filters
		initFilter()

		if err = os.Mkdir(outputPath, 0755); err != nil {
			fmt.Println("该文件夹已经存在或不可访问！")
			panic(err)
		}
		fmt.Printf("split file into %d files\n", splitFileCount)
		err = split(os.Args[1]) //拆分文件
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("done spliting file:", os.Args[1])
		}

		var FileInfo []os.FileInfo
		goer := parallelPool
		wg.Add(goer)

		if FileInfo, err = ioutil.ReadDir(outputPath); err != nil {
			fmt.Println("读取文件夹出错")
			return
		}
		var output []string
		for _, fileInfo := range FileInfo {
			n := fileInfo.Name()
			fullPath := fmt.Sprintf("%s/%s", outputPath, n)
			output = append(output, fullPath)

		}
		pathChan = make(chan string)
		for i := 0; i < goer; i++ {
			go trigger(dedup)
		}
		for _, s := range output {
			pathChan <- s
		}
		close(pathChan)
		wg.Wait()
		if err = mergeFile(output, os.Args[1]); err != nil {
			fmt.Println("合并文件出错")
			panic(err)
		} else {
			fmt.Println("文件处理完成。")
			return
		}
	} else {
		fmt.Println("请输入参数。参数含义 1.待处理文件名 2.拆分文件放入的路径 3.拆分文件数量 4.配置文件路径")

	}
}
