---
subcategory: "Apiaccesscontrol"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apiaccesscontrol_privileged_api_requests"
sidebar_current: "docs-oci-datasource-apiaccesscontrol-privileged_api_requests"
description: |-
  Provides the list of Privileged Api Requests in Oracle Cloud Infrastructure Apiaccesscontrol service
---

# Data Source: oci_apiaccesscontrol_privileged_api_requests
This data source provides the list of Privileged Api Requests in Oracle Cloud Infrastructure Apiaccesscontrol service.

Lists all privilegedApi requests in the compartment.


## Example Usage

```hcl
data "oci_apiaccesscontrol_privileged_api_requests" "test_privileged_api_requests" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.privileged_api_request_display_name
	id = var.privileged_api_request_id
	resource_id = oci_cloud_guard_resource.test_resource.id
	resource_type = var.privileged_api_request_resource_type
	state = var.privileged_api_request_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the PrivilegedApiRequest.
* `resource_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource .
* `resource_type` - (Optional) A filter to return only lists of resources that match the entire given service type.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `privileged_api_request_collection` - The list of privileged_api_request_collection.

### PrivilegedApiRequest Reference

The following attributes are exported:

* `approver_details` - Contains the approver details who have approved the privilegedApi Request during the initial request.
	* `approval_action` - The action done by the approver.
	* `approval_comment` - Comment specified by the approver of the request.
	* `approver_id` - The userId of the approver.
	* `time_approved_for_access` - Time for when the privilegedApi request should start that is authorized by the customer in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z' 
	* `time_of_authorization` - Time when the privilegedApi request was authorized by the customer in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.Example: '2020-05-22T21:10:29.600Z' 
* `closure_comment` - The comment entered by the operator while closing the request.
* `compartment_id` - The OCID of the compartment that contains the access request.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Name of the privilegedApi control. The name must be unique.
* `duration_in_hrs` - Duration in hours for which access is sought on the target resource.
* `entity_type` - entityType of resource for which the AccessRequest is applicable
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the privilegedApi request.
* `lifecycle_details` - more in detail about the lifeCycleState.
* `notification_topic_id` - The OCID of the Oracle Cloud Infrastructure Notification topic to publish messages related to this privileged api request.
* `number_of_approvers_required` - Number of approvers required to approve an privilegedApi request.
* `privileged_api_control_id` - The OCID of the privilegedApi control governing the target resource.
* `privileged_api_control_name` - Name of the privilegedApi control governing the target resource.
* `privileged_operation_list` - List of api names, attributes for which approval is sought by the user. 
	* `api_name` - name of the api which needs to be protected.
	* `attribute_names` - list of attributes belonging to the above api which needs to be protected.
* `reason_detail` - Reason in Detail for which the operator is requesting access on the target resource.
* `reason_summary` - Summary comment by the operator creating the access request.
* `request_id` - This is an automatic identifier generated by the system which is easier for human comprehension.
* `requested_by` - List of Users who has created this privilegedApiRequest. 
* `resource_id` - The OCID of the target resource associated with the access request. The operator raises an access request to get approval to access the target resource. 
* `resource_name` - resourceName for which the PrivilegedApiRequest is applicable
* `resource_type` - resourceType for which the AccessRequest is applicable
* `severity` - Priority assigned to the access request by the operator
* `state` - The current state of the PrivilegedApiRequest.
* `state_details` - A message that describes the current state of the PrivilegedApiControl in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `sub_resource_name_list` - The subresource names requested for approval.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `ticket_numbers` - A list of ticket numbers related to this Privileged Api Access Request, e.g. Service Request (SR) number and JIRA ticket number. 
* `time_created` - Time when the privilegedApi request was created in [RFC 3339](https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z' 
* `time_requested_for_future_access` - Time in future when the user for the privilegedApi request needs to be created in [RFC 3339](https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z' 
* `time_updated` - Time when the privilegedApi request was last modified in [RFC 3339](https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z' 

