/*
Copyright 2019 The Kubernetes Authors.

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

package main

import (
	"os"
	"os/exec"
	"time"

	"k8s.io/ingress-nginx/internal/nginx"
	"k8s.io/klog/v2"
)

func main() {
	err := exec.Command("bash", "-c", "sleep 5 && pkill -SIGTERM -f nginx-ingress-controller").Run()
	if err != nil {
		klog.ErrorS(err, "terminating ingress controller")
		os.Exit(1)
	}

	// wait for the NGINX process to terminate
	timer := time.NewTicker(time.Second * 1)
	for range timer.C {
		if !nginx.IsRunning() {
			timer.Stop()
			break
		}
	}
}
