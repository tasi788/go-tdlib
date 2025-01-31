package tdlib

import (
	"encoding/json"
	"fmt"
)

// PageBlock Describes a block of an instant view web page
type PageBlock interface {
	GetPageBlockEnum() PageBlockEnum
}

// PageBlockEnum Alias for abstract PageBlock 'Sub-Classes', used as constant-enum here
type PageBlockEnum string

// PageBlock enums
const (
	PageBlockTitleType           PageBlockEnum = "pageBlockTitle"
	PageBlockSubtitleType        PageBlockEnum = "pageBlockSubtitle"
	PageBlockAuthorDateType      PageBlockEnum = "pageBlockAuthorDate"
	PageBlockHeaderType          PageBlockEnum = "pageBlockHeader"
	PageBlockSubheaderType       PageBlockEnum = "pageBlockSubheader"
	PageBlockKickerType          PageBlockEnum = "pageBlockKicker"
	PageBlockParagraphType       PageBlockEnum = "pageBlockParagraph"
	PageBlockPreformattedType    PageBlockEnum = "pageBlockPreformatted"
	PageBlockFooterType          PageBlockEnum = "pageBlockFooter"
	PageBlockDividerType         PageBlockEnum = "pageBlockDivider"
	PageBlockAnchorType          PageBlockEnum = "pageBlockAnchor"
	PageBlockListType            PageBlockEnum = "pageBlockList"
	PageBlockBlockQuoteType      PageBlockEnum = "pageBlockBlockQuote"
	PageBlockPullQuoteType       PageBlockEnum = "pageBlockPullQuote"
	PageBlockAnimationType       PageBlockEnum = "pageBlockAnimation"
	PageBlockAudioType           PageBlockEnum = "pageBlockAudio"
	PageBlockPhotoType           PageBlockEnum = "pageBlockPhoto"
	PageBlockVideoType           PageBlockEnum = "pageBlockVideo"
	PageBlockVoiceNoteType       PageBlockEnum = "pageBlockVoiceNote"
	PageBlockCoverType           PageBlockEnum = "pageBlockCover"
	PageBlockEmbeddedType        PageBlockEnum = "pageBlockEmbedded"
	PageBlockEmbeddedPostType    PageBlockEnum = "pageBlockEmbeddedPost"
	PageBlockCollageType         PageBlockEnum = "pageBlockCollage"
	PageBlockSlideshowType       PageBlockEnum = "pageBlockSlideshow"
	PageBlockChatLinkType        PageBlockEnum = "pageBlockChatLink"
	PageBlockTableType           PageBlockEnum = "pageBlockTable"
	PageBlockDetailsType         PageBlockEnum = "pageBlockDetails"
	PageBlockRelatedArticlesType PageBlockEnum = "pageBlockRelatedArticles"
	PageBlockMapType             PageBlockEnum = "pageBlockMap"
)

