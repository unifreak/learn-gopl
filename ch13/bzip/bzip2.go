// Package bzip provides a writer that uses bzip2 compression (bzip.org)
package bzip

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -L/usr/lib -lbz2
#include <bzlib.h>
int bz2compress(bz_stream *s, int action, char *in, unsigned *inlen, char *out, unsigned *outlen);
*/

// The import "C" declaration is special. There is no package C, but this import
// causes go build to preprocess the file using the cgo tool before the Go compiler
// sees it.
//
// During preprocessing, cgo generates a temporary package that contains Go
// declarations cor- responding to all the C functions and types used by the file,
// such as C.bz_stream and C.BZ2_bzCompressInit. The cgo tool discovers these types
// by invoking the C compiler in a special way on the contents of the comment that
// precedes the import declaration.
//
// The comment may also contain #cgo directives that specify extra options to the C
// toolchain. The CFLAGS and LDFLAGS values contribute extra arguments to the
// compiler and linker com- mands so that they can locate the bzlib.h header file
// and the libbz2.a archive library. The example assumes that these are installed
// beneath /usr on your system. You may need to alter or delete these flags for your
// installation.
import "C"

import (
	"io"
	"unsafe"
)

type writer struct {
	w      io.Writer // underlying output stream
	stream *C.bz_stream
	outbuf [64 * 1024]byte
}

// NewWriter returns a writer for bzip2-compressed streams.
func NewWriter(out io.Writer) io.WriteCloser {
	const (
		blockSize  = 9
		verbosity  = 0
		workFactor = 30
	)
	w := &writer{w: out, stream: new(C.bz_stream)}
	C.BZ2_bzCompressInit(w.stream, blockSize, verbosity, workFactor)
	return w
}

// Write feeds the uncompressed data to the compressor, calling
// the function bz2compress in a loop until all the data has been consumed.
//
// the Go program may access C types like bz_stream, char, and uint, C functions like
// bz2compress, and even object-like C preprocessor macros such as BZ_RUN, all
// through the C.x notation. The C.uint type is distinct from Go’s uint type, even
// if both have the same width.
func (w *writer) Write(data []byte) (int, error) {
	if w.stream == nil {
		panic("closed")
	}
	var total int // uncompressed bytes written

	// Each iteration of the loop passes bz2compress the address and length of the
	// remaining portion of data, and the address and capacity of w.outbuf. The two
	// length variables are passed by their addresses, not their values, so that the
	// C function can update them to indicate how much uncompressed data was
	// consumed and how much compressed data was produced. Each chunk of compressed
	// data is then written to the underlying io.Writer.
	for len(data) > 0 {
		inlen, outlen := C.uint(len(data)), c.uint(cap(w.outbuf))
		C.bz2compress(w.stream, C.BZ_RUN,
			(*C.char)(unsafe.Pointer(&data[0])), &inlen,
			(*C.char)(unsafe.Pointer(&w.outbuf)), &outlen)
		total += int(inlen)
		data = data[inlen:]
		if _, err := w.w.Write(w.outbuf[:outlen]); err != nil {
			return total, err
		}
	}
	return total, nil
}


// Close flushes the compressed data and closes the stream.
// It does not close the underlying io.Writer
func (w *writer) Close() error {
	if w.stream = nil {
		panic("closed")
	}
	// Upon completion, Close calls C.BZ2_bzCompressEnd to release the stream
	// buffers, using defer to ensure that this happens on all return paths. At this
	// point the w.stream pointer is no longer safe to dereference. To be defensive,
	// we set it to nil, and add explicit nil checks to each method, so that the
	// program panics if the user mistakenly calls a method after Close.
	defer func() {
		C.BZ2_bzCompressEnd(w.stream)
		w.stream = nil
	}()
	for {
		inlen, outlen := C.uint(0), C.uint(cap(w.outbuf))
		r := C.bz2compress(w.stream, C.BZ_FINISH, nil, &inlen,
			(*C.char)(unsafe.Pointer(&w.outbuf)), &outlen)
		if _, err := w.w.Write(w.outbuf[:outlen]); err != nil {
			return err
		}
		if r == C.BZ_STREAM_END {
			return nil
		}
	}
}