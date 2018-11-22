package user

// User 用户表
type User struct {
	UserID   int64  `json:"user_id" db:"user_id"`     //	用户 ID 用户唯一标识符
	UserUUID string `json:"user_uuid" db:"user_uuid"` // 用户 UUID
	CardID   int64  `json:"card_id" db:"card_id"`     //	学生一卡通编号
	QQID     int64  `json:"qq_id" db:"qq_id"`         //	用户绑定 QQ 账号
	WeChatID int64  `json:"wechat_id" db:"wechat_id"` //	用户绑定微信账号
	StuNo    string `json:"stu_no" db:"stu_no"`       //	学生学号
	RealName string `json:"real_name" db:"real_name"` //	用户真实姓名
	NickName string `json:"nick_name" db:"nick_name"` //	用户昵称
	Gender   int64  `json:"gender" db:"gender"`       //	用户性别 0-未知 1-男 2-女
	UserType int64  `json:"user_type" db:"user_type"` //	用户类别 0-普通 10-VIP 20-管理 30-超级管理
	Pwd      string `json:"pwd" db:"pwd"`             //	用户密码
	Session  string `json:"session" db:"session"`     //	用户当前有效 session
	Mobile   int64  `json:"mobile" db:"mobile"`       //	用户手机号码
}

// Card 一卡通表
type Card struct {
	CardID    int64  `json:"card_id" db:"card_id"`       // 学生一卡通编号
	RealName  string `json:"real_name" db:"real_name"`   // 用户真实姓名
	Identity  int64  `json:"identity" db:"identity"`     // 用户身份 0-未知 1-本科生 2-硕士研究生 3-博士研究生 4-教师 5-临时卡
	StuNo     string `json:"stu_no" db:"stu_no"`         // 学生学号
	Class     string `json:"class" db:"class"`           // 班级
	DeptNo    string `json:"dept_no" db:"dept_no"`       // 学院编号
	DeptName  string `json:"dept_name" db:"dept_name"`   // 学院名称
	MajorName string `json:"major_name" db:"major_name"` // 专业名称
	Grade     int64  `json:"grade" db:"grade"`           // 年级
	PwdCard   string `json:"pwd_card" db:"pwd_card"`     // 一卡通密码
	PwdMoney  string `json:"pwd_money" db:"pwd_money"`   // 消费密码
	Gender    int64  `json:"gender" db:"gender"`         // 用户性别 0-未知 1-男 2-女
}
