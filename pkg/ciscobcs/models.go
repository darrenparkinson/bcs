package ciscobcs

import (
	"encoding/json"
	"time"
)

// Demo Data
// lines: 996
// types =
// done - device 300
// done - track_summary 20
// done - track_smupie_recommendation 2
// done - sw_eox_bulletin 34
// done - hw_eox_bulletin 184
// done - fn_bulletin 103
// done - psirt_bulletin 353
//

// BulkTypeChecker enables us to deserialize only the type field from a jsonlines object from the bulk endpoint
// in order to identify it's underlying type so that we can then unmarshal the remaining body appropriately.
type BulkTypeChecker struct {
	Type string `json:"type"`
}

// DateFormat represents the date only format provided in the Cisco results.
const DateFormat = "2006-01-02"

// Date represente a date provided in the Cisco results.
type Date struct {
	time.Time
}

// MarshalJSON will marshal the date format provided in the Cisco results.
func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format(DateFormat))
}

// UnmarshalJSON will unmarshal the date format provided in the Cisco results.
func (d *Date) UnmarshalJSON(data []byte) error {
	var dateStr string
	err := json.Unmarshal(data, &dateStr)
	if err != nil {
		return err
	}
	parsed, err := time.Parse(DateFormat, dateStr)
	if err != nil {
		return err
	}
	d.Time = parsed
	return nil
}

// String will return the date representation in the format provided by Cisco.
func (d Date) String() string {
	return d.Time.Format(DateFormat)
}

// DateTimeMinusTimezoneFormat represents the datetime field provided in the Cisco results which
// they have chosen not to provide any timezone information for.
const DateTimeMinusTimezoneFormat = "2006-01-02T15:04:05"

// DateTime represents the datetime field in the Cisco results which has no timezone information.
type DateTime struct {
	time.Time
}

// MarshalJSON will marshal the datetime format provided in the Cisco results.
func (d DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format(DateTimeMinusTimezoneFormat))
}

// UnmarshalJSON will unmarshal the datetime format provided in the Cisco results.
func (d *DateTime) UnmarshalJSON(data []byte) error {
	var dateStr string
	err := json.Unmarshal(data, &dateStr)
	if err != nil {
		return err
	}
	parsed, err := time.Parse(DateTimeMinusTimezoneFormat, dateStr)
	if err != nil {
		return err
	}
	d.Time = parsed
	return nil
}

// String will return the datetime representation in the format provided by Cisco.
func (d DateTime) String() string {
	return d.Time.Format(DateTimeMinusTimezoneFormat)
}

