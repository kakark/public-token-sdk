package token

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVerifyLogger(t *testing.T) {
	ctx := context.Background()

	Convey("be opaque error", t, func() {
		//todo@fangchao
		VerifyJWT(ctx, "ehrkjehfawe.feagaegwrgaw.opopopopopop")
	})

}
