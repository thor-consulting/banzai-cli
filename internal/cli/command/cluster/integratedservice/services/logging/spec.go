// Copyright © 2019 Banzai Cloud
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

package logging

import (
	"emperror.dev/errors"
)

type spec struct {
	Loki          lokiSpec          `json:"loki" mapstructure:"loki"`
	Logging       loggingSpec       `json:"logging" mapstructure:"logging"`
	ClusterOutput clusterOutputSpec `json:"clusterOutput" mapstructure:"clusterOutput"`
}

type lokiSpec struct {
	Enabled bool        `json:"enabled" mapstructure:"enabled"`
	Ingress ingressSpec `json:"ingress" mapstructure:"ingress"`
}

type ingressSpec struct {
	Enabled  bool   `json:"enabled" mapstructure:"enabled"`
	Domain   string `json:"domain" mapstructure:"domain"`
	Path     string `json:"path" mapstructure:"path"`
	SecretID string `json:"secretId" mapstructure:"secretId"`
}

type loggingSpec struct {
	Metrics bool `json:"metrics" mapstructure:"metrics"`
	TLS     bool `json:"tls" mapstructure:"tls"`
}

type clusterOutputSpec struct {
	Enabled  bool         `json:"enabled" mapstructure:"enabled"`
	Provider providerSpec `json:"provider" mapstructure:"provider"`
}

type providerSpec struct {
	Name     string     `json:"name" mapstructure:"name"`
	Bucket   bucketSpec `json:"bucket" mapstructure:"bucket"`
	SecretID string     `json:"secretId" mapstructure:"secretId"`
}

type bucketSpec struct {
	Name           string `json:"name" mapstructure:"name"`
	ResourceGroup  string `json:"resourceGroup" mapstructure:"resourceGroup"`
	StorageAccount string `json:"storageAccount" mapstructure:"storageAccount"`
}

func (s spec) Validate() error {
	if err := s.Loki.Validate(); err != nil {
		return err
	}

	if err := s.ClusterOutput.Validate(); err != nil {
		return err
	}

	return nil
}

func (s lokiSpec) Validate() error {
	if s.Enabled {
		if err := s.Ingress.Validate(); err != nil {
			return errors.WrapIf(err, "error during validating Loki ingress")
		}
	}

	return nil
}

func (s ingressSpec) Validate() error {
	if s.Enabled {
		if s.Path == "" {
			return requiredFieldError{name: "path"}
		}
	}

	return nil
}

func (s clusterOutputSpec) Validate() error {
	if s.Enabled {
		if err := s.Provider.Validate(); err != nil {
			return errors.WrapIf(err, "error during validating provider")
		}
	}

	return nil
}

func (s providerSpec) Validate() error {
	if s.SecretID == "" {
		return requiredFieldError{name: "secretId"}
	}

	if s.Name == "" {
		return requiredFieldError{name: "name"}
	}

	switch s.Name {
	case providerAmazonS3Key, providerAzureKey, providerGoogleGCSKey:
	default:
		return errors.New("invalid provider name")
	}

	if err := s.Bucket.Validate(s.Name); err != nil {
		return errors.WrapIf(err, "error during bucket validation")
	}

	return nil
}

func (s bucketSpec) Validate(provider string) error {
	if s.Name == "" {
		return requiredFieldError{name: "name"}
	}

	if provider == providerAzureKey {
		if s.ResourceGroup == "" {
			return requiredFieldError{name: "resourceGroup"}
		}

		if s.StorageAccount == "" {
			return requiredFieldError{name: "storageAccount"}
		}
	}

	return nil
}
