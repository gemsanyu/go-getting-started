package profile

import (
	"fmt"

	"github.com/heroku/go-getting-started/loclib"

	"github.com/heroku/go-getting-started/domain/repository"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 21:37
**/

type UploadAvatarUsecase interface {
	UploadAvatar(UploadAvatarRequest) (error, interface{})
}

type UploadAvatarRequest interface {
	GetUserID() string
	GetAvatar() string
}

type uploadAvatarUsecase struct {
	picRepo     repository.PictureRepository
	profileRepo repository.ProfileRepository
	userRepo    repository.UserRepository
}

func (au uploadAvatarUsecase) UploadAvatar(req UploadAvatarRequest) (error, interface{}) {

	err, profile := au.profileRepo.FindByUserID(req.GetUserID())
	if err != nil {
		return fmt.Errorf("Profile with user id: %d not found", req.GetUserID()), nil
	}

	err, user := au.userRepo.Find(profile.UserID)
	if err != nil {
		return fmt.Errorf("User id: %d not found from profile id: %d", profile.UserID, profile.ID), nil
	}

	avatarName := "avatar-" + user.Username + "-" + loclib.GenerateUUID()
	err, avatarURL := au.picRepo.Save(req.GetAvatar(), avatarName)
	if err != nil {
		return err, nil
	}

	tx, _ := au.profileRepo.DBConn().Begin()
	profile.AvatarURL = avatarURL
	err = au.profileRepo.Update(profile, tx)
	if err != nil {
		tx.Rollback()
		return err, nil
	}
	err = tx.Commit()
	return err, profile
}

func NewUploadAvatarUsecase(picRepo repository.PictureRepository, profileRepo repository.ProfileRepository, userRepo repository.UserRepository) UploadAvatarUsecase {
	return &uploadAvatarUsecase{picRepo, profileRepo, userRepo}
}
