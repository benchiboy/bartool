package util

import (
	"fmt"
	"testing"
	"time"
)

//type Farmer struct{}

//func (f Farmer) Run() {
//	fmt.Println("run..................... ")
//}

func TestSet_Plan(t *testing.T) {
	t.Log("hello-------------------")
	fmt.Println("=======")
	//	tim := Farmer{}
	sch := New_Schdule()

	sch.Add_TaskFunc(1000, "22ssss", false, "YYYY-MM-DD-W-HH-MI-01", func(taskId int) {
		fmt.Println(2000, "ssssssssssssssssssssssssss")
		time.Sleep(time.Second * 65)
		fmt.Println("end... ")
	})

	sch.Add_TaskFunc(1001, "22ssss", false, "YYYY-MM-DD-W-HH-MI-01", func(taskId int) {
		fmt.Println(2000, "ssssssssssssssssssssssssss")
		time.Sleep(time.Second * 65)
		fmt.Println("end... ")
	})
	sch.Del_AllTask()
	sch.Add_TaskFunc(1004, "22ssss", false, "YYYY-MM-DD-W-HH-MI-01", func(taskId int) {
		fmt.Println(2000, "ssssssssssssssssssssssssss")
		time.Sleep(time.Second * 65)
		fmt.Println("end... ")
	})
	sch.Run()
	time.Sleep(15000 * time.Second)

}

//func TestTimeTest(t *testing.T) {
//	//t.Log("hello")
//	//time.Now().tr

//	var d time.Duration

//	d = time.Microsecond

//	if d < time.Second {
//		t.Log("=========ssssssssssssssssssssss")
//	}
//	t.Log(time.Duration(d.Nanoseconds()))
//	t.Log("======>", d-time.Duration(d.Nanoseconds())%time.Second)
//	d = time.Nanosecond

//	//t.Log("======>", d.String())

//	//	t.Log("=====================》》》》", d.Seconds())
//	//	t.Log("=====================》》》》", d.Minutes())
//	//	t.Log("=====================》》》》", d.Nanoseconds())
//	//	t.Log("=====================")
//	//	t.Log("=====================-----", d-time.Duration(d.Nanoseconds())%time.Second)
//	//	t.Log("=====================", time.Now().Nanosecond())
//	//	t.Log("=====================", time.Now().Second())
//	//b := time.Date(2017, 0, 0, 0, 0, 0, 0, time.UTC)
//	//t.Log("=====================", b)

//}
