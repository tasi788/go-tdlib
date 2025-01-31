package tdlib

import (
	"encoding/json"
	"fmt"
)

// RichText Describes a text object inside an instant-view web page
type RichText interface {
	GetRichTextEnum() RichTextEnum
}

// RichTextEnum Alias for abstract RichText 'Sub-Classes', used as constant-enum here
type RichTextEnum string

// RichText enums
const (
	RichTextPlainType         RichTextEnum = "richTextPlain"
	RichTextBoldType          RichTextEnum = "richTextBold"
	RichTextItalicType        RichTextEnum = "richTextItalic"
	RichTextUnderlineType     RichTextEnum = "richTextUnderline"
	RichTextStrikethroughType RichTextEnum = "richTextStrikethrough"
	RichTextFixedType         RichTextEnum = "richTextFixed"
	RichTextUrlType           RichTextEnum = "richTextUrl"
	RichTextEmailAddressType  RichTextEnum = "richTextEmailAddress"
	RichTextSubscriptType     RichTextEnum = "richTextSubscript"
	RichTextSuperscriptType   RichTextEnum = "richTextSuperscript"
	RichTextMarkedType        RichTextEnum = "richTextMarked"
	RichTextPhoneNumberType   RichTextEnum = "richTextPhoneNumber"
	RichTextIconType          RichTextEnum = "richTextIcon"
	RichTextReferenceType     RichTextEnum = "richTextReference"
	RichTextAnchorType        RichTextEnum = "richTextAnchor"
	RichTextAnchorLinkType    RichTextEnum = "richTextAnchorLink"
	RichTextsType             RichTextEnum = "richTexts"
)

