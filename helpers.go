package imap_runner

import (
	"strings"
)

func isConnectionError(err error) bool {
	return strings.Contains(err.Error(), "imap: connection closed")
}
