package app

import "context"

type Server interface {
	Start() error
	Stop(context.Context) error
}