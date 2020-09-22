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
	cli = &myMock{}
	connectionTokenRun(nil, []string{"tokenName"})

	//pseudocode:
	assert(cli.ConnectorTokenCreateFile.calledwith(
	    token: "tokenName"
	    clientId: types.DefaultVanName
	))
}
