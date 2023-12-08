package db

import (
	"fmt"
	"testing"

	pb "github.com/emart.io/cross/zwan/service/go"
	"github.com/google/uuid"

	//"github.com/jmzwcn/db"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	attendantTable = "attendants"
)

func init() {
	Open("sqlite", "test.db")
}

func TestQuery(t *testing.T) {
	fmt.Println("begin")
	in := &pb.Attendant{Name: "nihao222"}
	in.Id = uuid.NewString()
	in.Created = timestamppb.Now() // timestamppb.Now()
	// if err := Insert(attendantTable, in); err != nil {
	// 	t.Error(err)
	// }

	// attendant := pb.Attendant{}
	// if err := GetById(attendantTable, in.Id, &attendant); err != nil {
	// 	//return nil, err
	// 	t.Error(err)
	// }
	// fmt.Println(attendant.Name)

	if err := Upsert(attendantTable, in.Id, in); err != nil {
		//return nil, err
		t.Error(err)
		fmt.Println(err)
	}

	attendant := pb.Attendant{}
	if err := GetById(attendantTable, in.Id, &attendant); err != nil {
		//return nil, err
		t.Error(err)
	}
	fmt.Println(attendant.Name)
	attendants := []*pb.Attendant{}
	if err := List(attendantTable, &attendants); err != nil {
		t.Error(err)
	}
	fmt.Println(attendants)
	t.Error("done")
}
