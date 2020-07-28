package response

// Message codes
const (
	StatusExist    = 9000
	StatusNotExist = 9001
	StatusCreated  = 9002
	StatusUpdated  = 9003
	StatusDeleted  = 9004
	// member
	StatusNotExistMember = 9011
	// etc
	StatusError = 9999
)

var statusText = map[int]string{
	StatusExist:    "존재함",
	StatusNotExist: "존재하지 않음",
	StatusCreated:  "생성됨",
	StatusUpdated:  "수정됨",
	StatusDeleted:  "삭제됨",
	// member
	StatusNotExistMember: "해당 멤버가 존재하지 않음",
	// etc
	StatusError: "에러",
}

// StatusText returns a text
func StatusText(code int) string {
	return statusText[code]
}
