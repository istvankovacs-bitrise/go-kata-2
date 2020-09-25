package main

import (
	"fmt"
	"os/exec"
	"path"
	"path/filepath"
	"io/ioutil"
	"os"
	"log"

	"gopkg.in/yaml.v2"
)

func main() {
	repo := "steps-ios-auto-provision"
	repoUrl := "https://github.com/bitrise-steplib/" + repo + ".git"
	fileName := "/bitrise.yml"

	cmd := exec.Command("git", "clone", repoUrl)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	ymlPath, _ := filepath.Abs(path.Join(repo + fileName))
    yamlFile, err := ioutil.ReadFile(ymlPath)

    if err != nil {
        log.Fatal(err)
	}
	
	var bitriseYml BitriseYml

    if err := bitriseYml.parse(yamlFile); err != nil {
		log.Fatal(err)
	}

	steps := bitriseYml.Workflows.Common.Steps

	scripts := []Steps{}
	result := []Steps{}

	for i := 0; i < len(steps); i++ {
		if steps[i].Script != nil {
			scripts = append(scripts, steps[len(scripts)])

			if len(scripts) == 2 {
				continue
			}
		}
		result = append(result, steps[i])
	}

	bitriseYml.Workflows.Common.Steps = result

	err = writeToFile(bitriseYml)

	if err != nil {
		log.Fatal(err)
	}
}

func (bitriseYml *BitriseYml) parse(data []byte) error {
	if err := yaml.Unmarshal(data, bitriseYml); err != nil {
		return err
	}
	 
	return nil
}

func (bitriseYml *BitriseYml) serialize() []byte {
	out, err := yaml.Marshal(&bitriseYml)
    if err != nil {
        log.Fatal(err)
    }

	return out
}

