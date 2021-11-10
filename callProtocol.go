package tdlib

// CallProtocol Specifies the supported call protocols
type CallProtocol struct {
	tdCommon
	UdpP2p          bool     `json:"udp_p2p"`          // True, if UDP peer-to-peer connections are supported
	UdpReflector    bool     `json:"udp_reflector"`    // True, if connection through UDP reflectors is supported
	MinLayer        int32    `json:"min_layer"`        // The minimum supported API layer; use 65
	MaxLayer        int32    `json:"max_layer"`        // The maximum supported API layer; use 65
	LibraryVersions []string `json:"library_versions"` // List of supported tgcalls versions
}

// MessageType return the string telegram-type of CallProtocol
func (callProtocol *CallProtocol) MessageType() string {
	return "callProtocol"
}

// NewCallProtocol creates a new CallProtocol
//
// @param udpP2p True, if UDP peer-to-peer connections are supported
// @param udpReflector True, if connection through UDP reflectors is supported
// @param minLayer The minimum supported API layer; use 65
// @param maxLayer The maximum supported API layer; use 65
// @param libraryVersions List of supported tgcalls versions
func NewCallProtocol(udpP2p bool, udpReflector bool, minLayer int32, maxLayer int32, libraryVersions []string) *CallProtocol {
	callProtocolTemp := CallProtocol{
		tdCommon:        tdCommon{Type: "callProtocol"},
		UdpP2p:          udpP2p,
		UdpReflector:    udpReflector,
		MinLayer:        minLayer,
		MaxLayer:        maxLayer,
		LibraryVersions: libraryVersions,
	}

	return &callProtocolTemp
}
