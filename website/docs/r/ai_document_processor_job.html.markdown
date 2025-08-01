---
subcategory: "Ai Document"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_document_processor_job"
sidebar_current: "docs-oci-resource-ai_document-processor_job"
description: |-
  Provides the Processor Job resource in Oracle Cloud Infrastructure Ai Document service
---

# oci_ai_document_processor_job
This resource provides the Processor Job resource in Oracle Cloud Infrastructure Ai Document service.

Create a processor job for document analysis.


## Example Usage

```hcl
resource "oci_ai_document_processor_job" "test_processor_job" {
	#Required
	compartment_id = var.compartment_id
	input_location {
		#Required
		source_type = var.processor_job_input_location_source_type

		#Optional
		data = var.processor_job_input_location_data
		object_locations {

			#Optional
			bucket = var.processor_job_input_location_object_locations_bucket
			namespace = var.processor_job_input_location_object_locations_namespace
			object = var.processor_job_input_location_object_locations_object
			page_range = var.processor_job_input_location_object_locations_page_range
		}
		page_range = var.processor_job_input_location_page_range
	}
	output_location {
		#Required
		bucket = var.processor_job_output_location_bucket
		namespace = var.processor_job_output_location_namespace
		prefix = var.processor_job_output_location_prefix
	}
	processor_config {
		#Required
		processor_type = var.processor_job_processor_config_processor_type

		#Optional
		document_type = var.processor_job_processor_config_document_type
		features {
			#Required
			feature_type = var.processor_job_processor_config_features_feature_type

			#Optional
			generate_searchable_pdf = var.processor_job_processor_config_features_generate_searchable_pdf
			max_results = var.processor_job_processor_config_features_max_results
			model_id = oci_ai_document_model.test_model.id
			selection_mark_detection = var.processor_job_processor_config_features_selection_mark_detection
			tenancy_id = oci_identity_tenancy.test_tenancy.id
		}
		is_zip_output_enabled = var.processor_job_processor_config_is_zip_output_enabled
		language = var.processor_job_processor_config_language
		model_id = oci_ai_document_model.test_model.id
		normalization_fields {

			#Optional
			map {

				#Optional
				normalization_type = var.processor_job_processor_config_normalization_fields_map_normalization_type
			}
		}
	}

	#Optional
	display_name = var.processor_job_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment identifier.
* `display_name` - (Optional) The display name of the processor job.
* `input_location` - (Required) The location of the inputs.
	* `data` - (Required when source_type=INLINE_DOCUMENT_CONTENT) Raw document data with Base64 encoding.
	* `object_locations` - (Required when source_type=OBJECT_STORAGE_LOCATIONS) The list of ObjectLocations.
		* `bucket` - (Required when source_type=OBJECT_STORAGE_LOCATIONS) The Object Storage bucket name.
		* `namespace` - (Required when source_type=OBJECT_STORAGE_LOCATIONS) The Object Storage namespace name.
		* `object` - (Required when source_type=OBJECT_STORAGE_LOCATIONS) The Object Storage object name.
		* `page_range` - (Applicable when source_type=OBJECT_STORAGE_LOCATIONS) The page ranges to be analysed.
	* `page_range` - (Applicable when source_type=INLINE_DOCUMENT_CONTENT) The page ranges to be analysed.
	* `source_type` - (Required) The type of input location. The allowed values are:
		* `OBJECT_STORAGE_LOCATIONS`: A list of object locations in Object Storage.
		* `INLINE_DOCUMENT_CONTENT`: The content of an inline document. 
* `output_location` - (Required) The object storage location where to store analysis results.
	* `bucket` - (Required) The Object Storage bucket name.
	* `namespace` - (Required) The Object Storage namespace.
	* `prefix` - (Required) The Object Storage folder name.
* `processor_config` - (Required) The configuration of a processor.
	* `document_type` - (Applicable when processor_type=GENERAL) The document type.
	* `features` - (Required when processor_type=GENERAL) The types of document analysis requested.
		* `feature_type` - (Required) The type of document analysis requested. The allowed values are:
			* `LANGUAGE_CLASSIFICATION`: Detect the language.
			* `TEXT_EXTRACTION`: Recognize text.
			* `TABLE_EXTRACTION`: Detect and extract data in tables.
			* `KEY_VALUE_EXTRACTION`: Extract form fields.
			* `DOCUMENT_CLASSIFICATION`: Identify the type of document.
			* `DOCUMENT_ELEMENTS_EXTRACTION`: Extract information from bar code 
		* `generate_searchable_pdf` - (Applicable when feature_type=TEXT_EXTRACTION) Whether or not to generate a searchable PDF file.
		* `max_results` - (Applicable when feature_type=DOCUMENT_CLASSIFICATION | LANGUAGE_CLASSIFICATION) The maximum number of results to return.
		* `model_id` - (Applicable when feature_type=DOCUMENT_CLASSIFICATION | DOCUMENT_ELEMENTS_EXTRACTION | KEY_VALUE_EXTRACTION | TABLE_EXTRACTION | TEXT_EXTRACTION) Unique identifier custom model OCID that should be used for inference.
		* `selection_mark_detection` - (Applicable when feature_type=TEXT_EXTRACTION) Whether checkbox detection feature is enabled or disabled.
		* `tenancy_id` - (Applicable when feature_type=DOCUMENT_CLASSIFICATION | KEY_VALUE_EXTRACTION) The custom model tenancy ID when modelId represents aliasName.
	* `is_zip_output_enabled` - (Applicable when processor_type=GENERAL) Whether or not to generate a ZIP file containing the results.
	* `language` - (Applicable when processor_type=GENERAL) The document language, abbreviated according to the BCP 47 Language-Tag syntax.
	* `model_id` - (Applicable when processor_type=INVOICE) Unique identifier custom model OCID that should be used for inference.
	* `normalization_fields` - (Required when processor_type=INVOICE) A string-to-object map where the key is the normalization field and the object contains information about the field.
		* `map` - (Applicable when processor_type=INVOICE) A wrapped map.
			* `normalization_type` - (Applicable when processor_type=INVOICE) A string mapping to the normalization type.
	* `processor_type` - (Required) The type of the processor.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The compartment identifier.
* `display_name` - The display name of the processor job.
* `id` - The id of the processor job.
* `input_location` - The location of the inputs.
	* `data` - Raw document data with Base64 encoding.
	* `object_locations` - The list of ObjectLocations.
		* `bucket` - The Object Storage bucket name.
		* `namespace` - The Object Storage namespace name.
		* `object` - The Object Storage object name.
		* `page_range` - The page ranges to be analysed.
	* `page_range` - The page ranges to be analysed.
	* `source_type` - The type of input location. The allowed values are:
		* `OBJECT_STORAGE_LOCATIONS`: A list of object locations in Object Storage.
		* `INLINE_DOCUMENT_CONTENT`: The content of an inline document. 
* `lifecycle_details` - The detailed status of FAILED state.
* `output_location` - The object storage location where to store analysis results.
	* `bucket` - The Object Storage bucket name.
	* `namespace` - The Object Storage namespace.
	* `prefix` - The Object Storage folder name.
* `percent_complete` - How much progress the operation has made, compared to the total amount of work to be performed.
* `processor_config` - The configuration of a processor.
	* `document_type` - The document type.
	* `features` - The types of document analysis requested.
		* `feature_type` - The type of document analysis requested. The allowed values are:
			* `LANGUAGE_CLASSIFICATION`: Detect the language.
			* `TEXT_EXTRACTION`: Recognize text.
			* `TABLE_EXTRACTION`: Detect and extract data in tables.
			* `KEY_VALUE_EXTRACTION`: Extract form fields.
			* `DOCUMENT_CLASSIFICATION`: Identify the type of document.
			* `DOCUMENT_ELEMENTS_EXTRACTION`: Extract information from bar code 
		* `generate_searchable_pdf` - Whether or not to generate a searchable PDF file.
		* `max_results` - The maximum number of results to return.
		* `model_id` - Unique identifier custom model OCID that should be used for inference.
		* `selection_mark_detection` - Whether checkbox detection feature is enabled or disabled.
		* `tenancy_id` - The custom model tenancy ID when modelId represents aliasName.
	* `is_zip_output_enabled` - Whether or not to generate a ZIP file containing the results.
	* `language` - The document language, abbreviated according to the BCP 47 Language-Tag syntax.
	* `model_id` - Unique identifier custom model OCID that should be used for inference.
	* `normalization_fields` - A string-to-object map where the key is the normalization field and the object contains information about the field.
		* `map` - A wrapped map.
			* `normalization_type` - A string mapping to the normalization type.
	* `processor_type` - The type of the processor.
* `state` - The current state of the processor job.
* `time_accepted` - The job acceptance time.
* `time_finished` - The job finish time.
* `time_started` - The job start time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Processor Job
	* `update` - (Defaults to 20 minutes), when updating the Processor Job
	* `delete` - (Defaults to 20 minutes), when destroying the Processor Job


## Import

ProcessorJobs can be imported using the `id`, e.g.

```
$ terraform import oci_ai_document_processor_job.test_processor_job "id"
```

