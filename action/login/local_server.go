// SPDX-License-Identifier: Apache-2.0

// mostly taken from https://github.com/cli/oauth/tree/v0.8.0/webapp

package login

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

type CodeResponse struct {
	Code  string
	State string
}

type localServer struct {
	CallbackPath     string
	WriteSuccessHTML func(w io.Writer)

	resultChan chan (CodeResponse)
	listener   net.Listener
}

// bindLocalServer initializes a LocalServer that will listen on a randomly available TCP port.
func bindLocalServer() (*localServer, error) {
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		return nil, err
	}

	return &localServer{
		listener:   listener,
		resultChan: make(chan CodeResponse, 1),
	}, nil
}

func (s *localServer) Port() int {
	return s.listener.Addr().(*net.TCPAddr).Port
}

func (s *localServer) Close() error {
	return s.listener.Close()
}

func (s *localServer) Serve() error {
	//nolint:gosec // TODO: add a way to timeout the local server
	return http.Serve(s.listener, s)
}

func (s *localServer) WaitForCode() (CodeResponse, error) {
	return <-s.resultChan, nil
}

// ServeHTTP implements http.Handler.
func (s *localServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.CallbackPath != "" && r.URL.Path != s.CallbackPath {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	defer func() {
		_ = s.Close()
	}()

	params := r.URL.Query()
	s.resultChan <- CodeResponse{
		Code:  params.Get("code"),
		State: params.Get("state"),
	}

	w.Header().Add("content-type", "text/html")

	if s.WriteSuccessHTML != nil {
		s.WriteSuccessHTML(w)
	} else {
		defaultSuccessHTML(w)
	}
}

func defaultSuccessHTML(w io.Writer) {
	fmt.Fprint(w, authSuccess)
}
