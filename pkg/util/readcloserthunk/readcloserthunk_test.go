package readcloserthunk

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func getTestReadCloserThunks() (rct1 ReadCloserThunk, rct2 ReadCloserThunk, rct3 ReadCloserThunk) {
	rct1 = func() (io.ReadCloser, error) {
		return os.Open("testdata/test1.txt")
	}
	rct2 = func() (io.ReadCloser, error) {
		return os.Open("testdata/test2.txt")
	}
	rct3 = func() (io.ReadCloser, error) {
		return os.Open("testdata/test3.txt")
	}
	return
}

func TestConstructReadCloserThunk(t *testing.T) {
	Convey("constructReadCloserThunk", t, func() {
		r := strings.NewReader("test")

		Convey("should construct ReadCloserThunk from io.Reader", func() {
			rct := Reader(r)
			rc, err1 := rct()
			// nolint:staticcheck
			defer rc.Close()
			b, err2 := io.ReadAll(rc)
			So(string(b), ShouldEqual, "test")
			So(err1, ShouldBeNil)
			So(err2, ShouldBeNil)
		})

	})
}

func TestMultiReadCloserThunk(t *testing.T) {
	Convey("multiReadCloserThunk", t, func() {
		Convey("should construct MultiReadCloserThunk from mutiple thunks", func() {
			rct1, rct2, rct3 := getTestReadCloserThunks()
			m := MultiReadCloserThunk(rct1, rct2, rct3)
			mrc, err1 := m()
			// nolint:staticcheck
			defer mrc.Close()
			b, err2 := io.ReadAll(mrc)
			So(string(b), ShouldEqual, "testing1testing2testing3")
			So(err1, ShouldBeNil)
			So(err2, ShouldBeNil)
		})

		Convey("should construct MultiReadCloserThunk from mutiple thunks and thunks generated by io.Reader", func() {
			rct1, rct2, rct3 := getTestReadCloserThunks()
			m := MultiReadCloserThunk(rct1, Reader(strings.NewReader("hi")), Reader(bytes.NewReader([]byte("123"))), rct2, rct3)
			mrc, err1 := m()
			// nolint:staticcheck
			defer mrc.Close()
			b, err2 := io.ReadAll(mrc)
			So(string(b), ShouldEqual, "testing1hi123testing2testing3")
			So(err1, ShouldBeNil)
			So(err2, ShouldBeNil)
		})
	})
}

func TestPerformance_Bytes(t *testing.T) {
	Convey("Performance_Bytes", t, func() {
		Convey("should output []byte from ReadCloserThunk", func() {
			rct, _, _ := getTestReadCloserThunks()
			b, err := Performance_Bytes(rct)
			So(string(b), ShouldEqual, "testing1")
			So(err, ShouldBeNil)
		})
	})
}

func TestCopy(t *testing.T) {
	Convey("Copy", t, func() {
		Convey("should copy the reader in thunk to the writter", func() {
			rct, _, _ := getTestReadCloserThunks()
			w := new(bytes.Buffer)
			written, err := Copy(w, rct)
			So(written, ShouldEqual, 8)
			So(err, ShouldBeNil)
		})
	})
}
