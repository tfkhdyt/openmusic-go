package album

type Entity struct {
	Name string `json:"name" binding:"required"`
	Year uint16 `json:"year" binding:"required"`
}
