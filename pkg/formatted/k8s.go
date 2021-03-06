// Copyright © 2019 The Tekton Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package formatted

import (
	"github.com/fatih/color"
	corev1 "k8s.io/api/core/v1"
	"knative.dev/pkg/apis/duck/v1beta1"
)

var ConditionColor = map[string]color.Attribute{
	"Failed":    color.FgHiRed,
	"Succeeded": color.FgHiGreen,
	"Running":   color.FgHiBlue,
	"Cancelled": color.FgHiMagenta,
}

func colorStatus(status string) string {
	return color.New(ConditionColor[status]).Sprint(status)
}

// Condition returns a human readable text based on the status of the Condition
func Condition(c v1beta1.Conditions) string {
	var status string
	if len(c) == 0 {
		return "---"
	}

	switch c[0].Status {
	case corev1.ConditionFalse:
		status = "Failed"
	case corev1.ConditionTrue:
		status = "Succeeded"
	case corev1.ConditionUnknown:
		status = "Running"
	}
	cstatus := colorStatus(status)

	if c[0].Reason != "" && c[0].Reason != status {

		if c[0].Reason == "PipelineRunCancelled" || c[0].Reason == "TaskRunCancelled" {
			status = colorStatus("Cancelled") + "(" + c[0].Reason + ")"
		} else if c[0].Reason != status {
			status = cstatus + "(" + c[0].Reason + ")"
		}
	} else {
		status = cstatus
	}

	return status
}
