package main


type server struct {
	GroupId	string `json:"GroupId,omitempty"`
	RegionNodeId string `json:"RegionNodeId,omitempty"`
	CacheSize struct {
		Value float64  `json:"Value,omitempty"`
	} `json:"CacheSize,omitempty"`
	State float64 `json:"State,omitempty"`
	SupportState float64 `json:"SupportState,omitempty"` 
	SnapshotMapStoreId string `json:"SnapshotMapStoreId,omitempty"`
	SnapshotMapStorePoolId string `json:"SnapshotMapStorePoolId,omitempty"`
	InstallPath string `json:"InstallPath,omitempty"` 
	ProductName string `json:"ProductName,omitempty"` 
	ProductType string `json:"ProductType,omitempty"`
	ProductVersion string `json:"ProductVersion,omitempty"`
	OsVersion string `json:"OsVersion,omitempty"`
	ProcessorInfo struct {
		CpuArchitecture float64 `json:"CpuArchitecture,omitempty"` 
		ProcessorName string `json:"ProcessorName,omitempty"` 
		NumberCores float64 `json:"NumberCores,omitempty"` 
		NumberPhysicalCores float64 `json:"NumberPhysicalCores,omitempty"`
	} `json:"ProcessorInfo,omitempty"`
	ProductBuild string `json:"ProductBuild,omitempty"` 
	BuildType string `json:"BuildType,omitempty"`
	DiagnosticMode float64 `json:"DiagnosticMode,omitempty"`
	LicenseRemaining float64 `json:"LicenseRemaining,omitempty"`
	ReplicationBufferFolder string `json:"ReplicationBufferFolder,omitempty"`
	TotalSystemMemory struct {
		Value float64 `json:"Value,omitempty"`
	} `json:"TotalSystemMemory,omitempty"`
	AvailableSystemMemory struct {
		Value float64 `json:"Value,omitempty"`
	} `json:"AvailableSystemMemory,omitempty"`
	LogStatus float64 `json:"LogStatus,omitempty"`
	LicenseSettings struct  {
		MaxServers float64 `json:"MaxServers,omitempty"`
		MaxPartnerGroups float64 `json:"MaxPartnerGroups,omitempty"`
		MaxMappedHosts float64 `json:"MaxMappedHosts,omitempty"`
		BidirectionalReplication float64 `json:"BidirectionalReplication,omitempty"`
		FiberChannel float64 `json:"FiberChannel,omitempty"` 
		ThinProvisioning float64 `json:"ThinProvisioning,omitempty"`
		Snapshot float64 `json:"Snapshot,omitempty"`
		iSCSI float64 `json:"iSCSI,omitempty"`
		StorageCapacity struct {
			Value float64 `json:"Value,omitempty"` 
		} `json:"StorageCapacity,omitempty"`
		LicensedBulkStorage struct {
			Value float64 `json:"Value,omitempty"`
		} `json:"LicensedBulkStorage,omitempty"`
		BulkEnabled float64 `json:"BulkEnabled,omitempty"` 
		RetentionTime float64 `json:"RetentionTime,omitempty"` 
		AutoTiering float64 `json:"AutoTiering,omitempty"` 
		HeatMaps float64 `json:"HeatMaps,omitempty"` 
		SharedStorage float64 `json:"SharedStorage,omitempty"` 
		PerformanceAnalysis float64 `json:"PerformanceAnalysis,omitempty"` 
		ResourceAuthorization float64 `json:"ResourceAuthorization,omitempty"` 
		SequentialStorage float64 `json:"SequentialStorage,omitempty"` 
		MaxBypassThreads float64 `json:"MaxBypassThreads,omitempty"`
		MaxPollerThreads float64 `json:"MaxPollerThreads,omitempty"`
		Mirroring float64 `json:"Mirroring,omitempty"`
		Witness float64 `json:"Witness,omitempty"`
		DataAtRestEncryption float64 `json:"DataAtRestEncryption,omitempty"`
	} `json:"LicenseSettings,omitempty"`
	IsLicensed bool `json:"IsLicensed,omitempty"`
	LicenseExceeded bool `json:"LicenseExceeded,omitempty"`
	StorageUsed struct {
		Value float64 `json:"Value,omitempty"`
	} `json:"StorageUsed,omitempty"`
	BulkStorageUsed struct {
		Value float64 `json:"Value,omitempty"`
	} `json:"BulkStorageUsed,omitempty"`
	DataCoreStorageUsed struct {
		Value float64 `json:"Value,omitempty"`
	} `json:"DataCoreStorageUsed,omitempty"`
	DataCoreBulkStorageUsed struct {
		Value float64 `json:"Value,omitempty"`
	} `json:"DataCoreBulkStorageUsed,omitempty"`
	PowerState float64 `json:"PowerState,omitempty"`
	CacheState float64 `json:"CacheState,omitempty"`
	BackupStorageFolder string `json:"BackupStorageFolder,omitempty"`
	IpAddresses   []string `json:"IpAddresses,omitempty"`
	LicenseNumber float64 `json:"LicenseNumber,omitempty"`
	AluaGroupId float64 `json:"AluaGroupId,omitempty"`
	IsVirtualMachine bool `json:"IsVirtualMachine,omitempty"`
	IsAzureVirtualMachine bool `json:"IsAzureVirtualMachine,omitempty"`
	IsPayGo bool `json:"IsPayGo,omitempty"`
	HypervisorHostId string `json:"HypervisorHostId,omitempty"`
	LogStorePoolId string `json:"LogStorePoolId,omitempty"`
	ConfigurationInfo struct {
		SequenceNumber float64 `json:"SequenceNumber,omitempty"`
		TimeStamp string `json:"TimeStamp,omitempty"`
	} `json:"ConfigurationInfo,omitempty"`
	Description string `json:"Description,omitempty"`
	HostName string `json:"HostName,omitempty"`
	MpioCapable bool `json:"MpioCapable,omitempty"`
	SequenceNumber float64 `json:"SequenceNumber,omitempty"`
	Id string `json:"Id,omitempty"`
	Caption string `json:"Caption,omitempty"`
	ExtendedCaption string `json:"ExtendedCaption,omitempty"`
	float64ernal bool `json:"float64ernal,omitempty"`
}