func unmarshalRichText(rawMsg *json.RawMessage) (RichText, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch RichTextEnum(objMap["@type"].(string)) {
	case RichTextPlainType:
		var richTextPlain RichTextPlain
		err := json.Unmarshal(*rawMsg, &richTextPlain)
		return &richTextPlain, err

	case RichTextBoldType:
		var richTextBold RichTextBold
		err := json.Unmarshal(*rawMsg, &richTextBold)
		return &richTextBold, err

	case RichTextItalicType:
		var richTextItalic RichTextItalic
		err := json.Unmarshal(*rawMsg, &richTextItalic)
		return &richTextItalic, err

	case RichTextUnderlineType:
		var richTextUnderline RichTextUnderline
		err := json.Unmarshal(*rawMsg, &richTextUnderline)
		return &richTextUnderline, err

	case RichTextStrikethroughType:
		var richTextStrikethrough RichTextStrikethrough
		err := json.Unmarshal(*rawMsg, &richTextStrikethrough)
		return &richTextStrikethrough, err

	case RichTextFixedType:
		var richTextFixed RichTextFixed
		err := json.Unmarshal(*rawMsg, &richTextFixed)
		return &richTextFixed, err

	case RichTextUrlType:
		var richTextUrl RichTextUrl
		err := json.Unmarshal(*rawMsg, &richTextUrl)
		return &richTextUrl, err

	case RichTextEmailAddressType:
		var richTextEmailAddress RichTextEmailAddress
		err := json.Unmarshal(*rawMsg, &richTextEmailAddress)
		return &richTextEmailAddress, err

	case RichTextSubscriptType:
		var richTextSubscript RichTextSubscript
		err := json.Unmarshal(*rawMsg, &richTextSubscript)
		return &richTextSubscript, err

	case RichTextSuperscriptType:
		var richTextSuperscript RichTextSuperscript
		err := json.Unmarshal(*rawMsg, &richTextSuperscript)
		return &richTextSuperscript, err

	case RichTextMarkedType:
		var richTextMarked RichTextMarked
		err := json.Unmarshal(*rawMsg, &richTextMarked)
		return &richTextMarked, err

	case RichTextPhoneNumberType:
		var richTextPhoneNumber RichTextPhoneNumber
		err := json.Unmarshal(*rawMsg, &richTextPhoneNumber)
		return &richTextPhoneNumber, err

	case RichTextIconType:
		var richTextIcon RichTextIcon
		err := json.Unmarshal(*rawMsg, &richTextIcon)
		return &richTextIcon, err

	case RichTextReferenceType:
		var richTextReference RichTextReference
		err := json.Unmarshal(*rawMsg, &richTextReference)
		return &richTextReference, err

	case RichTextAnchorType:
		var richTextAnchor RichTextAnchor
		err := json.Unmarshal(*rawMsg, &richTextAnchor)
		return &richTextAnchor, err

	case RichTextAnchorLinkType:
		var richTextAnchorLink RichTextAnchorLink
		err := json.Unmarshal(*rawMsg, &richTextAnchorLink)
		return &richTextAnchorLink, err

	case RichTextsType:
		var richTexts RichTexts
		err := json.Unmarshal(*rawMsg, &richTexts)
		return &richTexts, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// RichTextPlain A plain text
type RichTextPlain struct {
	tdCommon
	Text string `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextPlain
func (richTextPlain *RichTextPlain) MessageType() string {
	return "richTextPlain"
}

// NewRichTextPlain creates a new RichTextPlain
//
// @param text Text
func NewRichTextPlain(text string) *RichTextPlain {
	richTextPlainTemp := RichTextPlain{
		tdCommon: tdCommon{Type: "richTextPlain"},
		Text:     text,
	}

	return &richTextPlainTemp
}

// GetRichTextEnum return the enum type of this object
func (richTextPlain *RichTextPlain) GetRichTextEnum() RichTextEnum {
	return RichTextPlainType
}

// RichTextBold A bold rich text
type RichTextBold struct {
	tdCommon
	Text RichText `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextBold
func (richTextBold *RichTextBold) MessageType() string {
	return "richTextBold"
}

// NewRichTextBold creates a new RichTextBold
//
// @param text Text
func NewRichTextBold(text RichText) *RichTextBold {
	richTextBoldTemp := RichTextBold{
		tdCommon: tdCommon{Type: "richTextBold"},
		Text:     text,
	}

	return &richTextBoldTemp
}

// UnmarshalJSON unmarshal to json
func (richTextBold *RichTextBold) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	richTextBold.tdCommon = tempObj.tdCommon

	fieldText, _ := unmarshalRichText(objMap["text"])
	richTextBold.Text = fieldText

	return nil
}

// GetRichTextEnum return the enum type of this object
func (richTextBold *RichTextBold) GetRichTextEnum() RichTextEnum {
	return RichTextBoldType
}

// RichTextItalic An italicized rich text
type RichTextItalic struct {
	tdCommon
	Text RichText `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextItalic
func (richTextItalic *RichTextItalic) MessageType() string {
	return "richTextItalic"
}

// NewRichTextItalic creates a new RichTextItalic
//
// @param text Text
func NewRichTextItalic(text RichText) *RichTextItalic {
	richTextItalicTemp := RichTextItalic{
		tdCommon: tdCommon{Type: "richTextItalic"},
		Text:     text,
	}

	return &richTextItalicTemp
}

// UnmarshalJSON unmarshal to json
func (richTextItalic *RichTextItalic) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	richTextItalic.tdCommon = tempObj.tdCommon

	fieldText, _ := unmarshalRichText(objMap["text"])
	richTextItalic.Text = fieldText

	return nil
}

// GetRichTextEnum return the enum type of this object
func (richTextItalic *RichTextItalic) GetRichTextEnum() RichTextEnum {
	return RichTextItalicType
}

// RichTextUnderline An underlined rich text
type RichTextUnderline struct {
	tdCommon
	Text RichText `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextUnderline
func (richTextUnderline *RichTextUnderline) MessageType() string {
	return "richTextUnderline"
}

// NewRichTextUnderline creates a new RichTextUnderline
//
// @param text Text
func NewRichTextUnderline(text RichText) *RichTextUnderline {
	richTextUnderlineTemp := RichTextUnderline{
		tdCommon: tdCommon{Type: "richTextUnderline"},
		Text:     text,
	}

	return &richTextUnderlineTemp
}

// UnmarshalJSON unmarshal to json
func (richTextUnderline *RichTextUnderline) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	richTextUnderline.tdCommon = tempObj.tdCommon

	fieldText, _ := unmarshalRichText(objMap["text"])
	richTextUnderline.Text = fieldText

	return nil
}

// GetRichTextEnum return the enum type of this object
func (richTextUnderline *RichTextUnderline) GetRichTextEnum() RichTextEnum {
	return RichTextUnderlineType
}

// RichTextStrikethrough A strikethrough rich text
type RichTextStrikethrough struct {
	tdCommon
	Text RichText `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextStrikethrough
func (richTextStrikethrough *RichTextStrikethrough) MessageType() string {
	return "richTextStrikethrough"
}

// NewRichTextStrikethrough creates a new RichTextStrikethrough
//
// @param text Text
func NewRichTextStrikethrough(text RichText) *RichTextStrikethrough {
	richTextStrikethroughTemp := RichTextStrikethrough{
		tdCommon: tdCommon{Type: "richTextStrikethrough"},
		Text:     text,
	}

	return &richTextStrikethroughTemp
}

// UnmarshalJSON unmarshal to json
func (richTextStrikethrough *RichTextStrikethrough) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	richTextStrikethrough.tdCommon = tempObj.tdCommon

	fieldText, _ := unmarshalRichText(objMap["text"])
	richTextStrikethrough.Text = fieldText

	return nil
}

// GetRichTextEnum return the enum type of this object
func (richTextStrikethrough *RichTextStrikethrough) GetRichTextEnum() RichTextEnum {
	return RichTextStrikethroughType
}

// RichTextFixed A fixed-width rich text
type RichTextFixed struct {
	tdCommon
	Text RichText `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextFixed
func (richTextFixed *RichTextFixed) MessageType() string {
	return "richTextFixed"
}

// NewRichTextFixed creates a new RichTextFixed
//
// @param text Text
func NewRichTextFixed(text RichText) *RichTextFixed {
	richTextFixedTemp := RichTextFixed{
		tdCommon: tdCommon{Type: "richTextFixed"},
		Text:     text,
	}

	return &richTextFixedTemp
}

// UnmarshalJSON unmarshal to json
func (richTextFixed *RichTextFixed) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	richTextFixed.tdCommon = tempObj.tdCommon

	fieldText, _ := unmarshalRichText(objMap["text"])
	richTextFixed.Text = fieldText

	return nil
}

// GetRichTextEnum return the enum type of this object
func (richTextFixed *RichTextFixed) GetRichTextEnum() RichTextEnum {
	return RichTextFixedType
}

// RichTextUrl A rich text URL link
type RichTextUrl struct {
	tdCommon
	Text     RichText `json:"text"`      // Text
	Url      string   `json:"url"`       // URL
	IsCached bool     `json:"is_cached"` // True, if the URL has cached instant view server-side
}

// MessageType return the string telegram-type of RichTextUrl
func (richTextUrl *RichTextUrl) MessageType() string {
	return "richTextUrl"
}

// NewRichTextUrl creates a new RichTextUrl
//
// @param text Text
// @param url URL
// @param isCached True, if the URL has cached instant view server-side
func NewRichTextUrl(text RichText, url string, isCached bool) *RichTextUrl {
	richTextUrlTemp := RichTextUrl{
		tdCommon: tdCommon{Type: "richTextUrl"},
		Text:     text,
		Url:      url,
		IsCached: isCached,
	}

	return &richTextUrlTemp
}

// UnmarshalJSON unmarshal to json
func (richTextUrl *RichTextUrl) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Url      string `json:"url"`       // URL
		IsCached bool   `json:"is_cached"` // True, if the URL has cached instant view server-side
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	richTextUrl.tdCommon = tempObj.tdCommon
	richTextUrl.Url = tempObj.Url
	richTextUrl.IsCached = tempObj.IsCached

	fieldText, _ := unmarshalRichText(objMap["text"])
	richTextUrl.Text = fieldText

	return nil
}

// GetRichTextEnum return the enum type of this object
func (richTextUrl *RichTextUrl) GetRichTextEnum() RichTextEnum {
	return RichTextUrlType
}

// RichTextEmailAddress A rich text email link
type RichTextEmailAddress struct {
	tdCommon
	Text         RichText `json:"text"`          // Text
	EmailAddress string   `json:"email_address"` // Email address
}

// MessageType return the string telegram-type of RichTextEmailAddress
func (richTextEmailAddress *RichTextEmailAddress) MessageType() string {
	return "richTextEmailAddress"
}

// NewRichTextEmailAddress creates a new RichTextEmailAddress
//
// @param text Text
// @param emailAddress Email address
func NewRichTextEmailAddress(text RichText, emailAddress string) *RichTextEmailAddress {
	richTextEmailAddressTemp := RichTextEmailAddress{
		tdCommon:     tdCommon{Type: "richTextEmailAddress"},
		Text:         text,
		EmailAddress: emailAddress,
	}

	return &richTextEmailAddressTemp
}

// UnmarshalJSON unmarshal to json
func (richTextEmailAddress *RichTextEmailAddress) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		EmailAddress string `json:"email_address"` // Email address
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	richTextEmailAddress.tdCommon = tempObj.tdCommon
	richTextEmailAddress.EmailAddress = tempObj.EmailAddress

	fieldText, _ := unmarshalRichText(objMap["text"])
	richTextEmailAddress.Text = fieldText

	return nil
}

