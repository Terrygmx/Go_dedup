package main

// 定义处理客户用户文件的配置，先定义个人证件类型列表，6为电话号码索引
func initCustUser() int{
	IndividualPsptType = map[string]bool{
		"01":true,"02":true,"05":true,"08":true,
		"09":true,"10":true,"11":true,"12":true,
		"24":true,"33":true,"34":true,
	}
	return 6
}

