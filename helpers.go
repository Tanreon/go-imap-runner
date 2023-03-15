package imap_runner

import (
	"errors"
	"strings"
)

func isConnectionError(err error) bool {
	if strings.EqualFold(err.Error(), "EOF") {
		return true
	}

	wrappedErr := errors.Unwrap(err)
	if wrappedErr != nil && strings.EqualFold(wrappedErr.Error(), "EOF") {
		return true
	}

	if strings.Contains(err.Error(), "User is authenticated but not connected") {
		return true
	}

	if strings.Contains(err.Error(), "Server Unavailable") {
		return true
	}

	if strings.Contains(err.Error(), "Please try again later") {
		return true
	}

	if strings.Contains(err.Error(), "imap: connection closed") {
		return true
	}

	return false
}
