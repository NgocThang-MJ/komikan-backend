package gapi

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	db "github.com/liquiddev99/komikan-backend/db/sqlc"
	"github.com/liquiddev99/komikan-backend/pb"
)

func (server *Server) CreateHistory(
	ctx context.Context,
	req *pb.CreateHistoryRequest,
) (*pb.History, error) {
	payload, _ := server.authorizeUser(ctx)

	mtdt := server.extractMetadata(ctx)

	args := db.CreateHistoryParams{
		UserAgent:      req.UserAgent,
		ClientIp:       mtdt.ClientIp,
		UserID:         payload.UserId,
		MangadexID:     req.MangadexId,
		AlID:           req.AlId,
		Path:           req.Path,
		CoverImage:     req.CoverImage,
		Title:          req.Title,
		ReadingChapter: req.ReadingChapter,
	}

	history, err := server.db.CreateHistory(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create history, %s", err)
	}

	userId := convertBytesToString(payload.UserId.Bytes)

	res := &pb.History{
		UserAgent:      history.UserAgent,
		ClientIp:       history.ClientIp,
		MangadexId:     history.MangadexID,
		AlId:           history.AlID,
		Path:           history.Path,
		UserId:         userId,
		CoverImage:     history.CoverImage,
		Title:          req.Title,
		ReadingChapter: history.ReadingChapter,
		LastReadAt:     timestamppb.New(history.LastReadAt),
		CreatedAt:      timestamppb.New(history.CreatedAt),
	}
	return res, nil
}

// Get Histories by UserId or UserAgent
func (server *Server) GetHistories(
	ctx context.Context,
	req *pb.GetHistoriesRequest,
) (*pb.HistoriesResponse, error) {
	payload, _ := server.authorizeUser(ctx)
	var pgUUID pgtype.UUID

	arg := db.ListHistoriesParams{
		UserAgent: req.GetUserAgent(),
		UserID:    pgUUID,
		Limit:     req.GetLimit(),
		Offset:    req.GetOffset(),
	}

	if payload != nil {
		arg.UserID = payload.UserId
	}

	histories, err := server.db.ListHistories(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get histories, %s", err)
	}

	res := convertToHistories(histories)
	return res, nil
}

func (server *Server) UpdateHistory(
	ctx context.Context,
	req *pb.UpdateHistoryRequest,
) (*pb.UpdateHistoryResponse, error) {
	args := db.UpdateHistoryParams{
		MangadexID:     req.GetMangadexId(),
		UserAgent:      "",
		ReadingChapter: req.GetReadingChapter(),
		LastReadAt:     time.Now(),
	}

	_, err := server.db.UpdateHistory(ctx, args)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update history, %s", err)
	}

	res := &pb.UpdateHistoryResponse{
		Ok:      true,
		Message: "Update history successfully",
	}

	return res, nil
}

func (server *Server) UpsertHistory(
	ctx context.Context,
	req *pb.CreateHistoryRequest,
) (*pb.CreateHistoryResponse, error) {
	payload, _ := server.authorizeUser(ctx)
	mtdt := server.extractMetadata(ctx)

	var userId pgtype.UUID
	args := db.UpsertHistoryParams{
		UserAgent:      req.UserAgent,
		ClientIp:       mtdt.ClientIp,
		UserID:         userId,
		MangadexID:     req.MangadexId,
		AlId:           req.AlId,
		Path:           req.Path,
		CoverImage:     req.CoverImage,
		Title:          req.Title,
		ReadingChapter: req.ReadingChapter,
		LastReadAt:     time.Now(),
	}

	/*
	   We have 2 cases: logged in (has user_id) and not logged in (doesn't has user_id)
	   If has user_id, the Upsert will check the user_id conflict (the user_id param must come before user_agent)
	   and vice versa

	   The query will failed if there's already a user agent that not has user_id (guest)
	   and try to upsert with the same user agent and new user_id (user logout -> read manga -> login -> read that manga)
	   so we handle that case with UpdateHistoriesToUser.
	   If a guest already has some histories and signup, we should update all the histories with the new user_id
	*/
	if payload != nil {
		// Upsert prevent UserId constraint
		args.UserID = payload.UserId
		err := server.db.UpsertHistoryUserId(ctx, args)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				if pgErr.Code == "23505" {
					arg := db.UpdateHistoriesToUserParams{
						UserAgent: req.GetUserAgent(),
						UserID:    payload.UserId,
					}
					server.db.UpdateHistoriesToUser(ctx, arg)
				}
			}
			return nil, status.Errorf(codes.Internal, "Failed to upsert history, %s", err)
		}
	} else {
		// Upsert prevent UserAgent constraint
		var err error
		err = server.db.UpsertHistoryAgent(ctx, args)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to upsert history, %s", err)
		}
	}

	res := &pb.CreateHistoryResponse{
		Ok:      true,
		Message: "Upsert successfully",
	}
	return res, nil
}
