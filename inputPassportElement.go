package tdlib

import (
	"encoding/json"
	"fmt"
)

// InputPassportElement Contains information about a Telegram Passport element to be saved
type InputPassportElement interface {
	GetInputPassportElementEnum() InputPassportElementEnum
}

// InputPassportElementEnum Alias for abstract InputPassportElement 'Sub-Classes', used as constant-enum here
type InputPassportElementEnum string

// InputPassportElement enums
const (
	InputPassportElementPersonalDetailsType       InputPassportElementEnum = "inputPassportElementPersonalDetails"
	InputPassportElementPassportType              InputPassportElementEnum = "inputPassportElementPassport"
	InputPassportElementDriverLicenseType         InputPassportElementEnum = "inputPassportElementDriverLicense"
	InputPassportElementIdentityCardType          InputPassportElementEnum = "inputPassportElementIdentityCard"
	InputPassportElementInternalPassportType      InputPassportElementEnum = "inputPassportElementInternalPassport"
	InputPassportElementAddressType               InputPassportElementEnum = "inputPassportElementAddress"
	InputPassportElementUtilityBillType           InputPassportElementEnum = "inputPassportElementUtilityBill"
	InputPassportElementBankStatementType         InputPassportElementEnum = "inputPassportElementBankStatement"
	InputPassportElementRentalAgreementType       InputPassportElementEnum = "inputPassportElementRentalAgreement"
	InputPassportElementPassportRegistrationType  InputPassportElementEnum = "inputPassportElementPassportRegistration"
	InputPassportElementTemporaryRegistrationType InputPassportElementEnum = "inputPassportElementTemporaryRegistration"
	InputPassportElementPhoneNumberType           InputPassportElementEnum = "inputPassportElementPhoneNumber"
	InputPassportElementEmailAddressType          InputPassportElementEnum = "inputPassportElementEmailAddress"
)

