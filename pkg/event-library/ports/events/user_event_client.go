package events

import (
	in_event_port "github.com/wizact/go-todo-api/pkg/event-library/ports/input/events"
	out_event_port "github.com/wizact/go-todo-api/pkg/event-library/ports/output/events"
)

type UserEventClient interface {
	in_event_port.UserEventClientInput
	out_event_port.UserEventClientOutput
}
