package user

import (
	"context"
	"log"

	db "github.com/dkgv/dislikes/generated/sql"
)

type UserRepo interface {
	FindByID(ctx context.Context, id string) (db.OpenYoutubeDislikesUser, error)
	Insert(ctx context.Context, id string) error
}

type DislikeRepo interface {
	Insert(ctx context.Context, videoID string, userID string) error
	Delete(ctx context.Context, videoID string, userID string) error
	FindByID(ctx context.Context, videoID string, userID string) (db.OpenYoutubeDislikesDislike, error)
}

type LikeRepo interface {
	Insert(ctx context.Context, videoID string, userID string) error
	Delete(ctx context.Context, videoID string, userID string) error
	FindByID(ctx context.Context, videoID string, userID string) (db.OpenYoutubeDislikesLike, error)
}

type Service struct {
	dislikeRepo DislikeRepo
	likeRepo    LikeRepo
	userRepo    UserRepo
}

func New(userRepo UserRepo, likeRepo LikeRepo, dislikeRepo DislikeRepo) *Service {
	return &Service{
		userRepo:    userRepo,
		dislikeRepo: dislikeRepo,
		likeRepo:    likeRepo,
	}
}

func (s *Service) AddDislike(ctx context.Context, videoID string, userID string) error {
	err := s.userRepo.Insert(ctx, userID)
	if err != nil {
		return err
	}

	return s.dislikeRepo.Insert(ctx, videoID, userID)
}

func (s *Service) RemoveDislike(ctx context.Context, videoID string, userID string) error {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil || user == (db.OpenYoutubeDislikesUser{}) {
		return err
	}

	return s.dislikeRepo.Delete(ctx, videoID, userID)
}

func (s *Service) AddLike(ctx context.Context, videoID string, userID string) error {
	err := s.userRepo.Insert(ctx, userID)
	if err != nil {
		return err
	}

	return s.likeRepo.Insert(ctx, videoID, userID)
}

func (s *Service) RemoveLike(ctx context.Context, videoID string, userID string) error {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil || user == (db.OpenYoutubeDislikesUser{}) {
		return err
	}

	return s.likeRepo.Delete(ctx, videoID, userID)
}

func (s *Service) HasDislikedVideo(ctx context.Context, videoID string, userID string) (bool, error) {
	err := s.userRepo.Insert(ctx, userID)
	if err != nil {
		return false, err
	}

	_, err = s.dislikeRepo.FindByID(ctx, videoID, userID)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return false, nil
	}

	return err == nil, err
}

func (s *Service) HasLikedVideo(ctx context.Context, videoID string, userID string) (bool, error) {
	err := s.userRepo.Insert(ctx, userID)
	if err != nil {
		log.Printf("insert %v", err)
		return false, err
	}

	_, err = s.likeRepo.FindByID(ctx, videoID, userID)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return false, nil
	}

	return err == nil, err
}
