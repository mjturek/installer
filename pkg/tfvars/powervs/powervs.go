// Package aws contains AWS-specific Terraform-variable logic.
package aws

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"sigs.k8s.io/cluster-api-provider-powervs/pkg/apis/powervsprovider/v1alpha1"

	configpowervs "github.com/openshift/installer/pkg/asset/installconfig/powervs"
	"github.com/openshift/installer/pkg/types"
	typespowervs "github.com/openshift/installer/pkg/types/powervs"
	"github.com/openshift/installer/pkg/types/powervs/defaults"
)

type IBMCloudAuth struct {
        IBMCloudAPIKey		string		`json:"ibmcloud_api_key"`
        IBMCloudRegion		string		`json:"ibmcloud_region"`
        IBMCloudZone		string		`json:"ibmcloud_zone"`
}

type IBMCloudVPC struct {
	VPCName			string		`"powervs_vpc_name"`
	VPCSubnetName		string		`"powervs_vpc_subnet_name"`
}

type COSInfo struct {
	COSInstanceLocation	string		`json:"powervs_cos_instance_location"`
	COSBucketLocation	string		`json:"powervs_cos_bucket_location"`
	COSStorageClass		string		`json:"powervs_cos_storage_class"`
}

type config struct {
	IBMCloudAuth		string		`json:",inline"`
	COSInfo			string		`json:",inline"`
	IBMCloudVPC		string		`json:",inline"`

	PowerVSResourceGroup	string		`json:"powervs_resource_group"`

	ImageID			string		`json:"powervs_image_id"`
	NetworkName		string		`json:"powervs_network_ids"`

	BootstrapMemory		string		`json:"powervs_bootstrap_memory"`
	BootstrapProcessors	string		`json:"powervs_bootstrap_processors"`

	MasterMemory		string		`json:"powervs_master_memory"`
	MasterProcessors	string		`json:"powervs_master_processors"`

	ProcType		string		`json:"powervs_proc_type"`
	SysType			string		`json:"powervs_sys_type"`

}

// TFVarsSources contains the parameters to be converted into Terraform variables
type TFVarsSources struct {
	MasterConfigs, WorkerConfigs	[]*v1alpha1.PowerVSMachineProviderConfig
	IBMCloudAuth			IBMCloudAuth
	COSInfo				COSInfo
}

// TFVars generates AWS-specific Terraform variables launching the cluster.
func TFVars(sources TFVarsSources) ([]byte, error) {
	masterConfig := sources.MasterConfigs[0]

	cfg := &config{
		IBMCloudAuth
		COSInfo
		PowerVSResourceGroup
		ImageID:		masterConfig.ImageID
		NetworkName
		BootstrapMemory		masterConfig.Memory,
		BootstrapProcessors	masterConfig.Cores,
		MasterMemory:		masterConfig.Memory,
		MasterProcessors:	masterConfig.Cores,
		ProcType:		masterConfig.ProcessorType,
		SysType:		masterConfig.MachineType,
		VPCName
		VPCSubnetName
	}

	return json.MarshalIndent(cfg, "", "  ")
}
