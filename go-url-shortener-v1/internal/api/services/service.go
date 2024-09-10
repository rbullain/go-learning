package services

type ShortenerService interface {
	// GetUrl get the original url from stored shortened url.
	// It returns
	GetUrl(shortUrl string) (string, error)

	// InsertUrl shortens an url and stores it.
	// It returns the shortened url along with an error if the operation fails.
	InsertUrl(longUrl string) (string, error)
}
