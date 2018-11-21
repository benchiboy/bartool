// Interface1 project Interface1.go
package cron

import (
	"fmt"
	"sort"
	"time"
)

const (
	ATJOB = 1
	EVJOB = 2
)

type Schdule interface {
	Next(t time.Time) time.Time
}

type JobInfo struct {
	Schdule Schdule
	Job     Job
	JobNo   string
	JobName string
	JobType int
	Next    time.Time
	Prev    time.Time
}

// byTime is a handy wrapper to chronologically sort entries.
type byTime []*JobInfo

func (b byTime) Len() int      { return len(b) }
func (b byTime) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

// Less reports `earliest` time i should sort before j.
// zero time is not `earliest` time.
func (b byTime) Less(i, j int) bool {

	if b[i].Next.IsZero() {
		return false
	}
	if b[j].Next.IsZero() {
		return true
	}

	return b[i].Next.Before(b[j].Next)
}

// Job is the interface that wraps the basic Run method.
//
// Run executes the underlying func.
type Job interface {
	Run()
}

// todo: possibly func with params? maybe not needed.
type JobFunc func()

// Run calls j()
func (j JobFunc) Run() {

	fmt.Println("CALL FUNC ....")
	j()
}

type JobCtl struct {
	jobQueue []*JobInfo
	runing   bool
	stop     chan struct{}
	add      chan *JobInfo
}

func NewJobCtl() *JobCtl {
	return &JobCtl{stop: make(chan struct{}), add: make(chan *JobInfo)}
}

func (j *JobCtl) AddEveryJob(sch Schdule, job Job) {
	jobNode := JobInfo{Schdule: sch, Job: job, Next: time.Now(), JobType: EVJOB}
	if !j.runing {
		j.jobQueue = append(j.jobQueue, &jobNode)
		return
	}
	j.add <- &jobNode
}

func (j *JobCtl) AddAtJob(sch Schdule, job Job) {
	jobNode := JobInfo{Schdule: sch, Job: job, Next: time.Now(), JobType: ATJOB}
	if !j.runing {
		j.jobQueue = append(j.jobQueue, &jobNode)
		return
	}
	j.add <- &jobNode
}

func (j *JobCtl) AddEveryJobFunc(sch Schdule, f func()) {
	j.AddEveryJob(sch, JobFunc(f))
}

func (j *JobCtl) AddAtJobFunc(sch Schdule, f func()) {
	j.AddAtJob(sch, JobFunc(f))

}

func (j *JobCtl) Stop() {
	if !j.runing {
		return
	}
	j.runing = false
	j.stop <- struct{}{}

}

func (j *JobCtl) GetAllJobs() []*JobInfo {
	return j.jobQueue
}

func (j *JobCtl) Start() {
	j.runing = true

	var effective time.Time
	now := time.Now().Local()
	for _, e := range j.jobQueue {
		e.Next = e.Schdule.Next(now)
	}

	for {
		sort.Sort(byTime(j.jobQueue))
		if len(j.jobQueue) > 0 {
			effective = j.jobQueue[0].Next
		} else {
			effective = now.AddDate(15, 0, 0) // to prevent phantom jobs.
		}
		select {

		case now = <-time.After(effective.Sub(now)):
			// entries with same time gets run.
			for _, entry := range j.jobQueue {
				if entry.Next != effective {
					break
				}
				entry.Prev = now
				entry.Next = entry.Schdule.Next(now)
				go entry.Job.Run()
			}

		case e := <-j.add:
			e.Next = e.Schdule.Next(time.Now())
			j.jobQueue = append(j.jobQueue, e)
		case <-j.stop:
			return // terminate go-routine.
		}
		fmt.Println("start select ...", j.runing)
	}

}

type everySchedule struct {
	perild time.Duration
}

func (e everySchedule) Next(t time.Time) time.Time {
	return t.Add(e.perild)
}

func Every(d time.Duration) Schdule {
	return &everySchedule{perild: d}
}
