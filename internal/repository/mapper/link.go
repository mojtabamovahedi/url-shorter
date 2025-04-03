package mapper

import (
	"github.com/mojtabamovahedi/url-shorter/internal/model"
	"github.com/mojtabamovahedi/url-shorter/internal/repository/types"
)

func LinkStorage2Domain(url types.Link) model.Link {
	return model.Link{
		ID:    model.LinkID(url.Id),
		Url:   model.MainUrl(url.Url),
		Short: model.ShortUrl(url.Short),
	}
}

func LinkDomain2Storage(url model.Link) types.Link {
	return types.Link{
		Id:    uint(url.ID),
		Url:   string(url.Url),
		Short: string(url.Short),
	}
}