func writeToFile(bitriseYml BitriseYml) error {
	os.Mkdir("temp", 0777)
	tempPath, _ := filepath.Abs("temp/modified_bitrise.yml")

	out, err := os.Create(tempPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.Write(bitriseYml.serialize())

	if err != nil {
		return err
	}

	fmt.Println(tempPath)

	return nil
}


type BitriseYml struct {
	FormatVersion        int       `yaml:"format_version"`
	DefaultStepLibSource string    `yaml:"default_step_lib_source"`
	App                  App       `yaml:"app"`
	Workflows            Workflows `yaml:"workflows"`
}
type Envs struct {
	BITRISESTEPID                string `yaml:"BITRISE_STEP_ID,omitempty"`
	BITRISESTEPGITCLONEURL       string `yaml:"BITRISE_STEP_GIT_CLONE_URL,omitempty"`
	MYSTEPLIBREPOFORKGITURL      string `yaml:"MY_STEPLIB_REPO_FORK_GIT_URL,omitempty"`
	BITRISEBUILDURL              string `yaml:"BITRISE_BUILD_URL,omitempty"`
	BITRISEBUILDAPITOKEN         string `yaml:"BITRISE_BUILD_API_TOKEN,omitempty"`
	BITRISECERTIFICATEURL        string `yaml:"BITRISE_CERTIFICATE_URL,omitempty"`
	BITRISECERTIFICATEPASSPHRASE string `yaml:"BITRISE_CERTIFICATE_PASSPHRASE,omitempty"`
	TEAMID                       string `yaml:"TEAM_ID,omitempty"`
	BITRISEKEYCHAINPASSWORD      string `yaml:"BITRISE_KEYCHAIN_PASSWORD,omitempty"`
	APITOKEN                     string `yaml:"API_TOKEN,omitempty"`
	APPSLUG                      string `yaml:"APP_SLUG,omitempty"`
	ORIGBITRISESOURCEDIR         string `yaml:"ORIG_BITRISE_SOURCE_DIR,omitempty"`
	SAMPLEAPPURL                 string `yaml:"SAMPLE_APP_URL,omitempty"`
	BRANCH                       string `yaml:"BRANCH,omitempty"`
	BITRISEPROJECTPATH           string `yaml:"BITRISE_PROJECT_PATH,omitempty"`
	BITRISESCHEME                string `yaml:"BITRISE_SCHEME,omitempty"`
	BITRISECONFIGURATION         string `yaml:"BITRISE_CONFIGURATION,omitempty"`
	DISTRIBUTIONTYPE             string `yaml:"DISTRIBUTION_TYPE,omitempty"`
	GENERATEPROFILES             string `yaml:"GENERATE_PROFILES,omitempty"`
	INSTALLPODS          		 string `yaml:"INSTALL_PODS,omitempty"`
	BITRISESTEPVERSION      	 string `yaml:"BITRISE_STEP_VERSION,omitempty"`
}
type App struct {
	Envs []Envs `yaml:"envs"`
}
type Script struct {
	Title  string   `yaml:"title"`
	Inputs []Inputs `yaml:"inputs"`
}
type TriggerBitriseWorkflow struct {
	RunIf  string   `yaml:"run_if"`
	Inputs []Inputs `yaml:"inputs"`
}
type Steps struct {
	Script                 *Script                 `yaml:"script,omitempty"`
	TriggerBitriseWorkflow TriggerBitriseWorkflow `yaml:"trigger-bitrise-workflow,omitempty"`
	ChangeWorkdir     ChangeWorkdir     `yaml:"change-workdir,omitempty"`
	CocoapodsInstall  CocoapodsInstall  `yaml:"cocoapods-install,omitempty"`
	Path              Path              `yaml:"path::./,omitempty"`
	XcodeArchive      XcodeArchive      `yaml:"xcode-archive,omitempty"`
	DeployToBitriseIo DeployToBitriseIo `yaml:"deploy-to-bitrise-io,omitempty"`
}
type Test struct {
	BeforeRun []string `yaml:"before_run"`
	Steps     []Steps  `yaml:"steps"`
	AfterRun  []string `yaml:"after_run"`
}
type TestNewCertificates struct {
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Envs        []Envs   `yaml:"envs"`
	AfterRun    []string `yaml:"after_run"`
}
type TestBundleID struct {
	Envs     []Envs   `yaml:"envs"`
	AfterRun []string `yaml:"after_run"`
}
type TestXcodeManaged struct {
	Envs     []Envs   `yaml:"envs"`
	AfterRun []string `yaml:"after_run"`
}
type TestXcodeManagedGenerateEnabled struct {
	Envs     []Envs   `yaml:"envs"`
	AfterRun []string `yaml:"after_run"`
}
type TestEntitlements struct {
	Envs     []Envs   `yaml:"envs"`
	AfterRun []string `yaml:"after_run"`
}
type TestWorkspace struct {
	Envs     []Envs   `yaml:"envs"`
	AfterRun []string `yaml:"after_run"`
}
type TestTvos struct {
	Envs     []Envs   `yaml:"envs"`
	AfterRun []string `yaml:"after_run"`
}
type TestTvosDevelopment struct {
	Envs     []Envs   `yaml:"envs"`
	AfterRun []string `yaml:"after_run"`
}
type TestTvosManaged struct {
	Envs     []Envs   `yaml:"envs"`
	AfterRun []string `yaml:"after_run"`
}
type TestTvosDevelopmentManaged struct {
	Envs     []Envs   `yaml:"envs"`
	AfterRun []string `yaml:"after_run"`
}
type ChangeWorkdir struct {
	Title  string   `yaml:"title"`
	Inputs []Inputs `yaml:"inputs"`
}
type CocoapodsInstall struct {
	RunIf string `yaml:"run_if"`
	Title string `yaml:"title"`
}
type Inputs struct {
	BuildURL         string `yaml:"build_url,omitempty"`
	BuildAPIToken    string `yaml:"build_api_token,omitempty"`
	CertificateUrls  string `yaml:"certificate_urls,omitempty"`
	Passphrases      string `yaml:"passphrases,omitempty"`
	TeamID           string `yaml:"team_id,omitempty"`
	DistributionType string `yaml:"distribution_type,omitempty"`
	ProjectPath      string `yaml:"project_path,omitempty"`
	Scheme           string `yaml:"scheme,omitempty"`
	Configuration    string `yaml:"configuration,omitempty"`
	VerboseLog       string `yaml:"verbose_log,omitempty"`
	GenerateProfiles string `yaml:"generate_profiles,omitempty"`
	ExportMethod  string `yaml:"export_method,omitempty"`
	OutputTool    string `yaml:"output_tool,omitempty"`
	NotifyUserGroups string `yaml:"notify_user_groups,omitempty"`
	Path         string `yaml:"path,omitempty"`
	IsCreatePath string `yaml:"is_create_path,omitempty"`
	Content string `yaml:"content,omitempty"`
	APIToken   string `yaml:"api_token,omitempty"`
	AppSlug    string `yaml:"app_slug,omitempty"`
	WorkflowID string `yaml:"workflow_id,omitempty"`
}
type Path struct {
	Title  string   `yaml:"title"`
	RunIf  bool     `yaml:"run_if"`
	Inputs []Inputs `yaml:"inputs"`
}
type XcodeArchive struct {
	Title  string   `yaml:"title"`
	Inputs []Inputs `yaml:"inputs"`
}
type DeployToBitriseIo struct {
	Inputs []Inputs `yaml:"inputs"`
}
type Common struct {
	Steps []Steps `yaml:"steps"`
}
type CreateRelease struct {
	Steps []Steps `yaml:"steps"`
}
type AuditThisStep struct {
	Steps []Steps `yaml:"steps"`
}
type ShareThisStep struct {
	Envs        []Envs   `yaml:"envs"`
	Description string   `yaml:"description"`
	BeforeRun   []string `yaml:"before_run"`
	Steps       []Steps  `yaml:"steps"`
}
type Workflows struct {
	Test                            Test                            `yaml:"test"`
	TestNewCertificates             TestNewCertificates             `yaml:"test_new_certificates"`
	TestBundleID                    TestBundleID                    `yaml:"test_bundle_id"`
	TestXcodeManaged                TestXcodeManaged                `yaml:"test_xcode_managed"`
	TestXcodeManagedGenerateEnabled TestXcodeManagedGenerateEnabled `yaml:"test_xcode_managed_generate_enabled"`
	TestEntitlements                TestEntitlements                `yaml:"test_entitlements"`
	TestWorkspace                   TestWorkspace                   `yaml:"test_workspace"`
	TestTvos                        TestTvos                        `yaml:"test_tvos"`
	TestTvosDevelopment             TestTvosDevelopment             `yaml:"test_tvos_development"`
	TestTvosManaged                 TestTvosManaged                 `yaml:"test_tvos_managed"`
	TestTvosDevelopmentManaged      TestTvosDevelopmentManaged      `yaml:"test_tvos_development_managed"`
	Common                          Common                          `yaml:"_common"`
	CreateRelease                   CreateRelease                   `yaml:"create-release"`
	AuditThisStep                   AuditThisStep                   `yaml:"audit-this-step"`
	ShareThisStep                   ShareThisStep                   `yaml:"share-this-step"`
}