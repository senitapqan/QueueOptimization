package service

func (s *service) PredictTime(age, id int) (int, error) {
	return s.repos.PredictTime(age, id)
}