type pool struct {
	ServerId string `json:"ServerId,omitempty"` 
	SharedPoolId string `json:"SharedPoolId,omitempty"` 
	Alias string `json:"Alias,omitempty"` 
    Description string `json:"Description,omitempty"` 
    PresenceStatus float64 `json:"PresenceStatus,omitempty"`
    PoolStatus float64 `json:"PoolStatus,omitempty"`
    Type float64 `json:"Type,omitempty"`
    ChunkSize struct {
		Value float64 `json:"Value,omitempty"`
    } `json:"ChunkSize,omitempty"`
    SectorSize struct {
		Value float64 `json:"Value,omitempty"`
    } `json:"SectorSize,omitempty"`
	MaxTierNumber float64 `json:"MaxTierNumber,omitempty"`
	TierReservedPct float64 `json:"TierReservedPct,omitempty"`
    AutoTieringEnabled bool `json:"AutoTieringEnabled,omitempty"`
    IsAuthorized bool `json:"IsAuthorized,omitempty"`
    InSharedMode bool `json:"InSharedMode,omitempty"`
    SMPAApproved bool `json:"SMPAApproved,omitempty"`
    SupportsEncryption bool `json:"SupportsEncryption,omitempty"`
    HasEncryption bool `json:"HasEncryption,omitempty"`
    SequenceNumber float64 `json:"SequenceNumber,omitempty"`
    Id string `json:"Id,omitempty"` 
	Caption string `json:"Caption,omitempty"` 
	ExtendedCaption string `json:"ExtendedCaption,omitempty"` 
    Internal bool `json:"Internal,omitempty"`
    PoolMode float64 `json:"PoolMode,omitempty"`
}

