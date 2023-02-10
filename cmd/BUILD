load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_binary")

go_library(
  name = "hello_lib",
  srcs = ["main.go"],
  importpath = "bazel-go/cmd",
  visibility = ["//visibility:private"],
  deps = [
    "//calc",
  ]
)

go_binary(
    name = "hello",
    embed = [":hello_lib"],
)