// Device defines model for Device.
type Device struct {
	// The collector identifier, which can be either a 4 character collectorid or the applianceid.
	Collector *string `json:"collector,omitempty"`

	// The Configuration register of the device.
	ConfigRegister *string `json:"configRegister,omitempty"`

	// The status of Configuration collection.  Completed means the config was successfully collected.  NotAvailable means the config was not collected.  NotSupported means the device does not support collection of an ASCii config via CLI.
	ConfigStatus *string `json:"configStatus,omitempty"`

	// The time when the collector last successfully collected the configuration from the device.
	ConfigTime *DateTime `json:"configTime,omitempty"`

	// The date the record or rule was created in NP database.  For devices, a new record is created whenever a unique name+sysobjectid combination is seen in the collector.
	CreateDate *DateTime `json:"createDate,omitempty"`

	// The unique ID of the NP NetworkElement/Device.
	DeviceId *int `json:"deviceId,omitempty"`

	// The management IP address of the device.
	DeviceIp *string `json:"deviceIp,omitempty"`

	// The Network Element Name of the device.  When used as input, it can include % as wildcard.
	DeviceName *string `json:"deviceName,omitempty"`

	// The status of the device as reported by the collector.  Usually will be either ACTIVE or DEVICE NOT REACHABLE.
	DeviceStatus *string `json:"deviceStatus,omitempty"`

	// The hostname (SNMP sysName) of the device.  It will be fully-qualified name, if domain name is set on the device.
	DeviceSysName *string `json:"deviceSysName,omitempty"`

	// The type of the device.  Values include Managed Chassis, Managed Multi-Chassis, SDR, Contexts, and IOS-XR Admin
	DeviceType *string `json:"deviceType,omitempty"`

	// The name of the software feature set running on the device.  This data is primarily available for IOS.
	FeatureSetdesc *string `json:"featureSetdesc,omitempty"`

	// The Image Name of the software on the Network Element.
	ImageName *string `json:"imageName,omitempty"`

	// Indicates whether the device is in a collector seedfile (true) or has been logically created by NP (false).  This is important for some KPI measurements to be accurate.
	InSeedFile *bool `json:"inSeedFile,omitempty"`

	// The status of Inventory collection.  Completed means some SNMP inventory was successfully collected.  NotAvailable means SNMP inventory was not collected.  NotSupported means the device was not in CSPC to be collected.
	InventoryStatus *string `json:"inventoryStatus,omitempty"`

	// The time when the collector last successfully collected inventory from the device.
	InventoryTime *DateTime `json:"inventoryTime,omitempty"`

	// An IPv4 Address.
	IpAddress *string `json:"ipAddress,omitempty"`

	// The date timestamp of the last reset of the device as reported by the show version command.
	LastReset *DateTime `json:"lastReset,omitempty"`

	// The Cisco Product Family of the hardware.  Values come from MDF for Chassis.
	ProductFamily *string `json:"productFamily,omitempty"`

	// The Cisco Product ID (PID) of the hardware.
	ProductId *string `json:"productId,omitempty"`

	// The Cisco Product Type in COLD of the hardware.  Values usually come from MDF.
	ProductType *string `json:"productType,omitempty"`

	// The reason for the last system reset as reported in the show version output.
	ResetReason *string `json:"resetReason,omitempty"`

	// The Type of Software running on the NP Network Element.  Common values include IOS, IOS XR, IOS-XE, NX-OS, etc.
	SwType *string `json:"swType,omitempty"`

	// The Software Version of the device.
	SwVersion *string `json:"swVersion,omitempty"`

	// The SNMP sysContact of the device which is populated in most devices using a configuration command.
	SysContact *string `json:"sysContact,omitempty"`

	// The SNMP system description from the device.
	SysDescription *string `json:"sysDescription,omitempty"`

	// The SNMP sysLocation of the device which is populated in most devices using a configuration command.
	SysLocation *string `json:"sysLocation,omitempty"`

	// The SNMP sysObjectID of the device.
	SysObjectId *string `json:"sysObjectId,omitempty"`

	// The user field1 value populated in the collector seedfile.
	UserField1 *string `json:"userField1,omitempty"`

	// The user field2 value populated in the collector seedfile.
	UserField2 *string `json:"userField2,omitempty"`

	// The user field3 value populated in the collector seedfile.
	UserField3 *string `json:"userField3,omitempty"`

	// The user field4 value populated in the collector seedfile.
	UserField4 *string `json:"userField4,omitempty"`
}

