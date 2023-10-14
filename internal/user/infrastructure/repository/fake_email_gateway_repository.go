package repository

import (
	"context"
)

type FakeEmailGatewayRepository struct {
}

func NewFakeEmailGatewayRepository() *FakeEmailGatewayRepository {
	return &FakeEmailGatewayRepository{}
}

func (r *FakeEmailGatewayRepository) SendEmail(ctx context.Context, to, title, body string) error {
	return nil
}
