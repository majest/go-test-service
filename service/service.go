package service

type StringService struct{}

func (s *StringService) Count(value string) int {
	return len(value)
}
