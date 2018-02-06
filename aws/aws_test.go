package aws

import (
	"fmt"
	"testing"
)

const region_eu_central_1 = "eu-central-1"
const region_us_west_2 = "us-west-2"
const region = region_eu_central_1

func TestMain(m *testing.M) {
	m.Run()
	//	logger.TestMode = true
	//	logger.Level = 4
	//	testCluster = amazon.NewUbuntuCluster("aws-ubuntu-test")
}

func TestGetVolumes(t *testing.T) {
	volumes := GetVolumes(region)
	fmt.Println(volumes)
}

func TestGetVolumesForInstance(t *testing.T) {
	volumes := GetVolumesForInstance(region, "i-08a86a21b0b7c22a7")
	fmt.Println(volumes)
}

func TestInstancesByRegions(t *testing.T) {
	InstancesByRegions([]string{"running"}, []string{region})
}

func TestInstancesByTag(t *testing.T) {
	resp := InstancesByTag("KubernetesCluster", "kuber", region)
	t.Logf("%+v\n", *resp)
}

func TestListS3(t *testing.T) {
	ListS3("transics-datalake", region)
}
