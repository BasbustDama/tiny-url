package usecase

import (
	"log/slog"

	"github.com/BasbustDama/tiny-url/internal/errors"
)

type (
	ShortenerUsecase interface {
		Create(longUrl string) (string, error)
		Get(shortUrl string) (string, error)
	}

	ShortenerStorage interface {
		Put(shortUrl, longUrl string) error
		Get(shortUrl string) (string, error)
	}

	ShortenerGenerator interface {
		Generate(longUrl string) (string, error)
		Parse(shortUrl string) (bool, error)
	}
)

type shortenerUsecase struct {
	shortenerStorage   ShortenerStorage
	shortenerGenerator ShortenerGenerator
}

func New(
	shortenerStorage ShortenerStorage,
	shortenerGenerator ShortenerGenerator,
) ShortenerUsecase {
	return &shortenerUsecase{
		shortenerStorage:   shortenerStorage,
		shortenerGenerator: shortenerGenerator,
	}
}

func (usecase *shortenerUsecase) Create(longUrl string) (string, error) {
	shortUrl, err := usecase.shortenerGenerator.Generate(longUrl)
	if err != nil {
		slog.Error(err.Error(), slog.String("url", longUrl))
		return "", errors.ErrorInternal
	}

	err = usecase.shortenerStorage.Put(shortUrl, longUrl)
	if err != nil {
		slog.Error(err.Error(), slog.String("url", longUrl))
		return "", errors.ErrorInternal
	}

	return shortUrl, nil
}

func (usecase *shortenerUsecase) Get(shortUrl string) (string, error) {
	valid, err := usecase.shortenerGenerator.Parse(shortUrl)
	if err != nil {
		slog.Error(err.Error(), slog.String("shortUrl", shortUrl))
		return "", errors.ErrorInternal
	}

	if !valid {
		slog.Info("Url is not valid", slog.String("shortUrl", shortUrl))
		return "", errors.ErrorBadRequest
	}

	longUrl, err := usecase.shortenerStorage.Get(shortUrl)
	if err != nil {
		slog.Error(err.Error(), slog.String("shortUrl", shortUrl))
		return "", errors.ErrorInternal
	}

	if longUrl == "" {
		slog.Warn("empty valid short url", slog.String("shortUrl", shortUrl))
		return "", errors.ErrorNotFound
	}

	return longUrl, nil
}
