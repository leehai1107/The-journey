package xhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin/binding"
)

const contentTypeField = "Content-Type"

type builder struct {
	context     context.Context
	method      string
	url         string
	contentType string
	headers     map[string]string
	bodyData    interface{}
}

func NewRequestBuilder() *builder {
	headers := make(map[string]string)
	ctx := context.Background()
	return &builder{context: ctx, headers: headers}
}

func NewRequestBuilderWithCtx(ctx context.Context) *builder {
	headers := make(map[string]string)
	return &builder{context: ctx, headers: headers}
}

func (b *builder) WithMethod(method string) *builder {
	b.method = method
	return b
}

func (b *builder) WithBody(contentType string, data interface{}) *builder {
	b.contentType = contentType
	b.bodyData = data
	b.headers[contentTypeField] = contentType
	return b
}

func (b *builder) WithHeaders(headers map[string]string) *builder {
	for k, v := range headers {
		b.headers[k] = v
	}
	return b
}
func (b *builder) WithURL(url string) *builder {
	b.url = url
	return b
}
func (b *builder) WithContext(c context.Context) *builder {
	b.context = c
	return b
}

func (b *builder) Build() (req *http.Request, err error) {
	if b.method == http.MethodGet {
		return b.buildGetRequest()
	}
	bodyByte, err := b.buildBody()
	if err != nil {
		return nil, err
	}
	req, err = http.NewRequestWithContext(b.context, b.method, b.url, bytes.NewReader(bodyByte))
	if err != nil {
		return nil, err
	}
	for k, v := range b.headers {
		req.Header.Set(k, v)
	}
	return req, nil
}

func (b *builder) buildGetRequest() (req *http.Request, err error) {
	req, err = http.NewRequestWithContext(b.context, b.method, b.url, nil)
	if err != nil {
		return
	}
	for k, v := range b.headers {
		req.Header.Set(k, v)
	}
	return
}

func (b *builder) buildBody() ([]byte, error) {
	switch b.contentType {
	case binding.MIMEJSON:
		return json.Marshal(b.bodyData)
	case binding.MIMEPOSTForm:
		//TODO: transform bodyData to bodyData by form-urlencoded
		return nil, nil
	default:
		return json.Marshal(b.bodyData)
	}
}