func unmarshalPageBlock(rawMsg *json.RawMessage) (PageBlock, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch PageBlockEnum(objMap["@type"].(string)) {
	case PageBlockTitleType:
		var pageBlockTitle PageBlockTitle
		err := json.Unmarshal(*rawMsg, &pageBlockTitle)
		return &pageBlockTitle, err

	case PageBlockSubtitleType:
		var pageBlockSubtitle PageBlockSubtitle
		err := json.Unmarshal(*rawMsg, &pageBlockSubtitle)
		return &pageBlockSubtitle, err

	case PageBlockAuthorDateType:
		var pageBlockAuthorDate PageBlockAuthorDate
		err := json.Unmarshal(*rawMsg, &pageBlockAuthorDate)
		return &pageBlockAuthorDate, err

	case PageBlockHeaderType:
		var pageBlockHeader PageBlockHeader
		err := json.Unmarshal(*rawMsg, &pageBlockHeader)
		return &pageBlockHeader, err

	case PageBlockSubheaderType:
		var pageBlockSubheader PageBlockSubheader
		err := json.Unmarshal(*rawMsg, &pageBlockSubheader)
		return &pageBlockSubheader, err

	case PageBlockKickerType:
		var pageBlockKicker PageBlockKicker
		err := json.Unmarshal(*rawMsg, &pageBlockKicker)
		return &pageBlockKicker, err

	case PageBlockParagraphType:
		var pageBlockParagraph PageBlockParagraph
		err := json.Unmarshal(*rawMsg, &pageBlockParagraph)
		return &pageBlockParagraph, err

	case PageBlockPreformattedType:
		var pageBlockPreformatted PageBlockPreformatted
		err := json.Unmarshal(*rawMsg, &pageBlockPreformatted)
		return &pageBlockPreformatted, err

	case PageBlockFooterType:
		var pageBlockFooter PageBlockFooter
		err := json.Unmarshal(*rawMsg, &pageBlockFooter)
		return &pageBlockFooter, err

	case PageBlockDividerType:
		var pageBlockDivider PageBlockDivider
		err := json.Unmarshal(*rawMsg, &pageBlockDivider)
		return &pageBlockDivider, err

	case PageBlockAnchorType:
		var pageBlockAnchor PageBlockAnchor
		err := json.Unmarshal(*rawMsg, &pageBlockAnchor)
		return &pageBlockAnchor, err

	case PageBlockListType:
		var pageBlockList PageBlockList
		err := json.Unmarshal(*rawMsg, &pageBlockList)
		return &pageBlockList, err

	case PageBlockBlockQuoteType:
		var pageBlockBlockQuote PageBlockBlockQuote
		err := json.Unmarshal(*rawMsg, &pageBlockBlockQuote)
		return &pageBlockBlockQuote, err

	case PageBlockPullQuoteType:
		var pageBlockPullQuote PageBlockPullQuote
		err := json.Unmarshal(*rawMsg, &pageBlockPullQuote)
		return &pageBlockPullQuote, err

	case PageBlockAnimationType:
		var pageBlockAnimation PageBlockAnimation
		err := json.Unmarshal(*rawMsg, &pageBlockAnimation)
		return &pageBlockAnimation, err

	case PageBlockAudioType:
		var pageBlockAudio PageBlockAudio
		err := json.Unmarshal(*rawMsg, &pageBlockAudio)
		return &pageBlockAudio, err

	case PageBlockPhotoType:
		var pageBlockPhoto PageBlockPhoto
		err := json.Unmarshal(*rawMsg, &pageBlockPhoto)
		return &pageBlockPhoto, err

	case PageBlockVideoType:
		var pageBlockVideo PageBlockVideo
		err := json.Unmarshal(*rawMsg, &pageBlockVideo)
		return &pageBlockVideo, err

	case PageBlockVoiceNoteType:
		var pageBlockVoiceNote PageBlockVoiceNote
		err := json.Unmarshal(*rawMsg, &pageBlockVoiceNote)
		return &pageBlockVoiceNote, err

	case PageBlockCoverType:
		var pageBlockCover PageBlockCover
		err := json.Unmarshal(*rawMsg, &pageBlockCover)
		return &pageBlockCover, err

	case PageBlockEmbeddedType:
		var pageBlockEmbedded PageBlockEmbedded
		err := json.Unmarshal(*rawMsg, &pageBlockEmbedded)
		return &pageBlockEmbedded, err

	case PageBlockEmbeddedPostType:
		var pageBlockEmbeddedPost PageBlockEmbeddedPost
		err := json.Unmarshal(*rawMsg, &pageBlockEmbeddedPost)
		return &pageBlockEmbeddedPost, err

	case PageBlockCollageType:
		var pageBlockCollage PageBlockCollage
		err := json.Unmarshal(*rawMsg, &pageBlockCollage)
		return &pageBlockCollage, err

	case PageBlockSlideshowType:
		var pageBlockSlideshow PageBlockSlideshow
		err := json.Unmarshal(*rawMsg, &pageBlockSlideshow)
		return &pageBlockSlideshow, err

	case PageBlockChatLinkType:
		var pageBlockChatLink PageBlockChatLink
		err := json.Unmarshal(*rawMsg, &pageBlockChatLink)
		return &pageBlockChatLink, err

	case PageBlockTableType:
		var pageBlockTable PageBlockTable
		err := json.Unmarshal(*rawMsg, &pageBlockTable)
		return &pageBlockTable, err

	case PageBlockDetailsType:
		var pageBlockDetails PageBlockDetails
		err := json.Unmarshal(*rawMsg, &pageBlockDetails)
		return &pageBlockDetails, err

	case PageBlockRelatedArticlesType:
		var pageBlockRelatedArticles PageBlockRelatedArticles
		err := json.Unmarshal(*rawMsg, &pageBlockRelatedArticles)
		return &pageBlockRelatedArticles, err

	case PageBlockMapType:
		var pageBlockMap PageBlockMap
		err := json.Unmarshal(*rawMsg, &pageBlockMap)
		return &pageBlockMap, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// PageBlockTitle The title of a page
type PageBlockTitle struct {
	tdCommon
	Title RichText `json:"title"` // Title
}

// MessageType return the string telegram-type of PageBlockTitle
func (pageBlockTitle *PageBlockTitle) MessageType() string {
	return "pageBlockTitle"
}

// NewPageBlockTitle creates a new PageBlockTitle
//
// @param title Title
func NewPageBlockTitle(title RichText) *PageBlockTitle {
	pageBlockTitleTemp := PageBlockTitle{
		tdCommon: tdCommon{Type: "pageBlockTitle"},
		Title:    title,
	}

	return &pageBlockTitleTemp
}

// UnmarshalJSON unmarshal to json
func (pageBlockTitle *PageBlockTitle) UnmarshalJSON(b []byte) error {
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

	pageBlockTitle.tdCommon = tempObj.tdCommon

	fieldTitle, _ := unmarshalRichText(objMap["title"])
	pageBlockTitle.Title = fieldTitle

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockTitle *PageBlockTitle) GetPageBlockEnum() PageBlockEnum {
	return PageBlockTitleType
}

// PageBlockSubtitle The subtitle of a page
type PageBlockSubtitle struct {
	tdCommon
	Subtitle RichText `json:"subtitle"` // Subtitle
}

// MessageType return the string telegram-type of PageBlockSubtitle
func (pageBlockSubtitle *PageBlockSubtitle) MessageType() string {
	return "pageBlockSubtitle"
}

// NewPageBlockSubtitle creates a new PageBlockSubtitle
//
// @param subtitle Subtitle
func NewPageBlockSubtitle(subtitle RichText) *PageBlockSubtitle {
	pageBlockSubtitleTemp := PageBlockSubtitle{
		tdCommon: tdCommon{Type: "pageBlockSubtitle"},
		Subtitle: subtitle,
	}

	return &pageBlockSubtitleTemp
}

// UnmarshalJSON unmarshal to json
func (pageBlockSubtitle *PageBlockSubtitle) UnmarshalJSON(b []byte) error {
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

	pageBlockSubtitle.tdCommon = tempObj.tdCommon

	fieldSubtitle, _ := unmarshalRichText(objMap["subtitle"])
	pageBlockSubtitle.Subtitle = fieldSubtitle

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockSubtitle *PageBlockSubtitle) GetPageBlockEnum() PageBlockEnum {
	return PageBlockSubtitleType
}

// PageBlockAuthorDate The author and publishing date of a page
type PageBlockAuthorDate struct {
	tdCommon
	Author      RichText `json:"author"`       // Author
	PublishDate int32    `json:"publish_date"` // Point in time (Unix timestamp) when the article was published; 0 if unknown
}

// MessageType return the string telegram-type of PageBlockAuthorDate
func (pageBlockAuthorDate *PageBlockAuthorDate) MessageType() string {
	return "pageBlockAuthorDate"
}

// NewPageBlockAuthorDate creates a new PageBlockAuthorDate
//
// @param author Author
// @param publishDate Point in time (Unix timestamp) when the article was published; 0 if unknown
func NewPageBlockAuthorDate(author RichText, publishDate int32) *PageBlockAuthorDate {
	pageBlockAuthorDateTemp := PageBlockAuthorDate{
		tdCommon:    tdCommon{Type: "pageBlockAuthorDate"},
		Author:      author,
		PublishDate: publishDate,
	}

	return &pageBlockAuthorDateTemp
}

// UnmarshalJSON unmarshal to json
func (pageBlockAuthorDate *PageBlockAuthorDate) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		PublishDate int32 `json:"publish_date"` // Point in time (Unix timestamp) when the article was published; 0 if unknown
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockAuthorDate.tdCommon = tempObj.tdCommon
	pageBlockAuthorDate.PublishDate = tempObj.PublishDate

	fieldAuthor, _ := unmarshalRichText(objMap["author"])
	pageBlockAuthorDate.Author = fieldAuthor

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockAuthorDate *PageBlockAuthorDate) GetPageBlockEnum() PageBlockEnum {
	return PageBlockAuthorDateType
}

// PageBlockHeader A header
type PageBlockHeader struct {
	tdCommon
	Header RichText `json:"header"` // Header
}

// MessageType return the string telegram-type of PageBlockHeader
func (pageBlockHeader *PageBlockHeader) MessageType() string {
	return "pageBlockHeader"
}

// NewPageBlockHeader creates a new PageBlockHeader
//
// @param header Header
func NewPageBlockHeader(header RichText) *PageBlockHeader {
	pageBlockHeaderTemp := PageBlockHeader{
		tdCommon: tdCommon{Type: "pageBlockHeader"},
		Header:   header,
	}

	return &pageBlockHeaderTemp
}

// UnmarshalJSON unmarshal to json
func (pageBlockHeader *PageBlockHeader) UnmarshalJSON(b []byte) error {
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

	pageBlockHeader.tdCommon = tempObj.tdCommon

	fieldHeader, _ := unmarshalRichText(objMap["header"])
	pageBlockHeader.Header = fieldHeader

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockHeader *PageBlockHeader) GetPageBlockEnum() PageBlockEnum {
	return PageBlockHeaderType
}

// PageBlockSubheader A subheader
type PageBlockSubheader struct {
	tdCommon
	Subheader RichText `json:"subheader"` // Subheader
}

// MessageType return the string telegram-type of PageBlockSubheader
func (pageBlockSubheader *PageBlockSubheader) MessageType() string {
	return "pageBlockSubheader"
}

// NewPageBlockSubheader creates a new PageBlockSubheader
//
// @param subheader Subheader
func NewPageBlockSubheader(subheader RichText) *PageBlockSubheader {
	pageBlockSubheaderTemp := PageBlockSubheader{
		tdCommon:  tdCommon{Type: "pageBlockSubheader"},
		Subheader: subheader,
	}

	return &pageBlockSubheaderTemp
}

// UnmarshalJSON unmarshal to json
func (pageBlockSubheader *PageBlockSubheader) UnmarshalJSON(b []byte) error {
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

	pageBlockSubheader.tdCommon = tempObj.tdCommon

	fieldSubheader, _ := unmarshalRichText(objMap["subheader"])
	pageBlockSubheader.Subheader = fieldSubheader

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockSubheader *PageBlockSubheader) GetPageBlockEnum() PageBlockEnum {
	return PageBlockSubheaderType
}

// PageBlockKicker A kicker
type PageBlockKicker struct {
	tdCommon
	Kicker RichText `json:"kicker"` // Kicker
}

// MessageType return the string telegram-type of PageBlockKicker
func (pageBlockKicker *PageBlockKicker) MessageType() string {
	return "pageBlockKicker"
}

// NewPageBlockKicker creates a new PageBlockKicker
//
// @param kicker Kicker
func NewPageBlockKicker(kicker RichText) *PageBlockKicker {
	pageBlockKickerTemp := PageBlockKicker{
		tdCommon: tdCommon{Type: "pageBlockKicker"},
		Kicker:   kicker,
	}

	return &pageBlockKickerTemp
}

// UnmarshalJSON unmarshal to json
func (pageBlockKicker *PageBlockKicker) UnmarshalJSON(b []byte) error {
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

	pageBlockKicker.tdCommon = tempObj.tdCommon

	fieldKicker, _ := unmarshalRichText(objMap["kicker"])
	pageBlockKicker.Kicker = fieldKicker

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockKicker *PageBlockKicker) GetPageBlockEnum() PageBlockEnum {
	return PageBlockKickerType
}

// PageBlockParagraph A text paragraph
type PageBlockParagraph struct {
	tdCommon
	Text RichText `json:"text"` // Paragraph text
}

// MessageType return the string telegram-type of PageBlockParagraph
func (pageBlockParagraph *PageBlockParagraph) MessageType() string {
	return "pageBlockParagraph"
}

// NewPageBlockParagraph creates a new PageBlockParagraph
//
// @param text Paragraph text
func NewPageBlockParagraph(text RichText) *PageBlockParagraph {
	pageBlockParagraphTemp := PageBlockParagraph{
		tdCommon: tdCommon{Type: "pageBlockParagraph"},
		Text:     text,
	}

	return &pageBlockParagraphTemp
}

// UnmarshalJSON unmarshal to json
func (pageBlockParagraph *PageBlockParagraph) UnmarshalJSON(b []byte) error {
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

	pageBlockParagraph.tdCommon = tempObj.tdCommon

	fieldText, _ := unmarshalRichText(objMap["text"])
	pageBlockParagraph.Text = fieldText

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockParagraph *PageBlockParagraph) GetPageBlockEnum() PageBlockEnum {
	return PageBlockParagraphType
}

// PageBlockPreformatted A preformatted text paragraph
type PageBlockPreformatted struct {
	tdCommon
	Text     RichText `json:"text"`     // Paragraph text
	Language string   `json:"language"` // Programming language for which the text needs to be formatted
}

// MessageType return the string telegram-type of PageBlockPreformatted
func (pageBlockPreformatted *PageBlockPreformatted) MessageType() string {
	return "pageBlockPreformatted"
}

// NewPageBlockPreformatted creates a new PageBlockPreformatted
//
// @param text Paragraph text
// @param language Programming language for which the text needs to be formatted
func NewPageBlockPreformatted(text RichText, language string) *PageBlockPreformatted {
	pageBlockPreformattedTemp := PageBlockPreformatted{
		tdCommon: tdCommon{Type: "pageBlockPreformatted"},
		Text:     text,
		Language: language,
	}

	return &pageBlockPreformattedTemp
}

// UnmarshalJSON unmarshal to json
func (pageBlockPreformatted *PageBlockPreformatted) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Language string `json:"language"` // Programming language for which the text needs to be formatted
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockPreformatted.tdCommon = tempObj.tdCommon
	pageBlockPreformatted.Language = tempObj.Language

	fieldText, _ := unmarshalRichText(objMap["text"])
	pageBlockPreformatted.Text = fieldText

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockPreformatted *PageBlockPreformatted) GetPageBlockEnum() PageBlockEnum {
	return PageBlockPreformattedType
}

