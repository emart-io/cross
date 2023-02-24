package service

import (
	"context"

	db "github.com/emart.io/cross/zwan/internal/impl/biz"
	pb "github.com/emart.io/cross/zwan/service/go"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	attendantTable = "attendants"
)

type AttendantsImpl struct {
	pb.UnimplementedAttendantsServer
}

func (s *AttendantsImpl) Add(ctx context.Context, in *pb.Attendant) (*pb.Attendant, error) {
	in.Id = in.Telephone
	in.Created = timestamppb.Now() // timestamppb.Now()
	if _, err := db.Upsert(attendantTable, in.Id, in); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *AttendantsImpl) Get(ctx context.Context, in *pb.Attendant) (*pb.Attendant, error) {
	attendant := pb.Attendant{}
	if err := db.GetById(attendantTable, in.Id, &attendant); err != nil {
		return nil, err
	}
	return &attendant, nil
}

func (s *AttendantsImpl) Update(ctx context.Context, in *pb.Attendant) (*pb.Attendant, error) {
	attendant, err := s.Get(ctx, in)
	if err != nil {
		return nil, err
	}
	if in.Icon != "" {
		attendant.Icon = in.Icon
	}
	if in.Signature != "" {
		attendant.Signature = in.Signature
	}
	if in.Shops != nil {
		attendant.Shops = in.Shops
	}
	if _, err := db.Upsert(attendantTable, in.Id, attendant); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *AttendantsImpl) List(in *pb.Attendant, stream pb.Attendants_ListServer) error {
	attendants := []*pb.Attendant{}
	if err := db.List(attendantTable, &attendants, " order by data->'$.created' desc"); err != nil {
		return err
	}

	for _, v := range attendants {
		if err := stream.Send(v); err != nil {
			return err
		}
	}

	return nil
}

func (s *AttendantsImpl) Delete(ctx context.Context, in *pb.Attendant) (*emptypb.Empty, error) {
	if _, err := db.Delete(attendantTable, in.Id); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *AttendantsImpl) Login(ctx context.Context, in *pb.Attendant) (*pb.Attendant, error) {
	attendant := pb.Attendant{}
	err := db.Get(attendantTable, &attendant, "WHERE data->>'$.telephone'='%s' and data->>'$.password'='%s'", in.Telephone, in.Password)
	if err != nil {
		return nil, err
	}
	return &attendant, nil
}

func (s *AttendantsImpl) Certificate(ctx context.Context, in *pb.Attendant) (*pb.Attendant, error) {
	attendant, err := s.Get(ctx, in)
	if err != nil {
		return nil, err
	}
	attendant.Cert = in.Cert
	if _, err := db.Upsert(attendantTable, in.Id, attendant); err != nil {
		return nil, err
	}

	return in, nil
}
