---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_cloud_db_system_connector"
sidebar_current: "docs-oci-resource-database_management-cloud_db_system_connector"
description: |-
  Provides the Cloud Db System Connector resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_cloud_db_system_connector
This resource provides the Cloud Db System Connector resource in Oracle Cloud Infrastructure Database Management service.

Creates a new cloud connector.

## Example Usage

```hcl
resource "oci_database_management_cloud_db_system_connector" "test_cloud_db_system_connector" {
	#Required
	cloud_db_system_id = oci_database_management_cloud_db_system.test_cloud_db_system.id
	connector_type = var.cloud_db_system_connector_connector_type

	#Optional
	display_name = var.cloud_db_system_connector_display_name
}
```

## Argument Reference

The following arguments are supported:

* `cloud_db_system_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system.
* `connector_type` - (Required) (Updatable) The type of connector.
* `display_name` - (Optional) The user-friendly name for the cloud connector. The name does not have to be unique.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent used for the cloud DB system connector. 
* `cloud_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system that the connector is a part of.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `connection_failure_message` - The error message indicating the reason for connection failure or `null` if the connection was successful. 
* `connection_info` - The connection details required to connect to a cloud DB system component.
	* `component_type` - The component type.
	* `connection_credentials` - The credentials used to connect to the Cloud ASM instance. Currently only the `DETAILS` type is supported for creating MACS connector credentials. 
		* `credential_name` - The name of the credential information that used to connect to the DB system resource. The name should be in "x.y" format, where the length of "x" has a maximum of 64 characters, and length of "y" has a maximum of 199 characters. The name strings can contain letters, numbers and the underscore character only. Other characters are not valid, except for the "." character that separates the "x" and "y" portions of the name. *IMPORTANT* - The name must be unique within the Oracle Cloud Infrastructure region the credential is being created in. If you specify a name that duplicates the name of another credential within the same Oracle Cloud Infrastructure region, you may overwrite or corrupt the credential that is already using the name.

			For example: inventorydb.abc112233445566778899 
		* `credential_type` - The type of credential used to connect to the ASM instance.
		* `named_credential_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Named Credential where the database password metadata is stored. 
		* `password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
		* `role` - The role of the user connecting to the ASM instance.
		* `ssl_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the SSL keystore and truststore details.
		* `user_name` - The user name used to connect to the ASM instance.
	* `connection_string` - The Oracle Database connection string. 
		* `host_name` - The host name of the database or the SCAN name in case of a RAC database.
		* `hosts` - The list of host names of the ASM instances.
		* `port` - The port used to connect to the ASM instance.
		* `protocol` - The protocol used to connect to the ASM instance.
		* `service` - The service name of the ASM instance.
* `connection_status` - The status of connectivity to the cloud DB system component.
* `connector_type` - The type of connector.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name for the cloud connector. The name does not have to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system connector.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current lifecycle state of the cloud DB system connector.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_connection_status_last_updated` - The date and time the connectionStatus of the cloud DB system connector was last updated.
* `time_created` - The date and time the cloud DB system connector was created.
* `time_updated` - The date and time the cloud DB system connector was last updated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cloud Db System Connector
	* `update` - (Defaults to 20 minutes), when updating the Cloud Db System Connector
	* `delete` - (Defaults to 20 minutes), when destroying the Cloud Db System Connector


## Import

CloudDbSystemConnectors can be imported using the `id`, e.g.

```
$ terraform import oci_database_management_cloud_db_system_connector.test_cloud_db_system_connector "id"
```

