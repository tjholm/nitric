package service

import (
	"context"
	"net/http"
)

type Proxy interface {
	Handle(ctx context.Context, req *http.Request) (*http.Response, error)
}

type Service interface {
	// TODO: We will likely need to pass in additional plugins here for identification of ingressed events that are not http requests
	Start(Proxy) error
}
