package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_test(
    name = "go_default_test",
    srcs = [
        "regex_test.go",
        "replacer_test.go",
        "storage_internal_test.go",
    ],
    library = ":go_default_library",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/pkg/errors",
        "//vendor:github.com/stretchr/testify/assert",
        "//vendor:github.com/stretchr/testify/require",
    ],
)

go_library(
    name = "go_default_library",
    srcs = [
        "filesystem.go",
        "inmem.go",
        "regex.go",
        "s3.go",
        "storage.go",
        "utils.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/s3",
        "//vendor:github.com/pkg/errors",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "filesystem_test.go",
        "inmem_test.go",
        "s3_test.go",
        "storage_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//storage:go_default_library",
        "//storage/migrations:go_default_library",
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/s3",
        "//vendor:github.com/pkg/errors",
        "//vendor:github.com/stretchr/testify/assert",
        "//vendor:github.com/stretchr/testify/require",
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
        "//storage/migrations:all-srcs",
        "//storage/multistorage:all-srcs",
    ],
    tags = ["automanaged"],
)
