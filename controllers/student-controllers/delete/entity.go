package deleteStudent

type InputDeleteStudent struct {
	ID string `validate:"required,uuid"`
}
