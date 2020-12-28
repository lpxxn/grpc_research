package common

import (
	"context"
	"fmt"
	"io"
	"math/rand"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/lpxxn/grpc_research/protos"
	"github.com/lpxxn/grpc_research/protos/api"
	"github.com/lpxxn/grpc_research/protos/model"
)

type StudentSrv struct{ StudentList []*model.Student }

func (srv *StudentSrv) NewStudent(ctx context.Context, s *model.Student) (*protos.Result, error) {
	if s != nil {
		srv.StudentList = append(srv.StudentList, s)
	}
	return &protos.Result{
		Code: "OK",
		Desc: "hi!",
	}, nil
}

func (srv *StudentSrv) StudentByID(context.Context, *api.QueryStudent) (*api.QueryStudentResponse, error) {
	l := len(srv.StudentList)
	rev := &api.QueryStudentResponse{StudentList: srv.StudentList}
	srv.StudentList = srv.StudentList[l:]
	return rev, nil
}

func (srv *StudentSrv) AllStudent(e *empty.Empty, rev api.StudentSrv_AllStudentServer) error {
	const limit = 10
	data := &api.QueryStudentResponse{}
	curr := srv.StudentList
	for _, item := range curr {
		data.StudentList = append(data.StudentList, item)
		if len(data.StudentList) == limit {
			if err := rev.Send(data); err != nil {
				fmt.Printf("send error %#v", err)
			}
			data.StudentList = data.StudentList[:0]
		}
	}
	if len(data.StudentList) > 0 {
		if err := rev.Send(data); err != nil {
			fmt.Printf("send error %#v", err)
		}
	}
	return nil
}
func (srv *StudentSrv) StudentInfo(stream api.StudentSrv_StudentInfoServer) error {
	l := len(srv.StudentList)
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		fmt.Printf("get id: %d", in.Id)
		if l == 0 {
			fmt.Println("data is empty")
			return nil
		}
		if l > 0 {
			stream.Send(&api.QueryStudentResponse{StudentList: srv.StudentList[0:rand.Intn(l)]})
		}
	}
	return nil
}
