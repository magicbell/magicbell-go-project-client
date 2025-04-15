package notifications

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
)

type DeliveryPlanCollection struct {
	Data  []DeliveryPlan `json:"data,omitempty"`
	Links *shared.Links  `json:"links,omitempty"`
}

func (d *DeliveryPlanCollection) GetData() []DeliveryPlan {
	if d == nil {
		return nil
	}
	return d.Data
}

func (d *DeliveryPlanCollection) SetData(data []DeliveryPlan) {
	d.Data = data
}

func (d *DeliveryPlanCollection) GetLinks() *shared.Links {
	if d == nil {
		return nil
	}
	return d.Links
}

func (d *DeliveryPlanCollection) SetLinks(links shared.Links) {
	d.Links = &links
}

func (d DeliveryPlanCollection) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DeliveryPlanCollection to string"
	}
	return string(jsonData)
}
