package fileserver

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/heroku/go-getting-started/domain/repository"
)

/**
 * Created by Muhammad Muflih Kholidin
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 * at: 2019-03-09 21:15
**/

const (
	dirPermMode    os.FileMode = 0755
	formatErrorMsg string      = "%s format is not supported for the picture"
	pictureFolder  string      = "pictures"
)

type pictureService struct {
}

// Save an image to url = path+ name+'.imageType'. return url if success.
func (ps pictureService) Save(picEncodedBase64 string, name string) (error, string) {
	err, pictureDir := ps.createPictureDirectory()
	if err != nil {
		return err, ""
	}

	picEncodedBase64 = ps.removePicPrefix(picEncodedBase64)
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(picEncodedBase64))
	img, format, err := image.Decode(reader)
	if err != nil {
		return err, ""
	}

	imgURL := filepath.Join(pictureDir, name+"."+format)
	imgFile, err := os.Create(imgURL)
	if err != nil {
		return err, ""
	}
	defer imgFile.Close()

	err = ps.writePictureToFile(img, format, imgFile)
	if err != nil {
		return err, ""
	}

	return nil, imgURL
}

func (ps pictureService) createPictureDirectory() (error, string) {
	mainDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err, ""
	}
	pictureDirPath := filepath.Join(mainDir, pictureFolder)

	if _, err := os.Stat(pictureDirPath); err != nil {
		os.Mkdir(pictureDirPath, dirPermMode)
	}
	return nil, pictureDirPath
}

// removePicPrefix of 'image/type; base64,'
func (ps pictureService) removePicPrefix(picEncodedBase64 string) string {
	lastPrefixIdx := strings.Index(picEncodedBase64, ",")
	cleanPic := picEncodedBase64
	if lastPrefixIdx > 0 {
		cleanPic = cleanPic[lastPrefixIdx+1:]
	}
	return cleanPic
}

func (ps pictureService) writePictureToFile(img image.Image, format string, imgFile *os.File) error {
	switch format {
	case "png":
		if err := png.Encode(imgFile, img); err != nil {
			return err
		}
	case "jpeg":
		if err := jpeg.Encode(imgFile, img, nil); err != nil {
			return err
		}
	case "gif":
		if err := gif.Encode(imgFile, img, nil); err != nil {
			return err
		}
	default:
		return fmt.Errorf(formatErrorMsg, format)
	}

	return nil
}

func NewPictureRepo() repository.PictureRepository {
	return &pictureService{}
}
