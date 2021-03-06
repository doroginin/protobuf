// Code generated by protoc-gen-go-http-server.
// source: pb/strings.proto

package strings

import (
	"net/http"
)

func (s *StringsHTTPServer) ToUpper(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r, _, in, err := s.opts.cdc.ReadRequest(r)
	if err != nil {
		s.opts.cdc.WriteResponse(w, nil, err)
		return
	}
	resp, err := s.opts.srv.ToUpper(r.Context(), in.(*StringRequest))
	if err != nil {
		s.opts.cdc.WriteResponse(w, nil, err)
		return
	}
	s.opts.cdc.WriteResponse(w, resp, nil)
}

func (s *StringsHTTPServer) ToLower(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r, _, in, err := s.opts.cdc.ReadRequest(r)
	if err != nil {
		s.opts.cdc.WriteResponse(w, nil, err)
		return
	}
	resp, err := s.opts.srv.ToLower(r.Context(), in.(*StringRequest))
	if err != nil {
		s.opts.cdc.WriteResponse(w, nil, err)
		return
	}
	s.opts.cdc.WriteResponse(w, resp, nil)
}
