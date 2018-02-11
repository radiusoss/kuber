package slots

import (
	"github.com/datalayer/kuber/log"
)

type Slot struct {
	Id    int    `json:"id"`
	Start string `json:"start"`
	End   string `json:"end"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func PutSlots(slots []Slot) {
	log.Info("Slots: %v", slots)
}
