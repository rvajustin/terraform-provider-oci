---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_cloud_databases"
sidebar_current: "docs-oci-datasource-database_management-cloud_databases"
description: |-
  Provides the list of Cloud Databases in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_cloud_databases
This data source provides the list of Cloud Databases in Oracle Cloud Infrastructure Database Management service.

Lists the cloud databases in the specified compartment or in the specified DB system.

## Example Usage

```hcl
data "oci_database_management_cloud_databases" "test_cloud_databases" {
	#Required
	cloud_db_system_id = oci_database_management_cloud_db_system.test_cloud_db_system.id

	#Optional
	compartment_id = var.compartment_id
	display_name = var.cloud_database_display_name
}
```

## Argument Reference

The following arguments are supported:

* `cloud_db_system_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system.
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to only return the resources that match the entire display name.


## Attributes Reference

The following attributes are exported:

* `cloud_database_collection` - The list of cloud_database_collection.

### CloudDatabase Reference

The following attributes are exported:

* `items` - An array of cloud databases.
	* `cloud_db_home_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB home.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	* `database_platform_name` - The operating system of database.
	* `database_sub_type` - The subtype of Oracle Database. Indicates whether the database is a Container Database, Pluggable Database, or Non-container Database. 
	* `database_type` - The type of Oracle Database installation.
	* `database_version` - The Oracle database version.
	* `db_management_config` - The configuration details of Database Management for a cloud DB system.
		* `is_enabled` - The status of the associated service.
		* `metadata` - The associated service-specific inputs in JSON string format, which Database Management can identify.
	* `db_system_info` - The basic information about a cloud DB system.
		* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
		* `display_name` - The user-friendly name for the cloud DB system. The name does not have to be unique.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system.
	* `db_unique_name` - The `DB_UNIQUE_NAME` of the external database.
	* `dbmgmt_feature_configs` - The list of feature configurations
		* `connector_details` - The connector details required to connect to an Oracle cloud database.
			* `connector_type` - The list of supported connection types:
				* PE: Private endpoint
				* MACS: Management agent
				* EXTERNAL: External database connector
				* DIRECT: Direct connection 
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
	* `display_name` - The user-friendly name for the database. The name does not have to be unique.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system.
	* `instance_details` - The list of database instances if the database is a RAC database.
		* `host_name` - The name of the host machine.
		* `instance_name` - The name of the database instance.
		* `instance_number` - The instance number of the database instance.
	* `parent_container_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent Container Database (CDB) if this is a Pluggable Database (PDB). 
	* `state` - The current lifecycle state of the external database resource.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_created` - The date and time the external DB system was created.

