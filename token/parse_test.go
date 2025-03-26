package token

import (
	"context"
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/kakark/public-token-sdk/token/errno"
)

func TestParseJWT(t *testing.T) {
	ctx := context.Background()
	jwt := "eyJhbGciOiJFUzI1NiIsImZlYXR1cmVfY29kZSI6IkZlYXR1cmVPQXV0aEpXVFNpZ25fQk9FQ04iLCJraWQiOiI3MzYyNzAxNTYwOTAxNTMzNzE1IiwidHlwIjoiSldUIn0.eyJqdGkiOiI3MzYzMjg3OTgyODU5NjE2Mjc1IiwiaWF0IjoxNzE0Mzk5MDgwLCJleHAiOjE3MTQzOTkxNDAsImFpZCI6IjcyOTU5ODI4MzMyMDQ0OTQzNTYiLCJ1aWQiOiI3MjYzMDg2NDk0MDg3MTE4ODY3IiwidGlkIjoiMSIsInNpZCI6IkFBQUFBQUFBQUFObUsxd2tNd0FBRkE9PSIsInZlciI6InYxIiwidG9rZW5fdHlwZSI6ImFjY2Vzc190b2tlbiIsImNsaWVudF9pZCI6ImNsaV9hNWNhMzVhNjg1Yjg1MDFiIiwic2NvcGUiOiJhdXRoOnVzZXIuaWQ6cmVhZCBvZmZsaW5lX2FjY2VzcyB1c2VyX3Byb2ZpbGUiLCJhdXRoX2lkIjoiNzM2MzI4Nzk3ODA3ODEwOTcxNSIsImF1dGhfdGltZSI6MTcxNDM5OTA3OSwidW5pdCI6ImJvZWNuIn0.DqaNf8kwD9SVkP0yfp-IrwxfsKF3b3CQxtZp39kbCGxTLYfl8C5zIQuvMHNzUKD3PZHJMlJtyVpdZtYbK37eRw"
	targetUnit := "boecn"

	Convey("parse non-opaque token", t, func() {
		payload, err := ParseJWT(ctx, jwt)
		So(payload, ShouldNotBeNil)
		So(payload.Unit, ShouldEqual, targetUnit)
		So(err, ShouldBeNil)
	})

	Convey("parse jwt, but not access token", t, func() {
		nonAccessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwiYXVkIjpbXSwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.I_C0eytk41oWMZ65OkFKo9eOjN6o3FLvGpUO6eh_mmM"
		payload, err := ParseJWT(ctx, nonAccessToken)
		So(errors.Is(err, errno.ErrTokenMalformed), ShouldBeTrue)
		So(payload, ShouldBeNil)
	})
}
