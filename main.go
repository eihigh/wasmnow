package main

import (
	_ "embed"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	defaultAddr = "localhost:8080"
)

//go:embed basic-index.html
var basicIndexHTML []byte

func main() {
	log.SetPrefix("")
	log.SetFlags(0)
	buildOnly := flag.Bool("b", false, "build without starting the server")
	dir := flag.String("d", "wasmnow", "destination directory")
	addr := flag.String("http", defaultAddr, "HTTP service address")
	tags := flag.String("tags", "", "build tags")
	flag.Parse()

	if flag.NArg() > 0 {
		log.Print("unexpected arguments:", flag.Args())
		flag.Usage()
	}

	if err := build(*dir, *tags); err != nil {
		log.Fatalf("build failed: %v", err)
	}
	if *buildOnly {
		return
	}
	if err := serve(*dir, *addr); err != nil {
		log.Fatalf("serve failed: %v", err)
	}
}

func build(dir, tags string) error {
	// create dir if it doesn't exist
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}
		log.Print(welcomeToWasmnow.get())
		log.Printf(createdANewFolder.get(), dir)
		log.Print(noteUseTheDOptionToChangeTheDestinationDirectory.get())
		log.Print(forAnyExternalFilesYourProgramNeeds.get())
		log.Print(useGosEmbedPackageToIncludeThemRecommended.get())
		log.Printf(simplyCopyThemIntoTheDirectory.get(), dir)
		if runtime.GOOS == "windows" {
			log.Printf(inWindowsToCreateAZipFile.get(), dir)
		} else if runtime.GOOS == "darwin" {
			log.Printf(inMacOSToCreateAZipFile.get(), dir, filepath.Base(dir))
		}
	}

	// create basic index.html if it doesn't exist
	path := filepath.Join(dir, "index.html")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.WriteFile(path, basicIndexHTML, 0666); err != nil {
			return err
		}
		log.Printf(createdABasicWebpageIndexHtml.get(), dir)
	}

	// copy wasm_exec.js
	wasmExecJSPath, err := libWasmExecJSPath()
	if err != nil {
		return fmt.Errorf("failed to get wasm_exec.js path: %w", err)
	}
	src, err := os.Open(wasmExecJSPath)
	if err != nil {
		return fmt.Errorf("failed to open wasm_exec.js: %w", err)
	}
	defer src.Close()
	dst, err := os.Create(filepath.Join(dir, "wasm_exec.js"))
	if err != nil {
		return fmt.Errorf("failed to create wasm_exec.js: %w", err)
	}
	defer dst.Close()
	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Errorf("failed to copy wasm_exec.js: %w", err)
	}
	log.Printf(automaticallyCopiedARequiredFileWasmExecJs.get(), dir, wasmExecJSPath)

	// build main.wasm
	cmd := exec.Command("go", "build", "-o", filepath.Join(dir, "main.wasm"), "-tags", tags)
	cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm", "GOTOOLCHAIN=go1.24rc1+auto")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go build: %w", err)
	}
	log.Printf(builtMainWasm.get(), dir)

	return nil
}

func serve(dir, addr string) error {
	// serve directory
	log.Print(yourProgramIsReadyToTryOut.get())
	log.Printf(pleaseOpenHttpLocalhostInYourWebBrowser.get(), "http://"+addr)
	log.Print(noteUseTheBOptionToBuildWithoutSettingUpTheSpecialPage.get())
	if err := http.ListenAndServe(addr, http.FileServer(http.Dir(dir))); err != nil {
		return fmt.Errorf("http.ListenAndServe: %w", err)
	}
	return nil
}

func libWasmExecJSPath() (string, error) {
	// Since go1.24, the wasm-related files have been moved from GOROOT/misc to GOROOT/lib.
	// Unlike misc, the lib/ directory is included even in the automatically downloaded toolchain,
	// so set GOTOOLCHAIN to go1.24rc1+auto to get it from lib.
	cmd := exec.Command("go", "env", "GOROOT")
	cmd.Env = append(os.Environ(), "GOTOOLCHAIN=go1.24rc1+auto")
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return filepath.Join(strings.TrimSpace(string(out)), "lib", "wasm", "wasm_exec.js"), nil
}
