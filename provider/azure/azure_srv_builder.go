package azure

import (
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/dns/mgmt/2018-05-01/dns"

	"sigs.k8s.io/external-dns/endpoint"
)

func extractSrvTarget(endpoint *endpoint.Endpoint) *[]dns.SrvRecord {
	target := strings.Split(endpoint.Targets[0], " ")

	priority, _ := strconv.Atoi(target[0])
	prio32 := int32(priority)

	weighting, _ := strconv.Atoi(target[1])
	weight32 := int32(weighting)

	port, _ := strconv.Atoi(target[2])
	port32 := int32(port)

	targetname := target[3]

	return &[]dns.SrvRecord{
		{
			Priority: &prio32,
			Weight:   &weight32,
			Target:   &targetname,
			Port:     &port32,
		},
	}
}
