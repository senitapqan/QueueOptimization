package service

func (s *service) GetMostFreePlaceByCategoryId(id int) (int, error) {
	return s.repos.GetMostFreePlaceByCategoryId(id)
}
