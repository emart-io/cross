package service

import (
	"context"

	pb "github.com/emart.io/cross/zwan/service/go"
	"github.com/jmzwcn/db"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	AttendantTable = "Attendants"
)

type AttendantsImpl struct {
	pb.UnimplementedAttendantsServer
}

func (s *AttendantsImpl) Add(ctx context.Context, in *pb.Attendant) (*pb.Attendant, error) {
	in.Id = in.Telephone
	in.Created = timestamppb.Now() // timestamppb.Now()
	if err := db.Insert(AttendantTable, in); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *AttendantsImpl) Get(ctx context.Context, in *pb.Attendant) (*pb.Attendant, error) {
	Attendant := pb.Attendant{}
	if err := db.GetById(AttendantTable, in.Id, &Attendant); err != nil {
		return nil, err
	}
	return &Attendant, nil
}

func (s *AttendantsImpl) Update(ctx context.Context, in *pb.Attendant) (*pb.Attendant, error) {
	Attendant, err := s.Get(ctx, in)
	if err != nil {
		return nil, err
	}
	if in.Icon != "" {
		Attendant.Icon = in.Icon
	}
	if in.Signature != "" {
		Attendant.Signature = in.Signature
	}
	if in.Shops != nil {
		Attendant.Shops = in.Shops
	}
	if err := db.Update(AttendantTable, in.Id, Attendant); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *AttendantsImpl) List(in *pb.Attendant, stream pb.Attendants_ListServer) error {
	Attendants := []*pb.Attendant{}
	if err := db.List(AttendantTable, &Attendants, " order by data->'$.created' desc"); err != nil {
		return err
	}

	for _, v := range Attendants {
		if err := stream.Send(v); err != nil {
			return err
		}
	}

	return nil
}

func (s *AttendantsImpl) Delete(ctx context.Context, in *pb.Attendant) (*emptypb.Empty, error) {
	if err := db.Delete(AttendantTable, in.Id); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *AttendantsImpl) Login(ctx context.Context, in *pb.Attendant) (*pb.Attendant, error) {
	Attendant := pb.Attendant{}
	err := db.Get(AttendantTable, map[string]interface{}{"$.telephone": in.Telephone, "$.password": in.Password}, &Attendant)
	if err != nil {
		return nil, err
	}
	return &Attendant, nil
}

func (s *AttendantsImpl) Certificate(ctx context.Context, in *pb.Attendant) (*pb.Attendant, error) {
	Attendant, err := s.Get(ctx, in)
	if err != nil {
		return nil, err
	}
	Attendant.Cert = in.Cert
	if err := db.Update(AttendantTable, in.Id, Attendant); err != nil {
		return nil, err
	}

	return in, nil
}
