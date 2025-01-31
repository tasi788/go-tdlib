package tdlib

import (
	"encoding/json"
	"fmt"
)

// PassportElement Contains information about a Telegram Passport element
type PassportElement interface {
	GetPassportElementEnum() PassportElementEnum
}

// PassportElementEnum Alias for abstract PassportElement 'Sub-Classes', used as constant-enum here
type PassportElementEnum string

// PassportElement enums
const (
	PassportElementPersonalDetailsType       PassportElementEnum = "passportElementPersonalDetails"
	PassportElementPassportType              PassportElementEnum = "passportElementPassport"
	PassportElementDriverLicenseType         PassportElementEnum = "passportElementDriverLicense"
	PassportElementIdentityCardType          PassportElementEnum = "passportElementIdentityCard"
	PassportElementInternalPassportType      PassportElementEnum = "passportElementInternalPassport"
	PassportElementAddressType               PassportElementEnum = "passportElementAddress"
	PassportElementUtilityBillType           PassportElementEnum = "passportElementUtilityBill"
	PassportElementBankStatementType         PassportElementEnum = "passportElementBankStatement"
	PassportElementRentalAgreementType       PassportElementEnum = "passportElementRentalAgreement"
	PassportElementPassportRegistrationType  PassportElementEnum = "passportElementPassportRegistration"
	PassportElementTemporaryRegistrationType PassportElementEnum = "passportElementTemporaryRegistration"
	PassportElementPhoneNumberType           PassportElementEnum = "passportElementPhoneNumber"
	PassportElementEmailAddressType          PassportElementEnum = "passportElementEmailAddress"
)

func unmarshalPassportElement(rawMsg *json.RawMessage) (PassportElement, error) {

	if rawMsg == nil {
		return nil, nil
	}
	var objMap map[string]interface{}
	err := json.Unmarshal(*rawMsg, &objMap)
	if err != nil {
		return nil, err
	}

	switch PassportElementEnum(objMap["@type"].(string)) {
	case PassportElementPersonalDetailsType:
		var passportElementPersonalDetails PassportElementPersonalDetails
		err := json.Unmarshal(*rawMsg, &passportElementPersonalDetails)
		return &passportElementPersonalDetails, err

	case PassportElementPassportType:
		var passportElementPassport PassportElementPassport
		err := json.Unmarshal(*rawMsg, &passportElementPassport)
		return &passportElementPassport, err

	case PassportElementDriverLicenseType:
		var passportElementDriverLicense PassportElementDriverLicense
		err := json.Unmarshal(*rawMsg, &passportElementDriverLicense)
		return &passportElementDriverLicense, err

	case PassportElementIdentityCardType:
		var passportElementIdentityCard PassportElementIdentityCard
		err := json.Unmarshal(*rawMsg, &passportElementIdentityCard)
		return &passportElementIdentityCard, err

	case PassportElementInternalPassportType:
		var passportElementInternalPassport PassportElementInternalPassport
		err := json.Unmarshal(*rawMsg, &passportElementInternalPassport)
		return &passportElementInternalPassport, err

	case PassportElementAddressType:
		var passportElementAddress PassportElementAddress
		err := json.Unmarshal(*rawMsg, &passportElementAddress)
		return &passportElementAddress, err

	case PassportElementUtilityBillType:
		var passportElementUtilityBill PassportElementUtilityBill
		err := json.Unmarshal(*rawMsg, &passportElementUtilityBill)
		return &passportElementUtilityBill, err

	case PassportElementBankStatementType:
		var passportElementBankStatement PassportElementBankStatement
		err := json.Unmarshal(*rawMsg, &passportElementBankStatement)
		return &passportElementBankStatement, err

	case PassportElementRentalAgreementType:
		var passportElementRentalAgreement PassportElementRentalAgreement
		err := json.Unmarshal(*rawMsg, &passportElementRentalAgreement)
		return &passportElementRentalAgreement, err

	case PassportElementPassportRegistrationType:
		var passportElementPassportRegistration PassportElementPassportRegistration
		err := json.Unmarshal(*rawMsg, &passportElementPassportRegistration)
		return &passportElementPassportRegistration, err

	case PassportElementTemporaryRegistrationType:
		var passportElementTemporaryRegistration PassportElementTemporaryRegistration
		err := json.Unmarshal(*rawMsg, &passportElementTemporaryRegistration)
		return &passportElementTemporaryRegistration, err

	case PassportElementPhoneNumberType:
		var passportElementPhoneNumber PassportElementPhoneNumber
		err := json.Unmarshal(*rawMsg, &passportElementPhoneNumber)
		return &passportElementPhoneNumber, err

	case PassportElementEmailAddressType:
		var passportElementEmailAddress PassportElementEmailAddress
		err := json.Unmarshal(*rawMsg, &passportElementEmailAddress)
		return &passportElementEmailAddress, err

	default:
		return nil, fmt.Errorf("Error UnMarshaling, unknown type:" + objMap["@type"].(string))
	}
}

