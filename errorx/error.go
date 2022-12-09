package errorx

import "google.golang.org/grpc/status"

var (
	NoSuchCat       = status.Error(1001, "no such cat")
	GetManyCatError = status.Error(1002, "GetManyCatError")
	UpdateError     = status.Error(1003, "UpdateError")
)
