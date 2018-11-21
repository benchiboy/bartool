// ProcFlow.go
package control

import (
	"bartool/util"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const DATA_FILE = "./bars"
const IMAGE_PATH = "./public/static/"

/*
	二维码列表
*/
type BarList struct {
	Nodes     []Bar
	Count     int
	CurrPage  int
	TotalPage int
}

/*
	二维码节点
*/
type Bar struct {
	No         string
	Title      string
	FindText   string
	Context    string
	FileName   string
	CreateTime string
	Height     string
	Length     string
	PageIndex  string
	PageSize   string
}

/*
	构造函数
*/
func NewBars() *BarList {
	return &BarList{}
}

func gobWriteFile(path string, node interface{}) error {
	file, err := os.Create(path)
	enc := gob.NewEncoder(file)
	err = enc.Encode(node)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func gobReadFile(path string, node interface{}) error {
	file, err := os.OpenFile(path, os.O_RDWR, 0755)
	dec := gob.NewDecoder(file)
	err = dec.Decode(node)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

/*
	desc:增加节点
	return :返回增加后的长度
*/
func (r *BarList) AddNode(no string, title string, context string, length string, height string) (error, int) {
	nodeNo := fmt.Sprintf("%d", time.Now().UnixNano())
	n := Bar{No: nodeNo, Title: title, Context: context, Height: height, Length: length, CreateTime: time.Now().Format("2006-01-02 15:04:05")}
	var size int
	if isExist(DATA_FILE) {
		err := gobReadFile(DATA_FILE, &r.Nodes)
		if err != nil {
			log.Println(err)
			return err, 0
		}
		size = len(r.Nodes)
		r.Nodes = append(r.Nodes, n)
		err = gobWriteFile(DATA_FILE, &r.Nodes)
		if err != nil {
			log.Println(err)
			return err, size
		}
	} else {
		size = len(r.Nodes)
		r.Nodes = append(r.Nodes, n)
		err := gobWriteFile(DATA_FILE, &r.Nodes)
		if err != nil {
			log.Println(err)
			return err, size
		}
	}
	l, _ := strconv.Atoi(length)
	h, _ := strconv.Atoi(height)
	util.CreateQrcode(IMAGE_PATH+nodeNo+".jpg", "", context, l, h)
	return nil, size + 1
}

/*
	查找节点
*/
func (r *BarList) findNode(No string) int {
	var index int
	index = -1
	for i, v := range r.Nodes {
		if v.No == No {
			index = i
			break
		}
	}
	return index
}

/*
	根据标题模糊查找节点
*/
func (r *BarList) FindNodes(Title string) ([]Bar, int) {
	err := gobReadFile(DATA_FILE, &r.Nodes)
	if err != nil {
		return nil, 0
	}
	findBars := make([]Bar, 0)
	for _, v := range r.Nodes {

		if strings.Contains(v.Title, Title) {
			findBars = append(findBars, v)
		}
	}
	log.Println(findBars)
	return findBars, len(findBars)
}

/*
	修改节点
*/
func (r *BarList) ModiNode(No string, title string, context string, length string, height string) (error, int) {
	err := gobReadFile(DATA_FILE, &r.Nodes)
	if err != nil {
		log.Println(err.Error())
		return err, 0
	}
	index := r.findNode(No)
	if index < 0 {
		log.Println(No, "Node isn't Exist!")
		return fmt.Errorf("Node isn't Exist!"), 0
	}
	r.Nodes[index].No = No
	r.Nodes[index].Title = title
	r.Nodes[index].Context = context
	r.Nodes[index].Length = length
	r.Nodes[index].Height = height
	r.Nodes[index].CreateTime = time.Now().Format("2006-01-02 15:04:05")
	err = gobWriteFile(DATA_FILE, &r.Nodes)
	if err != nil {
		log.Println(err)
		return err, len(r.Nodes)
	}
	l, _ := strconv.Atoi(length)
	h, _ := strconv.Atoi(height)
	err = os.Remove(IMAGE_PATH + No + ".jpg")
	if err != nil {
		log.Println(err)
		return err, len(r.Nodes)
	}
	util.CreateQrcode(IMAGE_PATH+No+".jpg", title, context, l, h)
	return nil, len(r.Nodes)
}

/*
  删除节点
*/
func (r *BarList) DelNode(No string) (error, int) {
	err := gobReadFile(DATA_FILE, &r.Nodes)
	if err != nil {
		log.Println(err)
		return err, 0
	}
	index := r.findNode(No)
	if index < 0 {
		log.Println("Node isn't Exist!")
		return fmt.Errorf("Node isn't Exist!"), 0
	}
	r.Nodes = append(r.Nodes[:index], r.Nodes[index+1:]...)
	err = gobWriteFile(DATA_FILE, &r.Nodes)
	if err != nil {
		return err, len(r.Nodes)
	}
	os.Remove(IMAGE_PATH + No + ".jpg")
	return nil, len(r.Nodes)
}

/*
  得到某个节点信息
*/
func (r *BarList) GetNode(No string) *Bar {
	err := gobReadFile(DATA_FILE, &r.Nodes)
	if err != nil {
		return nil
	}
	index := r.findNode(No)
	if index < 0 {
		return nil
	}
	return &r.Nodes[index]
}

/*
  DESC:得到节点列表
	输入：pageIndex 页码
		 pageSize :每页的尺寸
	输出：

*/
func (r *BarList) GetNodes(title string, pageIndex int, pageSize int) ([]Bar, int) {
	err := gobReadFile(DATA_FILE, &r.Nodes)
	if err != nil {
		return nil, 0
	}
	findBars := make([]Bar, 0)
	for _, v := range r.Nodes {
		if strings.Contains(v.Title, title) {
			findBars = append(findBars, v)
		}
	}
	log.Println("findBars======>", findBars, len(findBars))
	var tmpBars []Bar
	for i := 0; i < pageSize; i++ {
		if pageIndex*pageSize+i < len(findBars) {
			tmpBars = append(tmpBars, findBars[pageIndex*pageSize+i])
		}
	}
	return tmpBars, len(findBars)
}

/*
  DESC:判断文件是否存在
	输入：path :
		 bool :true 存在，false 不存在
	输出：

*/
func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
