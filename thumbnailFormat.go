package tdlib

import (
	"encoding/json"
	"fmt"
)

// ThumbnailFormat Describes format of the thumbnail
type ThumbnailFormat interface {
	GetThumbnailFormatEnum() ThumbnailFormatEnum
}

// ThumbnailFormatEnum Alias for abstract ThumbnailFormat 'Sub-Classes', used as constant-enum here
type ThumbnailFormatEnum string

// ThumbnailFormat enums
const (
	ThumbnailFormatJpegType  ThumbnailFormatEnum = "thumbnailFormatJpeg"
	ThumbnailFormatPngType   ThumbnailFormatEnum = "thumbnailFormatPng"
	ThumbnailFormatWebpType  ThumbnailFormatEnum = "thumbnailFormatWebp"
	ThumbnailFormatGifType   ThumbnailFormatEnum = "thumbnailFormatGif"
	ThumbnailFormatTgsType   ThumbnailFormatEnum = "thumbnailFormatTgs"
	ThumbnailFormatMpeg4Type ThumbnailFormatEnum = "thumbnailFormatMpeg4"
)

func unmarshalThumbnailFormat(rawMsg *json.RawMessage) (ThumbnailFormat, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch ThumbnailFormatEnum(objMap["@type"].(string)) {
	case ThumbnailFormatJpegType:
		var thumbnailFormatJpeg ThumbnailFormatJpeg
		err := json.Unmarshal(*rawMsg, &thumbnailFormatJpeg)
		return &thumbnailFormatJpeg, err

	case ThumbnailFormatPngType:
		var thumbnailFormatPng ThumbnailFormatPng
		err := json.Unmarshal(*rawMsg, &thumbnailFormatPng)
		return &thumbnailFormatPng, err

	case ThumbnailFormatWebpType:
		var thumbnailFormatWebp ThumbnailFormatWebp
		err := json.Unmarshal(*rawMsg, &thumbnailFormatWebp)
		return &thumbnailFormatWebp, err

	case ThumbnailFormatGifType:
		var thumbnailFormatGif ThumbnailFormatGif
		err := json.Unmarshal(*rawMsg, &thumbnailFormatGif)
		return &thumbnailFormatGif, err

	case ThumbnailFormatTgsType:
		var thumbnailFormatTgs ThumbnailFormatTgs
		err := json.Unmarshal(*rawMsg, &thumbnailFormatTgs)
		return &thumbnailFormatTgs, err

	case ThumbnailFormatMpeg4Type:
		var thumbnailFormatMpeg4 ThumbnailFormatMpeg4
		err := json.Unmarshal(*rawMsg, &thumbnailFormatMpeg4)
		return &thumbnailFormatMpeg4, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// ThumbnailFormatJpeg The thumbnail is in JPEG format
type ThumbnailFormatJpeg struct {
	tdCommon
}

// MessageType return the string telegram-type of ThumbnailFormatJpeg
func (thumbnailFormatJpeg *ThumbnailFormatJpeg) MessageType() string {
	return "thumbnailFormatJpeg"
}

// NewThumbnailFormatJpeg creates a new ThumbnailFormatJpeg
//
func NewThumbnailFormatJpeg() *ThumbnailFormatJpeg {
	thumbnailFormatJpegTemp := ThumbnailFormatJpeg{
		tdCommon: tdCommon{Type: "thumbnailFormatJpeg"},
	}

	return &thumbnailFormatJpegTemp
}

// GetThumbnailFormatEnum return the enum type of this object
func (thumbnailFormatJpeg *ThumbnailFormatJpeg) GetThumbnailFormatEnum() ThumbnailFormatEnum {
	return ThumbnailFormatJpegType
}

// ThumbnailFormatPng The thumbnail is in PNG format. It will be used only for background patterns
type ThumbnailFormatPng struct {
	tdCommon
}

// MessageType return the string telegram-type of ThumbnailFormatPng
func (thumbnailFormatPng *ThumbnailFormatPng) MessageType() string {
	return "thumbnailFormatPng"
}

// NewThumbnailFormatPng creates a new ThumbnailFormatPng
//
func NewThumbnailFormatPng() *ThumbnailFormatPng {
	thumbnailFormatPngTemp := ThumbnailFormatPng{
		tdCommon: tdCommon{Type: "thumbnailFormatPng"},
	}

	return &thumbnailFormatPngTemp
}

// GetThumbnailFormatEnum return the enum type of this object
func (thumbnailFormatPng *ThumbnailFormatPng) GetThumbnailFormatEnum() ThumbnailFormatEnum {
	return ThumbnailFormatPngType
}

// ThumbnailFormatWebp The thumbnail is in WEBP format. It will be used only for some stickers
type ThumbnailFormatWebp struct {
	tdCommon
}

// MessageType return the string telegram-type of ThumbnailFormatWebp
func (thumbnailFormatWebp *ThumbnailFormatWebp) MessageType() string {
	return "thumbnailFormatWebp"
}

// NewThumbnailFormatWebp creates a new ThumbnailFormatWebp
//
func NewThumbnailFormatWebp() *ThumbnailFormatWebp {
	thumbnailFormatWebpTemp := ThumbnailFormatWebp{
		tdCommon: tdCommon{Type: "thumbnailFormatWebp"},
	}

	return &thumbnailFormatWebpTemp
}

// GetThumbnailFormatEnum return the enum type of this object
func (thumbnailFormatWebp *ThumbnailFormatWebp) GetThumbnailFormatEnum() ThumbnailFormatEnum {
	return ThumbnailFormatWebpType
}

// ThumbnailFormatGif The thumbnail is in static GIF format. It will be used only for some bot inline results
type ThumbnailFormatGif struct {
	tdCommon
}

// MessageType return the string telegram-type of ThumbnailFormatGif
func (thumbnailFormatGif *ThumbnailFormatGif) MessageType() string {
	return "thumbnailFormatGif"
}

// NewThumbnailFormatGif creates a new ThumbnailFormatGif
//
func NewThumbnailFormatGif() *ThumbnailFormatGif {
	thumbnailFormatGifTemp := ThumbnailFormatGif{
		tdCommon: tdCommon{Type: "thumbnailFormatGif"},
	}

	return &thumbnailFormatGifTemp
}

// GetThumbnailFormatEnum return the enum type of this object
func (thumbnailFormatGif *ThumbnailFormatGif) GetThumbnailFormatEnum() ThumbnailFormatEnum {
	return ThumbnailFormatGifType
}

// ThumbnailFormatTgs The thumbnail is in TGS format. It will be used only for animated sticker sets
type ThumbnailFormatTgs struct {
	tdCommon
}

// MessageType return the string telegram-type of ThumbnailFormatTgs
func (thumbnailFormatTgs *ThumbnailFormatTgs) MessageType() string {
	return "thumbnailFormatTgs"
}

// NewThumbnailFormatTgs creates a new ThumbnailFormatTgs
//
func NewThumbnailFormatTgs() *ThumbnailFormatTgs {
	thumbnailFormatTgsTemp := ThumbnailFormatTgs{
		tdCommon: tdCommon{Type: "thumbnailFormatTgs"},
	}

	return &thumbnailFormatTgsTemp
}

// GetThumbnailFormatEnum return the enum type of this object
func (thumbnailFormatTgs *ThumbnailFormatTgs) GetThumbnailFormatEnum() ThumbnailFormatEnum {
	return ThumbnailFormatTgsType
}

// ThumbnailFormatMpeg4 The thumbnail is in MPEG4 format. It will be used only for some animations and videos
type ThumbnailFormatMpeg4 struct {
	tdCommon
}

// MessageType return the string telegram-type of ThumbnailFormatMpeg4
func (thumbnailFormatMpeg4 *ThumbnailFormatMpeg4) MessageType() string {
	return "thumbnailFormatMpeg4"
}

// NewThumbnailFormatMpeg4 creates a new ThumbnailFormatMpeg4
//
func NewThumbnailFormatMpeg4() *ThumbnailFormatMpeg4 {
	thumbnailFormatMpeg4Temp := ThumbnailFormatMpeg4{
		tdCommon: tdCommon{Type: "thumbnailFormatMpeg4"},
	}

	return &thumbnailFormatMpeg4Temp
}

// GetThumbnailFormatEnum return the enum type of this object
func (thumbnailFormatMpeg4 *ThumbnailFormatMpeg4) GetThumbnailFormatEnum() ThumbnailFormatEnum {
	return ThumbnailFormatMpeg4Type
}
