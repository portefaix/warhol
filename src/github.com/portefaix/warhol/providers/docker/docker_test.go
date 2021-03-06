// Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package docker

import (
	"testing"

	"github.com/fsouza/go-dockerclient"

	"github.com/portefaix/warhol/pubsub"
)

func Test_NewBuilderWithoutDockerServer(t *testing.T) {
	_, err := NewBuilder(
		"unix:///var/run/docker.sock",
		false,
		"/tmp",
		"http://127.0.0.1:9999",
		&Authentication{
			Username: "foo",
			Password: "bar",
			Email:    "foo@bar.com",
		},
		&pubsub.Config{Type: pubsub.ZEROMQ})
	if err == nil {
		t.Fatalf("Invalid Docker builder.")
	}
}

func Test_DockerImageNameWithDefaultRegistry(t *testing.T) {
	url := REGISTRY
	client, _ := docker.NewClient(url)
	db := &Builder{
		Client:      client,
		RegistryURL: url,
		AuthConfig: docker.AuthConfiguration{
			Username:      "foo",
			Password:      "bar",
			Email:         "foo@bar.com",
			ServerAddress: url,
		},
		BuildChan: make(chan *Project),
		PushChan:  make(chan *Project),
	}
	name := db.GetImageName("foo")
	if name != "warhol/foo" {
		t.Fatalf("Invalid image name : %s", name)
	}

}

func Test_DockerImageNameWithPrivateRegistry(t *testing.T) {
	url := "registry.warhol.com"
	client, _ := docker.NewClient(url)
	db := &Builder{
		Client:      client,
		RegistryURL: url,
		AuthConfig: docker.AuthConfiguration{
			Username:      "foo",
			Password:      "bar",
			Email:         "foo@bar.com",
			ServerAddress: url,
		},
		BuildChan: make(chan *Project),
		PushChan:  make(chan *Project),
	}
	name := db.GetImageName("foo")
	if name != "registry.warhol.com/warhol/foo" {
		t.Fatalf("Invalid image name : %s", name)
	}
}
