package service

func (s *service) PredictTime(age, id int) (string, error) {
	return s.repos.PredictTime(age, id)
}