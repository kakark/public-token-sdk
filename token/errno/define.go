package errno

import (
	"errors"
)

var (
	ErrVerifyFailed     = errors.New("verify failed")
	ErrInvalidPayload   = errors.New("invalid payload")
	ErrInvalidTokenType = errors.New("invalid token type")
	ErrTokenRevoked     = errors.New("token revoked")
	ErrTokenMalformed   = errors.New("token malformed")
	// ErrTokenExpired 是Token过期错误
	ErrTokenExpired = errors.New("token expired")
	// ErrBadSignature 是签名错误
	ErrBadSignature = errors.New("bad signature")
	// ErrValidation 是Token字段验证错误
	ErrValidation = errors.New("token claim validation error")
	// ErrUnmarshalRaw 是Raw字段JSON反序列化错误
	ErrUnmarshalRaw = errors.New("failed to unmarshal raw")
	// ErrUnSupportFeatureCode 验签时 不支持该featureCode
	ErrUnSupportFeatureCode = errors.New("unSupport feature code")
	ErrTokenNotActive       = errors.New("token not yet valid")
	ErrInternal             = errors.New("internal error")
	ErrBadRequest           = errors.New("bad request")

	ErrRPC                  = errors.New("rpc error")
	ErrRPCWithoutBaseResp   = errors.New("rpc without base resp")
	ErrRPCStatusCodeNotZero = errors.New("rpc resp status code not zero")

	ErrBeOpaqueFailed = errors.New("be opaque failed")
)
