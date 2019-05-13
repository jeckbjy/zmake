package xcode

type xcPBXMap map[string]interface{}

type xcISA struct {
	ISA string `plist:"isa"`
}

type xcPBXProject struct {
	xcISA
	BuildConfigurationList string   `plist:"buildConfigurationList"`
	CompatibilityVersion   string   `plist:"compatibilityVersion"`
	HasScannedForEncodings string   `plist:"hasScannedForEncodings"`
	MainGroup              string   `plist:"mainGroup"`
	ProjectDirPath         string   `plist:"projectDirPath"`
	ProjectReferences      []string `plist:"projectReferences"`
	ProjectRoot            string   `plist:"projectRoot"`
	Targets                []string `plist:"targets"`
}

type xcPBXGroup struct {
	xcISA
	Children   []string `plist:"children"`
	Name       string   `plist:"name"`
	SourceTree string   `plist:"sourceTree"`
}

type xcPBXFileRef struct {
	xcISA
	LastKnownFileType string `plist:"lastKnownFileType"`
	Name              string `plist:"name"`
	Path              string `plist:"path"`
	SourceTree        string `plist:"sourceTree"`
}

type xcPBXNativeTarget struct {
	xcISA
	BuildConfigurationList string   `plist:"buildConfigurationList"`
	BuildPhases            []string `plist:"buildPhases"`
	BuildRules             []string `plist:"buildRules"`
	Dependencies           []string `plist:"dependencies"`
	Name                   string   `plist:"name"`
	ProductInstallPath     string   `plist:"productInstallPath"`
	ProductName            string   `plist:"productName"`
	ProductReference       string   `plist:"productReference"`
	ProductType            string   `plist:"productType"`
}

type xcPBXSourcesBuildPhase struct {
	xcISA
	BuildActionMask                    int64    `plist:"buildActionMask"`
	Files                              []string `plist:"files"`
	RunOnlyForDeploymentPostprocessing int      `plist:"runOnlyForDeploymentPostprocessing"`
}

type xcPBXBuildFile struct {
	xcISA
	LastKnownFileType string `plist:"lastKnownFileType"`
	Name              string `plist:"name"`
	Path              string `plist:"path"`
	SourceTree        string `plist:"sourceTree"`
}

type xcPBXContainerItemProxy struct {
	xcISA
	ContainerPortal      string `plist:"containerPortal"`
	ProxyType            int    `plist:"proxyType"`
	RemoteGlobalIDString string `plist:"remoteGlobalIDString"`
	RemoteInfo           string `plist:"remoteInfo"`
}

type xcPBXCopyFilesBuildPhase struct {
	xcISA
	BuildActionMask                    int64    `plist:"buildActionMask"`
	DstPath                            string   `plist:"dstPath"`
	DstSubfolderSpec                   int      `plist:"dstSubfolderSpec"`
	Files                              []string `plist:"files"`
	Name                               string   `plist:"name"`
	RunOnlyForDeploymentPostprocessing int      `plist:"runOnlyForDeploymentPostprocessing"`
}

type xcPBXFrameworksBuildPhase struct {
	xcISA
	BuildActionMask                    int64    `plist:"buildActionMask"`
	Files                              []string `plist:"files"`
	RunOnlyForDeploymentPostprocessing int      `plist:"runOnlyForDeploymentPostprocessing"`
}

type xcPBXHeadersBuildPhase struct {
	xcISA
	BuildActionMask                    int64    `plist:"buildActionMask"`
	Files                              []string `plist:"files"`
	RunOnlyForDeploymentPostprocessing int      `plist:"runOnlyForDeploymentPostprocessing"`
}

type xcPBXResourcesBuildPhase struct {
	xcISA
	BuildActionMask                    int64    `plist:"buildActionMask"`
	Files                              []string `plist:"files"`
	RunOnlyForDeploymentPostprocessing int      `plist:"runOnlyForDeploymentPostprocessing"`
}

type xcPBXShellScriptBuildPhase struct {
	xcISA
	BuildActionMask                    int64    `plist:"buildActionMask"` // string??
	Files                              []string `plist:"files"`
	InputPaths                         []string `plist:"inputPaths"`
	RunOnlyForDeploymentPostprocessing int      `plist:"runOnlyForDeploymentPostprocessing"`
	ShellPath                          string   `plist:"shellPath"`
	ShellScript                        string   `plist:"shellScript"`
}

type xcPBXTargetDependency struct {
	xcISA
	Target      string `plist:"target"`
	TargetProxy string `plist:"targetProxy"`
}

type xcPBXVariantGroup struct {
	xcISA
	Children   []string `plist:"children"`
	Name       string   `plist:"name"`
	SourceTree string   `plist:"sourceTree"`
}

type xcPBXReferenceProxy struct {
	xcISA
	FileType   string `plist:"fileType"`
	Path       string `plist:"path"`
	RemoteRef  string `plist:"remoteRef"`
	SourceTree string `plist:"sourceTree"`
}

type xcPBXAppleScriptBuildPhase struct {
	xcISA
	BuildActionMask                    string   `plist:"buildActionMask"`
	ContextName                        string   `plist:"contextName"`
	Files                              []string `plist:"files"`
	IsSharedContext                    int      `plist:"isSharedContext"`
	RunOnlyForDeploymentPostprocessing int      `plist:"runOnlyForDeploymentPostprocessing"`
}

type xcXCConfigurationList struct {
	xcISA
	BuildConfigurations           []string `plist:"buildConfigurations"`
	DefaultConfigurationIsVisible int      `plist:"defaultConfigurationIsVisible"`
	DefaultConfigurationName      string   `plist:"defaultConfigurationName"`
}

type xcXCBuildConfiguration struct {
	xcISA
	BuildSettings xcPBXMap `plist:"buildSettings"`
	Name          string   `plist:"name"`
}

type xcProject struct {
	ArchiveVersion int      `plist:"archiveVersion"`
	ObjectVersion  int      `plist:"objectVersion"`
	Classes        xcPBXMap `plist:"classes"`
	Objects        xcPBXMap `plist:"objects"`
	RootObject     string   `plist:"rootObject"`
}
