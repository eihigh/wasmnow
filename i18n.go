package main

import (
	"strings"

	"github.com/cloudfoundry-attic/jibber_jabber"
)

var lang string

func init() {
	t, err := jibber_jabber.DetectTerritory()
	if err != nil {
		lang = ""
	} else {
		lang = t
	}
}

type msgs map[string]string

func (m msgs) get() string {
	if msg, ok := m[strings.ToUpper(lang)]; ok {
		return msg
	}
	return m[""]
}

var (
	welcomeToWasmnow = msgs{
		"":   "😆 Welcome to wasmnow!",
		"JP": "😆 wasmnowへようこそ！",
	}
	wasmnowIsASimpleTool = msgs{
		"":   "🔧 wasmnow is a simple tool that helps you create Go programs that run in your web browser.",
		"JP": "🔧 wasmnowは、ブラウザで動くGoプログラムを簡単に作るためのツールです。",
	}
	createdANewFolderCalledWasmnow = msgs{
		"":   "📁 Created a new folder called 'wasmnow'. This is where all your program files will live.",
		"JP": "📁 'wasmnow'という新しいフォルダを作成しました。必要なファイルはこのフォルダの中に入れてください。",
	}
	forAnyExternalFilesYourProgramNeeds = msgs{
		"":   "📝 For any external files your program needs, you have two options:",
		"JP": "📝 プログラムで外部ファイルを読み込むには以下どちらかの方法を使います：",
	}
	useGosEmbedPackageToIncludeThemRecommended = msgs{
		"":   "  - Use Go's embed package to include them (recommended)",
		"JP": "  - embedパッケージを使って埋め込む（推奨）",
	}
	simplyCopyThemIntoTheWasmnowDirectory = msgs{
		"":   "  - Simply copy them into the wasmnow directory",
		"JP": "  - 単純にwasmnowフォルダの中にコピーする",
	}
	createdABasicWebpageWasmnowIndexHtml = msgs{
		"":   "📄 Created a basic webpage (wasmnow/index.html) that you can customize as you like.",
		"JP": "📄 基本的な内容のwasmnow/index.htmlを作成しました。このファイルは自由に編集できます。",
	}
	automaticallyCopiedARequiredFileWasmnowWasmExecJs = msgs{
		"":   "📄 Automatically copied a required file (wasmnow/wasm_exec.js) from %s",
		"JP": "📄 プログラムの実行に必要なwasmnow/wasm_exec.jsを %s から自動的にコピーしました。",
	}
	builtWasmnowMainWasm = msgs{
		"":   "🔧 Built wasmnow/main.wasm",
		"JP": "🔧 wasmnow/main.wasmをビルドしました。",
	}
	inWindowsToCreateAZipFile = msgs{
		"":   "📝 In Windows, to create a zip file, Right-click the wasmnow folder, go to 'Send to', and choose 'Compressed (zipped) folder'",
		"JP": "📝 Windowsの場合、zipファイルを作るには、wasmnowフォルダを右クリックして「送る」から「圧縮 (zip形式) フォルダー」を選んでください。",
	}
	inMacOSToCreateAZipFile = msgs{
		"":   "📝 In macOS, to create a zip file, Right-click the wasmnow folder and choose 'Compress 'wasmnow''",
		"JP": "📝 macOSの場合、zipファイルを作るには、wasmnowフォルダを右クリックして「\"wasmnow\" を圧縮」を選んでください。",
	}
	yourProgramIsReadyToTryOut = msgs{
		"":   "👀 Your program is ready to try out!",
		"JP": "👀 プログラムの動作確認ができる特別なページを準備しました。",
	}
	pleaseOpenHttpLocalhostInYourWebBrowser = msgs{
		"":   "🌏 Please open %s in your web browser. This is a special page that runs on your computer.",
		"JP": "🌏 ブラウザで %s を開いてください。",
	}
	noteUseTheBOptionToBuildWithoutSettingUpTheSpecialPage = msgs{
		"":   "📝 Note: Use the -b option to build without setting up the special page.",
		"JP": "📝 ※ビルドだけ行って特別なページは必要ない場合は、-b オプションを使ってください。",
	}
)
