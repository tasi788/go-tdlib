package tdlib

import (
	"encoding/json"
	"fmt"
)

// PassportElementType Contains the type of a Telegram Passport element
type PassportElementType interface {
	GetPassportElementTypeEnum() PassportElementTypeEnum
}

// PassportElementTypeEnum Alias for abstract PassportElementType 'Sub-Classes', used as constant-enum here
type PassportElementTypeEnum string

// PassportElementType enums
const (
	PassportElementTypePersonalDetailsType       PassportElementTypeEnum = "passportElementTypePersonalDetails"
	PassportElementTypePassportType              PassportElementTypeEnum = "passportElementTypePassport"
	PassportElementTypeDriverLicenseType         PassportElementTypeEnum = "passportElementTypeDriverLicense"
	PassportElementTypeIdentityCardType          PassportElementTypeEnum = "passportElementTypeIdentityCard"
	PassportElementTypeInternalPassportType      PassportElementTypeEnum = "passportElementTypeInternalPassport"
	PassportElementTypeAddressType               PassportElementTypeEnum = "passportElementTypeAddress"
	PassportElementTypeUtilityBillType           PassportElementTypeEnum = "passportElementTypeUtilityBill"
	PassportElementTypeBankStatementType         PassportElementTypeEnum = "passportElementTypeBankStatement"
	PassportElementTypeRentalAgreementType       PassportElementTypeEnum = "passportElementTypeRentalAgreement"
	PassportElementTypePassportRegistrationType  PassportElementTypeEnum = "passportElementTypePassportRegistration"
	PassportElementTypeTemporaryRegistrationType PassportElementTypeEnum = "passportElementTypeTemporaryRegistration"
	PassportElementTypePhoneNumberType           PassportElementTypeEnum = "passportElementTypePhoneNumber"
	PassportElementTypeEmailAddressType          PassportElementTypeEnum = "passportElementTypeEmailAddress"
)