// PageBlockFooter The footer of a page
type PageBlockFooter struct {
	tdCommon
	Footer RichText `json:"footer"` // Footer
}

// MessageType return the string telegram-type of PageBlockFooter
func (pageBlockFooter *PageBlockFooter) MessageType() string {
	return "pageBlockFooter"
}

// NewPageBlockFooter creates a new PageBlockFooter
//
// @param footer Footer
func NewPageBlockFooter(footer RichText) *PageBlockFooter {
	pageBlockFooterTemp := PageBlockFooter{
		tdCommon: tdCommon{Type: "pageBlockFooter"},
		Footer:   footer,
	}

	return &pageBlockFooterTemp
}

// UnmarshalJSON unmarshal to json
func (pageBlockFooter *PageBlockFooter) UnmarshalJSON(b []byte) error {
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

	pageBlockFooter.tdCommon = tempObj.tdCommon

	fieldFooter, _ := unmarshalRichText(objMap["footer"])
	pageBlockFooter.Footer = fieldFooter

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockFooter *PageBlockFooter) GetPageBlockEnum() PageBlockEnum {
	return PageBlockFooterType
}

// PageBlockDivider An empty block separating a page
type PageBlockDivider struct {
	tdCommon
}

// MessageType return the string telegram-type of PageBlockDivider
func (pageBlockDivider *PageBlockDivider) MessageType() string {
	return "pageBlockDivider"
}

