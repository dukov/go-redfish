/*
 * Redfish OAPI specification
 *
 * Partial Redfish OAPI specification for a limited client
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
type Health string

// List of Health
const (
	OK Health = "OK"
	WARNING Health = "Warning"
	CRITICAL Health = "Critical"
)