package service

import (
	"context"
	"github.com/mojtabamovahedi/url-shorter/internal/model"
	"github.com/mojtabamovahedi/url-shorter/internal/repository"
)

type LinkService interface {
	CreateLink(ctx context.Context, link model.Link) (model.LinkID, error)
	UpdateLink(ctx context.Context, link model.Link) (model.LinkID, error)
	FindLinkByUrl(ctx context.Context, url model.MainUrl) (*model.Link, error)
	FindLinkByShortUrl(ctx context.Context, short model.ShortUrl) (*model.Link, error)
}

type linkService struct {
	repo repo.LinkRepo
}

func NewLinkService(repo repo.LinkRepo) LinkService {
	return &linkService{repo: repo}
}

func (l *linkService) CreateLink(ctx context.Context, link model.Link) (model.LinkID, error) {
	return l.repo.Create(ctx, link)
}

func (l *linkService) UpdateLink(ctx context.Context, link model.Link) (model.LinkID, error) {
	return l.repo.Update(ctx, link)
}

func (l *linkService) FindLinkByUrl(ctx context.Context, url model.MainUrl) (*model.Link, error) {
	return l.repo.FindByUrl(ctx, url)
}

func (l *linkService) FindLinkByShortUrl(ctx context.Context, short model.ShortUrl) (*model.Link, error) {
	return l.repo.FindByShortUrl(ctx, short)
}
