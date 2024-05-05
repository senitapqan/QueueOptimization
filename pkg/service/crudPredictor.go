package service

func (s *service) PredictTime(id, age int) (int, error) {
	return s.repos.PredictTime(id, age)
}