// PassportElementPersonalDetails A Telegram Passport element containing the user's personal details
type PassportElementPersonalDetails struct {
	tdCommon
	PersonalDetails *PersonalDetails `json:"personal_details"` // Personal details of the user
}

// MessageType return the string telegram-type of PassportElementPersonalDetails
func (passportElementPersonalDetails *PassportElementPersonalDetails) MessageType() string {
	return "passportElementPersonalDetails"
}

// NewPassportElementPersonalDetails creates a new PassportElementPersonalDetails
//
// @param personalDetails Personal details of the user
func NewPassportElementPersonalDetails(personalDetails *PersonalDetails) *PassportElementPersonalDetails {
	passportElementPersonalDetailsTemp := PassportElementPersonalDetails{
		tdCommon:        tdCommon{Type: "passportElementPersonalDetails"},
		PersonalDetails: personalDetails,
	}

	return &passportElementPersonalDetailsTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementPersonalDetails *PassportElementPersonalDetails) GetPassportElementEnum() PassportElementEnum {
	return PassportElementPersonalDetailsType
}

// PassportElementPassport A Telegram Passport element containing the user's passport
type PassportElementPassport struct {
	tdCommon
	Passport *IdentityDocument `json:"passport"` // Passport
}

// MessageType return the string telegram-type of PassportElementPassport
func (passportElementPassport *PassportElementPassport) MessageType() string {
	return "passportElementPassport"
}

// NewPassportElementPassport creates a new PassportElementPassport
//
// @param passport Passport
func NewPassportElementPassport(passport *IdentityDocument) *PassportElementPassport {
	passportElementPassportTemp := PassportElementPassport{
		tdCommon: tdCommon{Type: "passportElementPassport"},
		Passport: passport,
	}

	return &passportElementPassportTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementPassport *PassportElementPassport) GetPassportElementEnum() PassportElementEnum {
	return PassportElementPassportType
}

// PassportElementDriverLicense A Telegram Passport element containing the user's driver license
type PassportElementDriverLicense struct {
	tdCommon
	DriverLicense *IdentityDocument `json:"driver_license"` // Driver license
}

// MessageType return the string telegram-type of PassportElementDriverLicense
func (passportElementDriverLicense *PassportElementDriverLicense) MessageType() string {
	return "passportElementDriverLicense"
}

// NewPassportElementDriverLicense creates a new PassportElementDriverLicense
//
// @param driverLicense Driver license
func NewPassportElementDriverLicense(driverLicense *IdentityDocument) *PassportElementDriverLicense {
	passportElementDriverLicenseTemp := PassportElementDriverLicense{
		tdCommon:      tdCommon{Type: "passportElementDriverLicense"},
		DriverLicense: driverLicense,
	}

	return &passportElementDriverLicenseTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementDriverLicense *PassportElementDriverLicense) GetPassportElementEnum() PassportElementEnum {
	return PassportElementDriverLicenseType
}

// PassportElementIdentityCard A Telegram Passport element containing the user's identity card
type PassportElementIdentityCard struct {
	tdCommon
	IdentityCard *IdentityDocument `json:"identity_card"` // Identity card
}

// MessageType return the string telegram-type of PassportElementIdentityCard
func (passportElementIdentityCard *PassportElementIdentityCard) MessageType() string {
	return "passportElementIdentityCard"
}

