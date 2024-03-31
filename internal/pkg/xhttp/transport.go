package xhttp

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/leehai1107/The-journey/internal/pkg/logger"
)

type Transport struct {
	transport http.RoundTripper
	opts      clientOptions
}

func NewTransport(opts clientOptions) http.RoundTripper {
	transport := getTransport(opts)
	return &Transport{transport: transport, opts: opts}
}

func getTransport(opts clientOptions) http.RoundTripper {
	transport := http.DefaultTransport
	if len(opts.proxyURL) != 0 {
		if proxyURL, err := url.Parse(opts.proxyURL); err == nil {
			transport = &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			}
		}
	}
	if !opts.promCfg.Enable {
		return transport
	}
	promCfg := opts.promCfg
	metrics := NewOutgoingMetrics(promCfg.Subsystem, promCfg.ConstLabel)
	promTransport := buildTraceTransport(transport, metrics)
	if err := promCfg.Register.Register(metrics); err != nil {
		logger.Errorw("failed to register http outgoing metrics", "error", err)
	}
	return promTransport
}

func (t *Transport) RoundTrip(req *http.Request) (res *http.Response, err error) {
	start := time.Now()
	t.dumpRequest(req)
	res, err = t.transport.RoundTrip(req)
	if err != nil {
		return
	}
	t.dumpResponse(res, start)
	return
}

func (t *Transport) dumpRequest(req *http.Request) {
	if t.opts.skipLog {
		return
	}
	ctx := req.Context()
	reqDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		logger.EnhanceWith(ctx).Errorf("failed to dump request %+v", err)
		return
	}
	reqDumpStr := string(reqDump)
	logger.EnhanceWith(ctx).Infof(
		"--) %s | url: %s | request: %s", req.Method, req.URL.String(), reqDumpStr)
}

func (t *Transport) dumpResponse(rsp *http.Response, start time.Time) {
	if t.opts.skipLog {
		return
	}
	ctx := rsp.Request.Context()
	lg := logger.EnhanceWith(ctx)
	rspDump, dumpErr := httputil.DumpResponse(rsp, true)
	if dumpErr != nil {
		lg.Errorf("failed to dump response %+v", dumpErr)
		return
	}
	method, url := rsp.Request.Method, rsp.Request.URL.String()
	logRsp := fmt.Sprintf("(-- END %s, url: %s, latencies.ms: %d,  bodyData: %s,",
		method, url, time.Since(start).Milliseconds(), string(rspDump))
	if !t.opts.splitLogBody {
		lg.Info(logRsp)
		return
	}
	if len(rspDump) <= t.opts.splitLogBodyLen {
		lg.Info(logRsp)
		return
	}
	rspLen := len(rspDump)
	limit := t.opts.splitLogBodyLen
	parts := rspLen / limit
	if rspLen%limit != 0 {
		parts++
	}
	for i := 0; i < parts; i++ {
		offset := i * limit
		end := offset + limit
		var dataStr string
		if end > rspLen {
			dataStr = string(rspDump[offset:])
		} else {
			dataStr = string(rspDump[offset:end])
		}
		lg.Infof(
			"(-- END %s, url: %s, latencies.ms: %d, PART: %d/%d, bodyData: %s,",
			method, url, time.Since(start).Milliseconds(), i+1, parts, dataStr)
	}
}
