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

// ModelMetrics Trained Model Metrics.
type ModelMetrics interface {
	GetDatasetSummary() *DatasetSummary
}

type modelmetrics struct {
	JsonData       []byte
	DatasetSummary *DatasetSummary `mandatory:"false" json:"datasetSummary"`
	ModelType      string          `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *modelmetrics) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermodelmetrics modelmetrics
	s := struct {
		Model Unmarshalermodelmetrics
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DatasetSummary = s.Model.DatasetSummary
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *modelmetrics) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION":
		mm := PreTrainedDocumentElementsExtractionModelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DOCUMENT_CLASSIFICATION":
		mm := DocumentClassificationModelMetrics{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRE_TRAINED_DOCUMENT_CLASSIFICATION":
		mm := PretrainedDocumentClassificationModelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRE_TRAINED_TABLE_EXTRACTION":
		mm := PretrainedTableExtractionModelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "KEY_VALUE_EXTRACTION":
		mm := KeyValueDetectionModelMetrics{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRE_TRAINED_KEY_VALUE_EXTRACTION":
		mm := PretrainedKeyValueExtractionModelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRE_TRAINED_TEXT_EXTRACTION":
		mm := PretrainedTextExtractionModelDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ModelMetrics: %s.", m.ModelType)
		return *m, nil
	}
}

// GetDatasetSummary returns DatasetSummary
func (m modelmetrics) GetDatasetSummary() *DatasetSummary {
	return m.DatasetSummary
}

func (m modelmetrics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m modelmetrics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModelMetricsModelTypeEnum Enum with underlying type: string
type ModelMetricsModelTypeEnum string

// Set of constants representing the allowable values for ModelMetricsModelTypeEnum
const (
	ModelMetricsModelTypeKeyValueExtraction                   ModelMetricsModelTypeEnum = "KEY_VALUE_EXTRACTION"
	ModelMetricsModelTypeDocumentClassification               ModelMetricsModelTypeEnum = "DOCUMENT_CLASSIFICATION"
	ModelMetricsModelTypePreTrainedTextExtraction             ModelMetricsModelTypeEnum = "PRE_TRAINED_TEXT_EXTRACTION"
	ModelMetricsModelTypePreTrainedTableExtraction            ModelMetricsModelTypeEnum = "PRE_TRAINED_TABLE_EXTRACTION"
	ModelMetricsModelTypePreTrainedKeyValueExtraction         ModelMetricsModelTypeEnum = "PRE_TRAINED_KEY_VALUE_EXTRACTION"
	ModelMetricsModelTypePreTrainedDocumentClassification     ModelMetricsModelTypeEnum = "PRE_TRAINED_DOCUMENT_CLASSIFICATION"
	ModelMetricsModelTypePreTrainedDocumentElementsExtraction ModelMetricsModelTypeEnum = "PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION"
)

var mappingModelMetricsModelTypeEnum = map[string]ModelMetricsModelTypeEnum{
	"KEY_VALUE_EXTRACTION":                     ModelMetricsModelTypeKeyValueExtraction,
	"DOCUMENT_CLASSIFICATION":                  ModelMetricsModelTypeDocumentClassification,
	"PRE_TRAINED_TEXT_EXTRACTION":              ModelMetricsModelTypePreTrainedTextExtraction,
	"PRE_TRAINED_TABLE_EXTRACTION":             ModelMetricsModelTypePreTrainedTableExtraction,
	"PRE_TRAINED_KEY_VALUE_EXTRACTION":         ModelMetricsModelTypePreTrainedKeyValueExtraction,
	"PRE_TRAINED_DOCUMENT_CLASSIFICATION":      ModelMetricsModelTypePreTrainedDocumentClassification,
	"PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION": ModelMetricsModelTypePreTrainedDocumentElementsExtraction,
}

var mappingModelMetricsModelTypeEnumLowerCase = map[string]ModelMetricsModelTypeEnum{
	"key_value_extraction":                     ModelMetricsModelTypeKeyValueExtraction,
	"document_classification":                  ModelMetricsModelTypeDocumentClassification,
	"pre_trained_text_extraction":              ModelMetricsModelTypePreTrainedTextExtraction,
	"pre_trained_table_extraction":             ModelMetricsModelTypePreTrainedTableExtraction,
	"pre_trained_key_value_extraction":         ModelMetricsModelTypePreTrainedKeyValueExtraction,
	"pre_trained_document_classification":      ModelMetricsModelTypePreTrainedDocumentClassification,
	"pre_trained_document_elements_extraction": ModelMetricsModelTypePreTrainedDocumentElementsExtraction,
}

// GetModelMetricsModelTypeEnumValues Enumerates the set of values for ModelMetricsModelTypeEnum
func GetModelMetricsModelTypeEnumValues() []ModelMetricsModelTypeEnum {
	values := make([]ModelMetricsModelTypeEnum, 0)
	for _, v := range mappingModelMetricsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetModelMetricsModelTypeEnumStringValues Enumerates the set of values in String for ModelMetricsModelTypeEnum
func GetModelMetricsModelTypeEnumStringValues() []string {
	return []string{
		"KEY_VALUE_EXTRACTION",
		"DOCUMENT_CLASSIFICATION",
		"PRE_TRAINED_TEXT_EXTRACTION",
		"PRE_TRAINED_TABLE_EXTRACTION",
		"PRE_TRAINED_KEY_VALUE_EXTRACTION",
		"PRE_TRAINED_DOCUMENT_CLASSIFICATION",
		"PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION",
	}
}

// GetMappingModelMetricsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModelMetricsModelTypeEnum(val string) (ModelMetricsModelTypeEnum, bool) {
	enum, ok := mappingModelMetricsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
