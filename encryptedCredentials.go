package tdlib

// EncryptedCredentials Contains encrypted Telegram Passport data credentials
type EncryptedCredentials struct {
	tdCommon
	Data   []byte `json:"data"`   // The encrypted credentials
	Hash   []byte `json:"hash"`   // The decrypted data hash
	Secret []byte `json:"secret"` // Secret for data decryption, encrypted with the service's public key
}

// MessageType return the string telegram-type of EncryptedCredentials
func (encryptedCredentials *EncryptedCredentials) MessageType() string {
	return "encryptedCredentials"
}

// NewEncryptedCredentials creates a new EncryptedCredentials
//
// @param data The encrypted credentials
// @param hash The decrypted data hash
// @param secret Secret for data decryption, encrypted with the service's public key
func NewEncryptedCredentials(data []byte, hash []byte, secret []byte) *EncryptedCredentials {
	encryptedCredentialsTemp := EncryptedCredentials{
		tdCommon: tdCommon{Type: "encryptedCredentials"},
		Data:     data,
		Hash:     hash,
		Secret:   secret,
	}

	return &encryptedCredentialsTemp
}
