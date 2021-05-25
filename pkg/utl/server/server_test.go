package server_test

import (
	"testing"

	"github.com/gabrielopesantos/myDrive-api/pkg/utl/server"
)

func TestNew(t *testing.T) {
	s, err := server.New()

	if s == nil {
		t.Errorf("Server should not be nil")
	}

	if err != nil {
		t.Errorf("Something wrong setting up the server")
	}
}
