package db

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var (
	DB *sql.DB
	//checked   = make(map[string]bool)
	marshaler = protojson.MarshalOptions{EmitUnpopulated: true}
)

// func InitDB(db *sql.DB, tables ...string) {
// 	DB = db
// 	for _, v := range tables {
// 		checkTable(v)
// 	}
// }

func init() {
	var err error
	for i := 1; i <= 5; i++ {
		DB, err = sql.Open("sqlite3", "./foo.db") //sql.Open("mysql", dsn)
		if err != nil {
			log.Errorf("Connect failed: %s, %d", err, i)
			time.Sleep(time.Second * 5)
			continue
		} else {
			break
		}
	}
	log.Debugln("DB init completely: ")
	DB.SetMaxIdleConns(0)
}

func checkTable(table string) (sql.Result, error) {
	// if checked[table] {
	// 	return nil, nil
	// }
	// checked[table] = true
	return DB.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (data JSON)", table))
}

func Insert(table string, obj proto.Message) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	if _, err := InsertTx(tx, table, obj); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func InsertTx(tx *sql.Tx, table string, obj proto.Message) (sql.Result, error) {
	checkTable(table)
	jsonv, err := marshaler.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return tx.Exec("INSERT INTO "+table+"(data) VALUES (?)", jsonv)
}

func InsertIfNotExist(table string, id interface{}, obj proto.Message) error {
	if GetById(table, id, proto.Clone(obj)) == sql.ErrNoRows {
		return Insert(table, obj)
	}
	return nil
}

// Update||Insert
func Upsert(table string, id interface{}, obj proto.Message) error {
	if err := GetById(table, id, proto.Clone(obj)); err == sql.ErrNoRows {
		return Insert(table, obj)
	}
	return Update(table, id, obj)
}

func GetById(table string, id interface{}, obj proto.Message) error {
	return Get(table, obj, fmt.Sprintf("WHERE data->>'$.id'='%s'", id))
}

func Get(table string, obj proto.Message, clause ...string) error {
	checkTable(table)
	query := fmt.Sprintf("SELECT data FROM %s %s", table, strings.Join(clause, " "))
	var data []byte
	if err := DB.QueryRow(query).Scan(&data); err != nil {
		return err
	}

	return protojson.Unmarshal(data, obj)
}

func List(table string, result interface{}, clause ...string) error {
	checkTable(table)
	resultv := reflect.ValueOf(result)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("result argument must be a slice address")
	}
	slicev := resultv.Elem()
	query := fmt.Sprintf("SELECT data FROM  %s %s", table, strings.Join(clause, " "))

	log.Debugln(query)
	rows, err := DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		elemp := reflect.New(slicev.Type().Elem().Elem())
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return err
		}
		protojson.Unmarshal(data, elemp.Interface().(proto.Message))
		slicev = reflect.Append(slicev, elemp)
	}
	reflect.ValueOf(result).Elem().Set(slicev)
	return nil
}

func Update(table string, id interface{}, newObj proto.Message) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	if _, err := UpdateTx(tx, table, id, newObj); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func UpdateTx(tx *sql.Tx, table string, id interface{}, newObj proto.Message) (sql.Result, error) {
	jsonv, err := marshaler.Marshal(newObj)
	if err != nil {
		return nil, err
	}
	return tx.Exec("UPDATE "+table+" SET data=? WHERE data->>'$.id'=?", jsonv, id)
}

func UpdateKVS(table string, id interface{}, kvs map[string]interface{}) (sql.Result, error) {
	var (
		keys   []string
		values []interface{}
	)
	for k, v := range kvs { // key should be [json-path], e.g:$.id
		keys = append(keys, ",'"+k+"',?")
		values = append(values, v)
	}
	sql := "UPDATE " + table + " SET data=" + "JSON_SET(data" + strings.Join(keys, "") + ") WHERE data->>'$.id'= ?"
	return DB.Exec(sql, append(values, id)...)
}

func Delete(table string, id interface{}) error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	if _, err := DeleteTx(tx, table, id); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func DeleteTx(tx *sql.Tx, table string, id interface{}) (sql.Result, error) {
	return tx.Exec("DELETE FROM "+table+" WHERE data->>'$.id' = ?", id)
}
