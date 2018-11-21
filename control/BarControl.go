// ProcFlow.go
package control

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const OPER_ADD = 1
const OPER_DEL = 2
const OPER_MOD = 3
const OPER_LOD = 4

func Index(w http.ResponseWriter, req *http.Request) {
	buf, _ := ioutil.ReadFile("./htmls/main.html")
	w.Write(buf)
}

/*

 */
func freshNodes(length int, operType int, r *BarList, findText string, pageIndex string, pageSize string, w http.ResponseWriter, req *http.Request) {

	pi, _ := strconv.Atoi(pageIndex)
	ps, _ := strconv.Atoi(pageSize)

	var totalPage int
	if length%ps == 0 {
		totalPage = length / ps
	} else {
		totalPage = length/ps + 1
	}
	if operType == OPER_ADD || operType == OPER_DEL {
		pi = totalPage - 1
	} else {
		pi = pi
	}
	log.Println("CurrPage,CurrSize", length, pi, ps)

	tabSlice, total := r.GetNodes(findText, pi, ps)
	barList := new(BarList)
	barList.Nodes = tabSlice
	barList.Count = total
	barList.TotalPage = totalPage
	barList.CurrPage = pi
	readBuf, _ := json.Marshal(barList)
	w.Write(readBuf)
}

/*

 */

func Baradd(w http.ResponseWriter, req *http.Request) {
	log.Println("===============>Baradd")
	var request Bar
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(w, fmt.Errorf("增加节点出错").Error(), http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()
	r := NewBars()
	err, length := r.AddNode(request.No, request.Title, request.Context, request.Length, request.Height)
	if err == nil {
		log.Println("length==", length)
		freshNodes(length, OPER_ADD, r, request.FindText, request.PageIndex, request.PageSize, w, req)
	} else {
		http.Error(w, fmt.Errorf("增加节点出错").Error(), http.StatusInternalServerError)
	}
}

/*

 */

func Barget(w http.ResponseWriter, req *http.Request) {
	log.Println("===============>Barget")
	var request Bar
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(w, fmt.Errorf("没有发现节点").Error(), http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()
	r := NewBars()
	tabSlice := r.GetNode(request.No)
	if tabSlice == nil {
		http.Error(w, fmt.Errorf("没有发现节点").Error(), http.StatusInternalServerError)
		return
	}
	buf, _ := json.Marshal(tabSlice)
	w.Write(buf)
}

/*

 */

func Barset(w http.ResponseWriter, req *http.Request) {
	log.Println("===============>Barset")
	var request Bar
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		log.Println("Barset......", err.Error())
		return
	}
	defer req.Body.Close()
	r := NewBars()
	err, length := r.ModiNode(request.No, request.Title, request.Context, request.Length, request.Height)
	log.Println("length===", length)
	freshNodes(length, OPER_MOD, r, request.FindText, request.PageIndex, request.PageSize, w, req)
}

/*

 */

func Bardel(w http.ResponseWriter, req *http.Request) {
	log.Println("===============>Bardel")
	var request Bar
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		log.Println("Bardel......", err.Error())
		return
	}
	defer req.Body.Close()
	r := NewBars()
	err, length := r.DelNode(request.No)
	log.Println("length====>", length)
	freshNodes(length, OPER_DEL, r, request.FindText, request.PageIndex, request.PageSize, w, req)
}

/*

 */

func BarLoad(w http.ResponseWriter, req *http.Request) {
	log.Println("===============>BarLoad")
	var request Bar
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		log.Println("Bardel......", err.Error())
		return
	}
	defer req.Body.Close()

	r := NewBars()
	pi, _ := strconv.Atoi(request.PageIndex)
	ps, _ := strconv.Atoi(request.PageSize)
	log.Println("FindText====>", request)
	tabSlice, total := r.GetNodes(request.FindText, pi, ps)
	var totalPage int
	if total%ps == 0 {
		totalPage = total / ps
	} else {
		totalPage = total/ps + 1
	}
	barList := new(BarList)
	barList.Nodes = tabSlice
	barList.Count = total
	barList.TotalPage = totalPage
	barList.CurrPage = pi
	readBuf, _ := json.Marshal(barList)
	w.Write(readBuf)
}
