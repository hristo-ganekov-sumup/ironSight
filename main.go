package main

import (
	"flag"
	"fmt"
	"github.com/hristo-ganekov-sumup/ironSight/internal/api"
	"github.com/hristo-ganekov-sumup/ironSight/internal/common"
	"github.com/hristo-ganekov-sumup/ironSight/internal/sg"
)

func difference(s1, s2 []common.TargetPair) ([]common.TargetPair, []common.TargetPair) {
	var onlyins1 []common.TargetPair
	var onlyins2 []common.TargetPair

	//Perspective S1
	found := false
	for i := 0; i < len(s1); i++ {
		for k := 0; k < len(s2); k++ {
			if (s1[i].Target == s2[k].Target) && (s1[i].Port == s2[k].Port) {
				found = true
				break
			}
		}
		if found {
			found = false
			continue
		} else {
			onlyins1 = append(onlyins1, s1[i])
		}
	}

	//Perspective S2
	found = false
	for i := 0; i < len(s2); i++ {
		for k := 0; k < len(s1); k++ {
			if (s2[i].Target == s1[k].Target) && (s2[i].Port == s1[k].Port) {
				found = true
				break
			}
		}
		if found {
			found = false
			continue
		} else {
			onlyins2 = append(onlyins2, s2[i])
		}
	}
	return onlyins1, onlyins2
}


var region = flag.String("region", "eu-west-1", "AWS Region")
var state_file = flag.String("state", "live.tfstate", "State file")

func init() {
	flag.Parse()
}

func main() {
	var sgs []string
	//Read the state

	fromState, err := sg.GetSGsfromStateRules(*state_file)
	if err != nil {
		panic(err)
	}

	//Get security groups that are processed from the state so we can pass them to the API parser
	for k, _ := range fromState {
		//sgs to get statement on from AWS API; This should be taken from the autosg state
		//so we are only verifying against those SGs
		sgs = append(sgs, k)
	}

	//Pull the statements in predefined exploded format; Should be the same as those red from the state file
	fromAPI, err := api.GetSGsfromAPI(*region, sgs)
	if err != nil {
		panic(err)
	}

	for k, _ := range fromState {

		onlyinStateIngress, onlyinAPIIngress := (difference(fromState[k].Ingress, fromAPI[k].Ingress))
		onlyinStateEgress, onlyinAPIEgress := (difference(fromState[k].Egress, fromAPI[k].Egress))

		if len(onlyinStateIngress) > 0 || len(onlyinAPIIngress) > 0 || len(onlyinStateEgress) > 0 || len(onlyinAPIEgress) > 0 {
			fmt.Println("*",k)
			//Print Ingress
			if len(onlyinStateIngress) > 0 {
				fmt.Println(" - Ingress Only in State: ")
				for _, instate := range onlyinStateIngress {
					fmt.Println("\t",instate)
				}
			}
			if len(onlyinAPIIngress) > 0 {
				fmt.Println(" - Ingress Only in API: ")
				for _, inapi := range onlyinAPIIngress {
					fmt.Println("\t",inapi)
				}
			}

			//Print Egress
			if len(onlyinStateEgress) > 0 {
				fmt.Println(" - Egress Only in State: ")
				for _, instate := range onlyinStateEgress {
					fmt.Println("\t",instate)
				}
			}
			if len(onlyinAPIEgress) > 0 {
				fmt.Println(" - Egress Only in API: ")
				for _, inapi := range onlyinAPIEgress {
					fmt.Println("\t",inapi)
				}
			}
		}

	}

}
