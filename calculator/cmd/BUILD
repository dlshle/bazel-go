load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "hello_lib",
    srcs = ["main.go"],
    importpath = "bazel-go/calculator/cmd",
    visibility = ["//visibility:private"],
    deps = [
        # use generated dependency from //external_dependencies
        "@com_github_dlshle_gommon//async",
        "//calculator/calc",
        "//proto/test:test_message_go_proto"
    ],
)

go_binary(
    name = "hello",
    embed = [":hello_lib"],
    visibility = ["//visibility:public"],
)
