package structs

//struct risk_profile dgn foreign key ke user ID
type Risk_profile struct {
	UserID        int     `json:"userid" gorm:"primaryKey"`
	MM_percent    float32 `json:"mm_percent"`
	Bond_percent  float32 `json:"bond_percent"`
	Stock_percent float32 `json:"stock_percent"`
}

//struct user
type User struct {
	ID           int          `json:"id" gorm:"primaryKey"`
	Name         string       `json:"name"`
	Age          int          `json:"age"`
	Password     string       `json:"password"`
	Risk_profile Risk_profile `json:"risk_profile" `
}

// Result is an array of post
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
