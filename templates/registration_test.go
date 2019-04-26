package templates

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"bou.ke/monkey"
)

func TestRegistrationTemplate_Load_Render_Simple(t *testing.T) {
	const data = "This is the template"
	monkey.Patch(ioutil.ReadFile, func(filename string) ([]byte, error) {
		return []byte(data), nil
	})

	tmp := NewRegistrationTemplate()
	if err := tmp.Load(); err != nil {
		t.Fail()
	}

	if tmp.Render(map[string]interface{}{}) != data {
		t.Fail()
	}
}

func TestRegistrationTemplate_Load_Render_Complex(t *testing.T) {
	const tVarName = "user"
	const tVarValue = "Mr. User"
	tmpVar := fmt.Sprintf("%s%s%s", StartTag, tVarName, EndTag)

	tmpTxt := fmt.Sprintf("Hello %s!!!", tmpVar)
	tmpValue := strings.ReplaceAll(tmpTxt, tmpVar, tVarValue)

	monkey.Patch(ioutil.ReadFile, func(filename string) ([]byte, error) {
		return []byte(tmpTxt), nil
	})

	tmp := NewRegistrationTemplate()
	if err := tmp.Load(); err != nil {
		t.Fail()
	}

	m := map[string]interface{}{tVarName: tVarValue}
	if tmp.Render(m) != tmpValue {
		t.Fail()
	}
}
