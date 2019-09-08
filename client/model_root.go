/*
 * Redfish OAPI specification
 *
 * Partial Redfish OAPI specification for a limited client
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// Root redfish path.
type Root struct {
	// The name of the resource.
	Id string `json:"Id,omitempty"`
	// The name of the resource.
	Name string `json:"Name"`
	// redfish version
	RedfishVersion string `json:"RedfishVersion,omitempty"`
	UUID string `json:"UUID,omitempty"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// redfish copyright
	RedfishCopyright string `json:"@Redfish.Copyright,omitempty"`
	Systems IdRef `json:"Systems,omitempty"`
	Managers IdRef `json:"Managers,omitempty"`
}
