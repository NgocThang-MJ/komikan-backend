package gapi

import (
	"log"

	db "github.com/liquiddev99/komikan-backend/db/sqlc"
	"github.com/liquiddev99/komikan-backend/pb"
	"github.com/liquiddev99/komikan-backend/token"
	"github.com/liquiddev99/komikan-backend/util"
)

type Server struct {
	pb.UnimplementedKomikanServer
	db     *db.Queries
	config util.Config
	token  token.Token
}

func NewServer(config util.Config, db *db.Queries) (*Server, error) {
	paseto, err := token.NewPaseto(config.SymmetricKey)
	if err != nil {
		log.Fatal("Failed to create token instance")
	}

	server := &Server{config: config, db: db, token: paseto}
	return server, nil
}
