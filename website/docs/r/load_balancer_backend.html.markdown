---
subcategory: "Load Balancer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_load_balancer_backend"
sidebar_current: "docs-oci-resource-load_balancer-backend"
description: |-
  Provides the Backend resource in Oracle Cloud Infrastructure Load Balancer service
---

# oci_load_balancer_backend
This resource provides the Backend resource in Oracle Cloud Infrastructure Load Balancer service.

Adds a backend server to a backend set.

## Example Usage

```hcl
resource "oci_load_balancer_backend" "test_backend" {
	#Required
	backendset_name = oci_load_balancer_backend_set.test_backend_set.name
	ip_address = var.backend_ip_address
	load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
	port = var.backend_port

	#Optional
	backup = var.backend_backup
	drain = var.backend_drain
	max_connections = var.backend_max_connections
	offline = var.backend_offline
	weight = var.backend_weight
}
```

## Argument Reference

The following arguments are supported:

* `backendset_name` - (Required) The name of the backend set to add the backend server to.  Example: `example_backend_set` 
* `backup` - (Optional) (Updatable) Whether the load balancer should treat this server as a backup unit. If `true`, the load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "backup" fail the health check policy.

	**Note:** You cannot add a backend server marked as `backup` to a backend set that uses the IP Hash policy.

	Example: `false` 
* `drain` - (Optional) (Updatable) Whether the load balancer should drain this server. Servers marked "drain" receive no new incoming traffic.  Example: `false` 
* `ip_address` - (Required) The IP address of the backend server.  Example: `10.0.0.3` 
* `load_balancer_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer associated with the backend set and servers.
* `max_connections` - (Optional) (Updatable) The maximum number of simultaneous connections the load balancer can make to the backend. If this is not set or set to 0 then the maximum number of simultaneous connections the load balancer can make to the backend is unlimited.

	If setting maxConnections to some value other than 0 then that value must be greater or equal to 256.

	Example: `300` 
* `offline` - (Optional) (Updatable) Whether the load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: `false` 
* `port` - (Required) The communication port for the backend server.  Example: `8080` 
* `weight` - (Optional) (Updatable) The load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives 3 times the number of new connections as a server weighted '1'. For more information on load balancing policies, see [How Load Balancing Policies Work](https://docs.cloud.oracle.com/iaas/Content/Balance/Reference/lbpolicies.htm).  Example: `3` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `backup` - Whether the load balancer should treat this server as a backup unit. If `true`, the load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "backup" fail the health check policy.

	**Note:** You cannot add a backend server marked as `backup` to a backend set that uses the IP Hash policy.

	Example: `false` 
* `drain` - Whether the load balancer should drain this server. Servers marked "drain" receive no new incoming traffic.  Example: `false` 
* `ip_address` - The IP address of the backend server.  Example: `10.0.0.3` 
* `max_connections` - The maximum number of simultaneous connections the load balancer can make to the backend. If this is not set or set to 0 then the maximum number of simultaneous connections the load balancer can make to the backend is unlimited.  Example: `300` 
* `name` - A read-only field showing the IP address and port that uniquely identify this backend server in the backend set.  Example: `10.0.0.3:8080` 
* `offline` - Whether the load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: `false` 
* `port` - The communication port for the backend server.  Example: `8080` 
* `weight` - The load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives 3 times the number of new connections as a server weighted '1'. For more information on load balancing policies, see [How Load Balancing Policies Work](https://docs.cloud.oracle.com/iaas/Content/Balance/Reference/lbpolicies.htm).  Example: `3` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Backend
	* `update` - (Defaults to 20 minutes), when updating the Backend
	* `delete` - (Defaults to 20 minutes), when destroying the Backend


## Import

Backends can be imported using the `id`, e.g.

```
$ terraform import oci_load_balancer_backend.test_backend "loadBalancers/{loadBalancerId}/backendSets/{backendSetName}/backends/{backendName}" 
```

