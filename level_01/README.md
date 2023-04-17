# level_01

## Description

These exercises are designed to get you familiar with the fluent Terraform testing module.
It will require some skills in Go, but the exercises are designed to be easy to follow.

In these early examples the boilerplate code has been provided for you.

You will need to edit the [module_test.go](./test/module_test.go) file.

## Task 1 - TestSimple

Modify the TestSimple function to check the `input` attribute of the `terraform_data.example` resource is equal to `"example"`.

## Task 2 - TestCondition

Add checks that the `terraform_data.example_condition` resource is present when `var.example_condition` is `true` and that the `input` attribute is correct.
Also add checks that the `terraform_data.example_condition` resource is not present when `var.example_condition` is `false`.

## Task 3 - TestForEach

Modify the TestForEach function to check the `input` attribute of the `terraform_data.example_for_each` resources is equal to the data in the map.
