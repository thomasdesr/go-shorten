load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_prefix")
load("@bazel_tools//tools/build_defs/pkg:pkg.bzl", "pkg_tar")

go_prefix("github.com/thomaso-mirodin/go-shorten")

go_binary(
    name = "go-shorten",
    library = ":go_default_library",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
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
        "//storage:go_default_library",
        "//storage/multistorage:go_default_library",
        "//vendor:github.com/codegangsta/negroni",
        "//vendor:github.com/google/shlex",
        "//vendor:github.com/jessevdk/go-flags",
        "//vendor:github.com/julienschmidt/httprouter",
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

filegroup(
    name = "static-src",
    srcs = glob(["static/**"]),
    visibility = ["//visibility:public"],
)

pkg_tar(
    name = "static-pkg",
    files = [":static-src"],
    strip_prefix = "/",
    visibility = ["//visibility:public"],
)

pkg_tar(
    name = "go-shorten-pkg",
    files = [":go-shorten"],
    package_dir = "go-shorten",
    strip_prefix = ".",
    deps = [":static-pkg"],
)
