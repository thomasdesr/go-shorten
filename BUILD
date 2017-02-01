load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_prefix")

go_prefix("github.com/thomaso-mirodin/go-shorten")

go_binary(
    name = "go-shorten",
    library = ":go_default_library",
    tags = ["automanaged"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "main.go",
        "options.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//handlers:go_default_library",
        "//handlers/templates:go_default_library",
        "//storage:go_default_library",
        "//storage/multistorage:go_default_library",
        "//vendor:github.com/GeertJohan/go.rice",
        "//vendor:github.com/codegangsta/negroni",
        "//vendor:github.com/google/shlex",
        "//vendor:github.com/guregu/kami",
        "//vendor:github.com/jessevdk/go-flags",
        "//vendor:github.com/pkg/errors",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"], exclude=["bazel-*/**", ".git/**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//handlers:all-srcs",
        "//storage:all-srcs",
        "//vendor:all-srcs",
    ],
    tags = ["automanaged"],
)
