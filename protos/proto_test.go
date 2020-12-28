package protos_test

import (
	"encoding/json"
	"testing"

	"github.com/golang/protobuf/proto"

	_ "github.com/grpc_study/protos/api"
	"github.com/grpc_study/protos/model"
)

func TestStudent1(t *testing.T) {
	s1 := &model.Student{
		Id:   1234567890,
		Name: "五六七",
		Age:  18,
	}
	sBody, err := proto.Marshal(s1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(sBody))
	revS1 := &model.Student{}
	if err := proto.Unmarshal(sBody, revS1); err != nil {
		t.Fatal(err)
	}
	t.Log(revS1)
}

func TestStudent2(t *testing.T) {
	s1 := &model.Student{
		Id:   1,
		Name: "孙悟空",
		Age:  300,
	}
	sBody, err := proto.Marshal(s1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%b", sBody)
	t.Logf("%d", sBody)
	t.Log(string(sBody))
	jBody, _ := json.Marshal(s1)
	t.Log(string(jBody))
	revS1 := &model.Student{}
	if err := proto.Unmarshal(sBody, revS1); err != nil {
		t.Fatal(err)
	}
	t.Log(revS1)
}

func TestStudentList1(t *testing.T) {
	s1 := &model.StudentList{
		Class: "三年级二班",
		Students: []*model.Student{
			&model.Student{Id: 123465, Name: "路飞", Age: 19},
			&model.Student{Id: 321, Name: "索龙", Age: 20},
			&model.Student{Id: 789, Name: "乔巴", Age: 6},
		},
		Teacher: "雷利",
		Score:   []int64{1, 2, 3, 4, 5, 6},
	}
	sBody, err := proto.Marshal(s1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(sBody))
	jBody, _ := json.Marshal(s1)
	t.Log(string(jBody))
	revS1 := &model.StudentList{}
	if err := proto.Unmarshal(sBody, revS1); err != nil {
		t.Fatal(err)
	}
	t.Log(revS1)
}
