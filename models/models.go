package models

import "bytes"

type F1Data interface {
	ToJson() (*bytes.Reader, error)
}
