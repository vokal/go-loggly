package loggly

import (
	"bytes"
	"fmt"
	"net/http"
)

var (
	logglyUrl = "https://logs-01.loggly.com/inputs/%s/tag/%s/"
)

type Loggly struct {
	token string
	tag   string
}

func New(token, tag string) *Loggly {
	l := Loggly{
		token: token,
		tag:   tag,
	}
	return &l
}

func (l *Loggly) Write(data []byte) (int, error) {
	buf := bytes.NewBuffer(data)

	resp, err := http.Post(fmt.Sprintf(logglyUrl, l.token, l.tag), "text/plain", buf)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	return len(data), nil
}
