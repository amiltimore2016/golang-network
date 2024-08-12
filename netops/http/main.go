package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"time"
)

type Build struct {
	Arch     string `json:"arch"`
	Filename string `json:"filename"`
	Name     string `json:"name"`
	Os       string `json:"os"`
	Url      string `json:"url"`
	Version  string `json:"version"`
}

type Releases struct {
	Releases map[string]Release `json:"versions"`
}

type Release struct {
	Builds            []Build  `json:"builds"`
	Name              string   `json:"name"`
	Shasums           string   `json:"shasums"`
	ShasumSignature   string   `json:"shasums_signature"`
	Version           string   `json:"version"`
	ShashumSignatures []string `json:"shasums_signatures"`
}

type Versions struct {
	Versions string `json:"versions"`
}

func main() {
	var releaseData Releases
	req, _ := http.NewRequest("GET", "https://releases.hashicorp.com/terraform/index.json", nil)
	req.Header.Set("Cache-Control", "no-cache")
	client := &http.Client{Timeout: time.Second * 10}
	resp, _ := client.Do(req)
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	if err := json.Unmarshal(body, &releaseData); err != nil {
		fmt.Println(err)
	}

	versionList := make([]string, len(releaseData.Releases))

	for k := range releaseData.Releases {
		fmt.Println(k)
		versionList = append(versionList, k)
	}
	sort.Strings(versionList)
	fmt.Println(versionList)
	// for range releaseData{
	// 	releaseList[releases.Releases["0.1.0"].Version] = releases.Releases["0.1.0"]
	// }
	// fmt.Printf("%v\n", releases.Releases["0.1.0"].Builds[0].Url)
	// for _, release := range releases.Releases {
	// 	fmt.Printf("%s\n", release.Version)
	// }

	// fmt.Printf("%v\n", releases.Releases)
}
