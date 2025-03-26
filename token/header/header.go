package header

type Header struct {
	Alg         string `json:"alg"`
	Typ         string `json:"typ"`
	Kid         string `json:"kid"`
	FeatureCode string `json:"feature_code,omitempty"` // KMS 的概念
}
