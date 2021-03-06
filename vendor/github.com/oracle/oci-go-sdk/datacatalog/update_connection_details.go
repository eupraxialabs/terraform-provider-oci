// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateConnectionDetails Properties used in connection update operations.
type UpdateConnectionDetails struct {

	// A description of the connection.
	Description *string `mandatory:"false" json:"description"`

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A map of maps that contains the properties which are specific to the connection type. Each connection type
	// definition defines it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// connections have required properties within the "default" category. To determine the set of optional and
	// required properties for a connection type, a query can be done on '/types?type=connection' that returns a
	// collection of all connection types. The appropriate connection type, which will include definitions of all
	// of it's properties, can be identified from this collection.
	// Example: `{"properties": { "default": { "username": "user1"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`

	// A map of maps that contains the encrypted values for sensitive properties which are specific to the
	// connection type. Each connection type definition defines it's set of required and optional properties.
	// The map keys are category names and the values are maps of property name to property value. Every property is
	// contained inside of a category. Most connections have required properties within the "default" category.
	// To determine the set of optional and required properties for a connection type, a query can be done
	// on '/types?type=connection' that returns a collection of all connection types. The appropriate connection
	// type, which will include definitions of all of it's properties, can be identified from this collection.
	// Example: `{"encProperties": { "default": { "password": "pwd"}}}`
	EncProperties map[string]map[string]string `mandatory:"false" json:"encProperties"`

	// Indicates whether this connection is the default connection.
	IsDefault *bool `mandatory:"false" json:"isDefault"`
}

func (m UpdateConnectionDetails) String() string {
	return common.PointerString(m)
}
