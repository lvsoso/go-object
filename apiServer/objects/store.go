package objects

import (
	"apiServer/locate"
	"bytes"
	"fmt"
	"go-object/lib/utils"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func storeObject(r io.Reader, hash string, size int64) (int, error) {
	if locate.Exist(url.PathEscape(hash)) {
		return http.StatusOK, nil
	}

	stream, e := putStream(url.PathEscape(hash), size)
	if e != nil {
		return http.StatusInternalServerError, e
	}

	body, _ := ioutil.ReadAll(r)
	buf := bytes.NewBuffer(body)
	fmt.Println(buf.String())

	r = ioutil.NopCloser(bytes.NewBuffer(body))

	reader := bytes.NewReader(body)
	// reader := io.TeeReader(r, stream)

	d := utils.CalculateHash(reader)
	if d != hash {
		stream.Commit(false)
		return http.StatusBadRequest, fmt.Errorf("object hash mismatch, calculated=%s, requested=%s", d, hash)
	}
	io.Copy(stream, r)
	stream.Commit(true)
	return http.StatusOK, nil
}
