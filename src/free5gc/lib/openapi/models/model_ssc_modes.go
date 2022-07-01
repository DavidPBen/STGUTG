/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type SscModes struct {
	DefaultSscMode  SscMode   `json:"defaultSscMode" yaml:"defaultSscMode" bson:"defaultSscMode" mapstructure:"DefaultSscMode"`
	AllowedSscModes []SscMode `json:"allowedSscModes,omitempty" yaml:"allowedSscModes" bson:"allowedSscModes" mapstructure:"AllowedSscModes"`
}