// TrackSummary defines model for summaryModel.
type TrackSummary struct {
	// The Type of Software running on the NP Network Element.  Common values include IOS, IOS XR, IOS-XE, NX-OS, etc.
	SwType *string `json:"swType,omitempty"`

	// The candidate/future recommended standard version for the NP Software Track.
	TrackCandidateSwVersion *string `json:"trackCandidateSwVersion,omitempty"`

	// NP Software Track Recommendation and Planning Comments.
	TrackComments *string `json:"trackComments,omitempty"`

	// Total number of devices in the NP Software Track that are running the Standard Recommended Version.
	TrackCompliantDevices *int `json:"trackCompliantDevices,omitempty"`

	// NP Software Track Description.
	TrackDescription *string `json:"trackDescription,omitempty"`

	// Internal NP Software Track identifier.  This is needed to join between various track API results.
	TrackId *int `json:"trackId,omitempty"`

	// The date when the software track was last edited or modified in Network Profile.
	TrackLastModifiedDate *Date `json:"trackLastModifiedDate,omitempty"`

	// NP Software Track Name.
	TrackName *string `json:"trackName,omitempty"`

	// Total number of devices in the NP Software Track that are not running the Standard Recommended Version.
	TrackNonCompliantDevices *int `json:"trackNonCompliantDevices,omitempty"`

	// The percent of devices running the standard recommended version.  Formula is trackCompliantDevices/trackTotalDevices.
	TrackPercentCompliant *float32 `json:"trackPercentCompliant,omitempty"`

	// The percent of devices running the standard recommended version or one of the two previous recommended versions.  Formula is (trackCompliantDevices+trackPrevCompliantDevices)/trackTotalDevices.
	TrackPercentFlexibleCompliant *float32 `json:"trackPercentFlexibleCompliant,omitempty"`

	// The PIE matching criteria for the previous standard version for the NP Software Track.  Valid values include:  Require Exact Match, Ignore For Conformance, Match but ignore extras
	TrackPrev1PieCriteria *string `json:"trackPrev1PieCriteria,omitempty"`

	// The SMU matching criteria for the previous standard version for the NP Software Track.  Valid values include:  Require Exact Match, Ignore For Conformance, Match but ignore extras
	TrackPrev1SmuCriteria *string `json:"trackPrev1SmuCriteria,omitempty"`

	// The PIE matching criteria for the 2nd previous standard version for the NP Software Track.  Valid values include:  Require Exact Match, Ignore For Conformance, Match but ignore extras
	TrackPrev2PieCriteria *string `json:"trackPrev2PieCriteria,omitempty"`

	// The SMU matching criteria for the 2nd previous standard version for the NP Software Track.  Valid values include:  Require Exact Match, Ignore For Conformance, Match but ignore extras
	TrackPrev2SmuCriteria *string `json:"trackPrev2SmuCriteria,omitempty"`

	// Total number of devices in the NP Software Track that are running the Standard Recommended Version.
	TrackPrevCompliantDevices *int `json:"trackPrevCompliantDevices,omitempty"`

	// The previous recommended standard version for the NP Software Track.
	TrackPrevSwVersion1 *string `json:"trackPrevSwVersion1,omitempty"`

	// The 2nd previous recommended standard version for the NP Software Track.
	TrackPrevSwVersion2 *string `json:"trackPrevSwVersion2,omitempty"`

	// The AS rating of the track compliance.  Results depend on whether account is using Absolute Compliance (default) or Flexible Compliance.  For Absolute, %Compliant 90 and above is Good, 60-90 is Fair, and below 60 is Poor.  For Flexible, it is the same thresholds, but %FlexibleCompliant is used.
	TrackRating *string `json:"trackRating,omitempty"`

	// The date when the last software recommendation was made.  This is manually set by the user in the NP Software Track.  This is used to measure the age of the recommendation.
	TrackRecommendationDate *Date `json:"trackRecommendationDate,omitempty"`

	// The overall compliance percentage of the devices to the recommended SMU list.  Extra SMUs are ignored for the calculation.
	TrackSmuCompliancePercent *float32 `json:"trackSmuCompliancePercent,omitempty"`

	// The PIE matching criteria for the recommended standard version for the NP Software Track.  Valid values include:  Require Exact Match, Ignore For Conformance, Match but ignore extras
	TrackStandardPieCriteria *string `json:"trackStandardPieCriteria,omitempty"`

	// The number of SMUs in the Standard Recommendation for the track.
	TrackStandardSmuCount *int `json:"trackStandardSmuCount,omitempty"`

	// The SMU matching criteria for the recommended standard version for the NP Software Track.  Valid values include:  Require Exact Match, Ignore For Conformance, Match but ignore extras
	TrackStandardSmuCriteria *string `json:"trackStandardSmuCriteria,omitempty"`

	// The current recommended standard version for the NP Software Track.
	TrackStandardSwVersion *string `json:"trackStandardSwVersion,omitempty"`

	// The status of the code deployment, as defined in the NP Software Track.  Values include Fully Deployed, In Migration, etc.
	TrackStatus *string `json:"trackStatus,omitempty"`

	// Total number of devices in the NP Software Track.
	TrackTotalDevices *int `json:"trackTotalDevices,omitempty"`

	// Total number of unique software versions in the NP Software Track.
	TrackTotalSwVersions *int `json:"trackTotalSwVersions,omitempty"`

	// The reason for the last change in software recommendation, as defined in the NP Software Track.  Values include New Software Implementation, Planned Maintenance, etc.
	TrackUpgradeReason *string `json:"trackUpgradeReason,omitempty"`
}

