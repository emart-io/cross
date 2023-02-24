package db

import (
	"fmt"
	"testing"

	pb "github.com/emart.io/cross/zwan/service/go"
)

func TestXxx(t *testing.T) {
	fmt.Println("db init")
	if _, err := checkTable("test"); err != nil {
		t.Error(err)
	}
	id := "abcde"
	address := &pb.Address{Id: id}
	if err := Upsert("test", id, address); err != nil {
		t.Error()
	}
	add2 := &pb.Address{}
	if err := GetById("test", id, add2); err != nil {
		t.Error(err)
	}
	fmt.Println(add2.Id)
}
