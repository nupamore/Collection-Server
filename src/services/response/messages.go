package response

// Message codes
const (
	StatusNotExist = 9001
	StatusCreated  = 9002
	StatusUpdated  = 9003
	StatusDeleted  = 9004
)

var statusText = map[int]string{
	StatusNotExist: "존재하지 않음",
	StatusCreated:  "생성됨",
	StatusUpdated:  "수정됨",
	StatusDeleted:  "삭제됨",
}

// StatusText returns a text
func StatusText(code int) string {
	return statusText[code]
}
