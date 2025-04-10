package repo

import (
	"context"
	"github.com/mojtabamovahedi/url-shorter/internal/model"
	"github.com/mojtabamovahedi/url-shorter/pkg/cache"
	"log"
)

type userCachedRepo struct {
	repo     LinkRepo
	provider cache.Provider
	oCacher  *cache.ObjectCacher[*model.Link]
}

func NewUserCachedRepo(r LinkRepo, p cache.Provider) LinkRepo {
	oc := cache.NewJsonObjectCacher[*model.Link](p)
	return &userCachedRepo{
		repo:     r,
		provider: p,
		oCacher:  oc,
	}
}

func (u *userCachedRepo) linkShortKey(sh model.ShortUrl) string {
	return "short." + string(sh)
}

func (u *userCachedRepo) Create(ctx context.Context, link model.Link) (model.LinkID, error) {
	lID, err := u.repo.Create(ctx, link)
	if err != nil {
		return 0, err
	}

	link.ID = lID

	err = u.oCacher.Set(ctx, u.linkShortKey(link.Short), &link)
	if err != nil {
		log.Printf("Error setting cache for link %s: %v", link.Short, err)
	}

	return lID, nil
}

func (u *userCachedRepo) Update(ctx context.Context, link model.Link) (model.LinkID, error) {
	return u.repo.Update(ctx, link)
}

func (u *userCachedRepo) FindByUrl(ctx context.Context, url model.MainUrl) (*model.Link, error) {
	return u.repo.FindByUrl(ctx, url)
}

func (u *userCachedRepo) FindByShortUrl(ctx context.Context, short model.ShortUrl) (*model.Link, error) {
	var (
		link *model.Link
		err  error
		key  = string(short)
	)
	link, err = u.oCacher.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if link != nil {
		return link, nil
	}

	link, err = u.repo.FindByShortUrl(ctx, short)
	if err != nil {
		return nil, err
	}

	err = u.oCacher.Set(ctx, key, link)
	if err != nil {
		log.Printf("Error on cache link with id %d %v", link.ID, err)
	}

	return link, nil
}
