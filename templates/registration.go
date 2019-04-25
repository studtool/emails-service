package templates

import (
	"github.com/valyala/fasttemplate"

	"github.com/studtool/emails-service/config"
)

type RegistrationTemplate struct {
	path     string
	template *fasttemplate.Template
}

func NewRegistrationTemplate() *RegistrationTemplate {
	return &RegistrationTemplate{
		path: config.RegEmailTemplatePath.Value(),
	}
}

func (t *RegistrationTemplate) Load() (err error) {
	t.template, err = loadFromFile(t.path)
	return err
}

func (t *RegistrationTemplate) Render(m map[string]interface{}) string {
	return t.template.ExecuteString(m)
}
