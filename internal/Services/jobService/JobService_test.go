package jobservice_test

import (
	"context"
	"strings"
	"testing"

	jobservice "github.com/Kosk0l/TgBotConverter/internal/Services/jobService"
	"github.com/Kosk0l/TgBotConverter/internal/Services/jobService/mocks"
	"github.com/Kosk0l/TgBotConverter/internal/domains"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateJob_OK(t *testing.T) {
	// Arange:
	ctx := context.Background()

	repoJob := mocks.NewJobRepository(t)
	repoFile := mocks.NewFileRepository(t)
	service := jobservice.NewJobService(repoJob, repoFile)

	job := domains.Job{}
	jobObj := domains.Object{}

	// Act:
	repoFile.On("SetObject", ctx, mock.Anything, jobObj.FlieURL, jobObj.Size, jobObj.ContentType).Return(nil).Once()
	repoJob.On("SetToHash", ctx, mock.Anything).Return(nil).Once()
	repoJob.On("SetToList", ctx, mock.Anything).Return(nil).Once()

	jobId, err := service.CreateJob(ctx, job, jobObj)

	// Assert:
	assert.NoError(t, err)
	assert.NotEmpty(t, jobId)
}

func TestGetJob_OK(t *testing.T) {
	// Arange:	
	ctx := context.Background()

	repoJob := mocks.NewJobRepository(t)
	repoFile := mocks.NewFileRepository(t)
	service := jobservice.NewJobService(repoJob, repoFile)

	job := domains.Job{
		JobID: "qwerty",
		ChatID: 123,
		FileTypeTo: domains.Pdf,
	}
	

	// Act:
	repoJob.On("GetFromList", ctx).Return(mock.Anything, nil).Once()
	repoFile.On("GetObject", ctx, mock.Anything).Return(strings.NewReader("reader"), nil).Once()
	repoJob.On("GetFromHash", ctx, mock.Anything).Return(job, nil).Once()

	job, reader, err := service.GetJob(ctx)

	// Assert:
	assert.NoError(t, err)
	assert.NotEmpty(t, job)
	assert.NotEmpty(t, reader)
}