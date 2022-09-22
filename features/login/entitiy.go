package login

type Core struct {
	ID       int
	Email    string
	Password string
	Role     string
}

type UsecaseInterface interface {
<<<<<<< HEAD
	LoginAuthorized(email, password string, role string) string
=======
	LoginAuthorized(email, password string) (string, string)
>>>>>>> 476986a755198c287082c23df981b94e38b8b7b9
}

type DataInterface interface {
	LoginUser(email string) (Core, error)
}
