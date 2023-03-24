package consts

// Error Response
const (
	AUTH_ErrorBind            string = "error bind data"
	AUTH_ErrorHash            string = "error hash password"
	AUTH_ErrorComparePassword string = "password not matched"
	AUTH_ErrorCreateToken     string = "error create token"
	AUTH_ErrorRole            string = "role must be student or mentor"
)

// Success Response
const (
	AUTH_SuccessCreate string = "success create account"
	AUTH_SuccessLogin  string = "login success"
)
