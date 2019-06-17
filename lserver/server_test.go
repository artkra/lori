package lserver

import (
	"bytes"
	"testing"
)

func TestLSplit(t *testing.T) {
	// add tests
	r0, r1, r2 := lSplit([]byte("+++idspispopd_____ABCAAAAAAAAAAAAAAZ789102jsAUoek99"), false)
	exp0, exp1, exp2 := 1, []byte("ABCAAAAAAAAAAAAAAZ789102jsAUoek99"), error(nil)

	if r0 != exp0 || !bytes.Equal(r1, exp1) || r2 != exp2 {
		t.Error(r0, string(r1), r2, " != ", exp0, string(exp1), exp2)
	}
}
