package repository

import "context"

type EmailGatewayRepository interface {
	SendEmail(ctx context.Context, to, title, body string) error
}
