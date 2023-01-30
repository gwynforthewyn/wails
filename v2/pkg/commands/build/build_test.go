package build

import (
	"github.com/wailsapp/wails/v2/internal/project"
	"github.com/wailsapp/wails/v2/pkg/clilogger"
	"os"
	"path/filepath"
	"testing"
)

func Test_execBuildApplicationDarwinUniversalFileLayout(t *testing.T) {
	var builder Builder

	var buildMode Mode = 1 // Testing in release mode

	outputFile := "execBuildApplicationTest"
	tempPath, err := os.MkdirTemp("", "")
	var outputLogger *clilogger.CLILogger

	if err != nil {
		panic(err)
	}

	projectData := &project.Project{
		Name:                   outputFile,
		AssetDirectory:         "",
		ReloadDirectories:      "",
		BuildCommand:           "npm run build",
		InstallCommand:         "npm install",
		DevCommand:             "",
		DevBuildCommand:        "",
		DevInstallCommand:      "",
		DevWatcherCommand:      "",
		FrontendDevServerURL:   "auto",
		WailsJSDir:             "frontend",
		Version:                "2",
		Path:                   tempPath,
		BuildDir:               "build",
		OutputFilename:         outputFile,
		OutputType:             "",
		Platform:               "",
		RunNonNativeBuildHooks: false,
		PostBuildHooks:         nil,
		PreBuildHooks:          nil,
		Author:                 project.Author{},
		Info:                   project.Info{},
		DebounceMS:             0,
		DevServer:              "",
		AppArgs:                "",
		NSISType:               "",
		Obfuscated:             false,
		GarbleArgs:             "",
		FrontendDir:            "",
		Bindings:               project.Bindings{},
	}

	options := &Options{
		LDFlags:           "",
		UserTags:          nil,
		Logger:            outputLogger,
		OutputType:        "desktop",
		Mode:              buildMode,
		ProjectData:       projectData,
		Pack:              true,
		Platform:          "",
		Arch:              "",
		Compiler:          "go",
		SkipModTidy:       false,
		IgnoreFrontend:    false,
		IgnoreApplication: false,
		OutputFile:        "",
		BinDirectory:      "",
		CleanBinDirectory: false,
		CompiledBinary:    "",
		KeepAssets:        false,
		Verbosity:         0,
		Compress:          false,
		CompressFlags:     "",
		WebView2Strategy:  "wv2runtime.download",
		RunDelve:          false,
		WailsJSDir:        "",
		ForceBuild:        false,
		BundleName:        "",
		TrimPath:          false,
		RaceDetector:      false,
		WindowsConsole:    false,
		Obfuscated:        false,
		GarbleArgs:        "-literals -tiny -seed=random",
		SkipBindings:      false,
	}

	builder = newDesktopBuilder(options)
	builder.SetProjectData(options.ProjectData)
	options.CompiledBinary = filepath.Join(options.BinDirectory, outputFile)

	execBuildApplication(builder, options)

}
