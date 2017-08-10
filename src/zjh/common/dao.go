package common

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Dao struct {
	CurrentDBHandle *sql.DB
	TotalDB         int
	MysqlHandle     []*sql.DB
	TableName       string
}

type MysqlConf struct {
	Ip       string "ip"
	Port     string "port"
	Database string "database"
	User     string "user"
	Password string "password"
}

// 获取当前key对应的服务器配置
func (this *Dao) setDBPoint(key int) {
	key = key % this.TotalDB
	this.CurrentDBHandle = this.MysqlHandle[key]
}

// 初始化数据库
func (this *Dao) InitConf() {
	ConfString, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var mysqlConf []MysqlConf
	err = json.Unmarshal(ConfString, &mysqlConf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, v := range mysqlConf {
		dbStr := v.User + ":" + v.Password + "@tcp(" + v.Ip + ":" + v.Port + ")/" + v.Database + "?charset=utf8"
		fmt.Println(v.Database, dbStr)

		dbHandle, err := sql.Open("mysql", dbStr)
		if err == nil {
			err = dbHandle.Ping()
			if err != nil {
				fmt.Println(err)
			}
			this.MysqlHandle = append(this.MysqlHandle, dbHandle)
		} else {
			fmt.Println(err)
		}
	}

	this.TotalDB = len(this.MysqlHandle)
}

func (this *Dao) GetRecord(table string, where string, key int) (res map[string]string) {
	sql := "select * from " + table + " where " + where + " limit 1"
	this.setDBPoint(key)
	// TODO 此处多协程里会有坑
	rows, err := this.CurrentDBHandle.Query(sql)
	if err != nil {
		fmt.Println("query", err)
		return nil
	}

	columns, _ := rows.Columns()

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}

		return record
	}

	return nil
}

func (this *Dao) AddRecord(table string, fields []string, values []string, key int) bool {

	var fieldString, valueString string

	fieldString = strings.Join(fields, ",")
	valueString = strings.Join(values, "','")

	sql := fmt.Sprintf("insert into "+table+"(%s) values('%s')", fieldString, valueString)
	this.setDBPoint(key)
	// TODO 此处多协程里会有坑
	_, err := this.CurrentDBHandle.Exec(sql)

	if err != nil {
		fmt.Println("query", err)
		return false
	}

	return true
}