// NewPassportElementIdentityCard creates a new PassportElementIdentityCard
//
// @param identityCard Identity card
func NewPassportElementIdentityCard(identityCard *IdentityDocument) *PassportElementIdentityCard {
	passportElementIdentityCardTemp := PassportElementIdentityCard{
		tdCommon:     tdCommon{Type: "passportElementIdentityCard"},
		IdentityCard: identityCard,
	}

	return &passportElementIdentityCardTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementIdentityCard *PassportElementIdentityCard) GetPassportElementEnum() PassportElementEnum {
	return PassportElementIdentityCardType
}

// PassportElementInternalPassport A Telegram Passport element containing the user's internal passport
type PassportElementInternalPassport struct {
	tdCommon
	InternalPassport *IdentityDocument `json:"internal_passport"` // Internal passport
}

// MessageType return the string telegram-type of PassportElementInternalPassport
func (passportElementInternalPassport *PassportElementInternalPassport) MessageType() string {
	return "passportElementInternalPassport"
}

// NewPassportElementInternalPassport creates a new PassportElementInternalPassport
//
// @param internalPassport Internal passport
func NewPassportElementInternalPassport(internalPassport *IdentityDocument) *PassportElementInternalPassport {
	passportElementInternalPassportTemp := PassportElementInternalPassport{
		tdCommon:         tdCommon{Type: "passportElementInternalPassport"},
		InternalPassport: internalPassport,
	}

	return &passportElementInternalPassportTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementInternalPassport *PassportElementInternalPassport) GetPassportElementEnum() PassportElementEnum {
	return PassportElementInternalPassportType
}

// PassportElementAddress A Telegram Passport element containing the user's address
type PassportElementAddress struct {
	tdCommon
	Address *Address `json:"address"` // Address
}

// MessageType return the string telegram-type of PassportElementAddress
func (passportElementAddress *PassportElementAddress) MessageType() string {
	return "passportElementAddress"
}

// NewPassportElementAddress creates a new PassportElementAddress
//
// @param address Address
func NewPassportElementAddress(address *Address) *PassportElementAddress {
	passportElementAddressTemp := PassportElementAddress{
		tdCommon: tdCommon{Type: "passportElementAddress"},
		Address:  address,
	}

	return &passportElementAddressTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementAddress *PassportElementAddress) GetPassportElementEnum() PassportElementEnum {
	return PassportElementAddressType
}

// PassportElementUtilityBill A Telegram Passport element containing the user's utility bill
type PassportElementUtilityBill struct {
	tdCommon
	UtilityBill *PersonalDocument `json:"utility_bill"` // Utility bill
}

// MessageType return the string telegram-type of PassportElementUtilityBill
func (passportElementUtilityBill *PassportElementUtilityBill) MessageType() string {
	return "passportElementUtilityBill"
}

// NewPassportElementUtilityBill creates a new PassportElementUtilityBill
//
// @param utilityBill Utility bill
func NewPassportElementUtilityBill(utilityBill *PersonalDocument) *PassportElementUtilityBill {
	passportElementUtilityBillTemp := PassportElementUtilityBill{
		tdCommon:    tdCommon{Type: "passportElementUtilityBill"},
		UtilityBill: utilityBill,
	}

	return &passportElementUtilityBillTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementUtilityBill *PassportElementUtilityBill) GetPassportElementEnum() PassportElementEnum {
	return PassportElementUtilityBillType
}

// PassportElementBankStatement A Telegram Passport element containing the user's bank statement
type PassportElementBankStatement struct {
	tdCommon
	BankStatement *PersonalDocument `json:"bank_statement"` // Bank statement
}

// MessageType return the string telegram-type of PassportElementBankStatement
func (passportElementBankStatement *PassportElementBankStatement) MessageType() string {
	return "passportElementBankStatement"
}

// NewPassportElementBankStatement creates a new PassportElementBankStatement
//
// @param bankStatement Bank statement
func NewPassportElementBankStatement(bankStatement *PersonalDocument) *PassportElementBankStatement {
	passportElementBankStatementTemp := PassportElementBankStatement{
		tdCommon:      tdCommon{Type: "passportElementBankStatement"},
		BankStatement: bankStatement,
	}

	return &passportElementBankStatementTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementBankStatement *PassportElementBankStatement) GetPassportElementEnum() PassportElementEnum {
	return PassportElementBankStatementType
}

