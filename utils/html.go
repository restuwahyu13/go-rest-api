package util

import (
	"bytes"
	"fmt"
	"path"
	"text/template"

	"github.com/sirupsen/logrus"
)

type BodyRequest struct {
	To    string
	Token string
}

func ParseHtml(fileName string, data map[string]string) string {
	html, errParse := template.ParseFiles(path.Dir("../templates/") + fileName)

	if errParse != nil {
		defer fmt.Println("parser file html failed")
		logrus.Fatal(errParse.Error())
	}

	body := BodyRequest{To: data["to"], Token: data["token"]}

	buf := new(bytes.Buffer)
	errExecute := html.Execute(buf, body)

	if errExecute != nil {
		defer fmt.Println("execute html file failed")
		logrus.Fatal(errExecute.Error())
	}

	return buf.String()
}
