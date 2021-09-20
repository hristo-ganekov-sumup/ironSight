package sg

type AwsSecurityGroup struct {
	Id                  string                    `json:"id"`
	Name                string                    `json:"name"`
	NamePrefix          string                    `json:"name_prefix"`
	OwnerId             string                    `json:"owner_id"`
	RevokeRulesOnDelete bool                      `json:"revoke_rules_on_delete"`
	VpcId               string                    `json:"vpc_id"`
	Arn                 string                    `json:"arn"`
	Timeouts            string                    `json:"timeouts"`
	Description         string                    `json:"description"`
	Egress              []AwsSecurityGroupEgress  `json:"egress"`
	Ingress             []AwsSecurityGroupIngress `json:"ingress"`
	Tags                map[string]string         `json:"tags,omitempty"`
}
type AwsSecurityGroupEgress struct {
	CidrBlocks     []interface{} `json:"cidr_blocks"`
	Description    string        `json:"description"`
	FromPort       int64         `json:"from_port"`
	Ipv6CidrBlocks []interface{} `json:"ipv6_cidr_blocks"`
	PrefixListIds  []interface{} `json:"prefix_list_ids"`
	Protocol       string        `json:"protocol"`
	SecurityGroups []string      `json:"security_groups"`
	Self           bool          `json:"self"`
	ToPort         int64         `json:"to_port"`
}

type AwsSecurityGroupIngress struct {
	CidrBlocks     []interface{} `json:"cidr_blocks"`
	Description    string        `json:"description"`
	FromPort       int64         `json:"from_port"`
	Ipv6CidrBlocks []interface{} `json:"ipv6_cidr_blocks"`
	PrefixListIds  []interface{} `json:"prefix_list_ids"`
	Protocol       string        `json:"protocol"`
	SecurityGroups []string      `json:"security_groups"`
	Self           bool          `json:"self"`
	ToPort         int64         `json:"to_port"`
}

type AwsSecurityGroupRuleAttributes struct {
	CidrBlocks            []string `json:"cidr_blocks"`
	Description           string   `json:"description"`
	FromPort              int64    `json:"from_port"`
	ID                    string   `json:"id"`
	Ipv6CidrBlocks        []string `json:"ipv6_cidr_blocks"`
	PrefixListIds         []string `json:"prefix_list_ids"`
	Protocol              string   `json:"protocol"`
	SecurityGroupID       string   `json:"security_group_id"`
	Self                  bool     `json:"self"`
	SourceSecurityGroupID string   `json:"source_security_group_id"`
	ToPort                int64    `json:"to_port"`
	Type                  string   `json:"type"`
}
