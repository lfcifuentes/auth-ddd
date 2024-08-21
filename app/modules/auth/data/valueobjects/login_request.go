package valueobjects

type LoginParams struct {
	Email    string `form:"email" validate:"required,email" documentation:"Email"`
	Password string `form:"password" validate:"required,password" documentation:"Contraseña"`
}
