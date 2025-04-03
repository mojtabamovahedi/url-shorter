package model

import (
	"errors"
	"regexp"
	"unicode/utf8"
)

// regex
var (
	urlRegex      *regexp.Regexp
	shortURLRegex *regexp.Regexp
)

// errors
var (
	InvalidID   = errors.New("invalid ID")
	NotValidURL = errors.New("not valid URL")
)

type (
	LinkID   uint
	MainUrl  string
	ShortUrl string
)

const (
	alphabets       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	alphabetsLength = 62
	shortURLLength  = 8
)

func init() {
	urlRegex = regexp.MustCompile(`((([A-Za-z]{3,9}:(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[A-Za-z0-9.-]+|(?:www.|[-;:&=\+\$,\w]+@)[A-Za-z0-9.-]+)((?:\/[\+~%\/.\w-_]*)?\??(?:[-\+=&;%@.\w_]*)#?(?:[.\!\/\\w]*))?)`)
	shortURLRegex = regexp.MustCompile(`^[a-zA-Z0-9]{8}$`)
}

type Link struct {
	ID    LinkID
	Url   MainUrl
	Short ShortUrl
}

func (l *Link) UrlValidation() bool {
	return urlRegex.MatchString(string(l.Url))
}

func (l *Link) ShortUrlValidation() bool {
	return shortURLRegex.MatchString(string(l.Short))
}

func (l *Link) MakeShort() error {
	id := uint(l.ID)

	if id <= 0 {
		return InvalidID
	}

	if !l.UrlValidation() {
		return NotValidURL
	}

	chars := []rune(alphabets)
	var short string

	for id > 0 {
		short += string(chars[id%alphabetsLength])
		id = id / alphabetsLength
	}

	short = expandURLLength(short)
	l.Short = ShortUrl(short)
	return nil
}

func expandURLLength(url string) string {
	var shortURL = ""
	var diff = shortURLLength - utf8.RuneCountInString(url)
	for i := 0; i < diff; i++ {
		shortURL += "M"
	}
	shortURL += url
	return shortURL
}
