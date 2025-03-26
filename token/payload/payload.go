package payload

import (
	"fmt"
	"time"

	"github.com/kakark/public-token-sdk/token/errex"
	"github.com/kakark/public-token-sdk/token/errno"
	"github.com/kakark/public-token-sdk/token/tokentype"
)

// StandardClaims https://datatracker.ietf.org/doc/html/rfc9068#name-data-structure
type StandardClaims struct {
	Jti int64  `json:"jti,omitempty,string"` // JWT ID
	Iss string `json:"iss,omitempty"`        // token endpoint
	Sub string `json:"sub,omitempty"`        // user's open_id for this app
	Aud string `json:"aud,omitempty"`        // Token 的目标接收者，client_id, follow microsoft's practice
	Iat int64  `json:"iat,omitempty"`        // Token 的创建时间
	Nbf int64  `json:"nbf,omitempty"`        // Token 的生效时间
	Exp int64  `json:"exp,omitempty"`        // Token 的失效时间
}

type SensitiveClaims struct {
	AppID    int64  `json:"aid,omitempty,string"`
	UserID   int64  `json:"uid,omitempty,string"`
	TenantID int64  `json:"tid,omitempty,string"` // follow microsoft's practice
	SID      string `json:"sid,omitempty"`        // Session UUID, https://www.iana.org/assignments/jwt/jwt.xhtml
}

type SessionExtra struct {
	// session信息
	UserType    int64  `json:"user_type,omitempty,string"` // 用户类型｜身份类型
	Brand       int64  `json:"brand,omitempty,string"`     // 品牌，1=feishu，2=lark
	UserBrand   string `json:"user_brand,omitempty"`       // 用户品牌，feishu|lark
	IsAnonymous bool   `json:"is_anonymous"`               // 是否匿名session
	// 设备信息
	DeviceID      int64  `json:"device_id,omitempty,string"`       // 设备ID
	DeviceLoginID int64  `json:"device_login_id,omitempty,string"` // 一个 session 容器内的 id，user id x device id x 单品 id
	WebDeviceID   int64  `json:"web_device_id,omitempty,string"`   // web设备ID
	TerminalType  int64  `json:"terminal_typ,omitempty,string"`    // 终端类型
	OSType        *int64 `json:"os_typ,omitempty,string"`          // [访问控制] session 实体 device os 映射得到, 实体定义同网关通参。原值为session解析后结果，类型是指针，不一定存在。不希望default为0，0有自己的含义unknown。这里为nil的时候，网关插件选择不透传该字段。
	// Client信息
	AppChannel string `json:"app_channel,omitempty"` // 登录创建Session时客户端的Channel
	// 其余信息
	NeedAuth       bool   `json:"need_auth"`        // 应用准入控制的结果
	ValidForAppID  bool   `json:"valid_for_app_id"` // 应用准入控制的结果
	Locale         string `json:"locale,omitempty"` // locale
	DPoPThumbprint string `json:"dpop_thumbprint"`  // session confirmation jwt key thumbprint
}

// Env 环境信息
type Env struct {
	ClientIP string `json:"client_ip,omitempty"` // 创建Token时，客户端IP
}

type Payload struct {
	StandardClaims
	SensitiveClaims
	Version      string              `json:"ver"`
	TokenType    tokentype.TokenType `json:"typ"`
	ClientID     string              `json:"client_id"`       // required by rfc9068
	Scope        string              `json:"scope,omitempty"` // required by rfc9068
	Unit         string              `json:"unit,omitempty"`
	TenantUnit   string              `json:"tenant_unit,omitempty"`
	SessionExtra *SessionExtra       `json:"session_extra,omitempty"` // 授权时session的额外信息
	Env          *Env                `json:"env,omitempty"`
}

// Valid 检查且只检查了 exp 以及 nbf 是否合法，以及 client id 是否为空
func (p *Payload) Valid() error {
	//if p.ClientID == "" { // 暂不检查
	//	return errex.WrapErr(errno.ErrInvalidPayload, errors.New("client_id is empty"))
	//}
	now := time.Now().Unix()
	if now > p.Exp {
		return errex.WrapErr(errno.ErrTokenExpired, fmt.Errorf("token expired: %d > %d", now, p.Exp))
	}
	if now < p.Nbf {
		return errex.WrapErr(errno.ErrTokenNotActive, fmt.Errorf("token not active: %d < %d", now, p.Nbf))
	}
	return nil
}

func (p *Payload) SetTime(createTime time.Time, lifeTime time.Duration) {
	p.Iat = createTime.Unix()
	p.Exp = createTime.Add(lifeTime).Unix()
}

func (p *Payload) GetRemainLifetime() time.Duration {
	return time.Duration(p.Exp-time.Now().Unix()) * time.Second
}
