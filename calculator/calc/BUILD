load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "calc",
    srcs = ["calc.go"],
    importpath = "bazel-go/calculator/calc",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_dlshle_gommon//async",
        "//proto/test:test_message_go_proto"
    ],
)
