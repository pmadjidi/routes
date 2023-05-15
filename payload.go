package main

type ApiPayload struct {
	route   Route
	inError bool
	msg     string
	body    []byte
	resp    chan *ApiPayload
}

func newApiPayload(src, dst latlong) *ApiPayload {
	return &ApiPayload{Route{src, dst}, false, "", nil, make(chan *ApiPayload, 1)}
}