func unmarshalPassportElementType(rawMsg *json.RawMessage) (PassportElementType, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch PassportElementTypeEnum(objMap["@type"].(string)) {
	case PassportElementTypePersonalDetailsType:
		var passportElementTypePersonalDetails PassportElementTypePersonalDetails
		err := json.Unmarshal(*rawMsg, &passportElementTypePersonalDetails)
		return &passportElementTypePersonalDetails, err

	case PassportElementTypePassportType:
		var passportElementTypePassport PassportElementTypePassport
		err := json.Unmarshal(*rawMsg, &passportElementTypePassport)
		return &passportElementTypePassport, err

	case PassportElementTypeDriverLicenseType:
		var passportElementTypeDriverLicense PassportElementTypeDriverLicense
		err := json.Unmarshal(*rawMsg, &passportElementTypeDriverLicense)
		return &passportElementTypeDriverLicense, err

	case PassportElementTypeIdentityCardType:
		var passportElementTypeIdentityCard PassportElementTypeIdentityCard
		err := json.Unmarshal(*rawMsg, &passportElementTypeIdentityCard)
		return &passportElementTypeIdentityCard, err

	case PassportElementTypeInternalPassportType:
		var passportElementTypeInternalPassport PassportElementTypeInternalPassport
		err := json.Unmarshal(*rawMsg, &passportElementTypeInternalPassport)
		return &passportElementTypeInternalPassport, err

	case PassportElementTypeAddressType:
		var passportElementTypeAddress PassportElementTypeAddress
		err := json.Unmarshal(*rawMsg, &passportElementTypeAddress)
		return &passportElementTypeAddress, err

	case PassportElementTypeUtilityBillType:
		var passportElementTypeUtilityBill PassportElementTypeUtilityBill
		err := json.Unmarshal(*rawMsg, &passportElementTypeUtilityBill)
		return &passportElementTypeUtilityBill, err

	case PassportElementTypeBankStatementType:
		var passportElementTypeBankStatement PassportElementTypeBankStatement
		err := json.Unmarshal(*rawMsg, &passportElementTypeBankStatement)
		return &passportElementTypeBankStatement, err

	case PassportElementTypeRentalAgreementType:
		var passportElementTypeRentalAgreement PassportElementTypeRentalAgreement
		err := json.Unmarshal(*rawMsg, &passportElementTypeRentalAgreement)
		return &passportElementTypeRentalAgreement, err

	case PassportElementTypePassportRegistrationType:
		var passportElementTypePassportRegistration PassportElementTypePassportRegistration
		err := json.Unmarshal(*rawMsg, &passportElementTypePassportRegistration)
		return &passportElementTypePassportRegistration, err

	case PassportElementTypeTemporaryRegistrationType:
		var passportElementTypeTemporaryRegistration PassportElementTypeTemporaryRegistration
		err := json.Unmarshal(*rawMsg, &passportElementTypeTemporaryRegistration)
		return &passportElementTypeTemporaryRegistration, err

	case PassportElementTypePhoneNumberType:
		var passportElementTypePhoneNumber PassportElementTypePhoneNumber
		err := json.Unmarshal(*rawMsg, &passportElementTypePhoneNumber)
		return &passportElementTypePhoneNumber, err

	case PassportElementTypeEmailAddressType:
		var passportElementTypeEmailAddress PassportElementTypeEmailAddress
		err := json.Unmarshal(*rawMsg, &passportElementTypeEmailAddress)
		return &passportElementTypeEmailAddress, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// PassportElementTypePersonalDetails A Telegram Passport element containing the user's personal details
type PassportElementTypePersonalDetails struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypePersonalDetails
func (passportElementTypePersonalDetails *PassportElementTypePersonalDetails) MessageType() string {
	return "passportElementTypePersonalDetails"
}

// NewPassportElementTypePersonalDetails creates a new PassportElementTypePersonalDetails
//
func NewPassportElementTypePersonalDetails() *PassportElementTypePersonalDetails {
	passportElementTypePersonalDetailsTemp := PassportElementTypePersonalDetails{
		tdCommon: tdCommon{Type: "passportElementTypePersonalDetails"},
	}

	return &passportElementTypePersonalDetailsTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypePersonalDetails *PassportElementTypePersonalDetails) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypePersonalDetailsType
}

// PassportElementTypePassport A Telegram Passport element containing the user's passport
type PassportElementTypePassport struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypePassport
func (passportElementTypePassport *PassportElementTypePassport) MessageType() string {
	return "passportElementTypePassport"
}

// NewPassportElementTypePassport creates a new PassportElementTypePassport
//
func NewPassportElementTypePassport() *PassportElementTypePassport {
	passportElementTypePassportTemp := PassportElementTypePassport{
		tdCommon: tdCommon{Type: "passportElementTypePassport"},
	}

	return &passportElementTypePassportTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypePassport *PassportElementTypePassport) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypePassportType
}

// PassportElementTypeDriverLicense A Telegram Passport element containing the user's driver license
type PassportElementTypeDriverLicense struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypeDriverLicense
func (passportElementTypeDriverLicense *PassportElementTypeDriverLicense) MessageType() string {
	return "passportElementTypeDriverLicense"
}

// NewPassportElementTypeDriverLicense creates a new PassportElementTypeDriverLicense
//
func NewPassportElementTypeDriverLicense() *PassportElementTypeDriverLicense {
	passportElementTypeDriverLicenseTemp := PassportElementTypeDriverLicense{
		tdCommon: tdCommon{Type: "passportElementTypeDriverLicense"},
	}

	return &passportElementTypeDriverLicenseTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypeDriverLicense *PassportElementTypeDriverLicense) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypeDriverLicenseType
}

// PassportElementTypeIdentityCard A Telegram Passport element containing the user's identity card
type PassportElementTypeIdentityCard struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypeIdentityCard
func (passportElementTypeIdentityCard *PassportElementTypeIdentityCard) MessageType() string {
	return "passportElementTypeIdentityCard"
}

