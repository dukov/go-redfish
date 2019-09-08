/*
 * Redfish OAPI specification
 *
 * Partial Redfish OAPI specification for a limited client
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
type State string

// List of State
const (
	ENABLED State = "Enabled"
	DISABLED State = "Disabled"
	STANDBY_OFFLINE State = "StandbyOffline"
	STANDBY_SPARE State = "StandbySpare"
	IN_TEST State = "InTest"
	STARTING State = "Starting"
	ABSENT State = "Absent"
	UNAVAILABLE_OFFLINE State = "UnavailableOffline"
	DEFERRING State = "Deferring"
	QUIESCED State = "Quiesced"
	UPDATING State = "Updating"
)