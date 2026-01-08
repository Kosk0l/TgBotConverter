package jobservice

type JobRepository interface {
	AddToList() ()
	GetFromList() ()
	AddToHash() ()
	GetFromHash() ()
}

type JobService struct {

}

func (js *JobService) CreateJob() () {

}

func (js *JobService) GetJob() () {

}

func (js *JobService) DeleteJob() () {
	
}