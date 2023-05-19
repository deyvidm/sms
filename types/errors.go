package types

import "fmt"

type ResponseParseError struct{}

func (e ResponseParseError) Error() string {
	return "could not parse a clear response"
}

type MissingKeyError struct {
	PendingInvites int
}

func (e MissingKeyError) Error() string {
	return fmt.Sprintf("user has %d pending invites and submitted no specifier", e.PendingInvites)
}
