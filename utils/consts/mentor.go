package consts

// Success Response
const (
	MENTOR_SuccessGetProfile     string = "success show mentor profile"
	MENTOR_SuccessUpdateProfile  string = "success update mentor profile"
	MENTOR_SuccessUpdatePassword string = "success update mentor password"
	MENTOR_SuccessAddCredential  string = "success add mentor credential"
	MENTOR_SuccessDelete         string = "succes deactivate mentor"
	MENTOR_SuccessGetAll         string = "success show all mentor"
)

// Error Response
const (
	AWS_ErrorUpload           string = "error upload to s3"
	MENTOR_NameOnlyLetters    string = "name must be filled only letters"
	MENTOR_ErrorQualification string = "qualification can only filled by International or National"
)
