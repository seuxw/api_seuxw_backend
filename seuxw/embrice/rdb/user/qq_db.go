package user

import (
	"fmt"
	"seuxw/embrice/entity/user"
)

// CreateQQDB QQ 用户创建 DB 操作
func (db *Database) CreateQQDB(qq *user.QQ) error {
	var (
		insertSQL      string
		selectCheckSQL string
		err            error
		count          int
	)

	selectCheckSQL = `
	select
		count(1) as count
	from
		sd_qq
	where
		qq_id = ? and deleted = 0
	`

	insertSQL = `
	insert into sd_qq (
		qq_id, address, birthday, gender, hometown,
		nick_name, rmk_name, vip, vip_level
	) values (
		?, ?, ?, ?, ?, ?, ?, ?, ?
	)
	`

	err = db.Get(&count, selectCheckSQL, qq.QQID)
	if err != nil {
		err = fmt.Errorf("数据预查询错误 err:%s", err)
		goto END
	}

	// 插入操作
	if count == 0 {
		_ = db.MustExec(
			insertSQL, qq.QQID, qq.Address, qq.Birthday, qq.Gender, qq.Hometown,
			qq.NickName, qq.RmkName, qq.VIP, qq.VIPLevel)
	} else {
		err = fmt.Errorf("QQ 用户 %d 已经存在！", qq.QQID)
	}

END:
	return err
}
