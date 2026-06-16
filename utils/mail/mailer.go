package mail

import (
	"context"
)

type Mailer interface {
	SendEmail(ctx context.Context, to string, subject *string, htmlBody *string, textBody *string) error
}
