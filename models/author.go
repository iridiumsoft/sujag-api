package models

// poloniex_buy_orders

type Author struct {
	Id        int `json:"_id"`
	FirstName int `json:"first_name"`
	LastName  int `json:"last_name"`
	Dp  string `json:"dp"`
	Dp_lg  string `json:"dp_lg"`
	Type  string `json:"type"`
	Intro  string `json:"intro"`
	ImgShow  bool `json:"imgShow"`
}
