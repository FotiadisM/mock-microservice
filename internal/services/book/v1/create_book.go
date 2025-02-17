package bookv1

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/google/uuid"

	bookv1 "github.com/FotiadisM/service-template/api/gen/go/book/v1"
	"github.com/FotiadisM/service-template/internal/services/book/v1/encoder"
	"github.com/FotiadisM/service-template/internal/services/book/v1/queries"
)

func (s *Service) CreateBook(ctx context.Context, req *connect.Request[bookv1.CreateBookRequest]) (*connect.Response[bookv1.CreateBookResponse], error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("failed to create uuid %w", err)
	}

	authorID, err := uuid.Parse(req.Msg.AuthorId)
	if err != nil {
		return nil, fmt.Errorf("failed to parse author id %w", err)
	}
	createParams := queries.CreateBookParams{
		ID:          id,
		Title:       req.Msg.Title,
		Description: req.Msg.Description,
		AuthorID:    authorID,
	}
	book, err := s.db.CreateBook(ctx, createParams)
	if err != nil {
		return nil, fmt.Errorf("failed to create author %w", err)
	}

	res := connect.NewResponse(&bookv1.CreateBookResponse{
		Book: encoder.DBBookToAPI(book),
	})

	return res, nil
}
