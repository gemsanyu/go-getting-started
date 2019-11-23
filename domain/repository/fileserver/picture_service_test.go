package fileserver_test

import (
	"strings"
	"testing"

	"github.com/heroku/go-getting-started/domain/repository/fileserver"
)

var tests []struct {
	dataBase64            string
	name                  string
	format                string
	expectedFormattedName string
}

func init() {
	tests = []struct {
		dataBase64            string
		name                  string
		format                string
		expectedFormattedName string
	}{
		{dataBase64: dataJpg, name: "img1", format: "jpeg", expectedFormattedName: "img1.jpeg"},
		{dataBase64: dataPng, name: "img2", format: "png", expectedFormattedName: "img2.png"},
		{dataBase64: dataGif, name: "img3", format: "gif", expectedFormattedName: "img3.gif"},
	}
}

func TestSave(t *testing.T) {

	// PREPARATION
	ps := fileserver.NewPictureRepo()

	for _, tc := range tests {
		// ACTION
		err, imgUrl := ps.Save(tc.dataBase64, tc.name)
		if err != nil {
			t.Error(err)
		}

		// ASSERTION
		lastDirIdx := strings.LastIndex(imgUrl, "/")
		resultFormattedName := imgUrl[lastDirIdx+1:]
		if resultFormattedName != tc.expectedFormattedName {
			t.Errorf("img url is %s, expected name = %s. but got %s", imgUrl, tc.expectedFormattedName, resultFormattedName)
		}

	}
}

const dataJpg = `
/9j/4AAQSkZJRgABAQEAYABgAAD//gA+Q1JFQVRPUjogZ2Qtan
BlZyB2MS4wICh1c2luZyBJSkcgSlBFRyB2ODApLCBkZWZhdWx0
IHF1YWxpdHkK/9sAQwAIBgYHBgUIBwcHCQkICgwUDQwLCwwZEh
MPFB0aHx4dGhwcICQuJyAiLCMcHCg3KSwwMTQ0NB8nOT04Mjwu
MzQy/9sAQwEJCQkMCwwYDQ0YMiEcITIyMjIyMjIyMjIyMjIyMj
IyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIy/8AA
EQgAHgAeAwEiAAIRAQMRAf/EAB8AAAEFAQEBAQEBAAAAAAAAAA
ABAgMEBQYHCAkKC//EALUQAAIBAwMCBAMFBQQEAAABfQECAwAE
EQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGB
kaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hp
anN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7
S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz
9PX29/j5+v/EAB8BAAMBAQEBAQEBAQEAAAAAAAABAgMEBQYHCA
kKC//EALURAAIBAgQEAwQHBQQEAAECdwABAgMRBAUhMQYSQVEH
YXETIjKBCBRCkaGxwQkjM1LwFWJy0QoWJDThJfEXGBkaJicoKS
o1Njc4OTpDREVGR0hJSlNUVVZXWFlaY2RlZmdoaWpzdHV2d3h5
eoKDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5us
LDxMXGx8jJytLT1NXW19jZ2uLj5OXm5+jp6vLz9PX29/j5+v/a
AAwDAQACEQMRAD8A+f6KK6XwrHplsl1qmtW4nsY3itQhGfmkb5
mHusaykejbfxAOaoq3qmny6Vqt3p85BktpWiYjo2DjI9j1Hsaq
UAFdffXWnaNomk6PdaWt5N5X2+djO8ZV5gCq4HpEsR+rGuQooA
6bxU0WqWml6/bweTHcRfZJkDl9s0IVep5OYzEcnkktXM0UUAf/
2Q==
`

const dataPng = `
data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAB4AAAAeBAMAAADJHrORAAAAG1
BMVEUAAAD///9/f38/Pz+fn5+/v78fHx9fX1/f398ZgXavAAAA
CXBIWXMAAA7EAAAOxAGVKw4bAAAAdUlEQVQYlWNgGPaAycTVgC
nFxYDZxV0BxGd1YAgD4oAUBqYAMJ+BwYgjgcFNmYEBxk8BYrVy
BjYHKJ8FxGctM02AqQfxjcwNwOazFjC4gfip5WYQfolqAYjPYF
wCsY8hAeYUmPlY+UVmJmADmaF8tvIyyj0PAMtsEAvsCfZwAAAA
AElFTkSuQmCC
`

const dataGif = `
R0lGODdhHgAeALMAAAAAAP///z8/P5+fn19fX39/fx8fH7+/v9
/f3wAAAAAAAAAAAAAAAAAAAAAAAAAAACwAAAAAHgAeAAAEbBDI
Sau9OOvNu/9gKI5kaZ5oqqLGChhDRQwFQbbU4cKfMPyEwks4LC
AKBcGGJynomC/DwYDTJHOv2AS6IdgkPC53cyAYBNMsZZwRJAlK
sFbukXoFOPG8IxDMhHokOi6DJUF4PnEmBl5mLo8lEQA7
`