// NewPageBlockDivider creates a new PageBlockDivider
//
func NewPageBlockDivider() *PageBlockDivider {
	pageBlockDividerTemp := PageBlockDivider{
		tdCommon: tdCommon{Type: "pageBlockDivider"},
	}

	return &pageBlockDividerTemp
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockDivider *PageBlockDivider) GetPageBlockEnum() PageBlockEnum {
	return PageBlockDividerType
}

// PageBlockAnchor An invisible anchor on a page, which can be used in a URL to open the page from the specified anchor
type PageBlockAnchor struct {
	tdCommon
	Name string `json:"name"` // Name of the anchor
}

// MessageType return the string telegram-type of PageBlockAnchor
func (pageBlockAnchor *PageBlockAnchor) MessageType() string {
	return "pageBlockAnchor"
}

// NewPageBlockAnchor creates a new PageBlockAnchor
//
// @param name Name of the anchor
func NewPageBlockAnchor(name string) *PageBlockAnchor {
	pageBlockAnchorTemp := PageBlockAnchor{
		tdCommon: tdCommon{Type: "pageBlockAnchor"},
		Name:     name,
	}

	return &pageBlockAnchorTemp
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockAnchor *PageBlockAnchor) GetPageBlockEnum() PageBlockEnum {
	return PageBlockAnchorType
}

// PageBlockList A list of data blocks
type PageBlockList struct {
	tdCommon
	Items []PageBlockListItem `json:"items"` // The items of the list
}

// MessageType return the string telegram-type of PageBlockList
func (pageBlockList *PageBlockList) MessageType() string {
	return "pageBlockList"
}

// NewPageBlockList creates a new PageBlockList
//
// @param items The items of the list
func NewPageBlockList(items []PageBlockListItem) *PageBlockList {
	pageBlockListTemp := PageBlockList{
		tdCommon: tdCommon{Type: "pageBlockList"},
		Items:    items,
	}

	return &pageBlockListTemp
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockList *PageBlockList) GetPageBlockEnum() PageBlockEnum {
	return PageBlockListType
}

// PageBlockBlockQuote A block quote
type PageBlockBlockQuote struct {
	tdCommon
	Text   RichText `json:"text"`   // Quote text
	Credit RichText `json:"credit"` // Quote credit
}

// MessageType return the string telegram-type of PageBlockBlockQuote
func (pageBlockBlockQuote *PageBlockBlockQuote) MessageType() string {
	return "pageBlockBlockQuote"
}

// NewPageBlockBlockQuote creates a new PageBlockBlockQuote
//
// @param text Quote text
// @param credit Quote credit
func NewPageBlockBlockQuote(text RichText, credit RichText) *PageBlockBlockQuote {
	pageBlockBlockQuoteTemp := PageBlockBlockQuote{
		tdCommon: tdCommon{Type: "pageBlockBlockQuote"},
		Text:     text,
		Credit:   credit,
	}

	return &pageBlockBlockQuoteTemp
}