// PassportElementRentalAgreement A Telegram Passport element containing the user's rental agreement
type PassportElementRentalAgreement struct {
	tdCommon
	RentalAgreement *PersonalDocument `json:"rental_agreement"` // Rental agreement
}

// MessageType return the string telegram-type of PassportElementRentalAgreement
func (passportElementRentalAgreement *PassportElementRentalAgreement) MessageType() string {
	return "passportElementRentalAgreement"
}

// NewPassportElementRentalAgreement creates a new PassportElementRentalAgreement
//
// @param rentalAgreement Rental agreement
func NewPassportElementRentalAgreement(rentalAgreement *PersonalDocument) *PassportElementRentalAgreement {
	passportElementRentalAgreementTemp := PassportElementRentalAgreement{
		tdCommon:        tdCommon{Type: "passportElementRentalAgreement"},
		RentalAgreement: rentalAgreement,
	}

	return &passportElementRentalAgreementTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementRentalAgreement *PassportElementRentalAgreement) GetPassportElementEnum() PassportElementEnum {
	return PassportElementRentalAgreementType
}

// PassportElementPassportRegistration A Telegram Passport element containing the user's passport registration pages
type PassportElementPassportRegistration struct {
	tdCommon
	PassportRegistration *PersonalDocument `json:"passport_registration"` // Passport registration pages
}

// MessageType return the string telegram-type of PassportElementPassportRegistration
func (passportElementPassportRegistration *PassportElementPassportRegistration) MessageType() string {
	return "passportElementPassportRegistration"
}

// NewPassportElementPassportRegistration creates a new PassportElementPassportRegistration
//
// @param passportRegistration Passport registration pages
func NewPassportElementPassportRegistration(passportRegistration *PersonalDocument) *PassportElementPassportRegistration {
	passportElementPassportRegistrationTemp := PassportElementPassportRegistration{
		tdCommon:             tdCommon{Type: "passportElementPassportRegistration"},
		PassportRegistration: passportRegistration,
	}

	return &passportElementPassportRegistrationTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementPassportRegistration *PassportElementPassportRegistration) GetPassportElementEnum() PassportElementEnum {
	return PassportElementPassportRegistrationType
}

// PassportElementTemporaryRegistration A Telegram Passport element containing the user's temporary registration
type PassportElementTemporaryRegistration struct {
	tdCommon
	TemporaryRegistration *PersonalDocument `json:"temporary_registration"` // Temporary registration
}

// MessageType return the string telegram-type of PassportElementTemporaryRegistration
func (passportElementTemporaryRegistration *PassportElementTemporaryRegistration) MessageType() string {
	return "passportElementTemporaryRegistration"
}

// NewPassportElementTemporaryRegistration creates a new PassportElementTemporaryRegistration
//
// @param temporaryRegistration Temporary registration
func NewPassportElementTemporaryRegistration(temporaryRegistration *PersonalDocument) *PassportElementTemporaryRegistration {
	passportElementTemporaryRegistrationTemp := PassportElementTemporaryRegistration{
		tdCommon:              tdCommon{Type: "passportElementTemporaryRegistration"},
		TemporaryRegistration: temporaryRegistration,
	}

	return &passportElementTemporaryRegistrationTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementTemporaryRegistration *PassportElementTemporaryRegistration) GetPassportElementEnum() PassportElementEnum {
	return PassportElementTemporaryRegistrationType
}

// PassportElementPhoneNumber A Telegram Passport element containing the user's phone number
type PassportElementPhoneNumber struct {
	tdCommon
	PhoneNumber string `json:"phone_number"` // Phone number
}

// MessageType return the string telegram-type of PassportElementPhoneNumber
func (passportElementPhoneNumber *PassportElementPhoneNumber) MessageType() string {
	return "passportElementPhoneNumber"
}

// NewPassportElementPhoneNumber creates a new PassportElementPhoneNumber
//
// @param phoneNumber Phone number
func NewPassportElementPhoneNumber(phoneNumber string) *PassportElementPhoneNumber {
	passportElementPhoneNumberTemp := PassportElementPhoneNumber{
		tdCommon:    tdCommon{Type: "passportElementPhoneNumber"},
		PhoneNumber: phoneNumber,
	}

	return &passportElementPhoneNumberTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementPhoneNumber *PassportElementPhoneNumber) GetPassportElementEnum() PassportElementEnum {
	return PassportElementPhoneNumberType
}

