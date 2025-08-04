package repository

type Repository interface {
	Save(code string, url string) error
	Find(code string) (string, error)
}
