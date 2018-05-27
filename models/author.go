package models

// poloniex_buy_orders

type Author struct {
	Id       int    `json:"_id,omitempty"`
	Name     string `json:"name,omitempty"`
	NameUrdu string `json:"name_urdu,omitempty"`
	UserName string `json:"username,omitempty"`
	Info     string `json:"info,omitempty"`
	Dp       string `json:"dp,omitempty"`
	Dp_lg    string `json:"dp_lg,omitempty"`
	Type     string `json:"type,omitempty"`
	Intro    string `json:"intro,omitempty"`
	ImgShow  bool   `json:"imgShow,omitempty"`
}
