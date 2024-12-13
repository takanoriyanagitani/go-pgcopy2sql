package reader2header

import (
	"context"
	"io"
	"os"

	util "github.com/takanoriyanagitani/go-pgcopy2sql/util"
)

type SimpleHeader struct {
	Signature       [11]byte
	Flags           [4]byte
	HeaderExtension [4]byte
}

// Reads PGCOPY's header(with no checks).
func ReaderToSimpleHeader(
	rdr io.Reader,
) (SimpleHeader, error) {
	var h SimpleHeader

	_, e := io.ReadFull(rdr, h.Signature[:])
	if nil != e {
		return h, e
	}

	_, e = io.ReadFull(rdr, h.Flags[:])
	if nil != e {
		return h, e
	}

	_, e = io.ReadFull(rdr, h.HeaderExtension[:])
	if nil != e {
		return h, e
	}

	return h, nil
}

func StdinToSimpleHeader(_ context.Context) (SimpleHeader, error) {
	return ReaderToSimpleHeader(os.Stdin)
}

var SimpleHeaderStdinDefault util.IO[SimpleHeader] = StdinToSimpleHeader
