package slots

import (
	"encoding/json"
	"fmt"

	"github.com/datalayer/kuber/log"
)

type Slot struct {
	Id    int    `json:"id"`
	Start string `json:"start"`
	End   string `json:"end"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

var Slots []Slot

func init() {

	var s = []byte(`[
		{"Id":1, "Start": "Mon Feb 12 2018 12:00:00 GMT+0100 (CET)", "End": "Mon Feb 12 2018 13:00:00 GMT+0100 (CET)", "Title": "Title 1", "Desc": "Desc 1"}
	]`)
	err := json.Unmarshal(s, &Slots)
	if err != nil {
		fmt.Println("error:", err)
	}

	log.Info("Initial Slots: %v", Slots)

}

func PutSlots(slots []Slot) {
	log.Info("Slots: %v", slots)
	Slots = slots
}
func GetSlots() []Slot {
	log.Info("Slots: %v", Slots)
	return Slots
}
