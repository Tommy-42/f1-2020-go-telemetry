package handler

import "github.com/pkg/errors"

var (
	// ErrUnknownPacket means that the packet couldnt be identified
	ErrUnknownPacket = errors.New("unknown packet")
	// ErrIgnorePacket means to avoid processing this packet
	ErrIgnorePacket = errors.New("ignore packet")
)
