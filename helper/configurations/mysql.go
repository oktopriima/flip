/*
* project flip
* created by oktopriima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 23:25
**/

package configurations

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type MysqlClient struct {
	Address  string `json:"address"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

func NewMysqlClient(env map[string]interface{}) (*gorm.DB, error) {
	jsonb, err := json.Marshal(env)
	if err != nil {
		return nil, err
	}

	client := new(MysqlClient)
	if err = json.Unmarshal(jsonb, &client); err != nil {
		return nil, err
	}

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		client.User,
		client.Pass,
		client.Address,
		client.Port,
		client.Database,
	)
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		return nil, err
	}

	db.SingularTable(true)

	return db, nil
}
