package control

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func OpenBoxById(id int64, isCus bool) error {
	resp, err := http.Post("http://10.128.199.229:8000", "text/plain", strings.NewReader(getCommand(id, isCus) + "On"))

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
