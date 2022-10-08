package powervs

// Metadata contains Power VS metadata (e.g. for uninstalling the cluster).
type Metadata struct {
	BaseDomain           string `json:"BaseDomain"`
	CISInstanceCRN       string `json:"cisInstanceCRN"`
	DNSInstanceCRN       string `json:"dnsInstanceCRN"`
	PowerVSResourceGroup string `json:"powerVSResourceGroup"`
	Region               string `json:"region"`
	VPCRegion            string `json:"vpcRegion"`
	VPCName              string `json:"vpcName,omitempty"`
	Zone                 string `json:"zone"`
	ServiceInstanceGUID  string `json:"serviceInstanceID"`
}
