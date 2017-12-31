package aws

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
	//	logger.TestMode = true
	//	logger.Level = 4
	//	testCluster = amazon.NewUbuntuCluster("aws-ubuntu-test")
}

func TestGetVolumes(t *testing.T) {
	volumes := GetVolumes("us-west-2")
	fmt.Println(volumes)
}

func TestGetVolumesForInstance(t *testing.T) {
	volumes := GetVolumesForInstance("us-west-2", "i-08a86a21b0b7c22a7")
	fmt.Println(volumes)
}

func TestInstancesByRegions(t *testing.T) {
	InstancesByRegions([]string{"running"}, []string{"us-west-2"})
}

func TestInstancesByTag(t *testing.T) {
	resp := InstancesByTag("KubernetesCluster", "kuber", "us-west-2")
	t.Logf("%+v\n", *resp)
}

func TestListS3(t *testing.T) {
	ListS3("transics-datalake", "eu-central-1")
}
