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

	sourcev1 "github.com/darkowlzz/flux-api-conversionwebhook/api/source/v1"
)

func (src *GitRepository) ConvertTo(dstRaw conversion.Hub) error {
	gitrepositorylog.Info("Converting GitRepository from v1beta2 to v1")

	dst := dstRaw.(*sourcev1.GitRepository)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.URL = src.Spec.URL
	dst.Spec.SecretRef = src.Spec.SecretRef
	dst.Spec.Interval = src.Spec.Interval
	dst.Spec.Timeout = src.Spec.Timeout
	dst.Spec.Reference = (*sourcev1.GitRepositoryRef)(src.Spec.Reference)
	dst.Spec.Verification = (*sourcev1.GitRepositoryVerification)(src.Spec.Verification)
	dst.Spec.Ignore = src.Spec.Ignore
	dst.Spec.Suspend = src.Spec.Suspend
	dst.Spec.RecurseSubmodules = src.Spec.RecurseSubmodules

	if len(src.Spec.Include) > 0 {
		newIncludes := []sourcev1.GitRepositoryInclude{}
		for _, inc := range src.Spec.Include {
			newIncludes = append(newIncludes, sourcev1.GitRepositoryInclude(inc))
		}
		dst.Spec.Include = newIncludes
	}

	// Status
	dst.Status.ObservedGeneration = src.Status.ObservedGeneration
	dst.Status.Conditions = src.Status.Conditions
	dst.Status.Artifact = src.Status.Artifact
	dst.Status.IncludedArtifacts = src.Status.IncludedArtifacts
	dst.Status.ObservedIgnore = src.Status.ObservedIgnore
	dst.Status.ObservedRecurseSubmodules = src.Status.ObservedRecurseSubmodules

	if len(src.Status.ObservedInclude) > 0 {
		newObservedIncludes := []sourcev1.GitRepositoryInclude{}
		for _, oinc := range src.Status.ObservedInclude {
			newObservedIncludes = append(newObservedIncludes, sourcev1.GitRepositoryInclude(oinc))
		}
		dst.Status.ObservedInclude = newObservedIncludes
	}
	dst.Status.ReconcileRequestStatus = src.Status.ReconcileRequestStatus

	return nil
}

func (dst *GitRepository) ConvertFrom(srcRaw conversion.Hub) error {
	gitrepositorylog.Info("Converting GitRepository from v1 to v1beta2")

	src := srcRaw.(*sourcev1.GitRepository)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.URL = src.Spec.URL
	dst.Spec.SecretRef = src.Spec.SecretRef
	dst.Spec.Interval = src.Spec.Interval
	dst.Spec.Timeout = src.Spec.Timeout
	dst.Spec.Reference = (*GitRepositoryRef)(src.Spec.Reference)
	dst.Spec.Verification = (*GitRepositoryVerification)(src.Spec.Verification)
	dst.Spec.Ignore = src.Spec.Ignore
	dst.Spec.Suspend = src.Spec.Suspend
	dst.Spec.RecurseSubmodules = src.Spec.RecurseSubmodules

	if len(src.Spec.Include) > 0 {
		newIncludes := []GitRepositoryInclude{}
		for _, inc := range src.Spec.Include {
			newIncludes = append(newIncludes, GitRepositoryInclude(inc))
		}
		dst.Spec.Include = newIncludes
	}

	// Status
	dst.Status.ObservedGeneration = src.Status.ObservedGeneration
	dst.Status.Conditions = src.Status.Conditions
	dst.Status.Artifact = src.Status.Artifact
	dst.Status.IncludedArtifacts = src.Status.IncludedArtifacts
	dst.Status.ObservedIgnore = src.Status.ObservedIgnore
	dst.Status.ObservedRecurseSubmodules = src.Status.ObservedRecurseSubmodules

	if len(src.Status.ObservedInclude) > 0 {
		newObservedIncludes := []GitRepositoryInclude{}
		for _, oinc := range src.Status.ObservedInclude {
			newObservedIncludes = append(newObservedIncludes, GitRepositoryInclude(oinc))
		}
		dst.Status.ObservedInclude = newObservedIncludes
	}
	dst.Status.ReconcileRequestStatus = src.Status.ReconcileRequestStatus

	return nil
}
