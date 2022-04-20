package imagestorage

import (
	"context"
	"io"
)

type Storage interface {
	Put(ctx context.Context, img Image) (ImageId, error)
	Get(ctx context.Context, id ImageId) (Image, error)
}

type Image interface {
	Content() io.Reader
	Type() ImgType
}

type image struct {
	c io.Reader
	t ImgType
}

func NewImage(content io.Reader, imgType ImgType) (Image, error) {
	return &image{
		c: content,
		t: imgType,
	}, nil
}

func (i *image) Content() io.Reader {
	return i.c
}

func (i *image) Type() ImgType {
	return i.t
}

type ImageId interface {
	Serialize() (string, error)
	Deserialize(string) (ImageId, error)
}

type ImgType uint8

const (
	ImgTypePng ImgType = iota
	ImgTypeJpg
)
