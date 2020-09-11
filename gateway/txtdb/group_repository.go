package txtdb

import (
	"encoding/json"
	"errors"
    "fmt"
	"log"
	"time"

    "github.com/andersonlira/group-api/domain"
	"github.com/andersonlira/goutils/io"
	"github.com/andersonlira/goutils/str"
	"sort"
)

//GetGroupList return all items 
func GetGroupList() []domain.Group {
	list := []domain.Group{}
    fileName := fmt.Sprintf("bd/%ss.json", "Group");
	listTxt, _ := io.ReadFile(fileName)
	json.Unmarshal([]byte(listTxt), &list)
	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})
	return list
}

//GetGroupByID return all items 
func GetGroupByID(ID string) (domain.Group, error) {
	list := GetGroupList()
	for idx, _ := range list {
		if(list[idx].ID == ID){
			return list[idx],nil
		}
	}
	return domain.Group{}, errors.New("NOT_FOUND")
}



//SaveGroup saves a Group object
func SaveGroup(it domain.Group) domain.Group {
	list := GetGroupList()
	it.ID = str.NewUUID()
	it.CreatedAt = time.Now()
	list = append(list, it)
	writeGroup(list)
	return it
}

//UpdateGroup( updates a Group object
func UpdateGroup(ID string, it domain.Group) domain.Group{
	list := GetGroupList()
	for idx, _ := range list {
		if(list[idx].ID == ID){
			list[idx] = it
			list[idx].ID = ID
			list[idx].UpdatedAt = time.Now()
			writeGroup(list)
			return list[idx]
		}
	}
	return it
}

//DeleteGroup delete object by giving ID
func DeleteGroup(ID string) bool {
	list := GetGroupList()
	for idx, _ := range list {
		if(list[idx].ID == ID){
			list = append(list[:idx], list[idx+1:]...)
			writeGroup(list)
			return true
		}
	}
	return false
}

func writeGroup(list []domain.Group) {
	b, err := json.Marshal(list)
	if err != nil {
		log.Println("Error while writiong file items")
		return
	}
	io.WriteFile(fmt.Sprintf("bd/%ss.json", "Group"), string(b))
}

