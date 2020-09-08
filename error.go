package oscrud

import "errors"

// Error Definition
var (
	ErrNotFound              = errors.New("oscrud: endpoint or service not found")
	ErrResponseFailed        = errors.New("oscrud: response doesn't return properly in transport")
	ErrSourceNotAddressable  = errors.New("oscrud: binder source must be addressable")
	ErrRequestTimeout        = errors.New("oscrud: request timeout")
	ErrMultipartNotSupported = errors.New("oscrud: multipart not support")
	ErrFormNotSupported      = errors.New("oscrud: form not supported")
	ErrTransportNotExists    = errors.New("oscrud: transport not exists")
	ErrTransportNotSupport   = errors.New("oscrud: transport not support")
)