func unmarshalInputPassportElement(rawMsg *json.RawMessage) (InputPassportElement, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch InputPassportElementEnum(objMap["@type"].(string)) {
	case InputPassportElementPersonalDetailsType:
		var inputPassportElementPersonalDetails InputPassportElementPersonalDetails
		err := json.Unmarshal(*rawMsg, &inputPassportElementPersonalDetails)
		return &inputPassportElementPersonalDetails, err

	case InputPassportElementPassportType:
		var inputPassportElementPassport InputPassportElementPassport
		err := json.Unmarshal(*rawMsg, &inputPassportElementPassport)
		return &inputPassportElementPassport, err

	case InputPassportElementDriverLicenseType:
		var inputPassportElementDriverLicense InputPassportElementDriverLicense
		err := json.Unmarshal(*rawMsg, &inputPassportElementDriverLicense)
		return &inputPassportElementDriverLicense, err

	case InputPassportElementIdentityCardType:
		var inputPassportElementIdentityCard InputPassportElementIdentityCard
		err := json.Unmarshal(*rawMsg, &inputPassportElementIdentityCard)
		return &inputPassportElementIdentityCard, err

	case InputPassportElementInternalPassportType:
		var inputPassportElementInternalPassport InputPassportElementInternalPassport
		err := json.Unmarshal(*rawMsg, &inputPassportElementInternalPassport)
		return &inputPassportElementInternalPassport, err

	case InputPassportElementAddressType:
		var inputPassportElementAddress InputPassportElementAddress
		err := json.Unmarshal(*rawMsg, &inputPassportElementAddress)
		return &inputPassportElementAddress, err

	case InputPassportElementUtilityBillType:
		var inputPassportElementUtilityBill InputPassportElementUtilityBill
		err := json.Unmarshal(*rawMsg, &inputPassportElementUtilityBill)
		return &inputPassportElementUtilityBill, err

	case InputPassportElementBankStatementType:
		var inputPassportElementBankStatement InputPassportElementBankStatement
		err := json.Unmarshal(*rawMsg, &inputPassportElementBankStatement)
		return &inputPassportElementBankStatement, err

	case InputPassportElementRentalAgreementType:
		var inputPassportElementRentalAgreement InputPassportElementRentalAgreement
		err := json.Unmarshal(*rawMsg, &inputPassportElementRentalAgreement)
		return &inputPassportElementRentalAgreement, err

	case InputPassportElementPassportRegistrationType:
		var inputPassportElementPassportRegistration InputPassportElementPassportRegistration
		err := json.Unmarshal(*rawMsg, &inputPassportElementPassportRegistration)
		return &inputPassportElementPassportRegistration, err

	case InputPassportElementTemporaryRegistrationType:
		var inputPassportElementTemporaryRegistration InputPassportElementTemporaryRegistration
		err := json.Unmarshal(*rawMsg, &inputPassportElementTemporaryRegistration)
		return &inputPassportElementTemporaryRegistration, err

	case InputPassportElementPhoneNumberType:
		var inputPassportElementPhoneNumber InputPassportElementPhoneNumber
		err := json.Unmarshal(*rawMsg, &inputPassportElementPhoneNumber)
		return &inputPassportElementPhoneNumber, err

	case InputPassportElementEmailAddressType:
		var inputPassportElementEmailAddress InputPassportElementEmailAddress
		err := json.Unmarshal(*rawMsg, &inputPassportElementEmailAddress)
		return &inputPassportElementEmailAddress, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// InputPassportElementPersonalDetails A Telegram Passport element to be saved containing the user's personal details
type InputPassportElementPersonalDetails struct {
	tdCommon
	PersonalDetails *PersonalDetails `json:"personal_details"` // Personal details of the user
}

// MessageType return the string telegram-type of InputPassportElementPersonalDetails
func (inputPassportElementPersonalDetails *InputPassportElementPersonalDetails) MessageType() string {
	return "inputPassportElementPersonalDetails"
}

// NewInputPassportElementPersonalDetails creates a new InputPassportElementPersonalDetails
//
// @param personalDetails Personal details of the user
func NewInputPassportElementPersonalDetails(personalDetails *PersonalDetails) *InputPassportElementPersonalDetails {
	inputPassportElementPersonalDetailsTemp := InputPassportElementPersonalDetails{
		tdCommon:        tdCommon{Type: "inputPassportElementPersonalDetails"},
		PersonalDetails: personalDetails,
	}

	return &inputPassportElementPersonalDetailsTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementPersonalDetails *InputPassportElementPersonalDetails) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementPersonalDetailsType
}

// InputPassportElementPassport A Telegram Passport element to be saved containing the user's passport
type InputPassportElementPassport struct {
	tdCommon
	Passport *InputIdentityDocument `json:"passport"` // The passport to be saved
}

// MessageType return the string telegram-type of InputPassportElementPassport
func (inputPassportElementPassport *InputPassportElementPassport) MessageType() string {
	return "inputPassportElementPassport"
}

// NewInputPassportElementPassport creates a new InputPassportElementPassport
//
// @param passport The passport to be saved
func NewInputPassportElementPassport(passport *InputIdentityDocument) *InputPassportElementPassport {
	inputPassportElementPassportTemp := InputPassportElementPassport{
		tdCommon: tdCommon{Type: "inputPassportElementPassport"},
		Passport: passport,
	}

	return &inputPassportElementPassportTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementPassport *InputPassportElementPassport) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementPassportType
}

// InputPassportElementDriverLicense A Telegram Passport element to be saved containing the user's driver license
type InputPassportElementDriverLicense struct {
	tdCommon
	DriverLicense *InputIdentityDocument `json:"driver_license"` // The driver license to be saved
}

// MessageType return the string telegram-type of InputPassportElementDriverLicense
func (inputPassportElementDriverLicense *InputPassportElementDriverLicense) MessageType() string {
	return "inputPassportElementDriverLicense"
}

