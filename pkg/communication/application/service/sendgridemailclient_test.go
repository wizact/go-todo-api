package service

import (
	"errors"
	"testing"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func TestSendGridEmailClient_Send_TransportFailurePreservesCause(t *testing.T) {
	t.Parallel()

	cause := errors.New("transport failed")
	client := SendGridEmailClient{
		sendGridFromName:  "TODO API",
		sendGridFromEmail: "todo@example.com",
		send: func(*mail.SGMailV3) (int, error) {
			return 0, cause
		},
	}

	err := client.Send("Ada", "ada@example.com", "Welcome", "Hello", "<p>Hello</p>")

	if !errors.Is(err, cause) {
		t.Fatalf("Send() error = %v, want transport cause", err)
	}
}

func TestSendGridEmailClient_Send_NonSuccessReturnsError(t *testing.T) {
	t.Parallel()

	client := SendGridEmailClient{
		sendGridFromName:  "TODO API",
		sendGridFromEmail: "todo@example.com",
		send: func(*mail.SGMailV3) (int, error) {
			return 400, nil
		},
	}

	err := client.Send("Ada", "ada@example.com", "Welcome", "Hello", "<p>Hello</p>")

	if !errors.Is(err, ErrSendGridRejected) {
		t.Fatalf("Send() error = %v, want rejection error", err)
	}
}
