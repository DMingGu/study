package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

type User struct {

}
func (u*User)Find(db *sql.DB)(string,error)  {
	var phone string
	err:=db.QueryRow("select phone from g_user where id=?",1).Scan(&phone)
	if err==sql.ErrNoRows {
		return "", nil
	}
	return phone,errors.Wrap(err,"g_user find error")
}