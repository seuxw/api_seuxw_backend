package user

import (
	"fmt"
	"seuxw/embrice/entity/user"
)

// CreateCardDB 校园卡用户创建 DB 操作
func (db *Database) CreateCardDB(card *user.Card) error {
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
		sd_card
	where
		card_id = ? and deleted = 0
	`

	insertSQL = `
	insert into sd_card (
		card_id, real_name, identity, stu_no,
		class, dept_no, dept_name, major_name,
		grade, pwd_card, pwd_money, gender
	) values (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	)
	`

	err = db.Get(&count, selectCheckSQL, card.CardID)
	if err != nil {
		err = fmt.Errorf("数据预查询错误 err:%s", err)
		goto END
	}

	// 插入操作
	if count == 0 {
		_ = db.MustExec(
			insertSQL, card.CardID, card.RealName, card.Identity, card.StuNo, card.Class,
			card.DeptNo, card.DeptName, card.MajorName, card.Grade, card.PwdCard,
			card.PwdMoney, card.Gender)
	} else {
		err = fmt.Errorf("卡用户 %d 已经存在！", card.CardID)
	}

END:
	return err
}
