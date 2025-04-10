package repo

import (
	"context"
	"errors"
	"github.com/mojtabamovahedi/url-shorter/internal/model"
	"github.com/mojtabamovahedi/url-shorter/internal/repository/mapper"
	"github.com/mojtabamovahedi/url-shorter/internal/repository/types"
	"github.com/mojtabamovahedi/url-shorter/pkg/cache"
	"gorm.io/gorm"
)

var (
	LinkNotFoundErr = errors.New("link not found")
)

type LinkRepo interface {
	Create(ctx context.Context, link model.Link) (model.LinkID, error)
	Update(ctx context.Context, link model.Link) (model.LinkID, error)
	FindByUrl(ctx context.Context, url model.MainUrl) (*model.Link, error)
	FindByShortUrl(ctx context.Context, short model.ShortUrl) (*model.Link, error)
}

type linkRepo struct {
	db *gorm.DB
}

func NewLinkRepo(db *gorm.DB, provider cache.Provider) LinkRepo {
	repo := &linkRepo{
		db: db,
	}

	if provider == nil {
		return repo
	}

	return NewUserCachedRepo(repo, provider)
}

func (l *linkRepo) Create(ctx context.Context, link model.Link) (model.LinkID, error) {
	rawLink := mapper.LinkDomain2Storage(link)
	return model.LinkID(rawLink.Id), l.db.Table("links").WithContext(ctx).Create(&rawLink).Error
}

func (l *linkRepo) Update(ctx context.Context, link model.Link) (model.LinkID, error) {
	rawLink := mapper.LinkDomain2Storage(link)
	return model.LinkID(rawLink.Id), l.db.Table("links").WithContext(ctx).Where("id = ?", rawLink.Id).Updates(&rawLink).Error
}

func (l *linkRepo) FindByUrl(ctx context.Context, url model.MainUrl) (*model.Link, error) {
	var rawLink types.Link
	err := l.db.Table("links").WithContext(ctx).Where("url = ?", url).First(&rawLink).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, LinkNotFoundErr
	}

	link := mapper.LinkStorage2Domain(rawLink)

	return &link, nil
}

func (l *linkRepo) FindByShortUrl(ctx context.Context, short model.ShortUrl) (*model.Link, error) {
	var rawLink types.Link
	err := l.db.Table("links").WithContext(ctx).Where("short = ?", short).First(&rawLink).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, LinkNotFoundErr
	}

	link := mapper.LinkStorage2Domain(rawLink)
	return &link, nil
}
