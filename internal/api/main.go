package api

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hristo-ganekov-sumup/ironSight/internal/common"
	"strconv"
)

func initAwsSvc(region string) *ec2.EC2 {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region)}))

	svc := ec2.New(sess, &aws.Config{})
	return svc
}

func produceEntries(perm *ec2.IpPermission) []common.TargetPair {
	out := []common.TargetPair{}
	entry := common.NewSgExplodedEntry()
	var port string
	if perm.FromPort != nil && perm.ToPort != nil {
		pto := perm.ToPort
		pfrom := perm.FromPort
		pto_value := strconv.Itoa(int(*pto))
		pfrom_value := strconv.Itoa(int(*pfrom))
		if pto_value == pfrom_value {
			if pto_value == "0" {
				port = "-1"
			} else {
				port = pto_value
			}
		} else {
			port = pfrom_value + ":" + pto_value
		}
	} else {
		port = "-1"
	}

	if len(perm.IpRanges) > 0 {
		for _, iprange := range perm.IpRanges {
			entry.Target = *iprange.CidrIp
			entry.Port = port
			out = append(
				out,
				*entry,
			)
		}
	}
	if len(perm.UserIdGroupPairs) > 0 {
		for _, sgpair := range perm.UserIdGroupPairs {
			entry.Target = *sgpair.GroupId
			entry.Port = port
			out = append(
				out,
				*entry,
			)
		}
	}

	return out
}

func GetSGsfromAPI(region string,sgIds []string) (common.Sgs,error){
	svc := initAwsSvc(region)
	dryrun := false
	exploded := common.NewSgExploded()
	AwsSgIds := common.PointersOf(sgIds)

	SGInput := ec2.DescribeSecurityGroupsInput{
		DryRun:     &dryrun,
		GroupIds:   AwsSgIds,
	}

	SGOutput, err := svc.DescribeSecurityGroups(&SGInput)
	if err != nil {
		fmt.Println(err)
		return *exploded,err
	}

	for _, sg := range SGOutput.SecurityGroups {
		//Ingresses
		entriesIngress := []common.TargetPair{}
		for _, perm := range sg.IpPermissions {
			entriesIngress = append(entriesIngress,produceEntries(perm)...)
		}

		//Egresses
		entriesEgress := []common.TargetPair{}
		for _, perm := range sg.IpPermissionsEgress {
			entriesEgress = append(entriesEgress,produceEntries(perm)...)
		}

		exploded_container := common.InEg{
			Ingress: entriesIngress,
			Egress:  entriesEgress,
		}
		(*exploded)[common.GetString(sg.GroupId)] = exploded_container


	}
	return *exploded,nil
}