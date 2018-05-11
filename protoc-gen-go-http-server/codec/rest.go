package codec

import (
	"encoding/json"
	"net/http"
	"github.com/doroginin/protobuf/protoc-gen-go-http-server/types"
)

type RESTCodec struct{}

func NewRESTCodec() types.Codec {
	return &RESTCodec{}
}

func (c *RESTCodec) Route(r *http.Request) (route string, err error) {
	return r.URL.Path, nil
}

func (c *RESTCodec) ReadRequest(r *http.Request, out interface{}) error {
	if r.Body == http.NoBody {
		return nil
	}
	var (
		decoder = json.NewDecoder(r.Body)
		err     = decoder.Decode(out)
	)
	r.Body.Close()

	return err
}

func (c *RESTCodec) WriteResponse(w http.ResponseWriter, resp interface{}) error {
	bResp, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	_, err = w.Write(bResp)
	return err
}

func (c *RESTCodec) WriteError(w http.ResponseWriter, err error) error {
	w.WriteHeader(http.StatusInternalServerError)
	bResp, _ := json.Marshal(&defaultError{Error: err.Error()})
	_, err = w.Write(bResp)
	return err
}

type defaultError struct {
	Error string `json:"error"`
}