type virtualdisk struct {
	VirtualDiskGroupId string `json:"VirtualDiskGroupId,omitempty"` 
	FirstHostId string `json:"FirstHostId,omitempty"` 
	SecondHostId string `json:"SecondHostId,omitempty"` 
	BackupHostId string `json:"BackupHostId,omitempty"`
	StorageProfileId string `json:"StorageProfileId,omitempty"`
	WitnessId string `json:"WitnessId,omitempty"`
	SnapshotPoolId string `json:"SnapshotPoolId,omitempty"`
	Alias string `json:"Alias,omitempty"`
	Description string `json:"Description,omitempty"` 
	Size struct  {
		Value float64 `json:"Value,omitempty"`
	} `json:"Size,omitempty"` 
	SectorSize struct {
		Value float64 `json:"Value,omitempty"`
	} `json:"SectorSize,omitempty"` 
	Type float64 `json:"Type,omitempty"`
	DiskStatus float64 `json:"DiskStatus,omitempty"`
	InquiryData struct {
		Vendor string `json:"Vendor,omitempty"`
		Product string `json:"Product,omitempty"`
		Revision string `json:"Revision,omitempty"`
		Serial string `json:"Serial,omitempty"`
	} `json:"InquiryData,omitempty"`
	ScsiDeviceId string `json:"ScsiDeviceId,omitempty"`
	ScsiDeviceIdString string `json:"ScsiDeviceIdString,omitempty"`
	RemovableMedia bool `json:"RemovableMedia,omitempty"`
	WriteThrough bool `json:"WriteThrough,omitempty"`
	NVMe bool `json:"NVMe,omitempty"`
	Offline bool `json:"Offline,omitempty"`
	Disabled bool `json:"Disabled,omitempty"`
	ManualRecovery bool `json:"ManualRecovery,omitempty"`
	DiskLayout struct {
		Cylinders float64 `json:"Cylinders,omitempty"`
		Heads float64 `json:"Heads,omitempty"`
		SectorsPerTrack float64 `json:"SectorsPerTrack,omitempty"`
	} `json:"DiskLayout,omitempty"`
	PersistentReserveEnabled bool `json:"PersistentReserveEnabled,omitempty"`
	RecoveryPriority float64 `json:"RecoveryPriority,omitempty"`
	TPThresholdsEnabled bool `json:"TPThresholdsEnabled,omitempty"`
	IsServed bool `json:"IsServed,omitempty"`
	SubType float64 `json:"SubType,omitempty"`
	MirrorTrunkMappingEnabled bool `json:"MirrorTrunkMappingEnabled,omitempty"`
	WitnessOption string `json:"WitnessOption,omitempty"`
	PreferredServer string `json:"PreferredServer,omitempty"`
	ResiliencyEnabled bool `json:"ResiliencyEnabled,omitempty"`
	EncryptionEnabled bool `json:"EncryptionEnabled,omitempty"`
	SequenceNumber float64 `json:"SequenceNumber,omitempty"`
	Id string `json:"Id,omitempty"`
	Caption string `json:"Caption,omitempty"`
	ExtendedCaption string `json:"ExtendedCaption,omitempty"`
	Internal bool `json:"Internal,omitempty"`
}

