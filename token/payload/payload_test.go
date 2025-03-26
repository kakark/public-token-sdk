package payload

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/kakark/public-token-sdk/token/errno"
	"github.com/kakark/public-token-sdk/token/tokentype"
)

func getDummySensitiveClaims() *SensitiveClaims {
	return &SensitiveClaims{
		AppID:    2,
		UserID:   3,
		SID:      "4",
		TenantID: 5,
	}
}
func getDummyPayload() *Payload {
	var payload Payload
	payload.SensitiveClaims = *getDummySensitiveClaims()
	payload.Iat = time.Now().Unix()
	payload.TokenType = tokentype.AccessToken
	payload.ClientID = "cli_test"
	return &payload
}

func TestPayloadMarshal(t *testing.T) {
	Convey("marshal opaque token", t, func() {
		payload := getDummyPayload()
		payloadBytes, err := json.Marshal(payload)
		So(err, ShouldBeNil)
		payloadStr := string(payloadBytes)
		So(payloadStr, ShouldNotContainSubstring, "enc")
		So(payloadStr, ShouldNotContainSubstring, "enc_ver")
	})

}

func TestValid(t *testing.T) {
	payload := getDummyPayload()
	now := time.Now()
	offset := time.Hour
	Convey("valid", t, func() {
		payload.SetTime(now, offset)
		err := payload.Valid()
		So(err, ShouldBeNil)
	})
	Convey("invalid, expired", t, func() {
		payload.SetTime(now.Add(-2*offset), offset)
		err := payload.Valid()
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "expired")
	})
	Convey("invalid, empty client id", t, func() {
		payload.ClientID = ""
		err := payload.Valid()
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "empty")
		payload.ClientID = "cli_test"
	})
	Convey("invalid, not active", t, func() {
		payload.SetTime(now, offset)
		payload.Nbf = now.Add(offset).Unix()
		err := payload.Valid()
		So(err, ShouldWrap, errno.ErrTokenNotActive)
	})
}