// NewPassportElementTypeIdentityCard creates a new PassportElementTypeIdentityCard
//
func NewPassportElementTypeIdentityCard() *PassportElementTypeIdentityCard {
	passportElementTypeIdentityCardTemp := PassportElementTypeIdentityCard{
		tdCommon: tdCommon{Type: "passportElementTypeIdentityCard"},
	}

	return &passportElementTypeIdentityCardTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypeIdentityCard *PassportElementTypeIdentityCard) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypeIdentityCardType
}

// PassportElementTypeInternalPassport A Telegram Passport element containing the user's internal passport
type PassportElementTypeInternalPassport struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypeInternalPassport
func (passportElementTypeInternalPassport *PassportElementTypeInternalPassport) MessageType() string {
	return "passportElementTypeInternalPassport"
}

// NewPassportElementTypeInternalPassport creates a new PassportElementTypeInternalPassport
//
func NewPassportElementTypeInternalPassport() *PassportElementTypeInternalPassport {
	passportElementTypeInternalPassportTemp := PassportElementTypeInternalPassport{
		tdCommon: tdCommon{Type: "passportElementTypeInternalPassport"},
	}

	return &passportElementTypeInternalPassportTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypeInternalPassport *PassportElementTypeInternalPassport) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypeInternalPassportType
}

// PassportElementTypeAddress A Telegram Passport element containing the user's address
type PassportElementTypeAddress struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypeAddress
func (passportElementTypeAddress *PassportElementTypeAddress) MessageType() string {
	return "passportElementTypeAddress"
}

// NewPassportElementTypeAddress creates a new PassportElementTypeAddress
//
func NewPassportElementTypeAddress() *PassportElementTypeAddress {
	passportElementTypeAddressTemp := PassportElementTypeAddress{
		tdCommon: tdCommon{Type: "passportElementTypeAddress"},
	}

	return &passportElementTypeAddressTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypeAddress *PassportElementTypeAddress) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypeAddressType
}

// PassportElementTypeUtilityBill A Telegram Passport element containing the user's utility bill
type PassportElementTypeUtilityBill struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypeUtilityBill
func (passportElementTypeUtilityBill *PassportElementTypeUtilityBill) MessageType() string {
	return "passportElementTypeUtilityBill"
}

// NewPassportElementTypeUtilityBill creates a new PassportElementTypeUtilityBill
//
func NewPassportElementTypeUtilityBill() *PassportElementTypeUtilityBill {
	passportElementTypeUtilityBillTemp := PassportElementTypeUtilityBill{
		tdCommon: tdCommon{Type: "passportElementTypeUtilityBill"},
	}

	return &passportElementTypeUtilityBillTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypeUtilityBill *PassportElementTypeUtilityBill) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypeUtilityBillType
}

// PassportElementTypeBankStatement A Telegram Passport element containing the user's bank statement
type PassportElementTypeBankStatement struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypeBankStatement
func (passportElementTypeBankStatement *PassportElementTypeBankStatement) MessageType() string {
	return "passportElementTypeBankStatement"
}

// NewPassportElementTypeBankStatement creates a new PassportElementTypeBankStatement
//
func NewPassportElementTypeBankStatement() *PassportElementTypeBankStatement {
	passportElementTypeBankStatementTemp := PassportElementTypeBankStatement{
		tdCommon: tdCommon{Type: "passportElementTypeBankStatement"},
	}

	return &passportElementTypeBankStatementTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypeBankStatement *PassportElementTypeBankStatement) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypeBankStatementType
}

// PassportElementTypeRentalAgreement A Telegram Passport element containing the user's rental agreement
type PassportElementTypeRentalAgreement struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypeRentalAgreement
func (passportElementTypeRentalAgreement *PassportElementTypeRentalAgreement) MessageType() string {
	return "passportElementTypeRentalAgreement"
}

