// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InvoiceProcessorConfig The configuration of an invoice processor.
type InvoiceProcessorConfig struct {
	NormalizationFields *NormalizationFields `mandatory:"true" json:"normalizationFields"`

	// Unique identifier custom model OCID that should be used for inference.
	ModelId *string `mandatory:"false" json:"modelId"`
}

func (m InvoiceProcessorConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InvoiceProcessorConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m InvoiceProcessorConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInvoiceProcessorConfig InvoiceProcessorConfig
	s := struct {
		DiscriminatorParam string `json:"processorType"`
		MarshalTypeInvoiceProcessorConfig
	}{
		"INVOICE",
		(MarshalTypeInvoiceProcessorConfig)(m),
	}

	return json.Marshal(&s)
}