// NewInputPassportElementDriverLicense creates a new InputPassportElementDriverLicense
//
// @param driverLicense The driver license to be saved
func NewInputPassportElementDriverLicense(driverLicense *InputIdentityDocument) *InputPassportElementDriverLicense {
	inputPassportElementDriverLicenseTemp := InputPassportElementDriverLicense{
		tdCommon:      tdCommon{Type: "inputPassportElementDriverLicense"},
		DriverLicense: driverLicense,
	}

	return &inputPassportElementDriverLicenseTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementDriverLicense *InputPassportElementDriverLicense) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementDriverLicenseType
}

// InputPassportElementIdentityCard A Telegram Passport element to be saved containing the user's identity card
type InputPassportElementIdentityCard struct {
	tdCommon
	IdentityCard *InputIdentityDocument `json:"identity_card"` // The identity card to be saved
}

// MessageType return the string telegram-type of InputPassportElementIdentityCard
func (inputPassportElementIdentityCard *InputPassportElementIdentityCard) MessageType() string {
	return "inputPassportElementIdentityCard"
}

// NewInputPassportElementIdentityCard creates a new InputPassportElementIdentityCard
//
// @param identityCard The identity card to be saved
func NewInputPassportElementIdentityCard(identityCard *InputIdentityDocument) *InputPassportElementIdentityCard {
	inputPassportElementIdentityCardTemp := InputPassportElementIdentityCard{
		tdCommon:     tdCommon{Type: "inputPassportElementIdentityCard"},
		IdentityCard: identityCard,
	}

	return &inputPassportElementIdentityCardTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementIdentityCard *InputPassportElementIdentityCard) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementIdentityCardType
}

// InputPassportElementInternalPassport A Telegram Passport element to be saved containing the user's internal passport
type InputPassportElementInternalPassport struct {
	tdCommon
	InternalPassport *InputIdentityDocument `json:"internal_passport"` // The internal passport to be saved
}

// MessageType return the string telegram-type of InputPassportElementInternalPassport
func (inputPassportElementInternalPassport *InputPassportElementInternalPassport) MessageType() string {
	return "inputPassportElementInternalPassport"
}

// NewInputPassportElementInternalPassport creates a new InputPassportElementInternalPassport
//
// @param internalPassport The internal passport to be saved
func NewInputPassportElementInternalPassport(internalPassport *InputIdentityDocument) *InputPassportElementInternalPassport {
	inputPassportElementInternalPassportTemp := InputPassportElementInternalPassport{
		tdCommon:         tdCommon{Type: "inputPassportElementInternalPassport"},
		InternalPassport: internalPassport,
	}

	return &inputPassportElementInternalPassportTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementInternalPassport *InputPassportElementInternalPassport) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementInternalPassportType
}

// InputPassportElementAddress A Telegram Passport element to be saved containing the user's address
type InputPassportElementAddress struct {
	tdCommon
	Address *Address `json:"address"` // The address to be saved
}

// MessageType return the string telegram-type of InputPassportElementAddress
func (inputPassportElementAddress *InputPassportElementAddress) MessageType() string {
	return "inputPassportElementAddress"
}

// NewInputPassportElementAddress creates a new InputPassportElementAddress
//
// @param address The address to be saved
func NewInputPassportElementAddress(address *Address) *InputPassportElementAddress {
	inputPassportElementAddressTemp := InputPassportElementAddress{
		tdCommon: tdCommon{Type: "inputPassportElementAddress"},
		Address:  address,
	}

	return &inputPassportElementAddressTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementAddress *InputPassportElementAddress) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementAddressType
}

// InputPassportElementUtilityBill A Telegram Passport element to be saved containing the user's utility bill
type InputPassportElementUtilityBill struct {
	tdCommon
	UtilityBill *InputPersonalDocument `json:"utility_bill"` // The utility bill to be saved
}

