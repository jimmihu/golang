package structs

//struct risk_profile dgn foreign key ke user ID
type Risk_profile struct {
	UserID        int     `json:"userid"`
	MM_percent    float32 `json:"mm_percent"`
	Bond_percent  float32 `json:"bond_percent"`
	Stock_percent float32 `json:"stock_percent"`
}

//struct user
type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Password     string `json:"password"`
	Risk_profile Risk_profile
}

// Result is an array of post
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
