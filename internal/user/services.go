package user

type Service struct {
	Repo *Repository
}

func (s *Service) RegisterIfNotExists(phone string) error {
	existing, err := s.Repo.FindByPhone(phone)
	if err != nil {
		return err
	}
	if existing == nil {
		return s.Repo.Create(phone)
	}
	return nil
}

func (s *Service) Get(phone string) (*User, error) {
	return s.Repo.FindByPhone(phone)
}

func (s *Service) List(limit, offset int, search string) ([]User, error) {
	return s.Repo.List(limit, offset, search)
}
