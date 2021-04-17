package util

import (
	"bytes"
	"fmt"
	"path"
	"text/template"

	"github.com/sirupsen/logrus"
)

type Request struct {
	Body string
}

func (r *Request) ParseHtml(fileName string, data interface{}) error {
	html, errParse := template.ParseFiles(path.Dir("../templates/") + fileName)

	if errParse != nil {
		defer fmt.Println("parser file html failed")
		logrus.Fatal(errParse.Error())
	}

	buf := new(bytes.Buffer)
	defer fmt.Println("execute html file failed")
	errExecute := html.Execute(buf, data)

	if errExecute != nil {
		logrus.Fatal(errExecute.Error())
	}

	r.Body = buf.String()
	return nil
}
