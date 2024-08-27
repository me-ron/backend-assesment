package usecase

import "loan_tracker/domain"

type LogUsecase struct {
	repo domain.LogRepository
}

func NewLogUsecase(repo domain.LogRepository) *LogUsecase {
	return &LogUsecase{repo: repo}
}

func (lu *LogUsecase) GetAllLogs() ([]domain.Log, error) {
	return lu.repo.GetAllLogs()
}