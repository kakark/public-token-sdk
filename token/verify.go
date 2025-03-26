package token

import (
	"context"

	"github.com/kakark/public-token-sdk/masker"
	"github.com/kakark/public-token-sdk/token/logs"
)

// VerifyJWT 只验签、校验 issuer & 校验有效期，不检查是否被撤销，也不关心 token 类型
func VerifyJWT(ctx context.Context, jwtString string) error {
	logs.Logger(ctx).Infof("VerifyJwt req:%v", masker.MaskJWT(jwtString))
	return nil // todo@fangchao 公钥验签名 verifyClient.VerifyStandard(ctx, jwtString)
}
