package consts

// Error Response
const (
	AUTH_ErrorBind            string = "error bind data"
	AUTH_ErrorHash            string = "error hash password"
	AUTH_ErrorComparePassword string = "password not matched"
	AUTH_ErrorCreateToken     string = "error create token"
)

// Success Response
const (
	AUTH_SuccessCreate string = "success create account"
)
