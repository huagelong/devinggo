package config

// ModuleConfig 模块配置结构 (.module.yaml)
type ModuleConfig struct {
	// 基本信息
	Name        string   `yaml:"name" json:"name"`               // 模块名称
	Version     string   `yaml:"version" json:"version"`         // 版本号
	Author      string   `yaml:"author" json:"author"`           // 作者
	License     string   `yaml:"license" json:"license"`         // 许可证
	Description string   `yaml:"description" json:"description"` // 描述
	Homepage    string   `yaml:"homepage" json:"homepage"`       // 主页
	GoVersion   string   `yaml:"goVersion" json:"goVersion"`     // Go版本要求
	Tags        []string `yaml:"tags" json:"tags"`               // 标签
	Keywords    []string `yaml:"keywords" json:"keywords"`       // 关键词

	// 依赖管理
	Dependencies map[string]string `yaml:"dependencies" json:"dependencies"` // Go模块依赖
	Modules      []string          `yaml:"modules" json:"modules"`           // DevingGo模块依赖

	// 文件管理
	Files ModuleFiles `yaml:"files" json:"files"` // 文件配置

	// 配置合并
	ConfigMerge ConfigMerge `yaml:"configMerge" json:"configMerge"` // 配置文件合并

	// 静态资源部署
	StaticDeploy StaticDeploy `yaml:"staticDeploy" json:"staticDeploy"` // 静态资源部署

	// 生命周期钩子
	Hooks ModuleHooks `yaml:"hooks" json:"hooks"` // 生命周期钩子

	// 模板变量
	Variables map[string]string `yaml:"variables" json:"variables"` // 模板变量定义

	// 安全配置
	Security SecurityConfig `yaml:"security" json:"security"` // 安全配置
}

// ModuleFiles 文件配置
type ModuleFiles struct {
	Go      []string `yaml:"go" json:"go"`           // Go源码文件
	SQL     []string `yaml:"sql" json:"sql"`         // SQL迁移文件
	Static  []string `yaml:"static" json:"static"`   // 静态资源文件
	Config  []string `yaml:"config" json:"config"`   // 配置文件
	Other   []string `yaml:"other" json:"other"`     // 其他文件
	Exclude []string `yaml:"exclude" json:"exclude"` // 排除文件
}

// ConfigMerge 配置文件合并设置
type ConfigMerge struct {
	Enabled bool              `yaml:"enabled" json:"enabled"` // 是否启用
	Files   []ConfigMergeFile `yaml:"files" json:"files"`     // 需要合并的配置文件
}

// ConfigMergeFile 单个配置文件合并设置
type ConfigMergeFile struct {
	Source    string            `yaml:"source" json:"source"`       // 源文件（模块内）
	Target    string            `yaml:"target" json:"target"`       // 目标文件（项目内）
	Strategy  string            `yaml:"strategy" json:"strategy"`   // 合并策略: merge, replace, skip
	Keys      []string          `yaml:"keys" json:"keys"`           // 需要合并的键（仅merge策略）
	Variables map[string]string `yaml:"variables" json:"variables"` // 变量替换
}

// StaticDeploy 静态资源部署配置
type StaticDeploy struct {
	Enabled bool               `yaml:"enabled" json:"enabled"` // 是否启用
	Rules   []StaticDeployRule `yaml:"rules" json:"rules"`     // 部署规则
}

// StaticDeployRule 单条部署规则
type StaticDeployRule struct {
	Source    string `yaml:"source" json:"source"`       // 源路径（模块内）
	Target    string `yaml:"target" json:"target"`       // 目标路径（项目内）
	Method    string `yaml:"method" json:"method"`       // 部署方式: copy, symlink, merge
	Overwrite bool   `yaml:"overwrite" json:"overwrite"` // 是否覆盖已存在文件
}

