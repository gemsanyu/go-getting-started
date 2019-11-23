package profile

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 21:51
**/

type uploadAvatarRequest struct {
	UserID string
	Avatar string `json:"avatar"`
}

func (ar uploadAvatarRequest) GetUserID() string {
	return ar.UserID
}

func (ar uploadAvatarRequest) GetAvatar() string {
	return ar.Avatar
}
