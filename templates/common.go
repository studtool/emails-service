package templates

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/valyala/fasttemplate"
)

const (
	StartTag = "{{"
	EndTag   = "}}"
)

type Template interface {
	Load() error
	Render(m map[string]interface{}) string
}

func loadFromFile(path string) (t *fasttemplate.Template, err error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()

	t = fasttemplate.New(string(b), StartTag, EndTag)
	return t, nil
}
