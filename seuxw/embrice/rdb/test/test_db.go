package test

import (
	"fmt"
	"seuxw/embrice/entity/test"
)

func (self *Database) Test(date string) (test.Test, error) {
	var (
		selectSQL      string
		selectCheckSQL string
		count          int
		Test           test.Test
		err            error
	)

	selectCheckSQL = `
	select
		count(1) as count
	from
		sunrisetime
	where
		date = ?
	`

	selectSQL = `
	select
		date, sun_rise_time, sun_down_time
	from
		sunrisetime
	where
		date = ?
	limit 1
	`

	if err = self.Get(&count, selectCheckSQL, date); err != nil {
		err = fmt.Errorf("数据库预查询错误 err:%s", err)
		goto END
	}

	if count <= 0 {
		err = fmt.Errorf("查询日期错误")
		goto END
	}

	if err = self.Get(&Test, selectSQL, date); err != nil {
		err = fmt.Errorf("数据库查询错误 err:%s", err)
		goto END
	}
	Test.Date = date

END:
	return Test, err
}
