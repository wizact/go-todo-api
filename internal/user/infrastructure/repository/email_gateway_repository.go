package repository

import (
	"context"
)

type EmailGatewayRepository struct {
}

func NewEmailGatewayRepository() *EmailGatewayRepository {
	return &EmailGatewayRepository{}
}

func (r *EmailGatewayRepository) SendEmail(ctx context.Context, to, title, body string) error {
	return nil
}