// NewPassportElementTypeRentalAgreement creates a new PassportElementTypeRentalAgreement
//
func NewPassportElementTypeRentalAgreement() *PassportElementTypeRentalAgreement {
	passportElementTypeRentalAgreementTemp := PassportElementTypeRentalAgreement{
		tdCommon: tdCommon{Type: "passportElementTypeRentalAgreement"},
	}

	return &passportElementTypeRentalAgreementTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypeRentalAgreement *PassportElementTypeRentalAgreement) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypeRentalAgreementType
}

// PassportElementTypePassportRegistration A Telegram Passport element containing the registration page of the user's passport
type PassportElementTypePassportRegistration struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypePassportRegistration
func (passportElementTypePassportRegistration *PassportElementTypePassportRegistration) MessageType() string {
	return "passportElementTypePassportRegistration"
}

// NewPassportElementTypePassportRegistration creates a new PassportElementTypePassportRegistration
//
func NewPassportElementTypePassportRegistration() *PassportElementTypePassportRegistration {
	passportElementTypePassportRegistrationTemp := PassportElementTypePassportRegistration{
		tdCommon: tdCommon{Type: "passportElementTypePassportRegistration"},
	}

	return &passportElementTypePassportRegistrationTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypePassportRegistration *PassportElementTypePassportRegistration) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypePassportRegistrationType
}

// PassportElementTypeTemporaryRegistration A Telegram Passport element containing the user's temporary registration
type PassportElementTypeTemporaryRegistration struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypeTemporaryRegistration
func (passportElementTypeTemporaryRegistration *PassportElementTypeTemporaryRegistration) MessageType() string {
	return "passportElementTypeTemporaryRegistration"
}

// NewPassportElementTypeTemporaryRegistration creates a new PassportElementTypeTemporaryRegistration
//
func NewPassportElementTypeTemporaryRegistration() *PassportElementTypeTemporaryRegistration {
	passportElementTypeTemporaryRegistrationTemp := PassportElementTypeTemporaryRegistration{
		tdCommon: tdCommon{Type: "passportElementTypeTemporaryRegistration"},
	}

	return &passportElementTypeTemporaryRegistrationTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypeTemporaryRegistration *PassportElementTypeTemporaryRegistration) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypeTemporaryRegistrationType
}

// PassportElementTypePhoneNumber A Telegram Passport element containing the user's phone number
type PassportElementTypePhoneNumber struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypePhoneNumber
func (passportElementTypePhoneNumber *PassportElementTypePhoneNumber) MessageType() string {
	return "passportElementTypePhoneNumber"
}

// NewPassportElementTypePhoneNumber creates a new PassportElementTypePhoneNumber
//
func NewPassportElementTypePhoneNumber() *PassportElementTypePhoneNumber {
	passportElementTypePhoneNumberTemp := PassportElementTypePhoneNumber{
		tdCommon: tdCommon{Type: "passportElementTypePhoneNumber"},
	}

	return &passportElementTypePhoneNumberTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypePhoneNumber *PassportElementTypePhoneNumber) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypePhoneNumberType
}

// PassportElementTypeEmailAddress A Telegram Passport element containing the user's email address
type PassportElementTypeEmailAddress struct {
	tdCommon
}

// MessageType return the string telegram-type of PassportElementTypeEmailAddress
func (passportElementTypeEmailAddress *PassportElementTypeEmailAddress) MessageType() string {
	return "passportElementTypeEmailAddress"
}

// NewPassportElementTypeEmailAddress creates a new PassportElementTypeEmailAddress
//
func NewPassportElementTypeEmailAddress() *PassportElementTypeEmailAddress {
	passportElementTypeEmailAddressTemp := PassportElementTypeEmailAddress{
		tdCommon: tdCommon{Type: "passportElementTypeEmailAddress"},
	}

	return &passportElementTypeEmailAddressTemp
}

// GetPassportElementTypeEnum return the enum type of this object
func (passportElementTypeEmailAddress *PassportElementTypeEmailAddress) GetPassportElementTypeEnum() PassportElementTypeEnum {
	return PassportElementTypeEmailAddressType
}
