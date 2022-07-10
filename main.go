//我们在数据库操作的时候，比如 `dao` 层中当遇到一个 `sql.ErrNoRows` 的时候，是否应该 `Wrap` 这个 `error`，抛给上层。为什么？应该怎么做请写出代码

//应该 Wrap error 抛给上层 具体错误判断及处理应当由调用方处理
package main

import (
	"database/sql"
	"github.com/pkg/errors"
)

type Dao struct{}

func New() *Dao {
	return &Dao{}
}

func (d *Dao) FindByID(id int) (user *model.User, err error) {

	err = Db.Table("t_user").Where("id = ?", userID).Find(user).Error
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.Wrapf(ierror.ErrNotFound, "query:%s", query)
	}
	return nil, errors.Wrapf(ierror.ErrDatabase, "query: %s error(%v)", query, err)
}