// UnmarshalJSON unmarshal to json
func (pageBlockBlockQuote *PageBlockBlockQuote) UnmarshalJSON(b []byte) error {
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

	pageBlockBlockQuote.tdCommon = tempObj.tdCommon

	fieldText, _ := unmarshalRichText(objMap["text"])
	pageBlockBlockQuote.Text = fieldText

	fieldCredit, _ := unmarshalRichText(objMap["credit"])
	pageBlockBlockQuote.Credit = fieldCredit

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockBlockQuote *PageBlockBlockQuote) GetPageBlockEnum() PageBlockEnum {
	return PageBlockBlockQuoteType
}

// PageBlockPullQuote A pull quote
type PageBlockPullQuote struct {
	tdCommon
	Text   RichText `json:"text"`   // Quote text
	Credit RichText `json:"credit"` // Quote credit
}

// MessageType return the string telegram-type of PageBlockPullQuote
func (pageBlockPullQuote *PageBlockPullQuote) MessageType() string {
	return "pageBlockPullQuote"
}

// NewPageBlockPullQuote creates a new PageBlockPullQuote
//
// @param text Quote text
// @param credit Quote credit
func NewPageBlockPullQuote(text RichText, credit RichText) *PageBlockPullQuote {
	pageBlockPullQuoteTemp := PageBlockPullQuote{
		tdCommon: tdCommon{Type: "pageBlockPullQuote"},
		Text:     text,
		Credit:   credit,
	}

	return &pageBlockPullQuoteTemp
}

// UnmarshalJSON unmarshal to json
func (pageBlockPullQuote *PageBlockPullQuote) UnmarshalJSON(b []byte) error {
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

	pageBlockPullQuote.tdCommon = tempObj.tdCommon

	fieldText, _ := unmarshalRichText(objMap["text"])
	pageBlockPullQuote.Text = fieldText

	fieldCredit, _ := unmarshalRichText(objMap["credit"])
	pageBlockPullQuote.Credit = fieldCredit

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockPullQuote *PageBlockPullQuote) GetPageBlockEnum() PageBlockEnum {
	return PageBlockPullQuoteType
}

// PageBlockAnimation An animation
type PageBlockAnimation struct {
	tdCommon
	Animation    *Animation        `json:"animation"`     // Animation file; may be null
	Caption      *PageBlockCaption `json:"caption"`       // Animation caption
	NeedAutoplay bool              `json:"need_autoplay"` // True, if the animation must be played automatically
}

// MessageType return the string telegram-type of PageBlockAnimation
func (pageBlockAnimation *PageBlockAnimation) MessageType() string {
	return "pageBlockAnimation"
}

// NewPageBlockAnimation creates a new PageBlockAnimation
//
// @param animation Animation file; may be null
// @param caption Animation caption
// @param needAutoplay True, if the animation must be played automatically
func NewPageBlockAnimation(animation *Animation, caption *PageBlockCaption, needAutoplay bool) *PageBlockAnimation {
	pageBlockAnimationTemp := PageBlockAnimation{
		tdCommon:     tdCommon{Type: "pageBlockAnimation"},
		Animation:    animation,
		Caption:      caption,
		NeedAutoplay: needAutoplay,
	}

	return &pageBlockAnimationTemp
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockAnimation *PageBlockAnimation) GetPageBlockEnum() PageBlockEnum {
	return PageBlockAnimationType
}

// PageBlockAudio An audio file
type PageBlockAudio struct {
	tdCommon
	Audio   *Audio            `json:"audio"`   // Audio file; may be null
	Caption *PageBlockCaption `json:"caption"` // Audio file caption
}

// MessageType return the string telegram-type of PageBlockAudio
func (pageBlockAudio *PageBlockAudio) MessageType() string {
	return "pageBlockAudio"
}

// NewPageBlockAudio creates a new PageBlockAudio
//
// @param audio Audio file; may be null
// @param caption Audio file caption
func NewPageBlockAudio(audio *Audio, caption *PageBlockCaption) *PageBlockAudio {
	pageBlockAudioTemp := PageBlockAudio{
		tdCommon: tdCommon{Type: "pageBlockAudio"},
		Audio:    audio,
		Caption:  caption,
	}

	return &pageBlockAudioTemp
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockAudio *PageBlockAudio) GetPageBlockEnum() PageBlockEnum {
	return PageBlockAudioType
}

// PageBlockPhoto A photo
type PageBlockPhoto struct {
	tdCommon
	Photo   *Photo            `json:"photo"`   // Photo file; may be null
	Caption *PageBlockCaption `json:"caption"` // Photo caption
	Url     string            `json:"url"`     // URL that needs to be opened when the photo is clicked
}

// MessageType return the string telegram-type of PageBlockPhoto
func (pageBlockPhoto *PageBlockPhoto) MessageType() string {
	return "pageBlockPhoto"
}

// NewPageBlockPhoto creates a new PageBlockPhoto
//
// @param photo Photo file; may be null
// @param caption Photo caption
// @param url URL that needs to be opened when the photo is clicked
func NewPageBlockPhoto(photo *Photo, caption *PageBlockCaption, url string) *PageBlockPhoto {
	pageBlockPhotoTemp := PageBlockPhoto{
		tdCommon: tdCommon{Type: "pageBlockPhoto"},
		Photo:    photo,
		Caption:  caption,
		Url:      url,
	}

	return &pageBlockPhotoTemp
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockPhoto *PageBlockPhoto) GetPageBlockEnum() PageBlockEnum {
	return PageBlockPhotoType
}