// MessageType return the string telegram-type of InputPassportElementUtilityBill
func (inputPassportElementUtilityBill *InputPassportElementUtilityBill) MessageType() string {
	return "inputPassportElementUtilityBill"
}

// NewInputPassportElementUtilityBill creates a new InputPassportElementUtilityBill
//
// @param utilityBill The utility bill to be saved
func NewInputPassportElementUtilityBill(utilityBill *InputPersonalDocument) *InputPassportElementUtilityBill {
	inputPassportElementUtilityBillTemp := InputPassportElementUtilityBill{
		tdCommon:    tdCommon{Type: "inputPassportElementUtilityBill"},
		UtilityBill: utilityBill,
	}

	return &inputPassportElementUtilityBillTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementUtilityBill *InputPassportElementUtilityBill) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementUtilityBillType
}

// InputPassportElementBankStatement A Telegram Passport element to be saved containing the user's bank statement
type InputPassportElementBankStatement struct {
	tdCommon
	BankStatement *InputPersonalDocument `json:"bank_statement"` // The bank statement to be saved
}

// MessageType return the string telegram-type of InputPassportElementBankStatement
func (inputPassportElementBankStatement *InputPassportElementBankStatement) MessageType() string {
	return "inputPassportElementBankStatement"
}

// NewInputPassportElementBankStatement creates a new InputPassportElementBankStatement
//
// @param bankStatement The bank statement to be saved
func NewInputPassportElementBankStatement(bankStatement *InputPersonalDocument) *InputPassportElementBankStatement {
	inputPassportElementBankStatementTemp := InputPassportElementBankStatement{
		tdCommon:      tdCommon{Type: "inputPassportElementBankStatement"},
		BankStatement: bankStatement,
	}

	return &inputPassportElementBankStatementTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementBankStatement *InputPassportElementBankStatement) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementBankStatementType
}

// InputPassportElementRentalAgreement A Telegram Passport element to be saved containing the user's rental agreement
type InputPassportElementRentalAgreement struct {
	tdCommon
	RentalAgreement *InputPersonalDocument `json:"rental_agreement"` // The rental agreement to be saved
}

// MessageType return the string telegram-type of InputPassportElementRentalAgreement
func (inputPassportElementRentalAgreement *InputPassportElementRentalAgreement) MessageType() string {
	return "inputPassportElementRentalAgreement"
}

// NewInputPassportElementRentalAgreement creates a new InputPassportElementRentalAgreement
//
// @param rentalAgreement The rental agreement to be saved
func NewInputPassportElementRentalAgreement(rentalAgreement *InputPersonalDocument) *InputPassportElementRentalAgreement {
	inputPassportElementRentalAgreementTemp := InputPassportElementRentalAgreement{
		tdCommon:        tdCommon{Type: "inputPassportElementRentalAgreement"},
		RentalAgreement: rentalAgreement,
	}

	return &inputPassportElementRentalAgreementTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementRentalAgreement *InputPassportElementRentalAgreement) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementRentalAgreementType
}

// InputPassportElementPassportRegistration A Telegram Passport element to be saved containing the user's passport registration
type InputPassportElementPassportRegistration struct {
	tdCommon
	PassportRegistration *InputPersonalDocument `json:"passport_registration"` // The passport registration page to be saved
}

// MessageType return the string telegram-type of InputPassportElementPassportRegistration
func (inputPassportElementPassportRegistration *InputPassportElementPassportRegistration) MessageType() string {
	return "inputPassportElementPassportRegistration"
}

