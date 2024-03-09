package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
)

type Request struct {
	Color string `json:"color"`
}

func main() {
	host := flag.String("host", "", "hostname/ip of the awtrix light")

	sourceName := flag.String("source", "@DEFAULT_SOURCE@", "pulseaudio source, default if empty")
	color := flag.String("color", "#FF0000", "hex code of indicator color, red if empty")
	indicator := flag.Int("indicator", 1, "position of indicator, 1 for top, 2 for center, 3 for bottom")
	onlyCheck := flag.Bool("onlyCheck", false, "don't toggle mute, only check and send to awtrix")

	flag.Parse()
	if host == nil || *host == "" {
		panic("missing host")
	}

	if !*onlyCheck {
		_, err := exec.Command("/usr/bin/pactl", "set-source-mute", *sourceName, "toggle").Output()
		if err != nil {
			panic("error muting pa source")
		}
	}
	out, err := exec.Command("/usr/bin/pactl", "get-source-mute", *sourceName).Output()
	if err != nil {
		fmt.Println(err.Error())
		panic("error getting pa source info")
	}

	payload := "0"
	if strings.Contains(string(out), "yes") {
		payload = *color
	}

	body, err := json.Marshal(&Request{Color: payload})
	if err != nil {
		panic("error building body")
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("http://%s/api/indicator%d", *host, *indicator), bytes.NewBuffer(body))
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		panic("error calling awtrix")
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		panic("error parsing response")
	}

	if string(resBody) != "OK" {
		panic("response not ok")
	}
}