type physicaldisk struct {
	RemoteLogicalDiskId string `json:"RemoteLogicalDiskId,omitempty"`
	LocalLogicalDiskId string `json:"LocalLogicalDiskId,omitempty"`
	PresenceStatus float64 `json:"PresenceStatus,omitempty"`
	Alias string `json:"Alias,omitempty"`
	Size struct {
		Value float64 `json:"Value,omitempty"`
	} `json:"Size,omitempty"`
	SectorSize struct {
		Value float64 `json:"Value,omitempty"`
	} `json:"SectorSize,omitempty"`
	FreeSpace struct {
		Value float64 `json:"Value,omitempty"`
	} `json:"FreeSpace,omitempty"`
	InquiryData struct {
		Vendor string `json:"Vendor,omitempty"`
		Product string `json:"Product,omitempty"`
		Revision string `json:"Revision,omitempty"`
		Serial string `json:"Serial,omitempty"`
	} `json:"InquiryData,omitempty"`
	ScsiPath struct {
		Port float64 `json:"Port,omitempty"`
		Bus float64 `json:"Bus,omitempty"`
		Target float64 `json:"Target,omitempty"`
		LUN float64 `json:"LUN,omitempty"`
	} `json:"ScsiPath,omitempty"`
	DiskIndex float64 `json:"DiskIndex,omitempty"`
	SystemName string `json:"SystemName,omitempty"`
	BusType float64 `json:"BusType,omitempty"`
	Type float64 `json:"Type,omitempty"`
	DiskStatus float64 `json:"DiskStatus,omitempty"`
	PoolMemberId string `json:"PoolMemberId,omitempty"`
	Partitioned bool `json:"Partitioned,omitempty"`
	InUse bool `json:"InUse,omitempty"`
	IsBootDisk bool `json:"IsBootDisk,omitempty"`
	Protected bool `json:"Protected,omitempty"`
	IsSolidState bool `json:"IsSolidState,omitempty"`
	HostId string `json:"HostId,omitempty"`
	UniqueIdentifier string `json:"UniqueIdentifier,omitempty"`
	SharedPhysicalDiskId string `json:"SharedPhysicalDiskId,omitempty"`
	IsDataCoreDisk bool `json:"IsDataCoreDisk,omitempty"`
	SequenceNumber float64 `json:"SequenceNumber,omitempty"`
	Id string `json:"Id,omitempty"`
	Caption string `json:"Caption,omitempty"`
	ExtendedCaption string `json:"ExtendedCaption,omitempty"`
	Internal bool `json:"Internal,omitempty"`
}

type port struct {
	PresenceStatus float64 `json:"PresenceStatus,omitempty"`
	PhysicalName string `json:"PhysicalName,omitempty"`
	ServerPortProperties struct  {
		Role float64 `json:"Role,omitempty"`
		PortGroup string `json:"PortGroup,omitempty"`
	} `json:"ServerPortProperties,omitempty"`
	RoleCapability float64 `json:"RoleCapability,omitempty"`
	AluaId float64 `json:"AluaId,omitempty"`
	PortName string `json:"PortName,omitempty"`
	Alias string `json:"Alias,omitempty"`
	Description string `json:"Description,omitempty"`
	PortType float64 `json:"PortType,omitempty"`
	PortMode float64 `json:"PortMode,omitempty"`
	HostId string `json:"HostId,omitempty"`
	Connected bool `json:"Connected,omitempty"`
	SequenceNumber float64 `json:"SequenceNumber,omitempty"`
	Id string `json:"Id,omitempty"`
	Caption string `json:"Caption,omitempty"`
	ExtendedCaption string `json:"ExtendedCaption,omitempty"`
	Internal bool `json:"Internal,omitempty"`
}

