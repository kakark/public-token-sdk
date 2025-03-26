package token

import (
	"context"
	"fmt"

	"github.com/kakark/public-token-sdk/token/errex"
	"github.com/kakark/public-token-sdk/token/errno"
	tokenpayload "github.com/kakark/public-token-sdk/token/payload"
	"github.com/kakark/public-token-sdk/token/tokentype"
)

// ValidateJWT 首先校验 JWT 是否合法，然后解析出 payload，最后返回 payload 本身。
func ValidateJWT(ctx context.Context, jwtString string, targetTokenType tokentype.TokenType) (*tokenpayload.Payload, error) {
	// 验证 JWT 有效性
	err := VerifyJWT(ctx, jwtString)
	if err != nil {
		return nil, errex.WrapErr(errno.ErrVerifyFailed, fmt.Errorf("verifyClient.VerifyStandard failed: %w", err))
	}
	payload, err := ParseJWT(ctx, jwtString)
	if err != nil {
		return nil, errex.WrapErr(errno.ErrInvalidPayload, fmt.Errorf("ParseJWT failed: %w", err))
	}
	// 接下来要校验 token 类型符合预期
	if payload.TokenType != targetTokenType {
		return nil, errex.WrapErr(errno.ErrInvalidTokenType, fmt.Errorf("target token type is %s, but got %s", targetTokenType, payload.TokenType))
	}

	return payload, nil
}
