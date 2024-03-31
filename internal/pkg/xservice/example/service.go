package example_test

import (
	"context"
	"fmt"

	"github.com/leehai1107/The-journey/internal/pkg/logger"
	"github.com/leehai1107/The-journey/internal/pkg/xhttp"
)

type IExampleSvc interface {
	ExampleGet(ctx context.Context, req *Request) (Response, error)
	ExamplePost(ctx context.Context, req *Request) (Response, error)
}

type exampleSvc struct {
	client  xhttp.Client
	baseUrl string
}

func NewExampleSvc(client xhttp.Client, baseUrl string) IExampleSvc {
	return &exampleSvc{
		client:  client,
		baseUrl: baseUrl,
	}
}

func (s *exampleSvc) ExampleGet(ctx context.Context, req *Request) (Response, error) {
	lg := logger.EnhanceWith(ctx)
	// lg.Infof("[Example]request: %+v", req)
	lg.Infow("[Example]", "request", req)
	var res Response
	path := fmt.Sprintf("%s/example", req.Data)
	url := fmt.Sprintf("%v%v", s.baseUrl, path)
	xOPT := xhttp.RequestOption{
		GroupPath: path,
	}
	_, err := s.client.Get(ctx, url, &res, xOPT)
	lg.Errorw("[Example]", "error", err)
	return res, nil
}

func (s *exampleSvc) ExamplePost(ctx context.Context, req *Request) (Response, error) {
	lg := logger.EnhanceWith(ctx)
	// lg.Infof("[Example]request: %+v", req)
	lg.Infow("[Example]", "request", req)
	var res Response
	path := fmt.Sprintf("%s/example", req.Data)
	url := fmt.Sprintf("%v%v", s.baseUrl, path)
	xOPT := xhttp.RequestOption{
		GroupPath: path,
	}
	_, err := s.client.PostJSON(ctx, url, req, &res, xOPT)
	lg.Errorw("[Example]", "error", err)
	return res, nil
}
