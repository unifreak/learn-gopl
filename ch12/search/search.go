package search

import (
	"fmt"
	"learn/gopl/ch12/params"
	"net/http"
)

// search implements the /search URL endpoint
//
// the variable data's fields correspond to the HTTP request parameters. The struct's
// field tags specify the parameter names. The Unpak function populates the struct
// from the request so that the parameters can be accessed conveniently and with an
// appropriate type.
func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels     []string `http:"1"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	data.MaxResults = 10 // set default
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest) // 400
		return
	}

	// ...rest of handler...
	fmt.Fprintf(resp, "Search: %+v\n", data)
}
