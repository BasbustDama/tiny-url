package generator

import (
	"math/rand"
	"regexp"
	"time"
)

type Generator interface {
	Generate(url string) (string, error)
	Parse(url string) (bool, error)
}

type urlGenerator struct {
	re     *regexp.Regexp
	length int
}

func New(length int) Generator {
	return &urlGenerator{
		re:     regexp.MustCompile("[a-zA-Z0-9]"),
		length: length,
	}
}

func (*urlGenerator) getCharset() []rune {
	return []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
}

func (generator *urlGenerator) Generate(url string) (string, error) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := generator.getCharset()

	shortUrl := make([]rune, generator.length)
	for index := range shortUrl {
		shortUrl[index] = chars[random.Intn(len(chars))]
	}

	return string(shortUrl), nil
}

func (generator *urlGenerator) Parse(url string) (bool, error) {
	if len(url) != generator.length {
		return false, nil
	}

	if len(generator.re.FindAllString(url, generator.length)) != generator.length {
		return false, nil
	}

	return true, nil
}
