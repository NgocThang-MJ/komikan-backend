package gapi

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	db "github.com/liquiddev99/komikan-backend/db/sqlc"
	"github.com/liquiddev99/komikan-backend/pb"
	"github.com/liquiddev99/komikan-backend/util"
	"github.com/liquiddev99/komikan-backend/validation"
)

func (server *Server) CreateUser(
	ctx context.Context,
	req *pb.CreateUserRequest,
) (*pb.CreateUserResponse, error) {
	err := validation.ValidateCreateUserRequest(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	hashed_password, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to hash password, %s", err)
	}

	args := db.CreateUserParams{
		FullName:       req.FullName,
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: hashed_password,
	}

	user, err := server.db.CreateUser(ctx, args)
	if err != nil {
		var pgErr *pgconn.PgError
		// Error unique username and email
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				errMsg := pgErr.Message
				if pgErr.ConstraintName == "users_username_key" {
					errMsg = "Username already exists"
				}
				if pgErr.ConstraintName == "users_email_key" {
					errMsg = "Email already exists"
				}
				return nil, status.Error(
					codes.InvalidArgument,
					errMsg,
				)
			}
		}
		return nil, status.Errorf(codes.Internal, "Failed to create user, %s", err)
	}

	arg := db.UpdateHistoriesToUserParams{
		UserAgent: req.GetUserAgent(),
		UserID:    user.ID,
	}

	server.db.UpdateHistoriesToUser(ctx, arg)

	token, err := server.token.CreateToken(user.ID, server.config.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create token, %s", err)
	}

	res := &pb.CreateUserResponse{
		User:        convertToUserRes(user),
		AccessToken: token,
	}

	return res, nil
}

func (server *Server) LoginUser(
	ctx context.Context,
	req *pb.LoginUserRequest,
) (*pb.LoginUserResponse, error) {
	user, err := server.db.GetUserByEmailOrUsername(ctx, req.GetCredential())
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, status.Error(codes.InvalidArgument, "Invalid credentials")
		}
		return nil, status.Errorf(codes.Internal, "Failed to get user, %s", err)
	}

	if err = util.CheckPassword(req.GetPassword(), user.HashedPassword); err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid credentials")
	}

	token, err := server.token.CreateToken(user.ID, server.config.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create token, %s", err)
	}

	res := &pb.LoginUserResponse{
		User:        convertToUserRes(user),
		AccessToken: token,
	}

	return res, nil
}
