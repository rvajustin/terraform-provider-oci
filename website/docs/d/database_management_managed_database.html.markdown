---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database"
sidebar_current: "docs-oci-datasource-database_management-managed_database"
description: |-
  Provides details about a specific Managed Database in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database
This data source provides details about a specific Managed Database resource in Oracle Cloud Infrastructure Database Management service.

Gets the details of the Managed Database specified by managedDatabaseId.


## Example Usage

```hcl
data "oci_database_management_managed_database" "test_managed_database" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.


## Attributes Reference

The following attributes are exported:

* `additional_details` - The additional details specific to a type of database defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_platform_name` - The operating system of database.
* `database_status` - The status of the Oracle Database. Indicates whether the status of the database is UP, DOWN, or UNKNOWN at the current time. 
* `database_sub_type` - The subtype of the Oracle Database. Indicates whether the database is a Container Database, Pluggable Database, Non-container Database, Autonomous Database, or Autonomous Container Database. 
* `database_type` - The type of Oracle Database installation.
* `database_version` - The Oracle Database version.
* `db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system that this Managed Database is part of. 
* `dbmgmt_feature_configs` - The list of feature configurations
	* `connector_details` - The connector details required to connect to an Oracle cloud database.
		* `connector_type` - The list of supported connection types:
			* PE: Private endpoint
			* MACS: Management agent
			* EXTERNAL: External database connector 
		* `database_connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external database connector.
		* `management_agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent.
		* `private_end_point_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint.
	* `database_connection_details` - The connection details required to connect to the database.
		* `connection_credentials` - The credentials used to connect to the database. Currently only the `DETAILS` type is supported for creating MACS connector credentials. 
			* `credential_name` - The name of the credential information that used to connect to the DB system resource. The name should be in "x.y" format, where the length of "x" has a maximum of 64 characters, and length of "y" has a maximum of 199 characters. The name strings can contain letters, numbers and the underscore character only. Other characters are not valid, except for the "." character that separates the "x" and "y" portions of the name. *IMPORTANT* - The name must be unique within the Oracle Cloud Infrastructure region the credential is being created in. If you specify a name that duplicates the name of another credential within the same Oracle Cloud Infrastructure region, you may overwrite or corrupt the credential that is already using the name.

				For example: inventorydb.abc112233445566778899 
			* `credential_type` - The type of credential used to connect to the database.
			* `named_credential_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Named Credential where the database password metadata is stored. 
			* `password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
			* `role` - The role of the user connecting to the database.
			* `ssl_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the SSL keystore and truststore details.
			* `user_name` - The user name used to connect to the database.
		* `connection_string` - The details of the Oracle Database connection string. 
			* `connection_type` - The list of supported connection types:
				* BASIC: Basic connection details 
			* `port` - The port number used to connect to the database.
			* `protocol` - The protocol used to connect to the database.
			* `service` - The service name of the database.
	* `feature` - The name of the Database Management feature.
	* `feature_status` - The list of statuses for Database Management features. 
	* `license_model` - The Oracle license model that applies to the external database. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `deployment_type` - The infrastructure used to deploy the Oracle Database.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `is_cluster` - Indicates whether the Oracle Database is part of a cluster.
* `managed_database_groups` - A list of Managed Database Groups that the Managed Database belongs to.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the Managed Database Group resides.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database Group.
	* `name` - The name of the Managed Database Group.
* `management_option` - The management option used when enabling Database Management.
* `name` - The name of the Managed Database.
* `parent_container_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent Container Database if Managed Database is a Pluggable Database. 
* `storage_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the storage DB system.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the Managed Database was created.
* `workload_type` - The workload type of the Autonomous Database.

