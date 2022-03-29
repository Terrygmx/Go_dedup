package main

// 定义处理断卡文件的配置，先定义个人证件类型列表，2为电话号码索引
func initDuanka() int{
	IndividualPsptType = map[string]bool{
		"0":true,"1":true,"2":true,"3":true,
		"5":true,"6":true,"7":true,"8":true,
		"9":true,"F":true,"G":true,"H":true,
		"J":true,"K":true,"P":true,
	}
	return 2
}