type host struct {
	HostId string `json:"HostId,omitempty"`
	PowerState float64 `json:"PowerState,omitempty"`
	ProvisionedSpace struct {
		Value float64 `json:"Value,omitempty"`
	} `json:"ProvisionedSpace,omitempty"`
	UsedSpace struct {
		Value float64 `json:"Value,omitempty"`
	} `json:"UsedSpace,omitempty"`
	Template bool `json:"Template,omitempty"`
	ConnectorId string `json:"ConnectorId,omitempty"`
	DatacenterId string `json:"DatacenterId,omitempty"`
	Datastores []string `json:"Datastores,omitempty"`
	HostGroupId string `json:"HostGroupId,omitempty"`
	Type float64 `json:"Type,omitempty"`
	State float64 `json:"State,omitempty"`
	AluaSupport bool `json:"AluaSupport,omitempty"`
	PathPolicy float64 `json:"PathPolicy,omitempty"`
	TelemetryData struct {
		DockerVersion string `json:"DockerVersion,omitempty"`
		KubernetesVersion string `json:"KubernetesVersion,omitempty"`
		NumberOfContainers float64 `json:"NumberOfContainers,omitempty"`
		PersistentVolumes string `json:"PersistentVolumes,omitempty"`
	} `json:"TelemetryData,omitempty"`
	Description string `json:"Description,omitempty"`
	HostName string `json:"HostName,omitempty"`
	MpioCapable bool `json:"MpioCapable,omitempty"`
	SequenceNumber float64 `json:"SequenceNumber,omitempty"`
	Id string `json:"Id,omitempty"`
	Caption string `json:"Caption,omitempty"`
	ExtendedCaption string `json:"ExtendedCaption,omitempty"`
	Internal bool `json:"Internal,omitempty"`
}

type preferredserver struct {
	ClientId string `json:"ClientId,omitempty"`
	ServerId string `json:"ServerId,omitempty"`
	SequenceNumber float64 `json:"SequenceNumber,omitempty"`
	Id string `json:"Id,omitempty"`
	Caption string `json:"Caption,omitempty"`
	ExtendedCaption string `json:"ExtendedCaption,omitempty"`
	Internal bool `json:"Internal,omitempty"`
}

