package repository

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 21:12
**/

// PictureRepository for uploading a picture, currently only for avatar
type PictureRepository interface {
	Save(picEncodedBase64 string, name string) (error, string)
}
