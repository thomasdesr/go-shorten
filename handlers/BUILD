package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "api.go",
        "static.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//handlers/templates:go_default_library",
        "//storage:go_default_library",
        "//vendor:github.com/GeertJohan/go.rice",
        "//vendor:github.com/guregu/kami",
        "//vendor:golang.org/x/net/context",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//handlers/templates:all-srcs",
    ],
    tags = ["automanaged"],
)
