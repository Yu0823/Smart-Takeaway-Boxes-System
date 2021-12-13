package control

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func OpenBoxById(id int64, isCus bool) error {
	//ip, err:= getIpFromTxt()
	//if err != nil{
	//	return err
	//}
	resp, err := http.Post("http://466w19n475.qicp.vip", "text/plain", strings.NewReader(getCommand(id, isCus) + "On"))

	if err != nil {
		fmt.Println(err)
	}

	if resp == nil{
		return fmt.Errorf("no res from hardware")
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func getCommand(id int64, IsCus bool) string{
	numOfDoor := id * 2
	if IsCus {
		return strconv.FormatInt(numOfDoor - 1, 10)
	} else{
		return strconv.FormatInt(numOfDoor, 10)
	}
}

func getIpFromTxt () (string, error) {
	fileName := "ip_config.txt"  // txt文件路径
	data, err := ioutil.ReadFile(fileName)   // 读取文件
	if err != nil {
		fmt.Println("文件读取失败！")
		return "", err
	}
	dataLine := strings.Split(string(data), "\n") // 将文件内容作为string按行切片
	var dataNameSlice [][]string                  // 用于存储每行内容的Slice
	for _, line := range dataLine {
		dataNameSlice = append(dataNameSlice, strings.Split(line, ":"))       // 每行内容按空格分割成切片
	}
	fmt.Println(dataNameSlice[0][1])
	return dataNameSlice[0][1], nil
}
