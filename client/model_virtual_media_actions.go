/*
 * Redfish OAPI specification
 *
 * Partial Redfish OAPI specification for a limited client
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

type VirtualMediaActions struct {
	VirtualMediaEjectMedia VirtualMediaActionsVirtualMediaEjectMedia `json:"#VirtualMedia.EjectMedia,omitempty"`
	VirtualMediaInsertMedia VirtualMediaActionsVirtualMediaEjectMedia `json:"#VirtualMedia.InsertMedia,omitempty"`
}