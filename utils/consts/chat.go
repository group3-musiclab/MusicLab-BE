package consts

// Success Response
const (
	CHAT_SuccessInsert        string = "success send message"
	CHAT_SuccessAllChat       string = "success show a chat by id student and id mentor"
	CHAT_SuccessChatByStudent string = "success show chat group by student"
)

// Error Response
const (
	CHAT_ErrorMentorID  string = "mentor id not found"
	CHAT_ErrorStudentID string = "student id not found"
)
