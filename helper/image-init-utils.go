package helper

import (
	"image"
	"log"
	"regexp"
)

var regexpUrlParse *regexp.Regexp

var noImg *image.RGBA

func init() {
	var err error

	regexpUrlParse, err = regexp.Compile("[a-z0-9]{32}")
	if err != nil {
		log.Fatalln("regexpUrlParse:", err)
	}

	noImg = image.NewRGBA(image.Rect(0, 0, 400, 400))

}
