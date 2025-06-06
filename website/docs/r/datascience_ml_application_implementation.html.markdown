---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_ml_application_implementation"
sidebar_current: "docs-oci-resource-datascience-ml_application_implementation"
description: |-
  Provides the Ml Application Implementation resource in Oracle Cloud Infrastructure Data Science service
---

# oci_datascience_ml_application_implementation
This resource provides the Ml Application Implementation resource in Oracle Cloud Infrastructure Data Science service.

Creates a new MlApplicationImplementation.


## Example Usage

```hcl
resource "oci_datascience_ml_application_implementation" "test_ml_application_implementation" {
	#Required
	compartment_id = var.compartment_id
	ml_application_id = oci_datascience_ml_application.test_ml_application.id
	name = var.ml_application_implementation_name

	#Optional
	ml_application_package = {
		source_type = var.ml_application_package_source_type
		uri = var.ml_application_package_uri
	}
	opc_ml_app_package_args = var.opc_ml_app_package_args
	allowed_migration_destinations = var.ml_application_implementation_allowed_migration_destinations
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	logging {

		#Optional
		aggregated_instance_view_log {

			#Optional
			enable_logging = var.ml_application_implementation_logging_aggregated_instance_view_log_enable_logging
			log_group_id = oci_logging_log_group.test_log_group.id
			log_id = oci_logging_log.test_log.id
		}
		implementation_log {

			#Optional
			enable_logging = var.ml_application_implementation_logging_implementation_log_enable_logging
			log_group_id = oci_logging_log_group.test_log_group.id
			log_id = oci_logging_log.test_log.id
		}
		trigger_log {

			#Optional
			enable_logging = var.ml_application_implementation_logging_trigger_log_enable_logging
			log_group_id = oci_logging_log_group.test_log_group.id
			log_id = oci_logging_log.test_log.id
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `allowed_migration_destinations` - (Optional) (Updatable) List of ML Application Implementation OCIDs for which migration from this implementation is allowed. Migration means that if consumers change implementation for their instances to implementation with OCID from this list, instance components will be updated in place otherwise new instance components are created based on the new implementation and old instance components are removed.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment where ML Application Implementation is created.
* `ml_application_package` - (Optional) (Updatable) Configuration of The ML Application Package to upload.
    * `source_type` - (Optional) ML Application Package source type. List of Source Types
      * `local` - Refers local file path
      * `object_storage_download` - Terraform downloads package from Object Storage Bucket on local and then uploads
      * `object_storage` - Terraform pass package reference to ML Application Service
    * `path` - (Optional) ML Application Package local file path (Applicable when source_type="local"). e.g. file://${path.root}/ml-app-package.zip
    * `uri` - (Optional) Object Storage Bucket path for ML Application Package(Applicable when source_type is "object_storage_download" or "object_storage").
* `opc_ml_app_package_args` - (Optional) (Updatable) ML Application package arguments required during ML Application package upload. Each argument is a simple key-value pair.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `logging` - (Optional) (Updatable) Configuration of Logging for ML Application Implementation.
	* `aggregated_instance_view_log` - (Optional) (Updatable) Log configuration details for particular areas of ML Application Implementation.
		* `enable_logging` - (Optional) (Updatable) If logging is enabled.
		* `log_group_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
		* `log_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
	* `implementation_log` - (Optional) (Updatable) Log configuration details for particular areas of ML Application Implementation.
		* `enable_logging` - (Optional) (Updatable) If logging is enabled.
		* `log_group_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
		* `log_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
	* `trigger_log` - (Optional) (Updatable) Log configuration details for particular areas of ML Application Implementation.
		* `enable_logging` - (Optional) (Updatable) If logging is enabled.
		* `log_group_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
		* `log_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
* `ml_application_id` - (Required) The OCID of the ML Application implemented by this ML Application Implementation
* `name` - (Required) ML Application Implementation name which is unique for given ML Application.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `allowed_migration_destinations` - List of ML Application Implementation OCIDs for which migration from this implementation is allowed. Migration means that if consumers change implementation for their instances to implementation with OCID from this list, instance components will be updated in place otherwise new instance components are created based on the new implementation and old instance components are removed.
* `application_components` - List of application components (OCI resources shared for all MlApplicationInstances). These have been created automatically based on their definitions in the ML Application package.
	* `application_id` - OCID of Data Flow Application
	* `component_name` - Name of application component
	* `id` - OCID of the resource
	* `job_id` - OCID of Data Science Job
	* `model_id` - OCID of Data Science Model
	* `name` - Name of referenced resource (generally resources do not have to have any name but most resources have name exposed as 'name' or 'displayName' field).
	* `pipeline_id` - OCID of Data Science Pipeline
	* `resource_type` - Type of the resource
	* `type` - Type of application component
* `compartment_id` - The OCID of the compartment where ML Application Implementation is created.
* `configuration_schema` - Schema of configuration which needs to be provided for each ML Application Instance. It is defined in the ML Application package descriptor.
	* `default_value` - The default value for the optional configuration property (it must not be specified for mandatory configuration properties)
	* `description` - Description of this configuration property
	* `is_mandatory` - If the value is true this configuration property is mandatory and visa versa. If not specified configuration property is optional.
	* `key_name` - Name of key (parameter name)
	* `sample_value` - Sample property value (it must match validationRegexp if it is specified)
	* `validation_regexp` - A regular expression will be used for the validation of property value.
	* `value_type` - Type of value
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - Description of ML Application Implementation defined in ML Application package descriptor
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the MlApplicationImplementation. Unique identifier that is immutable after creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `logging` - Configuration of Logging for ML Application Implementation.
	* `aggregated_instance_view_log` - Log configuration details for particular areas of ML Application Implementation.
		* `enable_logging` - If logging is enabled.
		* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
		* `log_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
	* `implementation_log` - Log configuration details for particular areas of ML Application Implementation.
		* `enable_logging` - If logging is enabled.
		* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
		* `log_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
	* `trigger_log` - Log configuration details for particular areas of ML Application Implementation.
		* `enable_logging` - If logging is enabled.
		* `log_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log group.
		* `log_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the log.
* `ml_application_id` - The OCID of the ML Application implemented by this ML Application Implementation.
* `ml_application_name` - The name of ML Application (based on mlApplicationId)
* `ml_application_package_arguments` - List of ML Application package arguments provided during ML Application package upload.
	* `arguments` - Array of the ML Application package arguments
		* `description` - short description of the argument
		* `is_mandatory` - argument is mandatory or not
		* `name` - Argument name
		* `type` - type of the argument
		* `value` - Argument value
* `name` - ML Application Implementation name which is unique for given ML Application.
* `package_version` - The version of ML Application Package (e.g. "1.2" or "2.0.4") defined in ML Application package descriptor. Value is not mandatory only for CREATING state otherwise it must be always presented.
* `state` - The current state of the MlApplicationImplementation.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Creation time of MlApplicationImplementation creation in the format defined by RFC 3339.
* `time_updated` - Time of last MlApplicationImplementation update in the format defined by RFC 3339.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Ml Application Implementation
	* `update` - (Defaults to 20 minutes), when updating the Ml Application Implementation
	* `delete` - (Defaults to 20 minutes), when destroying the Ml Application Implementation


## Import

MlApplicationImplementations can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_ml_application_implementation.test_ml_application_implementation "id"
```

