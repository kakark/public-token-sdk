package masker

import (
	"fmt"
	"strings"
)

const (
	showLength         = 5
	jwtShowLength      = 10
	uriShowLengthLeft  = 8 + 10
	uriShowLengthRight = 10
	urlMaskLength      = 10
	minJWTLength       = 100
)

func isJWT(token string) bool {
	return len(token) > minJWTLength
}

func MaskToken(token string) string {
	if token == "" {
		return "<empty>"
	}
	if isJWT(token) {
		return MaskJWT(token)
	}
	return MaskString(token)
}

func MaskStringPtr(str *string) string {
	if str == nil {
		return "<nil>"
	}
	return MaskString(*str)
}

func MaskString(str string) string {
	if str == "" {
		return "<empty>"
	}
	var prefix string
	if strings.HasPrefix(str, "u-") {
		prefix = "u-"
		str = strings.TrimPrefix(str, "u-")
	} else if strings.HasPrefix(str, "ur-") {
		prefix = "ur-"
		str = strings.TrimPrefix(str, "ur-")
	}
	l := len(str)
	if l <= showLength*2 {
		return prefix + str
	}
	maskedStr := str[:showLength] + strings.Repeat("*", l-showLength*2) + str[l-showLength:]
	return prefix + maskedStr
}

func MaskJWT(token string) string {
	if token == "" {
		return "<empty>"
	}
	l := len(token)
	if l <= jwtShowLength*2 {
		return MaskString(token)
	}
	return token[:jwtShowLength] + "*****<omitted>*****" + token[l-jwtShowLength:]
}

func MaskURI(uri string) string {
	if uri == "" {
		return "<empty>"
	}
	l := len(uri)
	if l <= urlMaskLength {
		return uri
	}
	start := (l - urlMaskLength) / 2
	end := start + urlMaskLength
	return uri[:start] + strings.Repeat("*", urlMaskLength) + uri[end:]
}

func MaskURIs(uris []string) string {
	var maskedURIs []string
	for _, uri := range uris {
		maskedURIs = append(maskedURIs, MaskURI(uri))
	}
	return fmt.Sprintf("[%s]", strings.Join(maskedURIs, ", "))
}
