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
	DB        *sql.DB
	marshaler = protojson.MarshalOptions{EmitUnpopulated: true}
)

func init() {
	var err error
	for i := 1; i <= 5; i++ {
		DB, err = sql.Open("sqlite3", "./foo.db")
		if err != nil {
			log.Errorf("Connect failed: %s, %d", err, i)
			time.Sleep(time.Second * 5)
			continue
		} else {
			break
		}
	}
	log.Debugln("DB init completely: ")
}

func checkTable(table string) (sql.Result, error) {
	return DB.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (data JSON)", table))
}

func Upsert(table, uid string, obj proto.Message) error {
	checkTable(table)
	jsonv, err := marshaler.Marshal(obj)
	if err != nil {
		return err
	}

	var result = ""
	row := DB.QueryRow(fmt.Sprintf("SELECT data FROM %s WHERE data->>'$.id'='%s'", table, uid))
	if err := row.Scan(&result); err == sql.ErrNoRows {
		_, err := DB.Exec(fmt.Sprintf("INSERT INTO %s (data) VALUES (?)", table), jsonv)
		return err
	}

	_, err = DB.Exec(fmt.Sprintf("UPDATE %s SET data=? WHERE data->>'$.id'='%s'", table, uid), jsonv)
	return err
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

func Delete(table string, id interface{}) (sql.Result, error) {
	return DB.Exec("DELETE FROM "+table+" WHERE data->>'$.id' = ?", id)
}
