package token

import (
	"context"
	"fmt"
	"strings"

	"github.com/kakark/public-token-sdk/token/errex"
	"github.com/kakark/public-token-sdk/token/errno"
	"github.com/kakark/public-token-sdk/token/header"
	"github.com/kakark/public-token-sdk/token/helper"
	"github.com/kakark/public-token-sdk/token/payload"
)

// ParseJWT 仅解析 payload，不做任何其他校验，除此之外对于 Opaque Token 会解密
func ParseJWT(ctx context.Context, jwt string) (*payload.Payload, error) {
	return parseJWTHelper(ctx, jwt)
}

// parseJWTHelper 仅解析 payload，不做任何其他校验
func parseJWTHelper(ctx context.Context, jwt string) (*payload.Payload, error) {
	parts := strings.Split(jwt, ".")
	if len(parts) != 3 {
		return nil, errex.WrapErr(errno.ErrTokenMalformed, fmt.Errorf("invalid jwt, len(parts) != 3: %d", len(parts)))
	}
	var p payload.Payload
	err := helper.DecodeAndUnmarshal(parts[1], &p)
	if err != nil {
		return nil, errex.WrapErr(errno.ErrTokenMalformed, fmt.Errorf("helper.DecodeAndUnmarshal failed: %w, jwt: %s.%s.******", err, parts[0], parts[1]))
	}
	return &p, nil
}

func ParseJWTForHeader(ctx context.Context, jwtString string) (*header.Header, error) {
	parts := strings.Split(jwtString, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid jwt, len(parts)!= 3: %d", len(parts))
	}
	var h header.Header
	err := helper.DecodeAndUnmarshal(parts[0], &h)
	if err != nil {
		return nil, fmt.Errorf("helper.DecodeAndUnmarshal failed: %w", err)
	}
	return &h, nil
}
