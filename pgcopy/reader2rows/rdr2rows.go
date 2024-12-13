package reader2rows

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"iter"
	"os"

	ps "github.com/takanoriyanagitani/go-pgcopy2sql"
	util "github.com/takanoriyanagitani/go-pgcopy2sql/util"

	pc "github.com/takanoriyanagitani/go-pgcopy2sql/pgcopy/convert"
)

func ReaderToRows(
	ctx context.Context,
	rdr io.Reader,
	converter pc.ConverterIx,
) iter.Seq2[[]ps.Value, error] {
	return func(yield func([]ps.Value, error) bool) {
		var br io.Reader = bufio.NewReader(rdr)

		var bcolcnt [2]byte
		var bcolsiz [4]byte
		var bcolumn bytes.Buffer

		var row []ps.Value

		for {
			_, e := io.ReadFull(br, bcolcnt[:])
			if nil != e {
				yield(nil, e)
				return
			}

			var colcnt ps.PgColumnCount = ps.PgColumnCountFromRawBytes(bcolcnt)
			var hasNext bool = !colcnt.LastRow()
			if !hasNext {
				return
			}

			clear(row)
			row = row[:0]
			var ix int16
			for ix = 0; ix < colcnt.Count(); ix++ {
				_, e := io.ReadFull(br, bcolsiz[:])
				if nil != e {
					yield(nil, e)
					return
				}

				var colsiz ps.PgColumnSize = ps.
					PgColumnSizeFromRawBytes(bcolsiz)
				var isNull bool = colsiz.IsNull()
				if isNull {
					val, e := converter(ix, nil)(ctx)
					if nil != e {
						yield(nil, e)
						return
					}
					row = append(row, val)
					continue
				}

				bcolumn.Reset()
				var icolsiz int64 = int64(colsiz.Size())
				limited := io.LimitedReader{
					R: br,
					N: icolsiz,
				}
				_, e = io.Copy(&bcolumn, &limited)
				if nil != e {
					yield(nil, e)
					return
				}

				val, e := converter(ix, bcolumn.Bytes())(ctx)
				if nil != e {
					yield(nil, e)
					return
				}
				row = append(row, val)
			}

			if !yield(row, nil) {
				return
			}
		}
	}
}

func StdinToRows(
	ctx context.Context,
	converter pc.ConverterIx,
) iter.Seq2[[]ps.Value, error] {
	return ReaderToRows(ctx, os.Stdin, converter)
}

func ConverterToStdinToRows(
	conv pc.ConverterIx,
) util.IO[iter.Seq2[[]ps.Value, error]] {
	return func(ctx context.Context) (iter.Seq2[[]ps.Value, error], error) {
		return StdinToRows(ctx, conv), nil
	}
}
