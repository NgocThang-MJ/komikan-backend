package gapi

import (
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

	db "github.com/liquiddev99/komikan-backend/db/sqlc"
	"github.com/liquiddev99/komikan-backend/pb"
)

func convertBytesToString(bytes [16]byte) string {
	return fmt.Sprintf(
		"%x-%x-%x-%x-%x",
		bytes[0:4],
		bytes[4:6],
		bytes[6:8],
		bytes[8:10],
		bytes[10:16],
	)
}

func convertToUserRes(user db.User) *pb.UserResponse {
	var userId string
	if user.ID.Valid {
		bytes := user.ID.Bytes
		userId = fmt.Sprintf(
			"%x-%x-%x-%x-%x",
			bytes[0:4],
			bytes[4:6],
			bytes[6:8],
			bytes[8:10],
			bytes[10:16],
		)
	}
	return &pb.UserResponse{
		UserId:    userId,
		FullName:  user.FullName,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}

func convertToHistories(histories []db.History) *pb.HistoriesResponse {
	pbHistories := make([]*pb.History, len(histories))

	for i, item := range histories {
		var userId string
		if item.UserID.Valid {
			bytes := item.UserID.Bytes
			userId = fmt.Sprintf(
				"%x-%x-%x-%x-%x",
				bytes[0:4],
				bytes[4:6],
				bytes[6:8],
				bytes[8:10],
				bytes[10:16],
			)
		}
		pbHistory := &pb.History{
			UserAgent:      item.UserAgent,
			ClientIp:       item.ClientIp,
			UserId:         userId,
			MangadexId:     item.MangadexID,
			AlId:           item.AlID,
			Path:           item.Path,
			CoverImage:     item.CoverImage,
			Title:          item.Title,
			ReadingChapter: item.ReadingChapter,
			LastReadAt:     timestamppb.New(item.LastReadAt),
			CreatedAt:      timestamppb.New(item.CreatedAt),
		}

		pbHistories[i] = pbHistory
	}

	return &pb.HistoriesResponse{Histories: pbHistories}
}
