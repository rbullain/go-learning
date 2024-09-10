package services

type dbShortenerService struct {
}

var _ ShortenerService = &dbShortenerService{}

func NewDBShortenerService() ShortenerService {
	return &dbShortenerService{}
}

func (s dbShortenerService) GetUrl(shortUrl string) (string, error) {
	panic("implement me")
}

func (s dbShortenerService) InsertUrl(longUrl string) (string, error) {
	panic("implement me")
}
