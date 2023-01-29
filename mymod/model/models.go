package model

var Orders []Order

// Models
type Order struct {
	OrderId    string  `json:"orderid"`
	OrderName  string  `json:"ordername"`
	OrderPrice int     `json:"price"`
	Client     *Client `json:"client"`
}

type Client struct {
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Password string `json:"-"`
}

func (c *Order) IsEmpty() bool {

	return c.OrderName == ""
}