// TrackSmupieRecommendation defines model for track_smupie_recommendation type.
type TrackSmupieRecommendation struct {
	// The Name of the Software running on the NP Network Element.  For System SW, the value is the Image Name.  For PIE it is the package name and for SMU the SMU name.
	SwName *string `json:"swName,omitempty"`

	// The Role of the Software running on the NP Network Element.  Values include SYSTEM, PKG, SMU
	SwRole *string `json:"swRole,omitempty"`

	// Internal NP Software Track identifier.  This is needed to join between various track API results.
	TrackId *int `json:"trackId,omitempty"`

	// NP Software Track Name.
	TrackName *string `json:"trackName,omitempty"`

	// This field indicates recommendation history value of the track SMUs/PIEs.  "Current" is for the standard recommendation and should be used in most cases.  "Previous1" is for the previous recommendation.  "Previous2" is for the 2nd previous recommendation. "Candidate" is for the future candidate recommendation.
	TrackRecHistory *string `json:"trackRecHistory,omitempty"`
}

// SWEOXBulletin defines model for SWEOXBulletins.
type SWEOXBulletin struct {
	// The Cisco.com bulletin number for an End-of-Life bulletin or Field Notice.
	BulletinNumber *string `json:"bulletinNumber,omitempty"`

	// The Cisco.com Title/Headline for the bulletin.
	BulletinTitle *string `json:"bulletinTitle,omitempty"`

	// The Cisco.com URL for the bulletin.
	BulletinUrl *string `json:"bulletinUrl,omitempty"`

	// The End-of-Life Announcement (Announced) Date.
	EoLifeAnnouncementDate *DateTime `json:"eoLifeAnnouncementDate,omitempty"`

	// The End-of-Sale (EoSale) Date.
	EoSaleDate *DateTime `json:"eoSaleDate,omitempty"`

	// The End of Vulnerability/Security Support (EoVSS) Date.
	EoSecurityVulSupportDate *DateTime `json:"eoSecurityVulSupportDate,omitempty"`

	// The End of SW Maintenance Releases (EoSWM) Date.
	EoSwMaintenanceReleasesDate *DateTime `json:"eoSwMaintenanceReleasesDate,omitempty"`

	// The Last Date of Support (LDoS).
	LastDateOfSupport *DateTime `json:"lastDateOfSupport,omitempty"`

	// Internal software end-of-life identifier to allow join with master sw_eox_bulletins API.
	SwEoxId *int `json:"swEoxId,omitempty"`

	// The maintenance version portion of the software version.  For example, in 12.4(21), it is "21"
	SwMaintenanceVersion *string `json:"swMaintenanceVersion,omitempty"`

	// The major version portion of the software version.
	SwMajorVersion *string `json:"swMajorVersion,omitempty"`

	// The Software Train, typically only applies to IOS.
	SwTrain *string `json:"swTrain,omitempty"`

	// The Type of Software running on the NP Network Element.  Common values include IOS, IOS XR, IOS-XE, NX-OS, etc.
	SwType *string `json:"swType,omitempty"`
}

// HWEOXBulletin defines model for HWEOXBulletins.
type HWEOXBulletin struct {
	// The Cisco.com bulletin number for an End-of-Life bulletin or Field Notice.
	BulletinNumber *string `json:"bulletinNumber,omitempty"`

	// The Cisco.com Title/Headline for the bulletin.
	BulletinTitle *string `json:"bulletinTitle,omitempty"`

	// The Cisco.com URL for the bulletin.
	BulletinUrl *string `json:"bulletinUrl,omitempty"`

	// The End-of-Life Announcement (Announced) Date.
	EoLifeAnnouncementDate *DateTime `json:"eoLifeAnnouncementDate,omitempty"`

	// The End of New Service Attachment Date.
	EoNewServiceAttachDate *DateTime `json:"eoNewServiceAttachDate,omitempty"`

	// The End of Routine Failure Analysis Date (EoRFA) Date.
	EoRoutineFailureAnalysisDate *DateTime `json:"eoRoutineFailureAnalysisDate,omitempty"`

	// The End-of-Sale (EoSale) Date.
	EoSaleDate *DateTime `json:"eoSaleDate,omitempty"`

	// The End of Vulnerability/Security Support (EoVSS) Date.
	EoSecurityVulSupportDate *DateTime `json:"eoSecurityVulSupportDate,omitempty"`

	// The End of Service Contract Renewal (EoSCR) Date.
	EoSoftwareContractRenewalDate *DateTime `json:"eoSoftwareContractRenewalDate,omitempty"`

	// The End of SW Maintenance Releases (EoSWM) Date.
	EoSwMaintenanceReleasesDate *DateTime `json:"eoSwMaintenanceReleasesDate,omitempty"`

	// Internal hardware end-of-life identifier to allow join with master hw_eox_bulletins API.
	HwEoxId *int `json:"hwEoxId,omitempty"`

	// The Last Date of Support (LDoS).
	LastDateOfSupport *DateTime `json:"lastDateOfSupport,omitempty"`

	// The Last Ship Date.
	LastShipDate *DateTime `json:"lastShipDate,omitempty"`

	// The Cisco Product ID (PID) of the hardware.
	ProductId *string `json:"productId,omitempty"`
}

