package model

type Response struct {
	Arguments interface{} `json:"arguments"`
	Result    int         `json:"result"`
	Text      string      `json:"text"`
}

type Lease4GetAllResponse struct {
	Leases []Lease4GetAllResponse_Lease `json:"leases"`
}

type Lease4GetAllResponse_Lease struct {
	ClientId  string `json:"client-id"`
	Cltt      int    `json:"cltnt"`
	FqdnFwd   bool   `json:"fqdn-fwd"`
	FqdnRev   bool   `json:"fqdn-rev"`
	Hostname  string `json:"hostname"`
	HwAddress string `json:"hw-address"`
	IpAddress string `json:"ip-address"`
	State     int    `json:"state"`
	SubnetId  int    `json:"subnet-id"`
	ValidLft  bool   `json:"valid-lft"`
}
