/*
Copyright 2017 The Kubernetes Authors.

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

package componentconfig

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	componentbaseconfig "k8s.io/component-base/config"
	registry "k8s.io/component-base/logs/api/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DeschedulerConfiguration struct {
	metav1.TypeMeta

	// Time interval for descheduler to run
	DeschedulingInterval time.Duration

	// KubeconfigFile is path to kubeconfig file with authorization and master
	// location information.
	KubeconfigFile string

	// PolicyConfigFile is the filepath to the descheduler policy configuration.
	PolicyConfigFile string

	// Dry run
	DryRun bool

	// Node selectors
	NodeSelector string

	// MaxNoOfPodsToEvictPerNode restricts maximum of pods to be evicted per node.
	MaxNoOfPodsToEvictPerNode int

	// EvictLocalStoragePods allows pods using local storage to be evicted.
	EvictLocalStoragePods bool

	// IgnorePVCPods sets whether PVC pods should be allowed to be evicted
	IgnorePVCPods bool

	// LeaderElection starts Deployment using leader election loop
	LeaderElection componentbaseconfig.LeaderElectionConfiguration

	// Logging specifies the options of logging.
	// Refer [Logs Options](https://github.com/kubernetes/component-base/blob/master/logs/api/v1/options.go) for more information.
	Logging registry.LoggingConfiguration
}
