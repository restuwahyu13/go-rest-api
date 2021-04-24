package deleteStudent

type InputDeleteStudent struct {
	ID uint64 `json:"id" binding:"required"`
}