// ModuleHooks 生命周期钩子
type ModuleHooks struct {
	PreInstall    []HookCommand `yaml:"preInstall" json:"preInstall"`       // 安装前
	PostInstall   []HookCommand `yaml:"postInstall" json:"postInstall"`     // 安装后
	PreUninstall  []HookCommand `yaml:"preUninstall" json:"preUninstall"`   // 卸载前
	PostUninstall []HookCommand `yaml:"postUninstall" json:"postUninstall"` // 卸载后
	PreUpgrade    []HookCommand `yaml:"preUpgrade" json:"preUpgrade"`       // 升级前
	PostUpgrade   []HookCommand `yaml:"postUpgrade" json:"postUpgrade"`     // 升级后
}

// HookCommand 钩子命令
type HookCommand struct {
	Name        string            `yaml:"name" json:"name"`               // 钩子名称
	Command     string            `yaml:"command" json:"command"`         // 执行命令
	WorkDir     string            `yaml:"workDir" json:"workDir"`         // 工作目录
	Env         map[string]string `yaml:"env" json:"env"`                 // 环境变量
	IgnoreError bool              `yaml:"ignoreError" json:"ignoreError"` // 忽略错误
}

// SecurityConfig 安全配置
type SecurityConfig struct {
	Signature      SignatureConfig   `yaml:"signature" json:"signature"`           // 数字签名
	Permissions    PermissionsConfig `yaml:"permissions" json:"permissions"`       // 权限要求
	SensitiveFiles []string          `yaml:"sensitiveFiles" json:"sensitiveFiles"` // 敏感文件
}

// SignatureConfig 数字签名配置
type SignatureConfig struct {
	Enabled   bool   `yaml:"enabled" json:"enabled"`     // 是否启用
	PublicKey string `yaml:"publicKey" json:"publicKey"` // 公钥
	Algorithm string `yaml:"algorithm" json:"algorithm"` // 算法: RSA, ECDSA
}

// PermissionsConfig 权限要求配置
type PermissionsConfig struct {
	FileSystem bool `yaml:"fileSystem" json:"fileSystem"` // 文件系统访问
	Network    bool `yaml:"network" json:"network"`       // 网络访问
	Database   bool `yaml:"database" json:"database"`     // 数据库访问
	Admin      bool `yaml:"admin" json:"admin"`           // 管理员权限
}

// ModuleMetadata 模块元数据（用于module.json向后兼容）
type ModuleMetadata struct {
	Name      string              `json:"name"`
	Author    string              `json:"author"`
	Version   string              `json:"version"`
	License   string              `json:"license"`
	GoVersion string              `json:"goVersion"`
	Mod       map[string]string   `json:"mod"`
	Files     map[string][]string `json:"files"`
}

// DefaultModuleConfig 返回默认模块配置
func DefaultModuleConfig(moduleName string) *ModuleConfig {
	return &ModuleConfig{
		Name:         moduleName,
		Version:      "1.0.0",
		Author:       "devinggo",
		License:      "MIT",
		Description:  moduleName + " module",
		GoVersion:    "1.23+",
		Tags:         []string{},
		Keywords:     []string{},
		Dependencies: make(map[string]string),
		Modules:      []string{},
		Files: ModuleFiles{
			Go:      []string{},
			SQL:     []string{},
			Static:  []string{},
			Config:  []string{},
			Other:   []string{},
			Exclude: []string{},
		},
		ConfigMerge: ConfigMerge{
			Enabled: false,
			Files:   []ConfigMergeFile{},
		},
		StaticDeploy: StaticDeploy{
			Enabled: false,
			Rules:   []StaticDeployRule{},
		},
		Hooks: ModuleHooks{
			PreInstall:    []HookCommand{},
			PostInstall:   []HookCommand{},
			PreUninstall:  []HookCommand{},
			PostUninstall: []HookCommand{},
			PreUpgrade:    []HookCommand{},
			PostUpgrade:   []HookCommand{},
		},
		Variables: make(map[string]string),
		Security: SecurityConfig{
			Signature: SignatureConfig{
				Enabled:   false,
				Algorithm: "RSA",
			},
			Permissions: PermissionsConfig{
				FileSystem: true,
				Database:   true,
			},
			SensitiveFiles: []string{},
		},
	}
}
