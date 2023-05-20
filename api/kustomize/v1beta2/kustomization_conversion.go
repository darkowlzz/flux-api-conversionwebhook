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

	kustomizev1 "github.com/darkowlzz/flux-api-conversionwebhook/api/kustomize/v1"
)

func (src *Kustomization) ConvertTo(dstRaw conversion.Hub) error {
	kustomizationlog.Info("Converting Kustomization from v1beta2 to v1")

	dst := dstRaw.(*kustomizev1.Kustomization)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.CommonMetadata = (*kustomizev1.CommonMetadata)(src.Spec.CommonMetadata)
	dst.Spec.DependsOn = src.Spec.DependsOn
	dst.Spec.Decryption = (*kustomizev1.Decryption)(src.Spec.Decryption)
	dst.Spec.Interval = src.Spec.Interval
	dst.Spec.RetryInterval = src.Spec.RetryInterval
	dst.Spec.KubeConfig = src.Spec.KubeConfig
	dst.Spec.Path = src.Spec.Path

	if src.Spec.PostBuild != nil {
		postBuild := kustomizev1.PostBuild{
			Substitute: src.Spec.PostBuild.Substitute,
		}

		substFrom := []kustomizev1.SubstituteReference{}
		for _, sf := range src.Spec.PostBuild.SubstituteFrom {
			substFrom = append(substFrom, kustomizev1.SubstituteReference(sf))
		}
		postBuild.SubstituteFrom = substFrom
	}

	dst.Spec.Prune = src.Spec.Prune
	dst.Spec.HealthChecks = src.Spec.HealthChecks
	dst.Spec.Patches = src.Spec.Patches
	dst.Spec.Images = src.Spec.Images
	dst.Spec.ServiceAccountName = src.Spec.ServiceAccountName
	dst.Spec.SourceRef = kustomizev1.CrossNamespaceSourceReference(src.Spec.SourceRef)
	dst.Spec.Suspend = src.Spec.Suspend
	dst.Spec.TargetNamespace = src.Spec.TargetNamespace
	dst.Spec.Timeout = src.Spec.Timeout
	dst.Spec.Force = src.Spec.Force
	dst.Spec.Wait = src.Spec.Wait
	dst.Spec.Components = src.Spec.Components

	// Status
	dst.Status.ReconcileRequestStatus = src.Status.ReconcileRequestStatus
	dst.Status.ObservedGeneration = src.Status.ObservedGeneration
	dst.Status.Conditions = src.Status.Conditions
	dst.Status.LastAppliedRevision = src.Status.LastAppliedRevision
	dst.Status.LastAttemptedRevision = src.Status.LastAttemptedRevision

	if src.Status.Inventory != nil {
		resRef := []kustomizev1.ResourceRef{}
		for _, rr := range src.Status.Inventory.Entries {
			resRef = append(resRef, kustomizev1.ResourceRef(rr))
		}
		dst.Status.Inventory = &kustomizev1.ResourceInventory{Entries: resRef}
	}

	return nil
}

func (dst *Kustomization) ConvertFrom(srcRaw conversion.Hub) error {
	kustomizationlog.Info("Converting Kustomization from v1 to v1beta2")

	src := srcRaw.(*kustomizev1.Kustomization)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.CommonMetadata = (*CommonMetadata)(src.Spec.CommonMetadata)
	dst.Spec.DependsOn = src.Spec.DependsOn
	dst.Spec.Decryption = (*Decryption)(src.Spec.Decryption)
	dst.Spec.Interval = src.Spec.Interval
	dst.Spec.RetryInterval = src.Spec.RetryInterval
	dst.Spec.KubeConfig = src.Spec.KubeConfig
	dst.Spec.Path = src.Spec.Path

	if src.Spec.PostBuild != nil {
		postBuild := kustomizev1.PostBuild{
			Substitute: src.Spec.PostBuild.Substitute,
		}

		substFrom := []kustomizev1.SubstituteReference{}
		for _, sf := range src.Spec.PostBuild.SubstituteFrom {
			substFrom = append(substFrom, kustomizev1.SubstituteReference(sf))
		}
		postBuild.SubstituteFrom = substFrom
	}

	dst.Spec.Prune = src.Spec.Prune
	dst.Spec.HealthChecks = src.Spec.HealthChecks
	dst.Spec.Patches = src.Spec.Patches
	dst.Spec.Images = src.Spec.Images
	dst.Spec.ServiceAccountName = src.Spec.ServiceAccountName
	dst.Spec.SourceRef = CrossNamespaceSourceReference(src.Spec.SourceRef)
	dst.Spec.Suspend = src.Spec.Suspend
	dst.Spec.TargetNamespace = src.Spec.TargetNamespace
	dst.Spec.Timeout = src.Spec.Timeout
	dst.Spec.Force = src.Spec.Force
	dst.Spec.Wait = src.Spec.Wait
	dst.Spec.Components = src.Spec.Components

	// Status
	dst.Status.ReconcileRequestStatus = src.Status.ReconcileRequestStatus
	dst.Status.ObservedGeneration = src.Status.ObservedGeneration
	dst.Status.Conditions = src.Status.Conditions
	dst.Status.LastAppliedRevision = src.Status.LastAppliedRevision
	dst.Status.LastAttemptedRevision = src.Status.LastAttemptedRevision

	if src.Status.Inventory != nil {
		resRef := []ResourceRef{}
		for _, rr := range src.Status.Inventory.Entries {
			resRef = append(resRef, ResourceRef(rr))
		}
		dst.Status.Inventory = &ResourceInventory{Entries: resRef}
	}

	return nil
}
