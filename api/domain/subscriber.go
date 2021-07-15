package domain

import (
	"errors"
	"net/url"
	"regexp"
)

type Subscriber struct {
	id   int
	name string
	url  string
}

func (s *Subscriber) SetName(name string) error {
	isValidName := regexp.MustCompile(`^[a-z0-9]*$`).MatchString
	if !isValidName(name) {
		return errors.New("invalid name: name must contain lower case alphabets and numbers only")
	}
	s.name = name
	return nil
}

func (s *Subscriber) SetUrl(urlStr string) error {
	_, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return errors.New("invalid url: please send correct one")
	}
	s.url = urlStr
	return nil
}

func (s *Subscriber) SetId(id int) {

	if id != 0 {
		s.id = id
	}

}
func (s *Subscriber) Id() int {

	return s.id

}

func (s *Subscriber) Name() string {

	return s.name
}

func (s *Subscriber) Url() string {

	return s.url
}