type servergroup struct {
	OurGroup bool `json:"OurGroup,omitempty"`
	Alias  string `json:"Alias,omitempty"`
	Description  string `json:"Description,omitempty"`
	State  float64 `json:"State,omitempty"`
	SmtpSettings struct {
		SmtpServer string `json:"SmtpServer,omitempty"`
		EmailAddress string `json:"EmailAddress,omitempty"`
		Username string `json:"Username,omitempty"`
		TcpPort string `json:"TcpPort,omitempty"`
		UseSSL bool `json:"UseSSL,omitempty"`
	} `json:"SmtpSettings,omitempty"`
	LicenseSettings  struct {
		MaxServers  float64 `json:"MaxServers,omitempty"`
		MaxPartnerGroups  float64 `json:"MaxPartnerGroups,omitempty"`
		MaxMappedHosts  float64 `json:"MaxMappedHosts,omitempty"`
		BidirectionalReplication  float64 `json:"BidirectionalReplication,omitempty"`
		FiberChannel  float64 `json:"FiberChannel,omitempty"`
		ThinProvisioning  float64 `json:"ThinProvisioning,omitempty"`
		Snapshot  float64 `json:"Snapshot,omitempty"`
		iSCSI  float64 `json:"iSCSI,omitempty"`
		StorageCapacity struct {
			Value  float64 `json:"Value,omitempty"`
		} `json:"StorageCapacity,omitempty"`
		LicensedBulkStorage  struct {
			Value float64 `json:"Value,omitempty"`
		} `json:"LicensedBulkStorage,omitempty"`
		BulkEnabled  float64 `json:"BulkEnabled,omitempty"`
		RetentionTime   float64 `json:"RetentionTime,omitempty"`
		AutoTiering   float64 `json:"AutoTiering,omitempty"`
		HeatMaps   float64 `json:"HeatMaps,omitempty"`
		SharedStorage   float64 `json:"SharedStorage,omitempty"`
		PerformanceAnalysis   float64 `json:"PerformanceAnalysis,omitempty"`
		ResourceAuthorization   float64 `json:"ResourceAuthorization,omitempty"`
		SequentialStorage   float64 `json:"SequentialStorage,omitempty"`
		MaxBypassThreads   float64 `json:"VaMaxBypassThreadslue,omitempty"`
		MaxPollerThreads   float64 `json:"MaxPollerThreads,omitempty"`
		Mirroring   float64 `json:"Mirroring,omitempty"`
		Witness   float64 `json:"Witness,omitempty"`
		DataAtRestEncryption   float64 `json:"DataAtRestEncryption,omitempty"`
	} `json:"LicenseSettings,omitempty"`
	LicenseType  float64 `json:"LicenseType,omitempty"`
	ContactData struct {
		ContactName  string `json:"ContactName,omitempty"`
		PhoneNumber  string `json:"PhoneNumber,omitempty"`
		EmailAddress  string `json:"EmailAddress,omitempty"`
		CompanyName  string `json:"CompanyName,omitempty"`
	} `json:"ContactData,omitempty"`
	StorageUsed  struct {
		Value  float64 `json:"Value,omitempty"`
	} `json:"StorageUsed,omitempty"`
	BulkStorageUsed struct {
		Value  float64 `json:"Value,omitempty"`
	} `json:"BulkStorageUsed,omitempty"`
	MaxStorage  struct {
		Value  float64 `json:"Value,omitempty"`
	} `json:"MaxStorage,omitempty"`
	RecoverySpeed  float64 `json:"RecoverySpeed,omitempty"`
	ExistingProductKeys  []struct {
		Key string `json:"Key,omitempty"`
        LastFive string `json:"LastFive,omitempty"`
        ServerId string `json:"ServerId,omitempty"`
        Capacity struct {
            Value float64 `json:"Value,omitempty"`
        } `json:"Capacity,omitempty"`
        ActualCapacity struct {
            Value float64 `json:"Value,omitempty"`
        } `json:"ActualCapacity,omitempty"`
        CapacityConsumed struct {
            Value float64 `json:"Value,omitempty"`
        } `json:"CapacityConsumed,omitempty"`
        KeyCode string `json:"KeyCode,omitempty"`
        ProductName string `json:"ProductName,omitempty"`
        SKU string `json:"SKU,omitempty"`
        ExpirationDate string `json:"ExpirationDate,omitempty"`
        KeyType float64 `json:"KeyType,omitempty"`
        Subscription bool `json:"Subscription,omitempty"`
        IsBaseLicense bool `json:"IsBaseLicense,omitempty"`
        Expired bool `json:"Expired,omitempty"`
	} `json:"ExistingProductKeys,omitempty"`
	DataCoreStorageUsed struct {
		Value  float64 `json:"Value,omitempty"`
	} `json:"DataCoreStorageUsed,omitempty"`
	SupportBundleRelayAddress  string `json:"SupportBundleRelayAddress,omitempty"`
	MirrorTrunkMappingEnabled  bool `json:"MirrorTrunkMappingEnabled,omitempty"`
	SelfHealingDelay  float64 `json:"SelfHealingDelay,omitempty"`
	DefaultWitness  string `json:"DefaultWitness,omitempty"`
	DefaultWitnessOption  float64 `json:"SelfHeaDefaultWitnessOptionlingDelay,omitempty"`
	WitnessAllowed  bool `json:"WitnessAllowed,omitempty"`
	Telemetry  string `json:"Telemetry,omitempty"`
	DefaultKmipEndpoint  string `json:"DefaultKmipEndpoint,omitempty"`
	CrashRecoveryCount  float64 `json:"CrashRecoveryCount,omitempty"`
	OutOfCompliance  bool `json:"OutOfCompliance,omitempty"`
	NextExpirationDate  string `json:"NextExpirationDate,omitempty"`
	LicenseRemaining  float64 `json:"LicenseRemaining,omitempty"`
	SequenceNumber  float64 `json:"SequenceNumber,omitempty"`
	Id  string `json:"Id,omitempty"`
	Caption  string `json:"Caption,omitempty"`
	ExtendedCaption  string `json:"ExtendedCaption,omitempty"`
	Internal  bool `json:"Internal,omitempty"`
}