package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type Metadata struct {
	ClientIp string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if clientIps := md.Get("x-forwarded-for"); len(clientIps) > 0 {
			mtdt.ClientIp = clientIps[0]
			return mtdt
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		mtdt.ClientIp = p.Addr.String()
	}

	return mtdt
}
