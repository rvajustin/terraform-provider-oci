// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_log_analytics_namespaces" "test_namespaces" {
  compartment_id = var.tenancy_ocid
}

resource "oci_log_analytics_namespace_storage_enable_disable_archiving" "test_namespace_storage_enable_archiving" {
  #Required
  namespace               = data.oci_log_analytics_namespaces.test_namespaces.namespace_collection.0.items.0.namespace
  enable_archiving_tenant = true
}

resource "oci_log_analytics_namespace_storage_enable_disable_archiving" "test_namespace_storage_disable_archiving" {
  #Required
  namespace               = data.oci_log_analytics_namespaces.test_namespaces.namespace_collection.0.items.0.namespace
  enable_archiving_tenant = false
}