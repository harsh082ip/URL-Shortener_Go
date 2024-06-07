package consts

import "time"

const (
	WEBPORT         = ":8002"
	LetterBytes     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	SessionTTL      = 30 * time.Second
	SessionIDlength = 16
)
