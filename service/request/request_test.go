package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var request Request
var ch chan string

func init() {
	ch = make(chan string)
	request = New(ch)
}

func getName(f func()) string {
	go f()
	name := <-ch
	ch <- ""
	return name
}

func TestSetScheme(t *testing.T) {
	assert.Equal(t, `kong.service.request.set_scheme:["http"]`, getName(func() { request.SetScheme("http") }))
}

func TestSetPath(t *testing.T) {
	assert.Equal(t, `kong.service.request.set_path:["/foo"]`, getName(func() { request.SetPath("/foo") }))
}

func TestSetRawQuery(t *testing.T) {
	assert.Equal(t, `kong.service.request.set_raw_query:["q1=v1\u0026q2=v2"]`, getName(func() { request.SetRawQuery("q1=v1&q2=v2") }))
}

func TestSetMethod(t *testing.T) {
	assert.Equal(t, `kong.service.request.set_method:["GET"]`, getName(func() { request.SetMethod("GET") }))
}

func TestSetQuery(t *testing.T) {
	assert.Equal(t, `kong.service.request.set_query:["q1=v2\u0026q2=v2"]`, getName(func() { request.SetQuery("q1=v2&q2=v2") }))
}

func TestSetHeader(t *testing.T) {
	assert.Equal(t, `kong.service.request.set_header:["q1","v1"]`, getName(func() { request.SetHeader("q1", "v1") }))
}

func TestAddHeader(t *testing.T) {
	assert.Equal(t, `kong.service.request.add_header:["q1","v1"]`, getName(func() { request.AddHeader("q1", "v1") }))
}

func TestClearHeader(t *testing.T) {
	assert.Equal(t, `kong.service.request.clear_header:["q1"]`, getName(func() { request.ClearHeader("q1") }))
}

func TestSetHeaders(t *testing.T) {
	assert.Equal(t, `kong.service.request.set_headers:[{"h1":"v1"}]`, getName(func() {
		request.SetHeaders(map[string]interface{}{
			"h1": "v1",
		})
	}))
}

func TestSetRawBody(t *testing.T) {
	assert.Equal(t, `kong.service.request.set_raw_body:["body"]`, getName(func() { request.SetRawBody("body") }))
}