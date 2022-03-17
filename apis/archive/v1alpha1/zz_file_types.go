/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FileObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OutputBase64Sha256 *string `json:"outputBase64Sha256,omitempty" tf:"output_base64sha256,omitempty"`

	OutputMd5 *string `json:"outputMd5,omitempty" tf:"output_md5,omitempty"`

	OutputSha *string `json:"outputSha,omitempty" tf:"output_sha,omitempty"`

	OutputSize *float64 `json:"outputSize,omitempty" tf:"output_size,omitempty"`
}

type FileParameters struct {

	// +kubebuilder:validation:Optional
	Excludes []*string `json:"excludes,omitempty" tf:"excludes,omitempty"`

	// +kubebuilder:validation:Optional
	OutputFileMode *string `json:"outputFileMode,omitempty" tf:"output_file_mode,omitempty"`

	// +kubebuilder:validation:Required
	OutputPath *string `json:"outputPath" tf:"output_path,omitempty"`

	// +kubebuilder:validation:Optional
	Source []SourceParameters `json:"source,omitempty" tf:"source,omitempty"`

	// +kubebuilder:validation:Optional
	SourceContent *string `json:"sourceContent,omitempty" tf:"source_content,omitempty"`

	// +kubebuilder:validation:Optional
	SourceContentFilename *string `json:"sourceContentFilename,omitempty" tf:"source_content_filename,omitempty"`

	// +kubebuilder:validation:Optional
	SourceDir *string `json:"sourceDir,omitempty" tf:"source_dir,omitempty"`

	// +kubebuilder:validation:Optional
	SourceFile *string `json:"sourceFile,omitempty" tf:"source_file,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SourceObservation struct {
}

type SourceParameters struct {

	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// +kubebuilder:validation:Required
	Filename *string `json:"filename" tf:"filename,omitempty"`
}

// FileSpec defines the desired state of File
type FileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FileParameters `json:"forProvider"`
}

// FileStatus defines the observed state of File.
type FileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// File is the Schema for the Files API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,archivejet}
type File struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FileSpec   `json:"spec"`
	Status            FileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FileList contains a list of Files
type FileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []File `json:"items"`
}

// Repository type metadata.
var (
	File_Kind             = "File"
	File_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: File_Kind}.String()
	File_KindAPIVersion   = File_Kind + "." + CRDGroupVersion.String()
	File_GroupVersionKind = CRDGroupVersion.WithKind(File_Kind)
)

func init() {
	SchemeBuilder.Register(&File{}, &FileList{})
}