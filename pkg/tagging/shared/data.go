package data

// *** Input configuration ***
type InputConfiguration struct {
	AccountsCoding            map[string]AccountCodingConfiguration
	AppEcosystem              string
	AppEnvironment            string
	AvailabilityZone          string
	CompanyNameLong           string
	CompanyNameShort          string
	Compliance                string
	ComputeType               string
	CostCentre                string
	CustomFQDN                string
	CustomName                string
	CustomNamePrefix          string
	CustomTags                map[string]string
	CustomTagsVerbatim        map[string]string
	Description               string
	InfraEnvironment          string
	IsApplicationLoadBalancer bool
	IsCreateBeforeDestroy     bool
	IsForceGeneratedName      bool
	IsInbound                 bool
	IsPropagateTagsAtLaunch   bool
	Account                   string
	Owner                     string
	PlatformName              string
	ProjectNameLong           string
	ProjectNameShort          string
	Region                    string
	ResourceSetLong           string
	ResourceSetShort          string
	Role                      string
	RoleExtend                string
	TerraformModule           string
	TerraformWorkspace        string
	VersioningSources         []string
}

type AccountCodingConfiguration struct {
	Class         string
	Name          string
	NameCanonical string
	NameEncoded   string
}

// *** Accounts and environment definitions ***
type EnvironmentClasses map[string]EnvironmentClassProperties

type EnvironmentClassProperties struct {
	NameEncoded string
}

// *** Resource properties ***
type ResourceProperties struct {
	AwsService string
	Id         ResourcePropertiesId
	Tags       ResourcePropertiesTags
	Name       ResourcePropertiesName
	Terraform  ResourcePropertiesTerraform
}

type ResourcePropertiesId struct {
	Key   string
	Long  string
	Short string
}

type ResourcePropertiesTags struct {
	Max        uint
	CustomList []string
}

type ResourcePropertiesName struct {
	IsLongPrefix        bool
	MaxLength           uint
	IsLowerCase         bool
	IsEnforceVersioning bool
	Components          []string
}

type ResourcePropertiesTerraform struct {
	ResourceName string
}

// *** Sanitised configuration ***

type SanitisedConfiguration struct {
	AccountId                 string
	AccountObject             AccountCodingConfiguration
	AppEnvironment            string
	CanonicalAvailabilityZone string
	CanonicalRegion           string
	CanonicalRole             string
	Constraints               Constraints
	CustomName                string
	EnvironmentClasses        EnvironmentClasses
	InfraEnvironment          string
	IsVersionedName           bool
	Keys                      Keys
	Prefixes                  Prefixes
	Separators                Separators
}

type AwsPlacements struct {
	Region           AwsPlacemenetEntity
	AvailabilityZone AwsPlacemenetEntity
}

type AwsPlacemenetEntity struct {
	Region    string
	Direction string
	Sequence  string
}

type Constraints struct {
	TagKeyMaxLength        uint
	TagValueMaxLength      uint
	TagValueReservedPrefix string
	EnvNameMaxLength       uint
	VersionStringMinLength uint
	TypePrefixMaxLength    uint
	NameTagKey             string
	CustomTagSection       string
	IsUseSectionsInNames   bool
}

type Keys struct {
	InfraEnvironment string
	AppEcosystem     string
	AppEnvironment   string
	AwsRegion        string
	AwsZone          string
}

type Prefixes struct {
	ReservedTagKey string
}

type Separators struct {
	TagComponent       string
	NameComponent      string
	ComponentWords     string
	LogstreamComponent string
}

// *** Computed resource Name and Type ***

type GeneratedName interface {
	GeneratedNameChild()
}

type Name struct {
	Safe       string
	SafePrefix string
	SafeSuffix string
	TypePrefix string
}

func (n *Name) GeneratedNameChild() {}

type Type struct {
	Long string
}

func (t *Type) GeneratedNameChild() {}
