package models

// poloniex_buy_orders

type Author struct {
	Id       int    `json:"_id,omitempty"`
	Name     int    `json:"name,omitempty"`
	NameUrdu int    `json:"name_urdu,omitempty"`
	UserName int    `json:"username,omitempty"`
	Info     int    `json:"info,omitempty"`
	Dp       string `json:"dp,omitempty"`
	Dp_lg    string `json:"dp_lg,omitempty"`
	Type     string `json:"type,omitempty"`
	Intro    string `json:"intro,omitempty"`
	ImgShow  bool   `json:"imgShow,omitempty"`
}
