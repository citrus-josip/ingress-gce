/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package app

import "os"

const (
	HostEnvVar      = "node"
	PodEnvVar       = "pod"
	NamespaceEnvVar = "namespace"
)

type Env struct {
	Node      string
	Pod       string
	Namespace string
}

var (
	E = Env{}
)

func init() {
	E.Node = os.Getenv(HostEnvVar)
	E.Pod = os.Getenv(PodEnvVar)
	E.Namespace = os.Getenv(NamespaceEnvVar)
}