// NewInputPassportElementPassportRegistration creates a new InputPassportElementPassportRegistration
//
// @param passportRegistration The passport registration page to be saved
func NewInputPassportElementPassportRegistration(passportRegistration *InputPersonalDocument) *InputPassportElementPassportRegistration {
	inputPassportElementPassportRegistrationTemp := InputPassportElementPassportRegistration{
		tdCommon:             tdCommon{Type: "inputPassportElementPassportRegistration"},
		PassportRegistration: passportRegistration,
	}

	return &inputPassportElementPassportRegistrationTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementPassportRegistration *InputPassportElementPassportRegistration) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementPassportRegistrationType
}

// InputPassportElementTemporaryRegistration A Telegram Passport element to be saved containing the user's temporary registration
type InputPassportElementTemporaryRegistration struct {
	tdCommon
	TemporaryRegistration *InputPersonalDocument `json:"temporary_registration"` // The temporary registration document to be saved
}

// MessageType return the string telegram-type of InputPassportElementTemporaryRegistration
func (inputPassportElementTemporaryRegistration *InputPassportElementTemporaryRegistration) MessageType() string {
	return "inputPassportElementTemporaryRegistration"
}

// NewInputPassportElementTemporaryRegistration creates a new InputPassportElementTemporaryRegistration
//
// @param temporaryRegistration The temporary registration document to be saved
func NewInputPassportElementTemporaryRegistration(temporaryRegistration *InputPersonalDocument) *InputPassportElementTemporaryRegistration {
	inputPassportElementTemporaryRegistrationTemp := InputPassportElementTemporaryRegistration{
		tdCommon:              tdCommon{Type: "inputPassportElementTemporaryRegistration"},
		TemporaryRegistration: temporaryRegistration,
	}

	return &inputPassportElementTemporaryRegistrationTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementTemporaryRegistration *InputPassportElementTemporaryRegistration) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementTemporaryRegistrationType
}

// InputPassportElementPhoneNumber A Telegram Passport element to be saved containing the user's phone number
type InputPassportElementPhoneNumber struct {
	tdCommon
	PhoneNumber string `json:"phone_number"` // The phone number to be saved
}

// MessageType return the string telegram-type of InputPassportElementPhoneNumber
func (inputPassportElementPhoneNumber *InputPassportElementPhoneNumber) MessageType() string {
	return "inputPassportElementPhoneNumber"
}

// NewInputPassportElementPhoneNumber creates a new InputPassportElementPhoneNumber
//
// @param phoneNumber The phone number to be saved
func NewInputPassportElementPhoneNumber(phoneNumber string) *InputPassportElementPhoneNumber {
	inputPassportElementPhoneNumberTemp := InputPassportElementPhoneNumber{
		tdCommon:    tdCommon{Type: "inputPassportElementPhoneNumber"},
		PhoneNumber: phoneNumber,
	}

	return &inputPassportElementPhoneNumberTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementPhoneNumber *InputPassportElementPhoneNumber) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementPhoneNumberType
}

// InputPassportElementEmailAddress A Telegram Passport element to be saved containing the user's email address
type InputPassportElementEmailAddress struct {
	tdCommon
	EmailAddress string `json:"email_address"` // The email address to be saved
}

// MessageType return the string telegram-type of InputPassportElementEmailAddress
func (inputPassportElementEmailAddress *InputPassportElementEmailAddress) MessageType() string {
	return "inputPassportElementEmailAddress"
}

// NewInputPassportElementEmailAddress creates a new InputPassportElementEmailAddress
//
// @param emailAddress The email address to be saved
func NewInputPassportElementEmailAddress(emailAddress string) *InputPassportElementEmailAddress {
	inputPassportElementEmailAddressTemp := InputPassportElementEmailAddress{
		tdCommon:     tdCommon{Type: "inputPassportElementEmailAddress"},
		EmailAddress: emailAddress,
	}

	return &inputPassportElementEmailAddressTemp
}

// GetInputPassportElementEnum return the enum type of this object
func (inputPassportElementEmailAddress *InputPassportElementEmailAddress) GetInputPassportElementEnum() InputPassportElementEnum {
	return InputPassportElementEmailAddressType
}
