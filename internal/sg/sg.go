package sg

import (
	"encoding/json"
	"github.com/hristo-ganekov-sumup/ironSight/internal/common"
	"github.com/hristo-ganekov-sumup/ironSight/internal/tfstate"
	"strconv"
)

func GetSGsfromState(stateFilename string) (common.Sgs, error) {
	exploded := common.NewSgExploded()
	state, err := tfstate.ParseTerraformStateFile(stateFilename)
	if err != nil {
		return nil, err
	}
	for _, resource := range state.Resources {
		if resource.Type == "aws_security_group" && resource.Mode == "managed" && resource.Name == "autosg" {
			for _, instance := range resource.Instances {
				awsSg := &AwsSecurityGroup{}
				err = json.Unmarshal(instance.AttributesRaw, awsSg)
				if err != nil {
					return nil, err
				}
				var port string
				//Ingress fill
				entriesIngress := []common.TargetPair{}
				for _, ingress := range awsSg.Ingress {
					//Set the port
					if ingress.FromPort == ingress.ToPort {
						if ingress.FromPort == 0 {
							port = "-1"
						} else {
							port = strconv.Itoa(int(ingress.ToPort))
						}
					} else {
						port = strconv.Itoa(int(ingress.FromPort)) + ":" + strconv.Itoa(int(ingress.ToPort))
					}
					//If we have SG references add entries with them
					if len(ingress.SecurityGroups) > 0 {
						for _, target := range ingress.SecurityGroups {
							entriesIngress = append(entriesIngress, common.TargetPair{
								Target: target,
								Port:   port,
							})
						}
					}
					//If we have Cidr references add entries with them
					if len(ingress.CidrBlocks) > 0 {
						for _, target := range ingress.CidrBlocks {
							entriesIngress = append(entriesIngress, common.TargetPair{
								Target: target.(string),
								Port:   port,
							})
						}
					}

					if ingress.Self == true {
						entriesIngress = append(entriesIngress, common.TargetPair{
							Target: awsSg.Id,
							Port:   port,
						})
					}
				}

				//Egress fill
				entriesEgress := []common.TargetPair{}
				for _, egress := range awsSg.Egress {
					//Set the port
					if egress.FromPort == egress.ToPort {
						if egress.FromPort == 0 {
							port = "-1"
						} else {
							port = strconv.Itoa(int(egress.ToPort))
						}
					} else {
						port = strconv.Itoa(int(egress.FromPort)) + ":" + strconv.Itoa(int(egress.ToPort))
					}
					//If we have SG references add entries with them
					if len(egress.SecurityGroups) > 0 {
						for _, target := range egress.SecurityGroups {
							entriesEgress = append(entriesEgress, common.TargetPair{
								Target: target,
								Port:   port,
							})
						}
					}
					//If we have Cidr references add entries with them
					if len(egress.CidrBlocks) > 0 {
						for _, target := range egress.CidrBlocks {
							entriesEgress = append(entriesEgress, common.TargetPair{
								Target: target.(string),
								Port:   port,
							})
						}
					}

					if egress.Self == true {
						entriesEgress = append(entriesEgress, common.TargetPair{
							Target: awsSg.Id,
							Port:   port,
						})
					}

				}
				exploded_container := common.InEg{
					Ingress: entriesIngress,
					Egress:  entriesEgress,
				}
				(*exploded)[awsSg.Id] = exploded_container
			}
		}
	}
	return *exploded, nil

}

//func GetSGsfromState(stateFilename string) (common.Sgs,error) {
//	exploded := common.NewSgExploded()
//	state, err := tfstate.ParseTerraformStateFile(stateFilename)
//	if err != nil {
//		return nil, err
//	}
//
//	//SG Rules
//	for _, resource := range state.Resources {
//		if resource.Type == "aws_security_group_rule" {
//			for _, instance := range resource.Instances {
//				awsSgRuleAttributes := &AwsSecurityGroupRuleAttributes{}
//				err = json.Unmarshal(instance.AttributesRaw, awsSgRuleAttributes)
//				if err != nil {
//					return nil, err
//				}
//				entriesIngress := []common.TargetPair{}
//				//TODO: get the ingress entries
//				entriesEgress := []common.TargetPair{}
//				//TODO: get the egress entries
//				exploded_container := common.InEg{
//					Ingress: entriesIngress,
//					Egress:  entriesEgress,
//				}
//				(*exploded)[awsSgRuleAttributes.SecurityGroupID] = exploded_container
//				//sgRuleMap[awsSgRuleAttributes.SecurityGroupID] = append(sgRuleMap[awsSgRuleAttributes.SecurityGroupID], awsSgRuleAttributes.ID)
//			}
//		}
//	}
//	return *exploded,nil
//}
