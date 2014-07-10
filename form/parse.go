package form

import (
	"errors"
	"net/http"
	"net/url"
)

// warning: modifies req by calling req.ParseForm()
func Parse(req *http.Request, values ...string) (form url.Values, err error) {
	req.ParseForm()
	form = req.PostForm
	err = Check(form, values...)
	return
}

func Check(data url.Values, values ...string) error {
	for _, s := range values {
		if len(data[s]) == 0 {
			return errors.New(s + " not passed")
		}
	}
	return nil
}
