package constuser

// DictDeptNumVal Dict for dept_num => dept_value
var DictDeptNumVal = map[int64]string{
	12: "材料",
}

// DictDeptNumStrVal Dict for dept_num_str => dept_value
var DictDeptNumStrVal = map[string]string{
	"12": "材料",
}

// DictDeptNumFullVal Dict for dept_num => dept_full_value
var DictDeptNumFullVal = map[int64]string{
	12: "材料科学与工程学院",
}

// DictDeptNumStrFullVal Dict for dept_num_str => dept_full_value
var DictDeptNumStrFullVal = map[string]string{
	"12": "材料科学与工程学院",
}

// DictCardType Dict for card type/ user identity
var DictCardType = map[string]int64{
	"213": 1, // 本科生
	"220": 2, // 硕士研究生
	"230": 3, // 博士研究生
	"110": 4, // 教师
}
