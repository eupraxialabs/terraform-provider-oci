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

// ImportGlossaryDetails Import glossary from the contents of the glossary definition file.
type ImportGlossaryDetails struct {

	// The file contents used for the import of glossary.
	GlossaryFileContents []byte `mandatory:"false" json:"glossaryFileContents"`
}

func (m ImportGlossaryDetails) String() string {
	return common.PointerString(m)
}
