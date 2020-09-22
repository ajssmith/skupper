package main

import (
	"testing"

	"github.com/skupperproject/skupper/client"
)

type myMock struct {
	//TODO implement the "real" mock
	client.VanClient
}

func TestConnectionTokenRun(t *testing.T) {

	cmd := NewCmdConnectionToken(nil)

	cli = &myMock{}

	cmd.RunE(nil, []string{"tokenName"})

	//pseudocode:
	assert(cli.ConnectorTokenCreateFile.calledwith(
	    token: "tokenName"
	    clientId: types.DefaultVanName
	))
}
