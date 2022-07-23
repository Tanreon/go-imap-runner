package imap_runner

import (
	"strings"
)

func isConnectionError(err error) bool {
	if strings.Contains(err.Error(), "User is authenticated but not connected") {
		return true
	}

	if strings.Contains(err.Error(), "imap: connection closed") {
		return true
	}

	return false
}
