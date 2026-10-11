// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"strings"
	"testing"
	"time"
)

const (
	ExampleLine   = exampleLine
	ExampleColumn = exampleColumn
	ExampleOffset = exampleOffset
)

// runServe runs the serve command with the given -listen address and returns
// its error, or nil if it is still running (i.e. listening) after a short wait.
func runServe(t *testing.T, address string) error {
	t.Helper()
	app := New("", nil)
	s := &Serve{Address: address, app: app}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	errc := make(chan error, 1)
	go func() { errc <- s.Run(ctx) }()
	select {
	case err := <-errc:
		return err
	case <-time.After(500 * time.Millisecond):
		return nil
	}
}

func TestServeListenRejectsImplicitAllInterfaces(t *testing.T) {
	for _, addr := range []string{":0", ":12345"} {
		err := runServe(t, addr)
		if err == nil || !strings.Contains(err.Error(), "implicitly binds all network interfaces") {
			t.Errorf("-listen=%s: got error %v, want implicit bind-all rejection", addr, err)
		}
	}
}

func TestServeListenAllowsExplicitHost(t *testing.T) {
	if err := runServe(t, "localhost:0"); err != nil && strings.Contains(err.Error(), "implicitly binds all network interfaces") {
		t.Errorf("-listen=localhost:0: unexpected rejection: %v", err)
	}
}
