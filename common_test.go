package main

import "testing"
import "log"
import "errors"
import "time"

type TestJobDispatcher struct {
}

func DispatchTh(jobd interface{}, resultChan chan interface{}) {

}

type TestErrorDispatcher struct {
}

func (testErrorDispatcher *TestErrorDispatcher) DispatchError(failedJob *FailedJob) error {
	log.Print("success dispatch error")
	return nil
}

func (testJobDispatcher *TestJobDispatcher) Dispatch(jobd interface{}) (interface{}, error) {
	job := jobd.(Job)
	if job.JobId == "erroid" {
		time.Sleep(time.Second * 2)
		return FailedJob{JobId: "erroid", ErrorData: errors.New("generated error")}, nil
	} else if job.JobId == "workid" {
		time.Sleep(time.Second * 1)
		return DoneJob{JobId: "workid"}, nil
	} else {
		return FailedJob{JobId: job.JobId, ErrorData: errors.New("generated error")}, nil
	}
	return errors.New(""), nil
}

func TestJobBallancer(t *testing.T) {
	testJobDispatcher := TestJobDispatcher{}
	testErrorDispatcher := TestErrorDispatcher{}
	jobBallancer := JobBallancer{}
	jobBallancer.Init(&testJobDispatcher, &testErrorDispatcher)

	if err := jobBallancer.PushJob(Job{JobId: "workid"}); err != nil {
		t.Errorf("error: push err job failed " + err.Error())
		return
	}

	if err := jobBallancer.PushJob(Job{JobId: "erroid"}); err != nil {
		t.Errorf("error: push err job failed " + err.Error())
		return
	}
	if err := jobBallancer.TerminateTakeJob(); err != nil {
		t.Errorf("error: terminate job failed " + err.Error())
	}

}

func TestJobBallancerNowait(t *testing.T) {
	testJobDispatcher := TestJobDispatcher{}
	testErrorDispatcher := TestErrorDispatcher{}
	jobBallancer := JobBallancer{}
	jobBallancer.Init(&testJobDispatcher, &testErrorDispatcher)

	jobBallancer.PushJob(Job{JobId: "erroid1"})
	jobBallancer.PushJob(Job{JobId: "erroid2"})
	jobBallancer.PushJob(Job{JobId: "erroid3"})
	jobBallancer.PushJob(Job{JobId: "erroid4"})
	jobBallancer.PushJob(Job{JobId: "erroid5"})
	jobBallancer.PushJob(Job{JobId: "erroid6"})
	jobBallancer.PushJob(Job{JobId: "erroid7"})
	jobBallancer.PushJob(Job{JobId: "erroid8"})
	jobBallancer.PushJob(Job{JobId: "erroid9"})
	if err := jobBallancer.TerminateTakeJob(); err != nil {
		t.Errorf("error: terminate job failed " + err.Error())
	}
}

func TestDicomClient(t *testing.T) {
	/*DCOMClient:=DCOMClient{
			Address :
	Port           uint16
	ServerAE_Title string
	CallerAE_Title string
	}*/
}
