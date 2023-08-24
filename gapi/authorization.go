package gapi

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"

	"github.com/liquiddev99/komikan-backend/token"
)

func (server *Server) authorizeUser(ctx context.Context) (*token.TokenPayload, error) {
	mtdt, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("Missing metadata")
	}

	values := mtdt.Get("authorization")
	if len(values) == 0 {
		return nil, fmt.Errorf("Missing authorization header")
	}

	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, fmt.Errorf("Invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if authType != "bearer" {
		return nil, fmt.Errorf("Unsupported authorization type: %s", authType)
	}

	accessToken := fields[1]
	payload, err := server.token.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("Invalid access token: %s", err)
	}

	return payload, nil
}
