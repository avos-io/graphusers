package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const (
	zipName    = "linux-x64.zip"
	zipBinary  = "kiota"
	binaryName = "kiota"

	kiotaUrl = "https://api.github.com/repos/microsoft/kiota/releases/latest"
)

var (
	latest   bool
	generate bool
	apply    bool
	info     bool
)

type kiotaAsset struct {
	Name        string `json:"name"`
	DownloadUrl string `json:"browser_download_url"`
}

type kiotaResp struct {
	Name   string       `json:"name"`
	Assets []kiotaAsset `json:"assets"`
}

func readAndWriteZip(zf *zip.File) error {
	fi, err := zf.Open()
	if err != nil {
		return err
	}
	defer fi.Close()

	b, err := io.ReadAll(fi)
	if err != nil {
		return err
	}

	fo, err := os.Create(binaryName)
	if err != nil {
		return err
	}
	defer func() {
		if err := fo.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err = os.Chmod(fo.Name(), 0755); err != nil {
		return err
	}

	_, err = fo.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func latestKiota() {
	log.Printf("Download latest kiota")
	r, err := http.Get(kiotaUrl)
	if err != nil {
		log.Fatal(err)
	}

	var resp kiotaResp
	if err = json.NewDecoder(r.Body).Decode(&resp); err != nil {
		log.Fatal(err)
	}
	r.Body.Close()

	log.Printf("Using kiota %s", resp.Name)

	var asset *kiotaAsset
	for _, a := range resp.Assets {
		if a.Name == zipName {
			asset = &a
			break
		}
	}

	if asset == nil {
		log.Fatal("Failed to find correct asset")
	}

	z, err := http.Get(asset.DownloadUrl)
	if err != nil {
		log.Fatal(err)
	}

	zipBody, err := io.ReadAll(z.Body)
	if err != nil {
		log.Fatal(err)
	}
	z.Body.Close()

	reader, err := zip.NewReader(bytes.NewReader(zipBody), int64(len(zipBody)))
	if err != nil {
		log.Fatal(err)
	}

	for _, zf := range reader.File {
		log.Printf("Reading file: %s", zf.Name)
		if zf.Name != zipBinary {
			continue
		}

		log.Printf("Writing file: %s", zf.Name)
		if err = readAndWriteZip(zf); err != nil {
			log.Fatal(err)
		}
		return
	}
}

func runKiota(name string) {
	fi, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	b, err := io.ReadAll(fi)
	if err != nil {
		log.Fatal(err)
	}

	cmd := strings.Split(strings.Join(strings.Split(string(b), "\n"), " "), " ")

	out, err := exec.Command("./kiota", cmd...).Output()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s\n", out)
}

func main() {
	flag.BoolVar(&latest, "latest", false, "Download latest kiota binary")
	flag.BoolVar(&generate, "generate", false, "Generate SDK")
	flag.BoolVar(&apply, "apply", false, "Apply generated SDK changes")
	flag.BoolVar(&info, "info", false, "List go dependencies needed")
	flag.Parse()

	if latest {
		latestKiota()
	}

	if generate {
		log.Printf("Generate SDK")
		runKiota("cmd/generate.txt")
	}

	if apply {
		log.Printf("Apply SDK changes")
		_, err := exec.Command("/bin/sh", "-c", "rm -rf model/ users/ groups/ && cp -r output/* ./").
			Output()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Success")
	}

	if info {
		log.Printf("List go dependencies")
		runKiota("cmd/info.txt")
	}
}
