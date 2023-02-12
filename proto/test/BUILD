load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "test_message_proto",
    srcs = [
        "message.proto",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "@com_google_protobuf//:timestamp_proto",
    ],
)

go_proto_library(
    name = "test_message_go_proto",
    importpath = "bazel-go/protos/test",
    protos = [":test_message_proto"],
    visibility = ["//visibility:public"],
)

# filegroup(
    # name = "test_proto_files",
    # srcs = glob(["*.proto"]),
    # visibility = [
        # "//visibility:public",
    # ],
# )