// PageBlockVideo A video
type PageBlockVideo struct {
	tdCommon
	Video        *Video            `json:"video"`         // Video file; may be null
	Caption      *PageBlockCaption `json:"caption"`       // Video caption
	NeedAutoplay bool              `json:"need_autoplay"` // True, if the video must be played automatically
	IsLooped     bool              `json:"is_looped"`     // True, if the video must be looped
}

// MessageType return the string telegram-type of PageBlockVideo
func (pageBlockVideo *PageBlockVideo) MessageType() string {
	return "pageBlockVideo"
}

// NewPageBlockVideo creates a new PageBlockVideo
//
// @param video Video file; may be null
// @param caption Video caption
// @param needAutoplay True, if the video must be played automatically
// @param isLooped True, if the video must be looped
func NewPageBlockVideo(video *Video, caption *PageBlockCaption, needAutoplay bool, isLooped bool) *PageBlockVideo {
	pageBlockVideoTemp := PageBlockVideo{
		tdCommon:     tdCommon{Type: "pageBlockVideo"},
		Video:        video,
		Caption:      caption,
		NeedAutoplay: needAutoplay,
		IsLooped:     isLooped,
	}

	return &pageBlockVideoTemp
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockVideo *PageBlockVideo) GetPageBlockEnum() PageBlockEnum {
	return PageBlockVideoType
}

// PageBlockVoiceNote A voice note
type PageBlockVoiceNote struct {
	tdCommon
	VoiceNote *VoiceNote        `json:"voice_note"` // Voice note; may be null
	Caption   *PageBlockCaption `json:"caption"`    // Voice note caption
}

// MessageType return the string telegram-type of PageBlockVoiceNote
func (pageBlockVoiceNote *PageBlockVoiceNote) MessageType() string {
	return "pageBlockVoiceNote"
}

// NewPageBlockVoiceNote creates a new PageBlockVoiceNote
//
// @param voiceNote Voice note; may be null
// @param caption Voice note caption
func NewPageBlockVoiceNote(voiceNote *VoiceNote, caption *PageBlockCaption) *PageBlockVoiceNote {
	pageBlockVoiceNoteTemp := PageBlockVoiceNote{
		tdCommon:  tdCommon{Type: "pageBlockVoiceNote"},
		VoiceNote: voiceNote,
		Caption:   caption,
	}

	return &pageBlockVoiceNoteTemp
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockVoiceNote *PageBlockVoiceNote) GetPageBlockEnum() PageBlockEnum {
	return PageBlockVoiceNoteType
}

// PageBlockCover A page cover
type PageBlockCover struct {
	tdCommon
	Cover PageBlock `json:"cover"` // Cover
}

// MessageType return the string telegram-type of PageBlockCover
func (pageBlockCover *PageBlockCover) MessageType() string {
	return "pageBlockCover"
}

// NewPageBlockCover creates a new PageBlockCover
//
// @param cover Cover
func NewPageBlockCover(cover PageBlock) *PageBlockCover {
	pageBlockCoverTemp := PageBlockCover{
		tdCommon: tdCommon{Type: "pageBlockCover"},
		Cover:    cover,
	}

	return &pageBlockCoverTemp
}

// UnmarshalJSON unmarshal to json
func (pageBlockCover *PageBlockCover) UnmarshalJSON(b []byte) error {
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

	pageBlockCover.tdCommon = tempObj.tdCommon

	fieldCover, _ := unmarshalPageBlock(objMap["cover"])
	pageBlockCover.Cover = fieldCover

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockCover *PageBlockCover) GetPageBlockEnum() PageBlockEnum {
	return PageBlockCoverType
}

// PageBlockEmbedded An embedded web page
type PageBlockEmbedded struct {
	tdCommon
	Url            string            `json:"url"`             // Web page URL, if available
	Html           string            `json:"html"`            // HTML-markup of the embedded page
	PosterPhoto    *Photo            `json:"poster_photo"`    // Poster photo, if available; may be null
	Width          int32             `json:"width"`           // Block width; 0 if unknown
	Height         int32             `json:"height"`          // Block height; 0 if unknown
	Caption        *PageBlockCaption `json:"caption"`         // Block caption
	IsFullWidth    bool              `json:"is_full_width"`   // True, if the block must be full width
	AllowScrolling bool              `json:"allow_scrolling"` // True, if scrolling needs to be allowed
}

// MessageType return the string telegram-type of PageBlockEmbedded
func (pageBlockEmbedded *PageBlockEmbedded) MessageType() string {
	return "pageBlockEmbedded"
}

// NewPageBlockEmbedded creates a new PageBlockEmbedded
//
// @param url Web page URL, if available
// @param html HTML-markup of the embedded page
// @param posterPhoto Poster photo, if available; may be null
// @param width Block width; 0 if unknown
// @param height Block height; 0 if unknown
// @param caption Block caption
// @param isFullWidth True, if the block must be full width
// @param allowScrolling True, if scrolling needs to be allowed
func NewPageBlockEmbedded(url string, html string, posterPhoto *Photo, width int32, height int32, caption *PageBlockCaption, isFullWidth bool, allowScrolling bool) *PageBlockEmbedded {
	pageBlockEmbeddedTemp := PageBlockEmbedded{
		tdCommon:       tdCommon{Type: "pageBlockEmbedded"},
		Url:            url,
		Html:           html,
		PosterPhoto:    posterPhoto,
		Width:          width,
		Height:         height,
		Caption:        caption,
		IsFullWidth:    isFullWidth,
		AllowScrolling: allowScrolling,
	}

	return &pageBlockEmbeddedTemp
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockEmbedded *PageBlockEmbedded) GetPageBlockEnum() PageBlockEnum {
	return PageBlockEmbeddedType
}