// GetRichTextEnum return the enum type of this object
func (richTextEmailAddress *RichTextEmailAddress) GetRichTextEnum() RichTextEnum {
	return RichTextEmailAddressType
}

// RichTextSubscript A subscript rich text
type RichTextSubscript struct {
	tdCommon
	Text RichText `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextSubscript
func (richTextSubscript *RichTextSubscript) MessageType() string {
	return "richTextSubscript"
}

// NewRichTextSubscript creates a new RichTextSubscript
//
// @param text Text
func NewRichTextSubscript(text RichText) *RichTextSubscript {
	richTextSubscriptTemp := RichTextSubscript{
		tdCommon: tdCommon{Type: "richTextSubscript"},
		Text:     text,
	}

	return &richTextSubscriptTemp
}

// UnmarshalJSON unmarshal to json
func (richTextSubscript *RichTextSubscript) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	richTextSubscript.tdCommon = tempObj.tdCommon

	fieldText, _ := unmarshalRichText(objMap["text"])
	richTextSubscript.Text = fieldText

	return nil
}

// GetRichTextEnum return the enum type of this object
func (richTextSubscript *RichTextSubscript) GetRichTextEnum() RichTextEnum {
	return RichTextSubscriptType
}

// RichTextSuperscript A superscript rich text
type RichTextSuperscript struct {
	tdCommon
	Text RichText `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextSuperscript
func (richTextSuperscript *RichTextSuperscript) MessageType() string {
	return "richTextSuperscript"
}

// NewRichTextSuperscript creates a new RichTextSuperscript
//
// @param text Text
func NewRichTextSuperscript(text RichText) *RichTextSuperscript {
	richTextSuperscriptTemp := RichTextSuperscript{
		tdCommon: tdCommon{Type: "richTextSuperscript"},
		Text:     text,
	}

	return &richTextSuperscriptTemp
}

// UnmarshalJSON unmarshal to json
func (richTextSuperscript *RichTextSuperscript) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	richTextSuperscript.tdCommon = tempObj.tdCommon

	fieldText, _ := unmarshalRichText(objMap["text"])
	richTextSuperscript.Text = fieldText

	return nil
}