// PassportElementEmailAddress A Telegram Passport element containing the user's email address
type PassportElementEmailAddress struct {
	tdCommon
	EmailAddress string `json:"email_address"` // Email address
}

// MessageType return the string telegram-type of PassportElementEmailAddress
func (passportElementEmailAddress *PassportElementEmailAddress) MessageType() string {
	return "passportElementEmailAddress"
}

// NewPassportElementEmailAddress creates a new PassportElementEmailAddress
//
// @param emailAddress Email address
func NewPassportElementEmailAddress(emailAddress string) *PassportElementEmailAddress {
	passportElementEmailAddressTemp := PassportElementEmailAddress{
		tdCommon:     tdCommon{Type: "passportElementEmailAddress"},
		EmailAddress: emailAddress,
	}

	return &passportElementEmailAddressTemp
}

// GetPassportElementEnum return the enum type of this object
func (passportElementEmailAddress *PassportElementEmailAddress) GetPassportElementEnum() PassportElementEnum {
	return PassportElementEmailAddressType
}

// GetPassportElement Returns one of the available Telegram Passport elements
// @param typeParam Telegram Passport element type
// @param password Password of the current user
func (client *Client) GetPassportElement(typeParam PassportElementType, password string) (PassportElement, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getPassportElement",
		"type":     typeParam,
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch PassportElementEnum(result.Data["@type"].(string)) {

	case PassportElementPersonalDetailsType:
		var passportElement PassportElementPersonalDetails
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPassportType:
		var passportElement PassportElementPassport
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementDriverLicenseType:
		var passportElement PassportElementDriverLicense
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementIdentityCardType:
		var passportElement PassportElementIdentityCard
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementInternalPassportType:
		var passportElement PassportElementInternalPassport
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementAddressType:
		var passportElement PassportElementAddress
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementUtilityBillType:
		var passportElement PassportElementUtilityBill
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementBankStatementType:
		var passportElement PassportElementBankStatement
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementRentalAgreementType:
		var passportElement PassportElementRentalAgreement
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPassportRegistrationType:
		var passportElement PassportElementPassportRegistration
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementTemporaryRegistrationType:
		var passportElement PassportElementTemporaryRegistration
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPhoneNumberType:
		var passportElement PassportElementPhoneNumber
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementEmailAddressType:
		var passportElement PassportElementEmailAddress
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// SetPassportElement Adds an element to the user's Telegram Passport. May return an error with a message "PHONE_VERIFICATION_NEEDED" or "EMAIL_VERIFICATION_NEEDED" if the chosen phone number or the chosen email address must be verified first
// @param element Input Telegram Passport element
// @param password Password of the current user
func (client *Client) SetPassportElement(element InputPassportElement, password string) (PassportElement, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setPassportElement",
		"element":  element,
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %v msg: %s", result.Data["code"], result.Data["message"])
	}

	switch PassportElementEnum(result.Data["@type"].(string)) {

	case PassportElementPersonalDetailsType:
		var passportElement PassportElementPersonalDetails
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPassportType:
		var passportElement PassportElementPassport
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementDriverLicenseType:
		var passportElement PassportElementDriverLicense
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementIdentityCardType:
		var passportElement PassportElementIdentityCard
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementInternalPassportType:
		var passportElement PassportElementInternalPassport
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementAddressType:
		var passportElement PassportElementAddress
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementUtilityBillType:
		var passportElement PassportElementUtilityBill
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementBankStatementType:
		var passportElement PassportElementBankStatement
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementRentalAgreementType:
		var passportElement PassportElementRentalAgreement
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPassportRegistrationType:
		var passportElement PassportElementPassportRegistration
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementTemporaryRegistrationType:
		var passportElement PassportElementTemporaryRegistration
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPhoneNumberType:
		var passportElement PassportElementPhoneNumber
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementEmailAddressType:
		var passportElement PassportElementEmailAddress
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
