---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_sql_firewall_policy"
sidebar_current: "docs-oci-resource-data_safe-sql_firewall_policy"
description: |-
  Provides the Sql Firewall Policy resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_sql_firewall_policy
This resource provides the Sql Firewall Policy resource in Oracle Cloud Infrastructure Data Safe service.

Updates the SQL Firewall policy.

## Example Usage

```hcl
resource "oci_data_safe_sql_firewall_policy" "test_sql_firewall_policy" {
	#Required
	sql_firewall_policy_id = oci_data_safe_sql_firewall_policy.test_sql_firewall_policy.id

	#Optional
	allowed_client_ips = var.sql_firewall_policy_allowed_client_ips
	allowed_client_os_usernames = var.sql_firewall_policy_allowed_client_os_usernames
	allowed_client_programs = var.sql_firewall_policy_allowed_client_programs
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.sql_firewall_policy_description
	display_name = var.sql_firewall_policy_display_name
	enforcement_scope = var.sql_firewall_policy_enforcement_scope
	freeform_tags = {"Department"= "Finance"}
	status = var.sql_firewall_policy_status
	violation_action = var.sql_firewall_policy_violation_action
	violation_audit = var.sql_firewall_policy_violation_audit
}
```

## Argument Reference

The following arguments are supported:

* `allowed_client_ips` - (Optional) (Updatable) List of allowed ip addresses for the SQL Firewall policy.
* `allowed_client_os_usernames` - (Optional) (Updatable) List of allowed operating system user names for the SQL Firewall policy.
* `allowed_client_programs` - (Optional) (Updatable) List of allowed client programs for the SQL Firewall policy.
* `compartment_id` - (Optional) (Updatable) The OCID of the compartment containing the SQL Firewall policy.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the SQL Firewall policy.
* `display_name` - (Optional) (Updatable) The display name of the SQL Firewall policy. The name does not have to be unique, and it is changeable.
* `enforcement_scope` - (Optional) (Updatable) Specifies the SQL Firewall policy enforcement option.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `sql_firewall_policy_id` - (Required) The OCID of the SQL Firewall policy resource.
* `status` - (Optional) (Updatable) Specifies whether the SQL Firewall policy is enabled or disabled.
* `violation_action` - (Optional) (Updatable) Specifies the SQL Firewall action based on detection of SQL Firewall violations.
* `violation_audit` - (Optional) (Updatable) Specifies whether a unified audit policy should be enabled for auditing the SQL Firewall policy violations.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `allowed_client_ips` - The list of allowed ip addresses for the SQL Firewall policy.
* `allowed_client_os_usernames` - The list of allowed operating system user names for the SQL Firewall policy.
* `allowed_client_programs` - The list of allowed client programs for the SQL Firewall policy.
* `compartment_id` - The OCID of the compartment containing the SQL Firewall policy.
* `db_user_name` - The database user name.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the SQL Firewall policy.
* `display_name` - The display name of the SQL Firewall policy.
* `enforcement_scope` - Specifies the SQL Firewall policy enforcement option.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the SQL Firewall policy.
* `lifecycle_details` - Details about the current state of the SQL Firewall policy in Data Safe.
* `security_policy_id` - The OCID of the security policy corresponding to the SQL Firewall policy.
* `sql_level` - Specifies the level of SQL included for this SQL Firewall policy. USER_ISSUED_SQL - User issued SQL statements only. ALL_SQL - Includes all SQL statements including SQL statement issued inside PL/SQL units. 
* `state` - The current state of the SQL Firewall policy.
* `status` - Specifies whether the SQL Firewall policy is enabled or disabled.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time that the SQL Firewall policy was created, in the format defined by RFC3339.
* `time_updated` - The date and time the SQL Firewall policy was last updated, in the format defined by RFC3339.
* `violation_action` - Specifies the mode in which the SQL Firewall policy is enabled.
* `violation_audit` - Specifies whether a unified audit policy should be enabled for auditing the SQL Firewall policy violations.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Sql Firewall Policy
	* `update` - (Defaults to 20 minutes), when updating the Sql Firewall Policy
	* `delete` - (Defaults to 20 minutes), when destroying the Sql Firewall Policy


## Import

SqlFirewallPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_sql_firewall_policy.test_sql_firewall_policy "id"
```

