package api

import (
	"context"
	ctl "local/fin/controllers"
	"local/fin/forms"
	pb "local/fin/protos"
	"local/fin/utils"
	"log"
)

type UserInfoServer struct {
	pb.UserInfoServer
}

func (s UserInfoServer) GetUser(c context.Context, req *pb.GetUserRequest) (res *pb.GetUserResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered: %v", r)
			utils.ErrLogging()
		}
	}()
	user, err := ctl.Get(req.Id)
	if err != nil {
		return res, err
	}
	obj := &pb.UserModel{
		Id:       user.Id,
		Email:    user.Email,
		UserName: user.UserName,
		Password: user.Password,
		Age:      user.Age,
		Phone:    user.Phone,
		Address:  user.Address,
	}
	return &pb.GetUserResponse{
		User: obj,
	}, nil
}

func (s UserInfoServer) ListUser(c context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("api.grpc.ListUser: %v", r)
			utils.ErrLogging()
		}
	}()
	users, err := ctl.List()
	if err != nil {
		return nil, err
	}
	var res []*pb.UserModel
	for _, user := range users {
		res = append(res, &pb.UserModel{
			Id:       user.Id,
			Email:    user.Email,
			UserName: user.UserName,
			Password: user.Password,
			Age:      user.Age,
			Phone:    user.Phone,
			Address:  user.Address,
		})
	}
	return &pb.ListUserResponse{
		Users: res,
	}, nil
}

func (s UserInfoServer) CreateUser(c context.Context, req *pb.CreateUserRequest) (res *pb.CreateUserResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("api.grpc.CreateUser: %v", r)
			utils.ErrLogging()
		}
	}()
	user := req.User
	obj := forms.UserForm{
		Email:    user.Email,
		UserName: user.UserName,
		Password: user.Password,
		Age:      user.Age,
		Phone:    user.Phone,
		Address:  user.Address,
	}
	if err := ctl.Create(&obj); err != nil {
		return res, err
	}
	log.Print("Create successed:", obj)
	return &pb.CreateUserResponse{
		Response: "create successed",
	}, nil
}

// func (s UserInfoServer) UpdateUser(c context.Context, req *pb.UpdateUserRequest) (res *pb.UpdateUserResponse, err error) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			log.Printf("api.grpc.UpdateUser: %v", r)
// 			utils.ErrLogging()
// 		}
// 	}()
// 	user := req.User
// 	obj := models.UserModel{
// 		Email:    user.Email,
// 		UserName: user.UserName,
// 		Password: user.Password,
// 		Age:      user.Age,
// 		Phone:    user.Phone,
// 		Address:  user.Address,
// 	}
// 	if err := ctl.Update(req.Id, &obj); err != nil {
// 		return res, err
// 	}
// 	return &pb.UpdateUserResponse{
// 		Response: "update successed",
// 	}, nil
// }

func (s UserInfoServer) DeleteUser(c context.Context, req *pb.DeleteUserRequest) (res *pb.DeleteUserResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("api.grpc.DeleteUser: %v", r)
			utils.ErrLogging()
		}
	}()
	if err := ctl.Delete(req.Id); err != nil {
		return res, err
	}
	return &pb.DeleteUserResponse{
		Response: "delete successsed",
	}, nil
}