// PageBlockEmbeddedPost An embedded post
type PageBlockEmbeddedPost struct {
	tdCommon
	Url         string            `json:"url"`          // Web page URL
	Author      string            `json:"author"`       // Post author
	AuthorPhoto *Photo            `json:"author_photo"` // Post author photo; may be null
	Date        int32             `json:"date"`         // Point in time (Unix timestamp) when the post was created; 0 if unknown
	PageBlocks  []PageBlock       `json:"page_blocks"`  // Post content
	Caption     *PageBlockCaption `json:"caption"`      // Post caption
}

// MessageType return the string telegram-type of PageBlockEmbeddedPost
func (pageBlockEmbeddedPost *PageBlockEmbeddedPost) MessageType() string {
	return "pageBlockEmbeddedPost"
}

// NewPageBlockEmbeddedPost creates a new PageBlockEmbeddedPost
//
// @param url Web page URL
// @param author Post author
// @param authorPhoto Post author photo; may be null
// @param date Point in time (Unix timestamp) when the post was created; 0 if unknown
// @param pageBlocks Post content
// @param caption Post caption
func NewPageBlockEmbeddedPost(url string, author string, authorPhoto *Photo, date int32, pageBlocks []PageBlock, caption *PageBlockCaption) *PageBlockEmbeddedPost {
	pageBlockEmbeddedPostTemp := PageBlockEmbeddedPost{
		tdCommon:    tdCommon{Type: "pageBlockEmbeddedPost"},
		Url:         url,
		Author:      author,
		AuthorPhoto: authorPhoto,
		Date:        date,
		PageBlocks:  pageBlocks,
		Caption:     caption,
	}

	return &pageBlockEmbeddedPostTemp
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockEmbeddedPost *PageBlockEmbeddedPost) GetPageBlockEnum() PageBlockEnum {
	return PageBlockEmbeddedPostType
}

// PageBlockCollage A collage
type PageBlockCollage struct {
	tdCommon
	PageBlocks []PageBlock       `json:"page_blocks"` // Collage item contents
	Caption    *PageBlockCaption `json:"caption"`     // Block caption
}

// MessageType return the string telegram-type of PageBlockCollage
func (pageBlockCollage *PageBlockCollage) MessageType() string {
	return "pageBlockCollage"
}

// NewPageBlockCollage creates a new PageBlockCollage
//
// @param pageBlocks Collage item contents
// @param caption Block caption
func NewPageBlockCollage(pageBlocks []PageBlock, caption *PageBlockCaption) *PageBlockCollage {
	pageBlockCollageTemp := PageBlockCollage{
		tdCommon:   tdCommon{Type: "pageBlockCollage"},
		PageBlocks: pageBlocks,
		Caption:    caption,
	}

	return &pageBlockCollageTemp
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockCollage *PageBlockCollage) GetPageBlockEnum() PageBlockEnum {
	return PageBlockCollageType
}

// PageBlockSlideshow A slideshow
type PageBlockSlideshow struct {
	tdCommon
	PageBlocks []PageBlock       `json:"page_blocks"` // Slideshow item contents
	Caption    *PageBlockCaption `json:"caption"`     // Block caption
}

// MessageType return the string telegram-type of PageBlockSlideshow
func (pageBlockSlideshow *PageBlockSlideshow) MessageType() string {
	return "pageBlockSlideshow"
}

// NewPageBlockSlideshow creates a new PageBlockSlideshow
//
// @param pageBlocks Slideshow item contents
// @param caption Block caption
func NewPageBlockSlideshow(pageBlocks []PageBlock, caption *PageBlockCaption) *PageBlockSlideshow {
	pageBlockSlideshowTemp := PageBlockSlideshow{
		tdCommon:   tdCommon{Type: "pageBlockSlideshow"},
		PageBlocks: pageBlocks,
		Caption:    caption,
	}

	return &pageBlockSlideshowTemp
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockSlideshow *PageBlockSlideshow) GetPageBlockEnum() PageBlockEnum {
	return PageBlockSlideshowType
}

// PageBlockChatLink A link to a chat
type PageBlockChatLink struct {
	tdCommon
	Title    string         `json:"title"`    // Chat title
	Photo    *ChatPhotoInfo `json:"photo"`    // Chat photo; may be null
	Username string         `json:"username"` // Chat username, by which all other information about the chat can be resolved
}

// MessageType return the string telegram-type of PageBlockChatLink
func (pageBlockChatLink *PageBlockChatLink) MessageType() string {
	return "pageBlockChatLink"
}

// NewPageBlockChatLink creates a new PageBlockChatLink
//
// @param title Chat title
// @param photo Chat photo; may be null
// @param username Chat username, by which all other information about the chat can be resolved
func NewPageBlockChatLink(title string, photo *ChatPhotoInfo, username string) *PageBlockChatLink {
	pageBlockChatLinkTemp := PageBlockChatLink{
		tdCommon: tdCommon{Type: "pageBlockChatLink"},
		Title:    title,
		Photo:    photo,
		Username: username,
	}

	return &pageBlockChatLinkTemp
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockChatLink *PageBlockChatLink) GetPageBlockEnum() PageBlockEnum {
	return PageBlockChatLinkType
}

// PageBlockTable A table
type PageBlockTable struct {
	tdCommon
	Caption    RichText               `json:"caption"`     // Table caption
	Cells      [][]PageBlockTableCell `json:"cells"`       // Table cells
	IsBordered bool                   `json:"is_bordered"` // True, if the table is bordered
	IsStriped  bool                   `json:"is_striped"`  // True, if the table is striped
}

// MessageType return the string telegram-type of PageBlockTable
func (pageBlockTable *PageBlockTable) MessageType() string {
	return "pageBlockTable"
}

