package jobservice

import (

)

type JobRepository interface {
	AddToList() ()
	GetFromList() ()
	AddToHash() ()
	GetFromHash() ()
}

type JobService struct {
	repo JobRepository
}

func NewJobService(repo JobRepository) (*JobService) {
	return &JobService{
		repo: repo,
	}
}

func (js *JobService) CreateJob() () {

}

func (js *JobService) GetJob() () {

}

func (js *JobService) DeleteJob() () {

}