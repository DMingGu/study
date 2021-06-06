package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"testing"
)

func TestUser_Find(t *testing.T) {
	db, err := sql.Open("mysql",
		"root:dongming.sihui@tcp(127.0.0.1:3306)/jike_study")
	if err != nil {
		fmt.Println("err==>",err)
		return
	}
	defer db.Close()
	phone,err:=(&User{}).Find(db)
	fmt.Println("==>",phone,"==>",err,"==>",errors.Cause(err))
}