// NewPageBlockTable creates a new PageBlockTable
//
// @param caption Table caption
// @param cells Table cells
// @param isBordered True, if the table is bordered
// @param isStriped True, if the table is striped
func NewPageBlockTable(caption RichText, cells [][]PageBlockTableCell, isBordered bool, isStriped bool) *PageBlockTable {
	pageBlockTableTemp := PageBlockTable{
		tdCommon:   tdCommon{Type: "pageBlockTable"},
		Caption:    caption,
		Cells:      cells,
		IsBordered: isBordered,
		IsStriped:  isStriped,
	}

	return &pageBlockTableTemp
}

// UnmarshalJSON unmarshal to json
func (pageBlockTable *PageBlockTable) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Cells      [][]PageBlockTableCell `json:"cells"`       // Table cells
		IsBordered bool                   `json:"is_bordered"` // True, if the table is bordered
		IsStriped  bool                   `json:"is_striped"`  // True, if the table is striped
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockTable.tdCommon = tempObj.tdCommon
	pageBlockTable.Cells = tempObj.Cells
	pageBlockTable.IsBordered = tempObj.IsBordered
	pageBlockTable.IsStriped = tempObj.IsStriped

	fieldCaption, _ := unmarshalRichText(objMap["caption"])
	pageBlockTable.Caption = fieldCaption

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockTable *PageBlockTable) GetPageBlockEnum() PageBlockEnum {
	return PageBlockTableType
}

// PageBlockDetails A collapsible block
type PageBlockDetails struct {
	tdCommon
	Header     RichText    `json:"header"`      // Always visible heading for the block
	PageBlocks []PageBlock `json:"page_blocks"` // Block contents
	IsOpen     bool        `json:"is_open"`     // True, if the block is open by default
}

// MessageType return the string telegram-type of PageBlockDetails
func (pageBlockDetails *PageBlockDetails) MessageType() string {
	return "pageBlockDetails"
}

// NewPageBlockDetails creates a new PageBlockDetails
//
// @param header Always visible heading for the block
// @param pageBlocks Block contents
// @param isOpen True, if the block is open by default
func NewPageBlockDetails(header RichText, pageBlocks []PageBlock, isOpen bool) *PageBlockDetails {
	pageBlockDetailsTemp := PageBlockDetails{
		tdCommon:   tdCommon{Type: "pageBlockDetails"},
		Header:     header,
		PageBlocks: pageBlocks,
		IsOpen:     isOpen,
	}

	return &pageBlockDetailsTemp
}

// UnmarshalJSON unmarshal to json
func (pageBlockDetails *PageBlockDetails) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		PageBlocks []PageBlock `json:"page_blocks"` // Block contents
		IsOpen     bool        `json:"is_open"`     // True, if the block is open by default
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockDetails.tdCommon = tempObj.tdCommon
	pageBlockDetails.PageBlocks = tempObj.PageBlocks
	pageBlockDetails.IsOpen = tempObj.IsOpen

	fieldHeader, _ := unmarshalRichText(objMap["header"])
	pageBlockDetails.Header = fieldHeader

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockDetails *PageBlockDetails) GetPageBlockEnum() PageBlockEnum {
	return PageBlockDetailsType
}

// PageBlockRelatedArticles Related articles
type PageBlockRelatedArticles struct {
	tdCommon
	Header   RichText                  `json:"header"`   // Block header
	Articles []PageBlockRelatedArticle `json:"articles"` // List of related articles
}

// MessageType return the string telegram-type of PageBlockRelatedArticles
func (pageBlockRelatedArticles *PageBlockRelatedArticles) MessageType() string {
	return "pageBlockRelatedArticles"
}

// NewPageBlockRelatedArticles creates a new PageBlockRelatedArticles
//
// @param header Block header
// @param articles List of related articles
func NewPageBlockRelatedArticles(header RichText, articles []PageBlockRelatedArticle) *PageBlockRelatedArticles {
	pageBlockRelatedArticlesTemp := PageBlockRelatedArticles{
		tdCommon: tdCommon{Type: "pageBlockRelatedArticles"},
		Header:   header,
		Articles: articles,
	}

	return &pageBlockRelatedArticlesTemp
}

// UnmarshalJSON unmarshal to json
func (pageBlockRelatedArticles *PageBlockRelatedArticles) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		Articles []PageBlockRelatedArticle `json:"articles"` // List of related articles
	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	pageBlockRelatedArticles.tdCommon = tempObj.tdCommon
	pageBlockRelatedArticles.Articles = tempObj.Articles

	fieldHeader, _ := unmarshalRichText(objMap["header"])
	pageBlockRelatedArticles.Header = fieldHeader

	return nil
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockRelatedArticles *PageBlockRelatedArticles) GetPageBlockEnum() PageBlockEnum {
	return PageBlockRelatedArticlesType
}

// PageBlockMap A map
type PageBlockMap struct {
	tdCommon
	Location *Location         `json:"location"` // Location of the map center
	Zoom     int32             `json:"zoom"`     // Map zoom level
	Width    int32             `json:"width"`    // Map width
	Height   int32             `json:"height"`   // Map height
	Caption  *PageBlockCaption `json:"caption"`  // Block caption
}

// MessageType return the string telegram-type of PageBlockMap
func (pageBlockMap *PageBlockMap) MessageType() string {
	return "pageBlockMap"
}

// NewPageBlockMap creates a new PageBlockMap
//
// @param location Location of the map center
// @param zoom Map zoom level
// @param width Map width
// @param height Map height
// @param caption Block caption
func NewPageBlockMap(location *Location, zoom int32, width int32, height int32, caption *PageBlockCaption) *PageBlockMap {
	pageBlockMapTemp := PageBlockMap{
		tdCommon: tdCommon{Type: "pageBlockMap"},
		Location: location,
		Zoom:     zoom,
		Width:    width,
		Height:   height,
		Caption:  caption,
	}

	return &pageBlockMapTemp
}

// GetPageBlockEnum return the enum type of this object
func (pageBlockMap *PageBlockMap) GetPageBlockEnum() PageBlockEnum {
	return PageBlockMapType
}