// FNBulletin defines model for FNBulletins.
type FNBulletin struct {
	// The date when the bulletin was first published to Cisco.com.  Most API calls will allow Regex input for this field.
	BulletinFirstPublished *string `json:"bulletinFirstPublished,omitempty"`

	// The date when the bulletin was last updated on Cisco.com.
	BulletinLastUpdated *DateTime `json:"bulletinLastUpdated,omitempty"`

	// The Bulletin Mapping Caveat gives any explanations why the automation may need additional review by the customer.
	BulletinMappingCaveat *string `json:"bulletinMappingCaveat,omitempty"`

	// The Cisco.com Title/Headline for the bulletin.
	BulletinTitle *string `json:"bulletinTitle,omitempty"`

	// The Cisco.com URL for the bulletin.
	BulletinUrl *string `json:"bulletinUrl,omitempty"`

	// Field Notice ID number.
	FieldNoticeId *string `json:"fieldNoticeId,omitempty"`

	// Type of Field Notice as defined from PLATO.  Valid values include: hardware, software, other
	FnType *string `json:"fnType,omitempty"`

	// The description of the problem on a Cisco bulletin.
	ProblemDescription *string `json:"problemDescription,omitempty"`
}

// PSIRTBulletin defines model for PSIRTBulletins.
type PSIRTBulletin struct {
	// The date when the bulletin was first published to Cisco.com.  Most API calls will allow Regex input for this field.
	BulletinFirstPublished *string `json:"bulletinFirstPublished,omitempty"`

	// The date when the bulletin was last updated on Cisco.com.
	BulletinLastUpdated *DateTime `json:"bulletinLastUpdated,omitempty"`

	// The Bulletin Mapping Caveat gives any explanations why the automation may need additional review by the customer.
	BulletinMappingCaveat *string `json:"bulletinMappingCaveat,omitempty"`

	// The Summary of a Cisco.com bulletin.
	BulletinSummary *string `json:"bulletinSummary,omitempty"`

	// The Cisco.com Title/Headline for the bulletin.
	BulletinTitle *string `json:"bulletinTitle,omitempty"`

	// The Cisco.com URL for the bulletin.
	BulletinUrl *string `json:"bulletinUrl,omitempty"`

	// The version # of the Cisco.com bulletin.
	BulletinVersion *string `json:"bulletinVersion,omitempty"`

	// Comma-separated list of Cisco Bug IDs.
	CiscoBugIds *string `json:"ciscoBugIds,omitempty"`

	// Common Vulnerabilities and Exposures (CVE) Identifier
	CveId *string `json:"cveId,omitempty"`

	// Common Vulnerability Scoring System (CVSS) Base Score
	CvssBase *string `json:"cvssBase,omitempty"`

	// Common Vulnerability Scoring System (CVSS) Temporal Score
	CvssTemporal *string `json:"cvssTemporal,omitempty"`

	// The Advisory ID of a PSIRT as seen on Cisco.com.
	PsirtAdvisoryId *string `json:"psirtAdvisoryId,omitempty"`

	// The internal COLD ID for a PSIRT.  This is useful for joining multiple data sources.
	PsirtColdId *int `json:"psirtColdId,omitempty"`

	// The Security Impact Rating (SIR) for Cisco PSIRTs.
	Sir *string `json:"sir,omitempty"`
}
