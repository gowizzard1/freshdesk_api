package models

import (
	"encoding/json"
	"time"
)

//CustomerResponse customer response
type CustomerResponse struct {
	Customer Customer
}

//Customer payload struct
type Customer struct {
	ID           int                    `json:"id,omitempty"`
	Name         string                 `json:"name,omitempty"`
	CustomerID   string                 `json:"cust_identifier,omitempty"`
	Description  string                 `json:"description,omitempty"`
	Domains      string                 `json:"domains,omitempty"`
	Email        string                 `json:"email,omitempty"`
	Note         string                 `json:"note,omitempty"`
	SLAPolicyID  int                    `json:"sla_policy_id,omitempty"`
	CreatedAt    time.Time              `json:"created_at,omitempty"`
	UpdatedAt    time.Time              `json:"updated_at,omitempty"`
	CustomFields map[string]interface{} `json:"custom_field,omitempty"`
}

//ConvertCustomer returns the json string of customer
func (c *CustomerResponse) ConvertCustomer() string {
	b, _ := json.Marshal(c)
	return string(b)
}
