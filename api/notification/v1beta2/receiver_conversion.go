/*
Copyright 2023 The Flux authors

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

package v1beta2

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	notificationv1 "github.com/darkowlzz/flux-api-conversionwebhook/api/notification/v1"
)

func (src *Receiver) ConvertTo(dstRaw conversion.Hub) error {
	receiverlog.Info("Converting Receiver from v1beta2 to v1")

	dst := dstRaw.(*notificationv1.Receiver)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.Type = src.Spec.Type
	dst.Spec.Interval = src.Spec.Interval
	dst.Spec.Events = src.Spec.Events
	dst.Spec.Resources = src.Spec.Resources
	dst.Spec.SecretRef = src.Spec.SecretRef
	dst.Spec.Suspend = src.Spec.Suspend

	// Status
	dst.Status.ReconcileRequestStatus = src.Status.ReconcileRequestStatus
	dst.Status.Conditions = src.Status.Conditions
	dst.Status.WebhookPath = src.Status.WebhookPath
	dst.Status.ObservedGeneration = src.Status.ObservedGeneration

	return nil
}

func (dst *Receiver) ConvertFrom(srcRaw conversion.Hub) error {
	receiverlog.Info("Converting Receiver from v1 to v1beta2")

	src := srcRaw.(*notificationv1.Receiver)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.Type = src.Spec.Type
	dst.Spec.Interval = src.Spec.Interval
	dst.Spec.Events = src.Spec.Events
	dst.Spec.Resources = src.Spec.Resources
	dst.Spec.SecretRef = src.Spec.SecretRef
	dst.Spec.Suspend = src.Spec.Suspend

	// Status
	dst.Status.ReconcileRequestStatus = src.Status.ReconcileRequestStatus
	dst.Status.Conditions = src.Status.Conditions
	dst.Status.WebhookPath = src.Status.WebhookPath
	dst.Status.ObservedGeneration = src.Status.ObservedGeneration

	return nil
}
