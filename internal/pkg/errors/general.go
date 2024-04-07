package errors

const (
	// General Errors: 0 -> -49
	// Processing indicate success but the object is being processed
	Processing ErrorType = 2
	// Success indicates no error
	Success ErrorType = 1
	// Unknown error indicates unknown state or step
	Unknown ErrorType = 0
	// BadRequest error
	BadRequestErr ErrorType = -1
	// NotFound error
	NotFound ErrorType = -2
	// AuthenFailed error
	AuthenticationFailed ErrorType = -3
	// Internal server error
	InternalServerError ErrorType = -4
	// IllegalStateError
	IllegalStateError ErrorType = -5
	// SendMessageError
	SendMessageError ErrorType = -6
	// Call Internal API Error
	CallInternalAPIError ErrorType = -7
	// Invalid Data
	InvalidData ErrorType = -8
	// SerializeError
	SerializingError ErrorType = -9
	// DeserializeError
	DeserializingError ErrorType = -10
	// CastingError
	CastingError ErrorType = -11
	// ParsingError
	ParsingError ErrorType = -12
	// ConflictError
	ConflictError ErrorType = -13
	// Call GRPC Internal API Error
	CallGRPCAPIError ErrorType = -14
	// EncryptError
	EncryptError ErrorType = -15
	// DecryptError
	DecryptError ErrorType = -16
	// MethodError
	MethodError ErrorType = -17

	//Failed
	Fail ErrorType = -49
)
