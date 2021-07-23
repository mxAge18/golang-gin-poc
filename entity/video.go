package entity

type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName string `json:"lastname" binding:"required"`
	Age int8 `json:"age" binding:"gte=1,lte=130"`
	Email string `json:"email" binding:"required,email"`
}

type Video struct {
	Title string `json:"title" binding:"min=2" validate:"is-cool"` 
	Description string `json:"description" binding:"min=6"`
	URL string `json:"url" binding:"required,url"`
	Author Person `json:"author" binging:"required"`
}