// GetRichTextEnum return the enum type of this object
func (richTextSuperscript *RichTextSuperscript) GetRichTextEnum() RichTextEnum {
	return RichTextSuperscriptType
}

// RichTextMarked A marked rich text
type RichTextMarked struct {
	tdCommon
	Text RichText `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextMarked
func (richTextMarked *RichTextMarked) MessageType() string {
	return "richTextMarked"
}

// NewRichTextMarked creates a new RichTextMarked
//
// @param text Text
func NewRichTextMarked(text RichText) *RichTextMarked {
	richTextMarkedTemp := RichTextMarked{
		tdCommon: tdCommon{Type: "richTextMarked"},
		Text:     text,
	}

	return &richTextMarkedTemp
}

// UnmarshalJSON unmarshal to json
func (richTextMarked *RichTextMarked) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	richTextMarked.tdCommon = tempObj.tdCommon

	fieldText, _ := unmarshalRichText(objMap["text"])
	richTextMarked.Text = fieldText

	return nil
}

// GetRichTextEnum return the enum type of this object
func (richTextMarked *RichTextMarked) GetRichTextEnum() RichTextEnum {
	return RichTextMarkedType
}

// RichTextPhoneNumber A rich text phone number
type RichTextPhoneNumber struct {
	tdCommon
	Text        RichText `json:"text"`         // Text
	PhoneNumber string   `json:"phone_number"` // Phone number
}

// MessageType return the string telegram-type of RichTextPhoneNumber
func (richTextPhoneNumber *RichTextPhoneNumber) MessageType() string {
	return "richTextPhoneNumber"
}

// NewRichTextPhoneNumber creates a new RichTextPhoneNumber
//
// @param text Text
// @param phoneNumber Phone number
func NewRichTextPhoneNumber(text RichText, phoneNumber string) *RichTextPhoneNumber {
	richTextPhoneNumberTemp := RichTextPhoneNumber{
		tdCommon:    tdCommon{Type: "richTextPhoneNumber"},
		Text:        text,
		PhoneNumber: phoneNumber,
	}

	return &richTextPhoneNumberTemp
}

// UnmarshalJSON unmarshal to json
func (richTextPhoneNumber *RichTextPhoneNumber) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		PhoneNumber string `json:"phone_number"` // Phone number
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	richTextPhoneNumber.tdCommon = tempObj.tdCommon
	richTextPhoneNumber.PhoneNumber = tempObj.PhoneNumber

	fieldText, _ := unmarshalRichText(objMap["text"])
	richTextPhoneNumber.Text = fieldText

	return nil
}

// GetRichTextEnum return the enum type of this object
func (richTextPhoneNumber *RichTextPhoneNumber) GetRichTextEnum() RichTextEnum {
	return RichTextPhoneNumberType
}

// RichTextIcon A small image inside the text
type RichTextIcon struct {
	tdCommon
	Document *Document `json:"document"` // The image represented as a document. The image can be in GIF, JPEG or PNG format
	Width    int32     `json:"width"`    // Width of a bounding box in which the image must be shown; 0 if unknown
	Height   int32     `json:"height"`   // Height of a bounding box in which the image must be shown; 0 if unknown
}

// MessageType return the string telegram-type of RichTextIcon
func (richTextIcon *RichTextIcon) MessageType() string {
	return "richTextIcon"
}

// NewRichTextIcon creates a new RichTextIcon
//
// @param document The image represented as a document. The image can be in GIF, JPEG or PNG format
// @param width Width of a bounding box in which the image must be shown; 0 if unknown
// @param height Height of a bounding box in which the image must be shown; 0 if unknown
func NewRichTextIcon(document *Document, width int32, height int32) *RichTextIcon {
	richTextIconTemp := RichTextIcon{
		tdCommon: tdCommon{Type: "richTextIcon"},
		Document: document,
		Width:    width,
		Height:   height,
	}

	return &richTextIconTemp
}

// GetRichTextEnum return the enum type of this object
func (richTextIcon *RichTextIcon) GetRichTextEnum() RichTextEnum {
	return RichTextIconType
}

// RichTextReference A reference to a richTexts object on the same web page
type RichTextReference struct {
	tdCommon
	Text       RichText `json:"text"`        // The text
	AnchorName string   `json:"anchor_name"` // The name of a richTextAnchor object, which is the first element of the target richTexts object
	Url        string   `json:"url"`         // An HTTP URL, opening the reference
}

// MessageType return the string telegram-type of RichTextReference
func (richTextReference *RichTextReference) MessageType() string {
	return "richTextReference"
}

// NewRichTextReference creates a new RichTextReference
//
// @param text The text
// @param anchorName The name of a richTextAnchor object, which is the first element of the target richTexts object
// @param url An HTTP URL, opening the reference
func NewRichTextReference(text RichText, anchorName string, url string) *RichTextReference {
	richTextReferenceTemp := RichTextReference{
		tdCommon:   tdCommon{Type: "richTextReference"},
		Text:       text,
		AnchorName: anchorName,
		Url:        url,
	}

	return &richTextReferenceTemp
}

// UnmarshalJSON unmarshal to json
func (richTextReference *RichTextReference) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		AnchorName string `json:"anchor_name"` // The name of a richTextAnchor object, which is the first element of the target richTexts object
		Url        string `json:"url"`         // An HTTP URL, opening the reference
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	richTextReference.tdCommon = tempObj.tdCommon
	richTextReference.AnchorName = tempObj.AnchorName
	richTextReference.Url = tempObj.Url

	fieldText, _ := unmarshalRichText(objMap["text"])
	richTextReference.Text = fieldText

	return nil
}

// GetRichTextEnum return the enum type of this object
func (richTextReference *RichTextReference) GetRichTextEnum() RichTextEnum {
	return RichTextReferenceType
}

// RichTextAnchor An anchor
type RichTextAnchor struct {
	tdCommon
	Name string `json:"name"` // Anchor name
}

// MessageType return the string telegram-type of RichTextAnchor
func (richTextAnchor *RichTextAnchor) MessageType() string {
	return "richTextAnchor"
}

// NewRichTextAnchor creates a new RichTextAnchor
//
// @param name Anchor name
func NewRichTextAnchor(name string) *RichTextAnchor {
	richTextAnchorTemp := RichTextAnchor{
		tdCommon: tdCommon{Type: "richTextAnchor"},
		Name:     name,
	}

	return &richTextAnchorTemp
}

// GetRichTextEnum return the enum type of this object
func (richTextAnchor *RichTextAnchor) GetRichTextEnum() RichTextEnum {
	return RichTextAnchorType
}

// RichTextAnchorLink A link to an anchor on the same web page
type RichTextAnchorLink struct {
	tdCommon
	Text       RichText `json:"text"`        // The link text
	AnchorName string   `json:"anchor_name"` // The anchor name. If the name is empty, the link must bring back to top
	Url        string   `json:"url"`         // An HTTP URL, opening the anchor
}

// MessageType return the string telegram-type of RichTextAnchorLink
func (richTextAnchorLink *RichTextAnchorLink) MessageType() string {
	return "richTextAnchorLink"
}

// NewRichTextAnchorLink creates a new RichTextAnchorLink
//
// @param text The link text
// @param anchorName The anchor name. If the name is empty, the link must bring back to top
// @param url An HTTP URL, opening the anchor
func NewRichTextAnchorLink(text RichText, anchorName string, url string) *RichTextAnchorLink {
	richTextAnchorLinkTemp := RichTextAnchorLink{
		tdCommon:   tdCommon{Type: "richTextAnchorLink"},
		Text:       text,
		AnchorName: anchorName,
		Url:        url,
	}

	return &richTextAnchorLinkTemp
}

// UnmarshalJSON unmarshal to json
func (richTextAnchorLink *RichTextAnchorLink) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		AnchorName string `json:"anchor_name"` // The anchor name. If the name is empty, the link must bring back to top
		Url        string `json:"url"`         // An HTTP URL, opening the anchor
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	richTextAnchorLink.tdCommon = tempObj.tdCommon
	richTextAnchorLink.AnchorName = tempObj.AnchorName
	richTextAnchorLink.Url = tempObj.Url

	fieldText, _ := unmarshalRichText(objMap["text"])
	richTextAnchorLink.Text = fieldText

	return nil
}

// GetRichTextEnum return the enum type of this object
func (richTextAnchorLink *RichTextAnchorLink) GetRichTextEnum() RichTextEnum {
	return RichTextAnchorLinkType
}

// RichTexts A concatenation of rich texts
type RichTexts struct {
	tdCommon
	Texts []RichText `json:"texts"` // Texts
}

// MessageType return the string telegram-type of RichTexts
func (richTexts *RichTexts) MessageType() string {
	return "richTexts"
}

// NewRichTexts creates a new RichTexts
//
// @param texts Texts
func NewRichTexts(texts []RichText) *RichTexts {
	richTextsTemp := RichTexts{
		tdCommon: tdCommon{Type: "richTexts"},
		Texts:    texts,
	}

	return &richTextsTemp
}

// GetRichTextEnum return the enum type of this object
func (richTexts *RichTexts) GetRichTextEnum() RichTextEnum {
	return RichTextsType
}
