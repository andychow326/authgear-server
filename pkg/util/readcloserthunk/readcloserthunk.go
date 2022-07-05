package readcloserthunk

import (
	"bytes"
	"errors"
	"hash"
	"io"
	"net/http"
)

type ReadCloserThunk = func() (io.ReadCloser, error)

// nolint:golint
func Performance_Bytes(thunk ReadCloserThunk) (b []byte, err error) {
	rc, err := thunk()
	if err != nil {
		return
	}
	defer rc.Close()
	b, err = io.ReadAll(rc)
	return
}

func HTTPDetectContentType(thunk ReadCloserThunk) (out string) {
	rc, err := thunk()
	if err != nil {
		return
	}
	defer rc.Close()
	r := &io.LimitedReader{R: rc, N: 512}
	var w bytes.Buffer
	_, err = io.Copy(&w, r)
	if err != nil {
		return
	}
	return http.DetectContentType(w.Bytes())
}

func Hash(thunk ReadCloserThunk, hash hash.Hash) ([]byte, error) {
	rc, err := thunk()
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	_, err = io.Copy(hash, rc)
	if err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}

func Copy(w io.Writer, thunk ReadCloserThunk) (written int64, err error) {
	rc, err := thunk()
	if err != nil {
		return
	}
	defer rc.Close()
	return io.Copy(w, rc)
}

func Reader(r io.Reader) ReadCloserThunk {
	return func() (io.ReadCloser, error) {
		return io.NopCloser(r), nil
	}
}

type multiReadCloserThunk struct {
	thunks     []ReadCloserThunk
	readCloser io.ReadCloser
}

func MultiReadCloserThunk(thunks ...ReadCloserThunk) ReadCloserThunk {
	return func() (io.ReadCloser, error) {
		return &multiReadCloserThunk{
			thunks:     thunks,
			readCloser: nil,
		}, nil
	}
}

func (m *multiReadCloserThunk) Read(b []byte) (n int, err error) {
	if m.readCloser == nil {
		if len(m.thunks) <= 0 {
			return 0, io.EOF
		}
		var rc io.ReadCloser
		rc, err = m.thunks[0]()
		if err != nil {
			return
		}
		m.readCloser = rc
		m.thunks = m.thunks[1:]
	}
	n, err = m.readCloser.Read(b)
	if errors.Is(err, io.EOF) {
		m.readCloser.Close()
		m.readCloser = nil
		err = nil
	}
	return
}

func (m *multiReadCloserThunk) Close() (err error) {
	if m.readCloser != nil {
		return m.readCloser.Close()
	}
	return
}
