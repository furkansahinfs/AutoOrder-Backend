package user_information

type Reader interface {
}

type Writer interface {
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}
