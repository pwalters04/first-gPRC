package service

import (
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const API_VERSION = "v1"

//toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
type toDoServiceServer struct {
	db *sql.DB
}

func NewtoDoServer(db *sql.DB) *toDoServiceServer {
	 return &toDoServiceServer{db: db}
}

func (s *toDoServiceServer) checkAPI(apiVersion string) error {
	if apiVersion == ""{
		return fmt.Errorf()
	}
	if len(api) > 0 {
		if API_VERSION != apiVersion {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// connect returns SQL database connection from the pool
func (s *toDoServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
	}
	return c, nil
}

