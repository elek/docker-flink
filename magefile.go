// +build mage

package main

import (
	"github.com/getlantern/errors"
	"io/ioutil"
	"fmt"
	"github.com/magefile/mage/sh"
	"net/http"
)

var name = "flink"
var downloadPath = "flink/flink-%s/flink-%s-bin-scala_2.11.tgz"

func versions() map[string][]string {
	versions := make(map[string][]string)
	versions["1.9.1"] = []string{"latest", "1.9.1","1.9","1"}
	return versions
}

func buildImage(tag string, dir string) {
	sh.Run("docker", "build", "-t", tag, dir)
}

func getApacheDownloadUrl(path string) (string, error) {
	url := fmt.Sprintf("https://www-eu.apache.org/dist/" + path)
	resp, err := http.Head(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 200 {
		return url, nil
	}
	url = fmt.Sprintf("https://archive.apache.org/dist/" + path)
	resp, err = http.Head(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 200 {
		return url, nil
	}
	return "", errors.New("Can't find download url for " + path)

}
func deployImage(tag string) {
	sh.Run("docker", "push", tag)
}

func buildContainer(version string, tags []string) error {
	url, err := getApacheDownloadUrl(fmt.Sprintf(downloadPath, version, version))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("url", []byte(url), 0644)
	if err != nil {
		return err
	}
	buildImage("flokkr/flink:build", ".")
	for _, tag := range tags {
		sh.Run("docker", "tag", "flokkr/flink:build", "flokkr/flink:"+tag)
		sh.Run("docker", "tag", "flokkr/flink:build", "quay.io/flokkr/flink:"+tag)
	}
	return nil
}

func Build() error {
	for version, tags := range versions() {
		err := buildContainer(version, tags)
		if err != nil {
			return err
		}
	}
	return nil
}

func Deploy() error {
	for _, tags := range versions() {
		for _, tag := range tags {
			deployImage("flokkr/flink:" + tag)
//			deployImage("quay.io/flokkr/flink:" + tag)
		}
	}
	return nil
}
