package application

import (
	"context"
	"sugarerpgo/internal/core"
)

type UserRepository interface {
	GetUsers(ctx context.Context, tenantID string) ([]core.User, error)
}

type GetUsersQuery struct {
	TenantID string
}

type GetUsersHandler struct {
	Repo     UserRepository
	EventBus EventBus
}

type EventBus interface {
	Publish(event interface{}) error
}

func (h *GetUsersHandler) Handle(ctx context.Context, query GetUsersQuery) ([]core.User, error) {
	users, err := h.Repo.GetUsers(ctx, query.TenantID)
	if err != nil {
		return nil, err
	}
	for _, u := range users {
		h.EventBus.Publish(core.UserRetrievedEvent{User: u})
	}
	return users, nil
}
