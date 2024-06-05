package entities

type SecurityRules struct {
	SecurityGroupId []string  `json:"security_group_ids"`
	IPs             []IPRange `json:"ips"`
}

type IPRange struct {
	CIDR        string `json:"cidr"`
	Description string `json:"description"`
}

func NewSecurityGroup() *SecurityRules {
	return &SecurityRules{}
}
