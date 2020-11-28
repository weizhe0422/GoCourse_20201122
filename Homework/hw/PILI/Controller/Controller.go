package PILI

import (
	Model "../Ｍodel"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const SourceData = "../Model/布袋戲資料.json"

type PILI struct {
	AllDramas []Model.Role
}

func New() *PILI {
	p := new(PILI)
	p.GetAllRecords()
	return p
}

//取得全部資料 [GET] http://localhost:8080/role
//新增單筆資料 [POST] http://localhost:8080/role
//取得單筆資料 [GET] http://localhost:8080/role/:id
//更新單筆資料 [PUT] http://localhost:8080/role/:id
//刪除單筆資料 [DELETE] http://localhost:8080/role/:id

// GetAllRecords will get all records
func (p *PILI) GetAllRecords() error {
	var records []Model.Role

	file, err := ioutil.ReadFile(SourceData)
	if err != nil {
		return fmt.Errorf("failed to get records: %v", err)
	}

	if err = json.Unmarshal(file, &records); err != nil {
		return fmt.Errorf("failed to decode source JSON string: %v", err)
	}
	p.AllDramas = records
	return nil
}

//GetSpecificRecord will get the specific id record
func (p *PILI) GetSpecificRecord(id int) (*Model.Role, error){
	if len(p.AllDramas) == 0{
		p.GetAllRecords()
	}

	for _, record := range p.AllDramas {
		if uint(id) == record.ID {
			return &record, nil
		}
	}
	return nil, fmt.Errorf("can not find id=%d record!!", id)
}

//InsertRecord will decode input string to insert, this is also support multi-request
func (p *PILI) InsertRecord(insertRecord []Model.Role) error{
	if len(p.AllDramas) == 0 {
		p.GetAllRecords()
	}

	for _, inputData := range insertRecord {
		p.AllDramas = append(p.AllDramas, inputData)
	}

	return nil
}

//DeleteOneRecord will delete the specific record
func (p *PILI) DeleteOneRecord(id uint) (bool, string){
	if len(p.AllDramas) == 0 {
		p.GetAllRecords()
	}
	for idx, record := range p.AllDramas {
		if record.ID == id {
			p.AllDramas = append(p.AllDramas[:idx], p.AllDramas[idx:]...)
			return true, ""
		}
	}
	return false, fmt.Sprintf("can not find id=%d record!", id)
}

func (p *PILI) UpdateRecord(id uint, name, summary string) (bool, string) {
	if len(p.AllDramas) == 0 {
		p.GetAllRecords()
	}

	for idx, record := range p.AllDramas {
		if record.ID == id {
			p.AllDramas[idx].Name = name
			p.AllDramas[idx].Summary = summary
			return true, ""
		}
	}
	return false, fmt.Sprintf("can not find id=%d record!", id)
}