package grpc

import (
	"context"
	"fmt"
	"smart-dokan/usersvc/internal/application/core/domain"

	pb "github.com/asadlive84/smart-dokan-pb/golang/user"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
)

func (a *Adapter) CreateUser(c context.Context, userReq *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	user, err := a.api.CreateUser(&domain.User{
		FirstName:   userReq.FirstName,
		LastName:    userReq.LastName,
		Email:       userReq.Email,
		PhoneNumber: userReq.PhoneNumber,
		Password:    userReq.Password,
	})

	fmt.Printf("Error: %+v\n", err)

	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		UserId: fmt.Sprintf("%s", user.ID),
	}, nil
}

func (a *Adapter) GetUser(ctx context.Context, userReq *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	userUUID, err := uuid.Parse(userReq.UserId)
	if err != nil {
		fmt.Println("invalid uuid:", err)
		return nil, err
	}
	user, err := a.api.GetUser(&domain.User{
		ID: userUUID,
	})

	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		UserId:      fmt.Sprintf("%s", user.ID),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   &timestamp.Timestamp{},
		UpdatedAt:   &timestamp.Timestamp{},
		DeletedAt:   &timestamp.Timestamp{},
		LastLogin:   &timestamp.Timestamp{},
		LastLogout:  &timestamp.Timestamp{},
	}, nil
}
