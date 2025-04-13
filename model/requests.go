package model

type Command struct {
	Command   string      `json:"command"`
	Arguments interface{} `json:"arguments"`
	Service   []string    `json:"service"`
}

type Lease4GetAllArguments struct {
	Subnets []int `json:"subnets"`
}
