package consts

// Success Response
const (
	MENTOR_SuccessGetProfile     string = "success show mentor profile"
	MENTOR_SuccessUpdateProfile  string = "success update mentor profile"
	MENTOR_SuccessUpdatePassword string = "success update mentor password"
	MENTOR_SuccessAddCredential  string = "success add mentor credential"
)

// Error Response
const (
	AWS_ErrorUpload           string = "error upload to s3"
	MENTOR_ErrorEmptyPassword string = "old password, new password and confirmation password field cannot be empty"
)
