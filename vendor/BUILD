package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
    "go_test",
)

go_test(
    name = "github.com/GeertJohan/go.incremental_test",
    srcs = [
        "github.com/GeertJohan/go.incremental/int16_test.go",
        "github.com/GeertJohan/go.incremental/int32_test.go",
        "github.com/GeertJohan/go.incremental/int64_test.go",
        "github.com/GeertJohan/go.incremental/int8_test.go",
        "github.com/GeertJohan/go.incremental/int_test.go",
        "github.com/GeertJohan/go.incremental/uint16_test.go",
        "github.com/GeertJohan/go.incremental/uint32_test.go",
        "github.com/GeertJohan/go.incremental/uint64_test.go",
        "github.com/GeertJohan/go.incremental/uint8_test.go",
        "github.com/GeertJohan/go.incremental/uint_test.go",
    ],
    library = ":github.com/GeertJohan/go.incremental",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/GeertJohan/go.incremental",
    srcs = [
        "github.com/GeertJohan/go.incremental/doc.go",
        "github.com/GeertJohan/go.incremental/int.go",
        "github.com/GeertJohan/go.incremental/int16.go",
        "github.com/GeertJohan/go.incremental/int32.go",
        "github.com/GeertJohan/go.incremental/int64.go",
        "github.com/GeertJohan/go.incremental/int8.go",
        "github.com/GeertJohan/go.incremental/uint.go",
        "github.com/GeertJohan/go.incremental/uint16.go",
        "github.com/GeertJohan/go.incremental/uint32.go",
        "github.com/GeertJohan/go.incremental/uint64.go",
        "github.com/GeertJohan/go.incremental/uint8.go",
    ],
    tags = ["automanaged"],
)

go_binary(
    name = "github.com/GeertJohan/go.incremental/example_bin",
    library = ":github.com/GeertJohan/go.incremental/example",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/GeertJohan/go.incremental/example",
    srcs = ["github.com/GeertJohan/go.incremental/example/example.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/GeertJohan/go.incremental"],
)

go_binary(
    name = "github.com/GeertJohan/go.incremental/gen_bin",
    library = ":github.com/GeertJohan/go.incremental/gen",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/GeertJohan/go.incremental/gen",
    srcs = ["github.com/GeertJohan/go.incremental/gen/generator.go"],
    tags = ["automanaged"],
)

go_test(
    name = "github.com/GeertJohan/go.rice_test",
    srcs = ["github.com/GeertJohan/go.rice/config_test.go"],
    library = ":github.com/GeertJohan/go.rice",
    tags = ["automanaged"],
    deps = ["//vendor:github.com/GeertJohan/go.rice/embedded"],
)

go_library(
    name = "github.com/GeertJohan/go.rice",
    srcs = [
        "github.com/GeertJohan/go.rice/appended.go",
        "github.com/GeertJohan/go.rice/box.go",
        "github.com/GeertJohan/go.rice/config.go",
        "github.com/GeertJohan/go.rice/debug.go",
        "github.com/GeertJohan/go.rice/embedded.go",
        "github.com/GeertJohan/go.rice/file.go",
        "github.com/GeertJohan/go.rice/http.go",
        "github.com/GeertJohan/go.rice/sort.go",
        "github.com/GeertJohan/go.rice/virtual.go",
        "github.com/GeertJohan/go.rice/walk.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/GeertJohan/go.rice/embedded",
        "//vendor:github.com/daaku/go.zipexe",
        "//vendor:github.com/kardianos/osext",
    ],
)

go_library(
    name = "github.com/GeertJohan/go.rice/embedded",
    srcs = ["github.com/GeertJohan/go.rice/embedded/embedded.go"],
    tags = ["automanaged"],
)

go_binary(
    name = "github.com/GeertJohan/go.rice/example_bin",
    library = ":github.com/GeertJohan/go.rice/example",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/GeertJohan/go.rice/example",
    srcs = ["github.com/GeertJohan/go.rice/example/example.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/GeertJohan/go.rice",
        "//vendor:github.com/davecgh/go-spew/spew",
    ],
)

go_binary(
    name = "github.com/GeertJohan/go.rice/rice_bin",
    library = ":github.com/GeertJohan/go.rice/rice",
    tags = ["automanaged"],
)

go_test(
    name = "github.com/GeertJohan/go.rice/rice_test",
    srcs = ["github.com/GeertJohan/go.rice/rice/find_test.go"],
    library = ":github.com/GeertJohan/go.rice/rice",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/GeertJohan/go.rice/rice",
    srcs = [
        "github.com/GeertJohan/go.rice/rice/append.go",
        "github.com/GeertJohan/go.rice/rice/clean.go",
        "github.com/GeertJohan/go.rice/rice/embed-go.go",
        "github.com/GeertJohan/go.rice/rice/embed-syso.go",
        "github.com/GeertJohan/go.rice/rice/find.go",
        "github.com/GeertJohan/go.rice/rice/flags.go",
        "github.com/GeertJohan/go.rice/rice/identifier.go",
        "github.com/GeertJohan/go.rice/rice/main.go",
        "github.com/GeertJohan/go.rice/rice/templates.go",
        "github.com/GeertJohan/go.rice/rice/util.go",
        "github.com/GeertJohan/go.rice/rice/writecoff.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/GeertJohan/go.incremental",
        "//vendor:github.com/GeertJohan/go.rice/embedded",
        "//vendor:github.com/akavel/rsrc/binutil",
        "//vendor:github.com/akavel/rsrc/coff",
        "//vendor:github.com/daaku/go.zipexe",
        "//vendor:github.com/jessevdk/go-flags",
    ],
)

go_library(
    name = "github.com/akavel/rsrc/binutil",
    srcs = [
        "github.com/akavel/rsrc/binutil/plain.go",
        "github.com/akavel/rsrc/binutil/sizedfile.go",
        "github.com/akavel/rsrc/binutil/walk.go",
        "github.com/akavel/rsrc/binutil/writer.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/akavel/rsrc/coff",
    srcs = ["github.com/akavel/rsrc/coff/coff.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/akavel/rsrc/binutil"],
)

go_test(
    name = "github.com/aws/aws-sdk-go/aws_test",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/config_test.go",
        "github.com/aws/aws-sdk-go/aws/convert_types_test.go",
        "github.com/aws/aws-sdk-go/aws/types_test.go",
    ],
    library = ":github.com/aws/aws-sdk-go/aws",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/aws",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/config.go",
        "github.com/aws/aws-sdk-go/aws/convert_types.go",
        "github.com/aws/aws-sdk-go/aws/errors.go",
        "github.com/aws/aws-sdk-go/aws/logger.go",
        "github.com/aws/aws-sdk-go/aws/types.go",
        "github.com/aws/aws-sdk-go/aws/version.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor:github.com/aws/aws-sdk-go/aws/endpoints",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/aws/awserr",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/awserr/error.go",
        "github.com/aws/aws-sdk-go/aws/awserr/types.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/aws/aws-sdk-go/aws/awsutil",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/awsutil/copy.go",
        "github.com/aws/aws-sdk-go/aws/awsutil/equal.go",
        "github.com/aws/aws-sdk-go/aws/awsutil/path_value.go",
        "github.com/aws/aws-sdk-go/aws/awsutil/prettify.go",
        "github.com/aws/aws-sdk-go/aws/awsutil/string_value.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/jmespath/go-jmespath"],
)

go_test(
    name = "github.com/aws/aws-sdk-go/aws/awsutil_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/awsutil/copy_test.go",
        "github.com/aws/aws-sdk-go/aws/awsutil/equal_test.go",
        "github.com/aws/aws-sdk-go/aws/awsutil/path_value_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/aws/client",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/client/client.go",
        "github.com/aws/aws-sdk-go/aws/client/default_retryer.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/aws/client/metadata",
    srcs = ["github.com/aws/aws-sdk-go/aws/client/metadata/client_info.go"],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/aws/aws-sdk-go/aws/corehandlers",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/corehandlers/handlers.go",
        "github.com/aws/aws-sdk-go/aws/corehandlers/param_validator.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/aws/corehandlers_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/corehandlers/handlers_test.go",
        "github.com/aws/aws-sdk-go/aws/corehandlers/param_validator_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/corehandlers",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/awstesting",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/service/kinesis",
        "//vendor:github.com/aws/aws-sdk-go/service/s3",
        "//vendor:github.com/stretchr/testify/assert",
        "//vendor:github.com/stretchr/testify/require",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/aws/credentials_test",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/credentials/chain_provider_test.go",
        "github.com/aws/aws-sdk-go/aws/credentials/credentials_test.go",
        "github.com/aws/aws-sdk-go/aws/credentials/env_provider_test.go",
        "github.com/aws/aws-sdk-go/aws/credentials/shared_credentials_provider_test.go",
        "github.com/aws/aws-sdk-go/aws/credentials/static_provider_test.go",
    ],
    library = ":github.com/aws/aws-sdk-go/aws/credentials",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/aws/credentials",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/credentials/chain_provider.go",
        "github.com/aws/aws-sdk-go/aws/credentials/credentials.go",
        "github.com/aws/aws-sdk-go/aws/credentials/env_provider.go",
        "github.com/aws/aws-sdk-go/aws/credentials/shared_credentials_provider.go",
        "github.com/aws/aws-sdk-go/aws/credentials/static_provider.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/go-ini/ini",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds",
    srcs = ["github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds/ec2_role_provider.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor:github.com/aws/aws-sdk-go/aws/ec2metadata",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds_xtest",
    srcs = ["github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds/ec2_role_provider_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds",
        "//vendor:github.com/aws/aws-sdk-go/aws/ec2metadata",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/aws/credentials/endpointcreds",
    srcs = ["github.com/aws/aws-sdk-go/aws/credentials/endpointcreds/provider.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/aws/credentials/endpointcreds_xtest",
    srcs = ["github.com/aws/aws-sdk-go/aws/credentials/endpointcreds/provider_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials/endpointcreds",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/aws/credentials/stscreds_test",
    srcs = ["github.com/aws/aws-sdk-go/aws/credentials/stscreds/assume_role_provider_test.go"],
    library = ":github.com/aws/aws-sdk-go/aws/credentials/stscreds",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/service/sts",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/aws/credentials/stscreds",
    srcs = ["github.com/aws/aws-sdk-go/aws/credentials/stscreds/assume_role_provider.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor:github.com/aws/aws-sdk-go/service/sts",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/aws/defaults_test",
    srcs = ["github.com/aws/aws-sdk-go/aws/defaults/defaults_test.go"],
    library = ":github.com/aws/aws-sdk-go/aws/defaults",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials/endpointcreds",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/aws/defaults",
    srcs = ["github.com/aws/aws-sdk-go/aws/defaults/defaults.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/corehandlers",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials/endpointcreds",
        "//vendor:github.com/aws/aws-sdk-go/aws/ec2metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/endpoints",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/aws/ec2metadata",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/ec2metadata/api.go",
        "github.com/aws/aws-sdk-go/aws/ec2metadata/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/aws/ec2metadata_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/ec2metadata/api_test.go",
        "github.com/aws/aws-sdk-go/aws/ec2metadata/service_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/ec2metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/aws/endpoints_test",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/endpoints/decode_test.go",
        "github.com/aws/aws-sdk-go/aws/endpoints/endpoints_test.go",
        "github.com/aws/aws-sdk-go/aws/endpoints/v3model_test.go",
    ],
    library = ":github.com/aws/aws-sdk-go/aws/endpoints",
    tags = ["automanaged"],
    deps = ["//vendor:github.com/stretchr/testify/assert"],
)

go_library(
    name = "github.com/aws/aws-sdk-go/aws/endpoints",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/endpoints/decode.go",
        "github.com/aws/aws-sdk-go/aws/endpoints/defaults.go",
        "github.com/aws/aws-sdk-go/aws/endpoints/doc.go",
        "github.com/aws/aws-sdk-go/aws/endpoints/endpoints.go",
        "github.com/aws/aws-sdk-go/aws/endpoints/v3model.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/aws/aws-sdk-go/aws/awserr"],
)

go_test(
    name = "github.com/aws/aws-sdk-go/aws/endpoints_xtest",
    srcs = ["github.com/aws/aws-sdk-go/aws/endpoints/example_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/endpoints",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/s3",
        "//vendor:github.com/aws/aws-sdk-go/service/sqs",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/aws/request_test",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/request/http_request_copy_test.go",
        "github.com/aws/aws-sdk-go/aws/request/offset_reader_test.go",
        "github.com/aws/aws-sdk-go/aws/request/request_1_7_test.go",
        "github.com/aws/aws-sdk-go/aws/request/request_internal_test.go",
        "github.com/aws/aws-sdk-go/aws/request/request_resetbody_test.go",
        "github.com/aws/aws-sdk-go/aws/request/retryer_test.go",
    ],
    library = ":github.com/aws/aws-sdk-go/aws/request",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/aws/request",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/request/handlers.go",
        "github.com/aws/aws-sdk-go/aws/request/http_request.go",
        "github.com/aws/aws-sdk-go/aws/request/offset_reader.go",
        "github.com/aws/aws-sdk-go/aws/request/request.go",
        "github.com/aws/aws-sdk-go/aws/request/request_1_7.go",
        "github.com/aws/aws-sdk-go/aws/request/request_pagination.go",
        "github.com/aws/aws-sdk-go/aws/request/retryer.go",
        "github.com/aws/aws-sdk-go/aws/request/validation.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/aws/request_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/request/handlers_test.go",
        "github.com/aws/aws-sdk-go/aws/request/http_request_retry_test.go",
        "github.com/aws/aws-sdk-go/aws/request/request_1_6_test.go",
        "github.com/aws/aws-sdk-go/aws/request/request_pagination_test.go",
        "github.com/aws/aws-sdk-go/aws/request/request_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor:github.com/aws/aws-sdk-go/aws/defaults",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/awstesting",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/mock",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/rest",
        "//vendor:github.com/aws/aws-sdk-go/service/dynamodb",
        "//vendor:github.com/aws/aws-sdk-go/service/route53",
        "//vendor:github.com/aws/aws-sdk-go/service/s3",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/aws/session_test",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/session/env_config_test.go",
        "github.com/aws/aws-sdk-go/aws/session/session_test.go",
        "github.com/aws/aws-sdk-go/aws/session/shared_config_test.go",
    ],
    library = ":github.com/aws/aws-sdk-go/aws/session",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor:github.com/aws/aws-sdk-go/aws/defaults",
        "//vendor:github.com/aws/aws-sdk-go/service/s3",
        "//vendor:github.com/go-ini/ini",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/aws/session",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/session/doc.go",
        "github.com/aws/aws-sdk-go/aws/session/env_config.go",
        "github.com/aws/aws-sdk-go/aws/session/session.go",
        "github.com/aws/aws-sdk-go/aws/session/shared_config.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/corehandlers",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials/stscreds",
        "//vendor:github.com/aws/aws-sdk-go/aws/defaults",
        "//vendor:github.com/aws/aws-sdk-go/aws/endpoints",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/go-ini/ini",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/aws/signer/v4_test",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/signer/v4/header_rules_test.go",
        "github.com/aws/aws-sdk-go/aws/signer/v4/v4_test.go",
    ],
    library = ":github.com/aws/aws-sdk-go/aws/signer/v4",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/awstesting",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/aws/signer/v4",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/signer/v4/header_rules.go",
        "github.com/aws/aws-sdk-go/aws/signer/v4/uri_path.go",
        "github.com/aws/aws-sdk-go/aws/signer/v4/v4.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/rest",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/aws/signer/v4_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/signer/v4/functional_1_5_test.go",
        "github.com/aws/aws-sdk-go/aws/signer/v4/functional_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/service/s3",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/awstesting",
    srcs = [
        "github.com/aws/aws-sdk-go/awstesting/assert.go",
        "github.com/aws/aws-sdk-go/awstesting/client.go",
        "github.com/aws/aws-sdk-go/awstesting/util.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/defaults",
        "//vendor:github.com/aws/aws-sdk-go/private/util",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/awstesting_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/awstesting/assert_test.go",
        "github.com/aws/aws-sdk-go/awstesting/util_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/awstesting",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/awstesting/integration/customizations/s3",
    srcs = ["github.com/aws/aws-sdk-go/awstesting/integration/customizations/s3/stub.go"],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/aws/aws-sdk-go/awstesting/integration/customizations/s3/s3manager",
    srcs = ["github.com/aws/aws-sdk-go/awstesting/integration/customizations/s3/s3manager/stub.go"],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/aws/aws-sdk-go/awstesting/mock",
    srcs = ["github.com/aws/aws-sdk-go/awstesting/mock/mock.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/awstesting/unit",
    srcs = ["github.com/aws/aws-sdk-go/awstesting/unit/unit.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/private/protocol",
    srcs = [
        "github.com/aws/aws-sdk-go/private/protocol/idempotency.go",
        "github.com/aws/aws-sdk-go/private/protocol/unmarshal.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/aws/aws-sdk-go/aws/request"],
)

go_test(
    name = "github.com/aws/aws-sdk-go/private/protocol_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/private/protocol/idempotency_test.go",
        "github.com/aws/aws-sdk-go/private/protocol/protocol_test.go",
        "github.com/aws/aws-sdk-go/private/protocol/unmarshal_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/awstesting",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/ec2query",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/query",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/rest",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/restjson",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/restxml",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/private/protocol/ec2query",
    srcs = [
        "github.com/aws/aws-sdk-go/private/protocol/ec2query/build.go",
        "github.com/aws/aws-sdk-go/private/protocol/ec2query/unmarshal.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/query/queryutil",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/private/protocol/ec2query_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/private/protocol/ec2query/build_test.go",
        "github.com/aws/aws-sdk-go/private/protocol/ec2query/unmarshal_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/awstesting",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/ec2query",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil",
        "//vendor:github.com/aws/aws-sdk-go/private/util",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/private/protocol/json/jsonutil",
    srcs = [
        "github.com/aws/aws-sdk-go/private/protocol/json/jsonutil/build.go",
        "github.com/aws/aws-sdk-go/private/protocol/json/jsonutil/unmarshal.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/aws/aws-sdk-go/private/protocol"],
)

go_test(
    name = "github.com/aws/aws-sdk-go/private/protocol/json/jsonutil_xtest",
    srcs = ["github.com/aws/aws-sdk-go/private/protocol/json/jsonutil/build_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/json/jsonutil",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    srcs = ["github.com/aws/aws-sdk-go/private/protocol/jsonrpc/jsonrpc.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/json/jsonutil",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/rest",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/private/protocol/jsonrpc_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/private/protocol/jsonrpc/build_test.go",
        "github.com/aws/aws-sdk-go/private/protocol/jsonrpc/unmarshal_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/awstesting",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil",
        "//vendor:github.com/aws/aws-sdk-go/private/util",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/private/protocol/query",
    srcs = [
        "github.com/aws/aws-sdk-go/private/protocol/query/build.go",
        "github.com/aws/aws-sdk-go/private/protocol/query/unmarshal.go",
        "github.com/aws/aws-sdk-go/private/protocol/query/unmarshal_error.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/query/queryutil",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/private/protocol/query_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/private/protocol/query/build_test.go",
        "github.com/aws/aws-sdk-go/private/protocol/query/unmarshal_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/awstesting",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/query",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil",
        "//vendor:github.com/aws/aws-sdk-go/private/util",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/private/protocol/query/queryutil",
    srcs = ["github.com/aws/aws-sdk-go/private/protocol/query/queryutil/queryutil.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/aws/aws-sdk-go/private/protocol"],
)

go_test(
    name = "github.com/aws/aws-sdk-go/private/protocol/rest_test",
    srcs = ["github.com/aws/aws-sdk-go/private/protocol/rest/build_test.go"],
    library = ":github.com/aws/aws-sdk-go/private/protocol/rest",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/private/protocol/rest",
    srcs = [
        "github.com/aws/aws-sdk-go/private/protocol/rest/build.go",
        "github.com/aws/aws-sdk-go/private/protocol/rest/payload.go",
        "github.com/aws/aws-sdk-go/private/protocol/rest/unmarshal.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/private/protocol/restjson",
    srcs = ["github.com/aws/aws-sdk-go/private/protocol/restjson/restjson.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/rest",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/private/protocol/restjson_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/private/protocol/restjson/build_test.go",
        "github.com/aws/aws-sdk-go/private/protocol/restjson/unmarshal_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/awstesting",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/restjson",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil",
        "//vendor:github.com/aws/aws-sdk-go/private/util",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/private/protocol/restxml",
    srcs = ["github.com/aws/aws-sdk-go/private/protocol/restxml/restxml.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/query",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/rest",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/private/protocol/restxml_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/private/protocol/restxml/build_test.go",
        "github.com/aws/aws-sdk-go/private/protocol/restxml/unmarshal_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/awstesting",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/restxml",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil",
        "//vendor:github.com/aws/aws-sdk-go/private/util",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil",
    srcs = [
        "github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil/build.go",
        "github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil/unmarshal.go",
        "github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil/xml_to_struct.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/aws/aws-sdk-go/private/protocol"],
)

go_test(
    name = "github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil_xtest",
    srcs = ["github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil/xml_to_struct_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/service/s3",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/private/signer/v2_test",
    srcs = ["github.com/aws/aws-sdk-go/private/signer/v2/v2_test.go"],
    library = ":github.com/aws/aws-sdk-go/private/signer/v2",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/awstesting",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/private/signer/v2",
    srcs = ["github.com/aws/aws-sdk-go/private/signer/v2/v2.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/private/util",
    srcs = [
        "github.com/aws/aws-sdk-go/private/util/sort_keys.go",
        "github.com/aws/aws-sdk-go/private/util/util.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil"],
)

go_library(
    name = "github.com/aws/aws-sdk-go/private/waiter",
    srcs = ["github.com/aws/aws-sdk-go/private/waiter/waiter.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/private/waiter_xtest",
    srcs = ["github.com/aws/aws-sdk-go/private/waiter/waiter_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/awstesting",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/acm",
    srcs = [
        "github.com/aws/aws-sdk-go/service/acm/api.go",
        "github.com/aws/aws-sdk-go/service/acm/errors.go",
        "github.com/aws/aws-sdk-go/service/acm/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/acm_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/acm/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/acm",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/acm/acmiface",
    srcs = ["github.com/aws/aws-sdk-go/service/acm/acmiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/acm",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/apigateway",
    srcs = [
        "github.com/aws/aws-sdk-go/service/apigateway/api.go",
        "github.com/aws/aws-sdk-go/service/apigateway/customization.go",
        "github.com/aws/aws-sdk-go/service/apigateway/errors.go",
        "github.com/aws/aws-sdk-go/service/apigateway/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/restjson",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/apigateway_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/apigateway/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/apigateway",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/apigateway/apigatewayiface",
    srcs = ["github.com/aws/aws-sdk-go/service/apigateway/apigatewayiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/apigateway",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/applicationdiscoveryservice",
    srcs = [
        "github.com/aws/aws-sdk-go/service/applicationdiscoveryservice/api.go",
        "github.com/aws/aws-sdk-go/service/applicationdiscoveryservice/errors.go",
        "github.com/aws/aws-sdk-go/service/applicationdiscoveryservice/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/applicationdiscoveryservice_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/applicationdiscoveryservice/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/applicationdiscoveryservice",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/applicationdiscoveryservice/applicationdiscoveryserviceiface",
    srcs = ["github.com/aws/aws-sdk-go/service/applicationdiscoveryservice/applicationdiscoveryserviceiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/applicationdiscoveryservice",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/autoscaling",
    srcs = [
        "github.com/aws/aws-sdk-go/service/autoscaling/api.go",
        "github.com/aws/aws-sdk-go/service/autoscaling/errors.go",
        "github.com/aws/aws-sdk-go/service/autoscaling/service.go",
        "github.com/aws/aws-sdk-go/service/autoscaling/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/query",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/autoscaling_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/autoscaling/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/autoscaling",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface",
    srcs = ["github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/autoscaling",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudformation",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cloudformation/api.go",
        "github.com/aws/aws-sdk-go/service/cloudformation/errors.go",
        "github.com/aws/aws-sdk-go/service/cloudformation/service.go",
        "github.com/aws/aws-sdk-go/service/cloudformation/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/query",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/cloudformation_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudformation/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudformation",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudformation",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudfront",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cloudfront/api.go",
        "github.com/aws/aws-sdk-go/service/cloudfront/errors.go",
        "github.com/aws/aws-sdk-go/service/cloudfront/service.go",
        "github.com/aws/aws-sdk-go/service/cloudfront/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/restxml",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/cloudfront_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudfront/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudfront",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudfront/cloudfrontiface",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudfront/cloudfrontiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudfront",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/cloudfront/sign_test",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cloudfront/sign/policy_test.go",
        "github.com/aws/aws-sdk-go/service/cloudfront/sign/privkey_test.go",
        "github.com/aws/aws-sdk-go/service/cloudfront/sign/sign_cookie_example_test.go",
        "github.com/aws/aws-sdk-go/service/cloudfront/sign/sign_cookie_test.go",
        "github.com/aws/aws-sdk-go/service/cloudfront/sign/sign_url_test.go",
    ],
    library = ":github.com/aws/aws-sdk-go/service/cloudfront/sign",
    tags = ["automanaged"],
    deps = ["//vendor:github.com/stretchr/testify/assert"],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudfront/sign",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cloudfront/sign/policy.go",
        "github.com/aws/aws-sdk-go/service/cloudfront/sign/privkey.go",
        "github.com/aws/aws-sdk-go/service/cloudfront/sign/randomreader.go",
        "github.com/aws/aws-sdk-go/service/cloudfront/sign/sign_cookie.go",
        "github.com/aws/aws-sdk-go/service/cloudfront/sign/sign_url.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudhsm",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cloudhsm/api.go",
        "github.com/aws/aws-sdk-go/service/cloudhsm/errors.go",
        "github.com/aws/aws-sdk-go/service/cloudhsm/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/cloudhsm_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudhsm/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudhsm",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudhsm/cloudhsmiface",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudhsm/cloudhsmiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudhsm",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudsearch",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cloudsearch/api.go",
        "github.com/aws/aws-sdk-go/service/cloudsearch/errors.go",
        "github.com/aws/aws-sdk-go/service/cloudsearch/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/query",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/cloudsearch_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudsearch/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudsearch",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudsearch/cloudsearchiface",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudsearch/cloudsearchiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudsearch",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudsearchdomain",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cloudsearchdomain/api.go",
        "github.com/aws/aws-sdk-go/service/cloudsearchdomain/errors.go",
        "github.com/aws/aws-sdk-go/service/cloudsearchdomain/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/restjson",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/cloudsearchdomain_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cloudsearchdomain/customizations_test.go",
        "github.com/aws/aws-sdk-go/service/cloudsearchdomain/examples_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudsearchdomain",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudsearchdomain/cloudsearchdomainiface",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudsearchdomain/cloudsearchdomainiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudsearchdomain",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudtrail",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cloudtrail/api.go",
        "github.com/aws/aws-sdk-go/service/cloudtrail/errors.go",
        "github.com/aws/aws-sdk-go/service/cloudtrail/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/cloudtrail_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudtrail/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudtrail",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudtrail/cloudtrailiface",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudtrail/cloudtrailiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudtrail",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudwatch",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cloudwatch/api.go",
        "github.com/aws/aws-sdk-go/service/cloudwatch/errors.go",
        "github.com/aws/aws-sdk-go/service/cloudwatch/service.go",
        "github.com/aws/aws-sdk-go/service/cloudwatch/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/query",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/cloudwatch_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudwatch/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudwatch",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudwatch",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudwatchevents",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cloudwatchevents/api.go",
        "github.com/aws/aws-sdk-go/service/cloudwatchevents/errors.go",
        "github.com/aws/aws-sdk-go/service/cloudwatchevents/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/cloudwatchevents_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudwatchevents/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudwatchevents",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudwatchevents/cloudwatcheventsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudwatchevents/cloudwatcheventsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudwatchevents",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudwatchlogs",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cloudwatchlogs/api.go",
        "github.com/aws/aws-sdk-go/service/cloudwatchlogs/errors.go",
        "github.com/aws/aws-sdk-go/service/cloudwatchlogs/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/cloudwatchlogs_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudwatchlogs/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudwatchlogs",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudwatchlogs",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/codecommit",
    srcs = [
        "github.com/aws/aws-sdk-go/service/codecommit/api.go",
        "github.com/aws/aws-sdk-go/service/codecommit/errors.go",
        "github.com/aws/aws-sdk-go/service/codecommit/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/codecommit_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/codecommit/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/codecommit",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/codecommit/codecommitiface",
    srcs = ["github.com/aws/aws-sdk-go/service/codecommit/codecommitiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/codecommit",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/codedeploy",
    srcs = [
        "github.com/aws/aws-sdk-go/service/codedeploy/api.go",
        "github.com/aws/aws-sdk-go/service/codedeploy/errors.go",
        "github.com/aws/aws-sdk-go/service/codedeploy/service.go",
        "github.com/aws/aws-sdk-go/service/codedeploy/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/codedeploy_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/codedeploy/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/codedeploy",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/codedeploy/codedeployiface",
    srcs = ["github.com/aws/aws-sdk-go/service/codedeploy/codedeployiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/codedeploy",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/codepipeline",
    srcs = [
        "github.com/aws/aws-sdk-go/service/codepipeline/api.go",
        "github.com/aws/aws-sdk-go/service/codepipeline/errors.go",
        "github.com/aws/aws-sdk-go/service/codepipeline/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/codepipeline_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/codepipeline/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/codepipeline",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/codepipeline/codepipelineiface",
    srcs = ["github.com/aws/aws-sdk-go/service/codepipeline/codepipelineiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/codepipeline",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cognitoidentity",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cognitoidentity/api.go",
        "github.com/aws/aws-sdk-go/service/cognitoidentity/customizations.go",
        "github.com/aws/aws-sdk-go/service/cognitoidentity/errors.go",
        "github.com/aws/aws-sdk-go/service/cognitoidentity/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/cognitoidentity_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cognitoidentity/customizations_test.go",
        "github.com/aws/aws-sdk-go/service/cognitoidentity/examples_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/service/cognitoidentity",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cognitoidentity/cognitoidentityiface",
    srcs = ["github.com/aws/aws-sdk-go/service/cognitoidentity/cognitoidentityiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cognitoidentity",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cognitosync",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cognitosync/api.go",
        "github.com/aws/aws-sdk-go/service/cognitosync/errors.go",
        "github.com/aws/aws-sdk-go/service/cognitosync/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/restjson",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/cognitosync_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/cognitosync/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/cognitosync",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cognitosync/cognitosynciface",
    srcs = ["github.com/aws/aws-sdk-go/service/cognitosync/cognitosynciface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cognitosync",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/configservice",
    srcs = [
        "github.com/aws/aws-sdk-go/service/configservice/api.go",
        "github.com/aws/aws-sdk-go/service/configservice/errors.go",
        "github.com/aws/aws-sdk-go/service/configservice/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/configservice_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/configservice/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/configservice",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/configservice/configserviceiface",
    srcs = ["github.com/aws/aws-sdk-go/service/configservice/configserviceiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/configservice",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/datapipeline",
    srcs = [
        "github.com/aws/aws-sdk-go/service/datapipeline/api.go",
        "github.com/aws/aws-sdk-go/service/datapipeline/errors.go",
        "github.com/aws/aws-sdk-go/service/datapipeline/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/datapipeline_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/datapipeline/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/datapipeline",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/datapipeline/datapipelineiface",
    srcs = ["github.com/aws/aws-sdk-go/service/datapipeline/datapipelineiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/datapipeline",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/devicefarm",
    srcs = [
        "github.com/aws/aws-sdk-go/service/devicefarm/api.go",
        "github.com/aws/aws-sdk-go/service/devicefarm/errors.go",
        "github.com/aws/aws-sdk-go/service/devicefarm/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/devicefarm_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/devicefarm/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/devicefarm",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/devicefarm/devicefarmiface",
    srcs = ["github.com/aws/aws-sdk-go/service/devicefarm/devicefarmiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/devicefarm",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/directconnect",
    srcs = [
        "github.com/aws/aws-sdk-go/service/directconnect/api.go",
        "github.com/aws/aws-sdk-go/service/directconnect/errors.go",
        "github.com/aws/aws-sdk-go/service/directconnect/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/directconnect_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/directconnect/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/directconnect",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/directconnect/directconnectiface",
    srcs = ["github.com/aws/aws-sdk-go/service/directconnect/directconnectiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/directconnect",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/directoryservice",
    srcs = [
        "github.com/aws/aws-sdk-go/service/directoryservice/api.go",
        "github.com/aws/aws-sdk-go/service/directoryservice/errors.go",
        "github.com/aws/aws-sdk-go/service/directoryservice/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/directoryservice_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/directoryservice/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/directoryservice",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/directoryservice/directoryserviceiface",
    srcs = ["github.com/aws/aws-sdk-go/service/directoryservice/directoryserviceiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/directoryservice",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/dynamodb",
    srcs = [
        "github.com/aws/aws-sdk-go/service/dynamodb/api.go",
        "github.com/aws/aws-sdk-go/service/dynamodb/customizations.go",
        "github.com/aws/aws-sdk-go/service/dynamodb/errors.go",
        "github.com/aws/aws-sdk-go/service/dynamodb/service.go",
        "github.com/aws/aws-sdk-go/service/dynamodb/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/dynamodb_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/service/dynamodb/customizations_test.go",
        "github.com/aws/aws-sdk-go/service/dynamodb/examples_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/service/dynamodb",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute_test",
    srcs = [
        "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute/converter_test.go",
        "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute/decode_test.go",
        "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute/encode_test.go",
        "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute/field_test.go",
        "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute/marshaler_test.go",
        "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute/shared_test.go",
        "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute/tag_test.go",
    ],
    library = ":github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/service/dynamodb",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute",
    srcs = [
        "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute/converter.go",
        "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute/decode.go",
        "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute/doc.go",
        "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute/encode.go",
        "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute/field.go",
        "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute/tag.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/service/dynamodb",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute/converter_examples_test.go",
        "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute/marshaler_examples_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/service/dynamodb",
        "//vendor:github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface",
    srcs = ["github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/dynamodb",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/dynamodbstreams",
    srcs = [
        "github.com/aws/aws-sdk-go/service/dynamodbstreams/api.go",
        "github.com/aws/aws-sdk-go/service/dynamodbstreams/errors.go",
        "github.com/aws/aws-sdk-go/service/dynamodbstreams/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
        "//vendor:github.com/aws/aws-sdk-go/service/dynamodb",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/dynamodbstreams_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/dynamodbstreams/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/dynamodbstreams",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/dynamodbstreams/dynamodbstreamsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/dynamodbstreams/dynamodbstreamsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/dynamodbstreams",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/ec2",
    srcs = [
        "github.com/aws/aws-sdk-go/service/ec2/api.go",
        "github.com/aws/aws-sdk-go/service/ec2/customizations.go",
        "github.com/aws/aws-sdk-go/service/ec2/errors.go",
        "github.com/aws/aws-sdk-go/service/ec2/service.go",
        "github.com/aws/aws-sdk-go/service/ec2/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/endpoints",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/ec2query",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/ec2_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/service/ec2/customizations_test.go",
        "github.com/aws/aws-sdk-go/service/ec2/examples_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/service/ec2",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/ec2/ec2iface",
    srcs = ["github.com/aws/aws-sdk-go/service/ec2/ec2iface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/ec2",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/ecr",
    srcs = [
        "github.com/aws/aws-sdk-go/service/ecr/api.go",
        "github.com/aws/aws-sdk-go/service/ecr/errors.go",
        "github.com/aws/aws-sdk-go/service/ecr/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/ecr_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/ecr/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/ecr",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/ecr/ecriface",
    srcs = ["github.com/aws/aws-sdk-go/service/ecr/ecriface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/ecr",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/ecs",
    srcs = [
        "github.com/aws/aws-sdk-go/service/ecs/api.go",
        "github.com/aws/aws-sdk-go/service/ecs/errors.go",
        "github.com/aws/aws-sdk-go/service/ecs/service.go",
        "github.com/aws/aws-sdk-go/service/ecs/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/ecs_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/ecs/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/ecs",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/ecs/ecsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/ecs/ecsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/ecs",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/efs",
    srcs = [
        "github.com/aws/aws-sdk-go/service/efs/api.go",
        "github.com/aws/aws-sdk-go/service/efs/errors.go",
        "github.com/aws/aws-sdk-go/service/efs/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/restjson",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/efs_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/efs/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/efs",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/efs/efsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/efs/efsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/efs",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/elasticache",
    srcs = [
        "github.com/aws/aws-sdk-go/service/elasticache/api.go",
        "github.com/aws/aws-sdk-go/service/elasticache/errors.go",
        "github.com/aws/aws-sdk-go/service/elasticache/service.go",
        "github.com/aws/aws-sdk-go/service/elasticache/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/query",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/elasticache_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/elasticache/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/elasticache",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/elasticache/elasticacheiface",
    srcs = ["github.com/aws/aws-sdk-go/service/elasticache/elasticacheiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/elasticache",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/elasticbeanstalk",
    srcs = [
        "github.com/aws/aws-sdk-go/service/elasticbeanstalk/api.go",
        "github.com/aws/aws-sdk-go/service/elasticbeanstalk/errors.go",
        "github.com/aws/aws-sdk-go/service/elasticbeanstalk/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/query",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/elasticbeanstalk_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/elasticbeanstalk/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/elasticbeanstalk",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/elasticbeanstalk/elasticbeanstalkiface",
    srcs = ["github.com/aws/aws-sdk-go/service/elasticbeanstalk/elasticbeanstalkiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/elasticbeanstalk",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/elasticsearchservice",
    srcs = [
        "github.com/aws/aws-sdk-go/service/elasticsearchservice/api.go",
        "github.com/aws/aws-sdk-go/service/elasticsearchservice/errors.go",
        "github.com/aws/aws-sdk-go/service/elasticsearchservice/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/restjson",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/elasticsearchservice_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/elasticsearchservice/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/elasticsearchservice",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/elasticsearchservice/elasticsearchserviceiface",
    srcs = ["github.com/aws/aws-sdk-go/service/elasticsearchservice/elasticsearchserviceiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/elasticsearchservice",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/elastictranscoder",
    srcs = [
        "github.com/aws/aws-sdk-go/service/elastictranscoder/api.go",
        "github.com/aws/aws-sdk-go/service/elastictranscoder/errors.go",
        "github.com/aws/aws-sdk-go/service/elastictranscoder/service.go",
        "github.com/aws/aws-sdk-go/service/elastictranscoder/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/restjson",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/elastictranscoder_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/elastictranscoder/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/elastictranscoder",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/elastictranscoder/elastictranscoderiface",
    srcs = ["github.com/aws/aws-sdk-go/service/elastictranscoder/elastictranscoderiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/elastictranscoder",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/elb",
    srcs = [
        "github.com/aws/aws-sdk-go/service/elb/api.go",
        "github.com/aws/aws-sdk-go/service/elb/errors.go",
        "github.com/aws/aws-sdk-go/service/elb/service.go",
        "github.com/aws/aws-sdk-go/service/elb/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/query",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/elb_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/elb/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/elb",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/elb/elbiface",
    srcs = ["github.com/aws/aws-sdk-go/service/elb/elbiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/elb",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/emr",
    srcs = [
        "github.com/aws/aws-sdk-go/service/emr/api.go",
        "github.com/aws/aws-sdk-go/service/emr/errors.go",
        "github.com/aws/aws-sdk-go/service/emr/service.go",
        "github.com/aws/aws-sdk-go/service/emr/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/emr_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/emr/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/emr",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/emr/emriface",
    srcs = ["github.com/aws/aws-sdk-go/service/emr/emriface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/emr",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/firehose",
    srcs = [
        "github.com/aws/aws-sdk-go/service/firehose/api.go",
        "github.com/aws/aws-sdk-go/service/firehose/errors.go",
        "github.com/aws/aws-sdk-go/service/firehose/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/firehose_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/firehose/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/firehose",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/firehose/firehoseiface",
    srcs = ["github.com/aws/aws-sdk-go/service/firehose/firehoseiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/firehose",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/glacier",
    srcs = [
        "github.com/aws/aws-sdk-go/service/glacier/api.go",
        "github.com/aws/aws-sdk-go/service/glacier/customizations.go",
        "github.com/aws/aws-sdk-go/service/glacier/errors.go",
        "github.com/aws/aws-sdk-go/service/glacier/service.go",
        "github.com/aws/aws-sdk-go/service/glacier/treehash.go",
        "github.com/aws/aws-sdk-go/service/glacier/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/restjson",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/glacier_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/service/glacier/customizations_test.go",
        "github.com/aws/aws-sdk-go/service/glacier/examples_test.go",
        "github.com/aws/aws-sdk-go/service/glacier/treehash_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/service/glacier",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/glacier/glacieriface",
    srcs = ["github.com/aws/aws-sdk-go/service/glacier/glacieriface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/glacier",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/iam",
    srcs = [
        "github.com/aws/aws-sdk-go/service/iam/api.go",
        "github.com/aws/aws-sdk-go/service/iam/errors.go",
        "github.com/aws/aws-sdk-go/service/iam/service.go",
        "github.com/aws/aws-sdk-go/service/iam/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/query",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/iam_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/iam/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/iam",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/iam/iamiface",
    srcs = ["github.com/aws/aws-sdk-go/service/iam/iamiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/iam",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/inspector",
    srcs = [
        "github.com/aws/aws-sdk-go/service/inspector/api.go",
        "github.com/aws/aws-sdk-go/service/inspector/errors.go",
        "github.com/aws/aws-sdk-go/service/inspector/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/inspector_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/inspector/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/inspector",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/inspector/inspectoriface",
    srcs = ["github.com/aws/aws-sdk-go/service/inspector/inspectoriface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/inspector",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/iot",
    srcs = [
        "github.com/aws/aws-sdk-go/service/iot/api.go",
        "github.com/aws/aws-sdk-go/service/iot/errors.go",
        "github.com/aws/aws-sdk-go/service/iot/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/restjson",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/iot_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/iot/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/iot",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/iot/iotiface",
    srcs = ["github.com/aws/aws-sdk-go/service/iot/iotiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/iot",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/iotdataplane",
    srcs = [
        "github.com/aws/aws-sdk-go/service/iotdataplane/api.go",
        "github.com/aws/aws-sdk-go/service/iotdataplane/errors.go",
        "github.com/aws/aws-sdk-go/service/iotdataplane/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/restjson",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/iotdataplane_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/service/iotdataplane/customizations_test.go",
        "github.com/aws/aws-sdk-go/service/iotdataplane/examples_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/service/iotdataplane",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/iotdataplane/iotdataplaneiface",
    srcs = ["github.com/aws/aws-sdk-go/service/iotdataplane/iotdataplaneiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/iotdataplane",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/kinesis",
    srcs = [
        "github.com/aws/aws-sdk-go/service/kinesis/api.go",
        "github.com/aws/aws-sdk-go/service/kinesis/errors.go",
        "github.com/aws/aws-sdk-go/service/kinesis/service.go",
        "github.com/aws/aws-sdk-go/service/kinesis/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/kinesis_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/kinesis/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/kinesis",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/kinesis/kinesisiface",
    srcs = ["github.com/aws/aws-sdk-go/service/kinesis/kinesisiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/kinesis",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/kms",
    srcs = [
        "github.com/aws/aws-sdk-go/service/kms/api.go",
        "github.com/aws/aws-sdk-go/service/kms/errors.go",
        "github.com/aws/aws-sdk-go/service/kms/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/kms_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/kms/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/kms",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/kms/kmsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/kms/kmsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/kms",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/lambda",
    srcs = [
        "github.com/aws/aws-sdk-go/service/lambda/api.go",
        "github.com/aws/aws-sdk-go/service/lambda/errors.go",
        "github.com/aws/aws-sdk-go/service/lambda/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/restjson",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/lambda_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/lambda/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/lambda",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/lambda/lambdaiface",
    srcs = ["github.com/aws/aws-sdk-go/service/lambda/lambdaiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/lambda",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/machinelearning",
    srcs = [
        "github.com/aws/aws-sdk-go/service/machinelearning/api.go",
        "github.com/aws/aws-sdk-go/service/machinelearning/customizations.go",
        "github.com/aws/aws-sdk-go/service/machinelearning/errors.go",
        "github.com/aws/aws-sdk-go/service/machinelearning/service.go",
        "github.com/aws/aws-sdk-go/service/machinelearning/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/machinelearning_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/service/machinelearning/customizations_test.go",
        "github.com/aws/aws-sdk-go/service/machinelearning/examples_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/service/machinelearning",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/machinelearning/machinelearningiface",
    srcs = ["github.com/aws/aws-sdk-go/service/machinelearning/machinelearningiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/machinelearning",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics",
    srcs = [
        "github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics/api.go",
        "github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics/errors.go",
        "github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics/marketplacecommerceanalyticsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics/marketplacecommerceanalyticsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/mobileanalytics",
    srcs = [
        "github.com/aws/aws-sdk-go/service/mobileanalytics/api.go",
        "github.com/aws/aws-sdk-go/service/mobileanalytics/errors.go",
        "github.com/aws/aws-sdk-go/service/mobileanalytics/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/restjson",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/mobileanalytics_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/mobileanalytics/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/mobileanalytics",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/mobileanalytics/mobileanalyticsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/mobileanalytics/mobileanalyticsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/mobileanalytics",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/opsworks",
    srcs = [
        "github.com/aws/aws-sdk-go/service/opsworks/api.go",
        "github.com/aws/aws-sdk-go/service/opsworks/errors.go",
        "github.com/aws/aws-sdk-go/service/opsworks/service.go",
        "github.com/aws/aws-sdk-go/service/opsworks/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/opsworks_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/opsworks/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/opsworks",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/opsworks/opsworksiface",
    srcs = ["github.com/aws/aws-sdk-go/service/opsworks/opsworksiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/opsworks",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/rds_test",
    srcs = ["github.com/aws/aws-sdk-go/service/rds/customizations_test.go"],
    library = ":github.com/aws/aws-sdk-go/service/rds",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/rds",
    srcs = [
        "github.com/aws/aws-sdk-go/service/rds/api.go",
        "github.com/aws/aws-sdk-go/service/rds/customizations.go",
        "github.com/aws/aws-sdk-go/service/rds/errors.go",
        "github.com/aws/aws-sdk-go/service/rds/service.go",
        "github.com/aws/aws-sdk-go/service/rds/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/endpoints",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/query",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/rds_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/rds/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/rds",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/rds/rdsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/rds/rdsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/rds",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/redshift",
    srcs = [
        "github.com/aws/aws-sdk-go/service/redshift/api.go",
        "github.com/aws/aws-sdk-go/service/redshift/errors.go",
        "github.com/aws/aws-sdk-go/service/redshift/service.go",
        "github.com/aws/aws-sdk-go/service/redshift/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/query",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/redshift_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/redshift/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/redshift",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/redshift/redshiftiface",
    srcs = ["github.com/aws/aws-sdk-go/service/redshift/redshiftiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/redshift",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/route53_test",
    srcs = ["github.com/aws/aws-sdk-go/service/route53/unmarshal_error_leak_test.go"],
    library = ":github.com/aws/aws-sdk-go/service/route53",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/awstesting",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/route53",
    srcs = [
        "github.com/aws/aws-sdk-go/service/route53/api.go",
        "github.com/aws/aws-sdk-go/service/route53/customizations.go",
        "github.com/aws/aws-sdk-go/service/route53/errors.go",
        "github.com/aws/aws-sdk-go/service/route53/service.go",
        "github.com/aws/aws-sdk-go/service/route53/unmarshal_error.go",
        "github.com/aws/aws-sdk-go/service/route53/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/restxml",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/route53_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/service/route53/customizations_test.go",
        "github.com/aws/aws-sdk-go/service/route53/examples_test.go",
        "github.com/aws/aws-sdk-go/service/route53/unmarshal_error_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/service/route53",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/route53/route53iface",
    srcs = ["github.com/aws/aws-sdk-go/service/route53/route53iface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/route53",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/route53domains",
    srcs = [
        "github.com/aws/aws-sdk-go/service/route53domains/api.go",
        "github.com/aws/aws-sdk-go/service/route53domains/errors.go",
        "github.com/aws/aws-sdk-go/service/route53domains/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/route53domains_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/route53domains/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/route53domains",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/route53domains/route53domainsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/route53domains/route53domainsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/route53domains",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/s3_test",
    srcs = ["github.com/aws/aws-sdk-go/service/s3/unmarshal_error_leak_test.go"],
    library = ":github.com/aws/aws-sdk-go/service/s3",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/awstesting",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/s3",
    srcs = [
        "github.com/aws/aws-sdk-go/service/s3/api.go",
        "github.com/aws/aws-sdk-go/service/s3/bucket_location.go",
        "github.com/aws/aws-sdk-go/service/s3/content_md5.go",
        "github.com/aws/aws-sdk-go/service/s3/customizations.go",
        "github.com/aws/aws-sdk-go/service/s3/errors.go",
        "github.com/aws/aws-sdk-go/service/s3/host_style_bucket.go",
        "github.com/aws/aws-sdk-go/service/s3/platform_handlers_go1.6.go",
        "github.com/aws/aws-sdk-go/service/s3/service.go",
        "github.com/aws/aws-sdk-go/service/s3/sse.go",
        "github.com/aws/aws-sdk-go/service/s3/statusok_error.go",
        "github.com/aws/aws-sdk-go/service/s3/unmarshal_error.go",
        "github.com/aws/aws-sdk-go/service/s3/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/restxml",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/s3_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/service/s3/bucket_location_test.go",
        "github.com/aws/aws-sdk-go/service/s3/customizations_test.go",
        "github.com/aws/aws-sdk-go/service/s3/examples_test.go",
        "github.com/aws/aws-sdk-go/service/s3/host_style_bucket_test.go",
        "github.com/aws/aws-sdk-go/service/s3/platform_handlers_go1.6_test.go",
        "github.com/aws/aws-sdk-go/service/s3/sse_test.go",
        "github.com/aws/aws-sdk-go/service/s3/statusok_error_test.go",
        "github.com/aws/aws-sdk-go/service/s3/unmarshal_error_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/service/s3",
        "//vendor:github.com/stretchr/testify/assert",
        "//vendor:github.com/stretchr/testify/require",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/s3/s3crypto_test",
    srcs = [
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/aes_gcm_test.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/cipher_util_test.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/hash_io_test.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/helper_test.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/key_handler_test.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/kms_key_handler_test.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/mat_desc_test.go",
    ],
    library = ":github.com/aws/aws-sdk-go/service/s3/s3crypto",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/service/kms",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/s3/s3crypto",
    srcs = [
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/aes_gcm.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/aes_gcm_content_cipher.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/cipher.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/cipher_builder.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/cipher_util.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/decryption_client.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/doc.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/encryption_client.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/envelope.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/hash_io.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/helper.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/key_handler.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/kms_key_handler.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/mat_desc.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/strategy.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/kms",
        "//vendor:github.com/aws/aws-sdk-go/service/kms/kmsiface",
        "//vendor:github.com/aws/aws-sdk-go/service/s3",
        "//vendor:github.com/aws/aws-sdk-go/service/s3/s3iface",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/s3/s3crypto_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/aes_gcm_content_cipher_test.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/cipher_test.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/decryption_client_test.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/encryption_client_test.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/mock_test.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/strategy_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/service/kms",
        "//vendor:github.com/aws/aws-sdk-go/service/s3",
        "//vendor:github.com/aws/aws-sdk-go/service/s3/s3crypto",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/s3/s3iface",
    srcs = ["github.com/aws/aws-sdk-go/service/s3/s3iface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/s3",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/s3/s3manager",
    srcs = [
        "github.com/aws/aws-sdk-go/service/s3/s3manager/doc.go",
        "github.com/aws/aws-sdk-go/service/s3/s3manager/download.go",
        "github.com/aws/aws-sdk-go/service/s3/s3manager/upload.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/s3",
        "//vendor:github.com/aws/aws-sdk-go/service/s3/s3iface",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/s3/s3manager_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/service/s3/s3manager/download_test.go",
        "github.com/aws/aws-sdk-go/service/s3/s3manager/shared_test.go",
        "github.com/aws/aws-sdk-go/service/s3/s3manager/upload_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/service/s3",
        "//vendor:github.com/aws/aws-sdk-go/service/s3/s3manager",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/s3/s3manager/s3manageriface",
    srcs = ["github.com/aws/aws-sdk-go/service/s3/s3manager/s3manageriface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/service/s3",
        "//vendor:github.com/aws/aws-sdk-go/service/s3/s3manager",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/ses",
    srcs = [
        "github.com/aws/aws-sdk-go/service/ses/api.go",
        "github.com/aws/aws-sdk-go/service/ses/errors.go",
        "github.com/aws/aws-sdk-go/service/ses/service.go",
        "github.com/aws/aws-sdk-go/service/ses/waiters.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/query",
        "//vendor:github.com/aws/aws-sdk-go/private/waiter",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/ses_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/ses/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/ses",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/ses/sesiface",
    srcs = ["github.com/aws/aws-sdk-go/service/ses/sesiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/ses",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/simpledb_test",
    srcs = ["github.com/aws/aws-sdk-go/service/simpledb/unmarshal_error_leak_test.go"],
    library = ":github.com/aws/aws-sdk-go/service/simpledb",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/awstesting",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/simpledb",
    srcs = [
        "github.com/aws/aws-sdk-go/service/simpledb/api.go",
        "github.com/aws/aws-sdk-go/service/simpledb/customizations.go",
        "github.com/aws/aws-sdk-go/service/simpledb/errors.go",
        "github.com/aws/aws-sdk-go/service/simpledb/service.go",
        "github.com/aws/aws-sdk-go/service/simpledb/unmarshall_error.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/corehandlers",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/query",
        "//vendor:github.com/aws/aws-sdk-go/private/signer/v2",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/simpledb_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/service/simpledb/examples_test.go",
        "github.com/aws/aws-sdk-go/service/simpledb/unmarshall_error_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/service/simpledb",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/simpledb/simpledbiface",
    srcs = ["github.com/aws/aws-sdk-go/service/simpledb/simpledbiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/simpledb",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/sns",
    srcs = [
        "github.com/aws/aws-sdk-go/service/sns/api.go",
        "github.com/aws/aws-sdk-go/service/sns/errors.go",
        "github.com/aws/aws-sdk-go/service/sns/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/query",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/sns_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/sns/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/sns",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/sns/snsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/sns/snsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/sns",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/sqs",
    srcs = [
        "github.com/aws/aws-sdk-go/service/sqs/api.go",
        "github.com/aws/aws-sdk-go/service/sqs/checksums.go",
        "github.com/aws/aws-sdk-go/service/sqs/customizations.go",
        "github.com/aws/aws-sdk-go/service/sqs/errors.go",
        "github.com/aws/aws-sdk-go/service/sqs/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/query",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/sqs_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/service/sqs/checksums_test.go",
        "github.com/aws/aws-sdk-go/service/sqs/examples_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/service/sqs",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/sqs/sqsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/sqs/sqsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/sqs",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/ssm",
    srcs = [
        "github.com/aws/aws-sdk-go/service/ssm/api.go",
        "github.com/aws/aws-sdk-go/service/ssm/errors.go",
        "github.com/aws/aws-sdk-go/service/ssm/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/ssm_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/ssm/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/ssm",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/ssm/ssmiface",
    srcs = ["github.com/aws/aws-sdk-go/service/ssm/ssmiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/ssm",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/storagegateway",
    srcs = [
        "github.com/aws/aws-sdk-go/service/storagegateway/api.go",
        "github.com/aws/aws-sdk-go/service/storagegateway/errors.go",
        "github.com/aws/aws-sdk-go/service/storagegateway/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/storagegateway_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/storagegateway/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/storagegateway",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/storagegateway/storagegatewayiface",
    srcs = ["github.com/aws/aws-sdk-go/service/storagegateway/storagegatewayiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/storagegateway",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/sts",
    srcs = [
        "github.com/aws/aws-sdk-go/service/sts/api.go",
        "github.com/aws/aws-sdk-go/service/sts/customizations.go",
        "github.com/aws/aws-sdk-go/service/sts/errors.go",
        "github.com/aws/aws-sdk-go/service/sts/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/query",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/sts_xtest",
    srcs = [
        "github.com/aws/aws-sdk-go/service/sts/customizations_test.go",
        "github.com/aws/aws-sdk-go/service/sts/examples_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/awstesting/unit",
        "//vendor:github.com/aws/aws-sdk-go/service/sts",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/sts/stsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/sts/stsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/sts",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/support",
    srcs = [
        "github.com/aws/aws-sdk-go/service/support/api.go",
        "github.com/aws/aws-sdk-go/service/support/errors.go",
        "github.com/aws/aws-sdk-go/service/support/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/support_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/support/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/support",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/support/supportiface",
    srcs = ["github.com/aws/aws-sdk-go/service/support/supportiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/support",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/swf",
    srcs = [
        "github.com/aws/aws-sdk-go/service/swf/api.go",
        "github.com/aws/aws-sdk-go/service/swf/errors.go",
        "github.com/aws/aws-sdk-go/service/swf/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/swf_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/swf/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/swf",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/swf/swfiface",
    srcs = ["github.com/aws/aws-sdk-go/service/swf/swfiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/swf",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/waf",
    srcs = [
        "github.com/aws/aws-sdk-go/service/waf/api.go",
        "github.com/aws/aws-sdk-go/service/waf/errors.go",
        "github.com/aws/aws-sdk-go/service/waf/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/waf_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/waf/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/waf",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/waf/wafiface",
    srcs = ["github.com/aws/aws-sdk-go/service/waf/wafiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/waf",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/workspaces",
    srcs = [
        "github.com/aws/aws-sdk-go/service/workspaces/api.go",
        "github.com/aws/aws-sdk-go/service/workspaces/errors.go",
        "github.com/aws/aws-sdk-go/service/workspaces/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/service/workspaces_xtest",
    srcs = ["github.com/aws/aws-sdk-go/service/workspaces/examples_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/session",
        "//vendor:github.com/aws/aws-sdk-go/service/workspaces",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/workspaces/workspacesiface",
    srcs = ["github.com/aws/aws-sdk-go/service/workspaces/workspacesiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/workspaces",
    ],
)

go_test(
    name = "github.com/aws/aws-sdk-go/vendor/github.com/go-ini/ini_test",
    srcs = [
        "github.com/aws/aws-sdk-go/vendor/github.com/go-ini/ini/ini_test.go",
        "github.com/aws/aws-sdk-go/vendor/github.com/go-ini/ini/struct_test.go",
    ],
    library = ":github.com/aws/aws-sdk-go/vendor/github.com/go-ini/ini",
    tags = ["automanaged"],
    deps = ["//vendor:github.com/smartystreets/goconvey/convey"],
)

go_library(
    name = "github.com/aws/aws-sdk-go/vendor/github.com/go-ini/ini",
    srcs = [
        "github.com/aws/aws-sdk-go/vendor/github.com/go-ini/ini/ini.go",
        "github.com/aws/aws-sdk-go/vendor/github.com/go-ini/ini/struct.go",
    ],
    tags = ["automanaged"],
)

go_test(
    name = "github.com/aws/aws-sdk-go/vendor/github.com/jmespath/go-jmespath_test",
    srcs = [
        "github.com/aws/aws-sdk-go/vendor/github.com/jmespath/go-jmespath/api_test.go",
        "github.com/aws/aws-sdk-go/vendor/github.com/jmespath/go-jmespath/compliance_test.go",
        "github.com/aws/aws-sdk-go/vendor/github.com/jmespath/go-jmespath/interpreter_test.go",
        "github.com/aws/aws-sdk-go/vendor/github.com/jmespath/go-jmespath/lexer_test.go",
        "github.com/aws/aws-sdk-go/vendor/github.com/jmespath/go-jmespath/parser_test.go",
        "github.com/aws/aws-sdk-go/vendor/github.com/jmespath/go-jmespath/util_test.go",
    ],
    library = ":github.com/aws/aws-sdk-go/vendor/github.com/jmespath/go-jmespath",
    tags = ["automanaged"],
    deps = ["//vendor:github.com/stretchr/testify/assert"],
)

go_library(
    name = "github.com/aws/aws-sdk-go/vendor/github.com/jmespath/go-jmespath",
    srcs = [
        "github.com/aws/aws-sdk-go/vendor/github.com/jmespath/go-jmespath/api.go",
        "github.com/aws/aws-sdk-go/vendor/github.com/jmespath/go-jmespath/astnodetype_string.go",
        "github.com/aws/aws-sdk-go/vendor/github.com/jmespath/go-jmespath/functions.go",
        "github.com/aws/aws-sdk-go/vendor/github.com/jmespath/go-jmespath/interpreter.go",
        "github.com/aws/aws-sdk-go/vendor/github.com/jmespath/go-jmespath/lexer.go",
        "github.com/aws/aws-sdk-go/vendor/github.com/jmespath/go-jmespath/parser.go",
        "github.com/aws/aws-sdk-go/vendor/github.com/jmespath/go-jmespath/toktype_string.go",
        "github.com/aws/aws-sdk-go/vendor/github.com/jmespath/go-jmespath/util.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/aws/aws-sdk-go/vendor/github.com/jmespath/go-jmespath/fuzz",
    srcs = ["github.com/aws/aws-sdk-go/vendor/github.com/jmespath/go-jmespath/fuzz/jmespath.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/jmespath/go-jmespath"],
)

go_test(
    name = "github.com/codegangsta/negroni_test",
    srcs = [
        "github.com/codegangsta/negroni/logger_test.go",
        "github.com/codegangsta/negroni/negroni_test.go",
        "github.com/codegangsta/negroni/recovery_test.go",
        "github.com/codegangsta/negroni/response_writer_test.go",
        "github.com/codegangsta/negroni/static_test.go",
    ],
    library = ":github.com/codegangsta/negroni",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/codegangsta/negroni",
    srcs = [
        "github.com/codegangsta/negroni/doc.go",
        "github.com/codegangsta/negroni/logger.go",
        "github.com/codegangsta/negroni/negroni.go",
        "github.com/codegangsta/negroni/recovery.go",
        "github.com/codegangsta/negroni/response_writer.go",
        "github.com/codegangsta/negroni/static.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/daaku/go.zipexe",
    srcs = ["github.com/daaku/go.zipexe/zipexe.go"],
    tags = ["automanaged"],
)

go_test(
    name = "github.com/davecgh/go-spew/spew_test",
    srcs = [
        "github.com/davecgh/go-spew/spew/internal_test.go",
        "github.com/davecgh/go-spew/spew/internalunsafe_test.go",
    ],
    library = ":github.com/davecgh/go-spew/spew",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/davecgh/go-spew/spew",
    srcs = [
        "github.com/davecgh/go-spew/spew/bypass.go",
        "github.com/davecgh/go-spew/spew/common.go",
        "github.com/davecgh/go-spew/spew/config.go",
        "github.com/davecgh/go-spew/spew/doc.go",
        "github.com/davecgh/go-spew/spew/dump.go",
        "github.com/davecgh/go-spew/spew/format.go",
        "github.com/davecgh/go-spew/spew/spew.go",
    ],
    tags = ["automanaged"],
)

go_test(
    name = "github.com/davecgh/go-spew/spew_xtest",
    srcs = [
        "github.com/davecgh/go-spew/spew/common_test.go",
        "github.com/davecgh/go-spew/spew/dump_test.go",
        "github.com/davecgh/go-spew/spew/dumpnocgo_test.go",
        "github.com/davecgh/go-spew/spew/example_test.go",
        "github.com/davecgh/go-spew/spew/format_test.go",
        "github.com/davecgh/go-spew/spew/spew_test.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/davecgh/go-spew/spew"],
)

go_test(
    name = "github.com/dimfeld/httppath_test",
    srcs = ["github.com/dimfeld/httppath/path_test.go"],
    library = ":github.com/dimfeld/httppath",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/dimfeld/httppath",
    srcs = ["github.com/dimfeld/httppath/path.go"],
    tags = ["automanaged"],
)

go_test(
    name = "github.com/dimfeld/httptreemux_test",
    srcs = [
        "github.com/dimfeld/httptreemux/context_test.go",
        "github.com/dimfeld/httptreemux/fallthrough_test.go",
        "github.com/dimfeld/httptreemux/group_test.go",
        "github.com/dimfeld/httptreemux/path_test.go",
        "github.com/dimfeld/httptreemux/router_test.go",
        "github.com/dimfeld/httptreemux/tree_test.go",
    ],
    library = ":github.com/dimfeld/httptreemux",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/dimfeld/httptreemux",
    srcs = [
        "github.com/dimfeld/httptreemux/context.go",
        "github.com/dimfeld/httptreemux/group.go",
        "github.com/dimfeld/httptreemux/panichandler.go",
        "github.com/dimfeld/httptreemux/path.go",
        "github.com/dimfeld/httptreemux/router.go",
        "github.com/dimfeld/httptreemux/tree.go",
        "github.com/dimfeld/httptreemux/treemux_17.go",
    ],
    tags = ["automanaged"],
)

go_test(
    name = "github.com/golang/protobuf/proto_test",
    srcs = [
        "github.com/golang/protobuf/proto/message_set_test.go",
        "github.com/golang/protobuf/proto/size2_test.go",
    ],
    library = ":github.com/golang/protobuf/proto",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/golang/protobuf/proto",
    srcs = [
        "github.com/golang/protobuf/proto/clone.go",
        "github.com/golang/protobuf/proto/decode.go",
        "github.com/golang/protobuf/proto/encode.go",
        "github.com/golang/protobuf/proto/equal.go",
        "github.com/golang/protobuf/proto/extensions.go",
        "github.com/golang/protobuf/proto/lib.go",
        "github.com/golang/protobuf/proto/message_set.go",
        "github.com/golang/protobuf/proto/pointer_unsafe.go",
        "github.com/golang/protobuf/proto/properties.go",
        "github.com/golang/protobuf/proto/text.go",
        "github.com/golang/protobuf/proto/text_parser.go",
    ],
    tags = ["automanaged"],
)

go_test(
    name = "github.com/golang/protobuf/proto_xtest",
    srcs = [
        "github.com/golang/protobuf/proto/all_test.go",
        "github.com/golang/protobuf/proto/any_test.go",
        "github.com/golang/protobuf/proto/clone_test.go",
        "github.com/golang/protobuf/proto/decode_test.go",
        "github.com/golang/protobuf/proto/encode_test.go",
        "github.com/golang/protobuf/proto/equal_test.go",
        "github.com/golang/protobuf/proto/extensions_test.go",
        "github.com/golang/protobuf/proto/map_test.go",
        "github.com/golang/protobuf/proto/proto3_test.go",
        "github.com/golang/protobuf/proto/size_test.go",
        "github.com/golang/protobuf/proto/text_parser_test.go",
        "github.com/golang/protobuf/proto/text_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:github.com/golang/protobuf/proto/proto3_proto",
        "//vendor:github.com/golang/protobuf/proto/testdata",
        "//vendor:github.com/golang/protobuf/ptypes",
        "//vendor:github.com/golang/protobuf/ptypes/any",
        "//vendor:golang.org/x/sync/errgroup",
    ],
)

go_library(
    name = "github.com/golang/protobuf/proto/proto3_proto",
    srcs = ["github.com/golang/protobuf/proto/proto3_proto/proto3.pb.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:github.com/golang/protobuf/proto/testdata",
        "//vendor:github.com/golang/protobuf/ptypes/any",
    ],
)

go_test(
    name = "github.com/golang/protobuf/proto/testdata_test",
    srcs = ["github.com/golang/protobuf/proto/testdata/golden_test.go"],
    library = ":github.com/golang/protobuf/proto/testdata",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/golang/protobuf/proto/testdata",
    srcs = ["github.com/golang/protobuf/proto/testdata/test.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "github.com/golang/protobuf/protoc-gen-go/descriptor",
    srcs = ["github.com/golang/protobuf/protoc-gen-go/descriptor/descriptor.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_test(
    name = "github.com/golang/protobuf/ptypes_test",
    srcs = [
        "github.com/golang/protobuf/ptypes/any_test.go",
        "github.com/golang/protobuf/ptypes/duration_test.go",
        "github.com/golang/protobuf/ptypes/timestamp_test.go",
    ],
    library = ":github.com/golang/protobuf/ptypes",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:github.com/golang/protobuf/protoc-gen-go/descriptor",
        "//vendor:github.com/golang/protobuf/ptypes/any",
        "//vendor:github.com/golang/protobuf/ptypes/duration",
        "//vendor:github.com/golang/protobuf/ptypes/timestamp",
    ],
)

go_library(
    name = "github.com/golang/protobuf/ptypes",
    srcs = [
        "github.com/golang/protobuf/ptypes/any.go",
        "github.com/golang/protobuf/ptypes/doc.go",
        "github.com/golang/protobuf/ptypes/duration.go",
        "github.com/golang/protobuf/ptypes/timestamp.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:github.com/golang/protobuf/ptypes/any",
        "//vendor:github.com/golang/protobuf/ptypes/duration",
        "//vendor:github.com/golang/protobuf/ptypes/timestamp",
    ],
)

go_library(
    name = "github.com/golang/protobuf/ptypes/any",
    srcs = ["github.com/golang/protobuf/ptypes/any/any.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "github.com/golang/protobuf/ptypes/duration",
    srcs = ["github.com/golang/protobuf/ptypes/duration/duration.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "github.com/golang/protobuf/ptypes/empty",
    srcs = ["github.com/golang/protobuf/ptypes/empty/empty.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "github.com/golang/protobuf/ptypes/struct",
    srcs = ["github.com/golang/protobuf/ptypes/struct/struct.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "github.com/golang/protobuf/ptypes/timestamp",
    srcs = ["github.com/golang/protobuf/ptypes/timestamp/timestamp.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "github.com/golang/protobuf/ptypes/wrappers",
    srcs = ["github.com/golang/protobuf/ptypes/wrappers/wrappers.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_test(
    name = "github.com/google/shlex_test",
    srcs = ["github.com/google/shlex/shlex_test.go"],
    library = ":github.com/google/shlex",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/google/shlex",
    srcs = ["github.com/google/shlex/shlex.go"],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/gopherjs/gopherjs/js",
    srcs = ["github.com/gopherjs/gopherjs/js/js.go"],
    tags = ["automanaged"],
)

go_test(
    name = "github.com/gucumber/gucumber_test",
    srcs = ["github.com/gucumber/gucumber/builder_test.go"],
    library = ":github.com/gucumber/gucumber",
    tags = ["automanaged"],
    deps = ["//vendor:github.com/stretchr/testify/assert"],
)

go_library(
    name = "github.com/gucumber/gucumber",
    srcs = [
        "github.com/gucumber/gucumber/builder.go",
        "github.com/gucumber/gucumber/run_main.go",
        "github.com/gucumber/gucumber/runner.go",
        "github.com/gucumber/gucumber/stepdef.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/gucumber/gucumber/gherkin",
        "//vendor:github.com/shiena/ansicolor",
    ],
)

go_test(
    name = "github.com/gucumber/gucumber_xtest",
    srcs = ["github.com/gucumber/gucumber/stepdef_test.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/gucumber/gucumber",
        "//vendor:github.com/gucumber/gucumber/gherkin",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_binary(
    name = "github.com/gucumber/gucumber/cmd/gucumber_bin",
    library = ":github.com/gucumber/gucumber/cmd/gucumber",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/gucumber/gucumber/cmd/gucumber",
    srcs = ["github.com/gucumber/gucumber/cmd/gucumber/gucumber.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/gucumber/gucumber"],
)

go_test(
    name = "github.com/gucumber/gucumber/gherkin_test",
    srcs = [
        "github.com/gucumber/gucumber/gherkin/parser_test.go",
        "github.com/gucumber/gucumber/gherkin/types_test.go",
    ],
    library = ":github.com/gucumber/gucumber/gherkin",
    tags = ["automanaged"],
    deps = ["//vendor:github.com/stretchr/testify/assert"],
)

go_library(
    name = "github.com/gucumber/gucumber/gherkin",
    srcs = [
        "github.com/gucumber/gucumber/gherkin/i18n.go",
        "github.com/gucumber/gucumber/gherkin/parser.go",
        "github.com/gucumber/gucumber/gherkin/types.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/gucumber/gucumber/internal/features/basic",
    srcs = ["github.com/gucumber/gucumber/internal/features/basic/step_definitions.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/gucumber/gucumber",
        "//vendor:github.com/stretchr/testify/assert",
    ],
)

go_library(
    name = "github.com/guregu/kami",
    srcs = [
        "github.com/guregu/kami/defaults.go",
        "github.com/guregu/kami/globalmux.go",
        "github.com/guregu/kami/globalmux_17.go",
        "github.com/guregu/kami/handler_17.go",
        "github.com/guregu/kami/kami_17.go",
        "github.com/guregu/kami/middleware.go",
        "github.com/guregu/kami/middleware_17.go",
        "github.com/guregu/kami/mux_17.go",
        "github.com/guregu/kami/params.go",
        "github.com/guregu/kami/serve.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/dimfeld/httptreemux",
        "//vendor:github.com/guregu/kami/treemux",
        "//vendor:github.com/zenazn/goji/bind",
        "//vendor:github.com/zenazn/goji/graceful",
        "//vendor:github.com/zenazn/goji/web/mutil",
        "//vendor:golang.org/x/net/context",
    ],
)

go_test(
    name = "github.com/guregu/kami_xtest",
    srcs = [
        "github.com/guregu/kami/bench_test.go",
        "github.com/guregu/kami/kami_17_test.go",
        "github.com/guregu/kami/middleware_test.go",
        "github.com/guregu/kami/mux_test.go",
        "github.com/guregu/kami/params_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/guregu/kami",
        "//vendor:github.com/zenazn/goji/web/mutil",
        "//vendor:golang.org/x/net/context",
    ],
)

go_test(
    name = "github.com/guregu/kami/treemux_test",
    srcs = ["github.com/guregu/kami/treemux/tree_test.go"],
    library = ":github.com/guregu/kami/treemux",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/guregu/kami/treemux",
    srcs = [
        "github.com/guregu/kami/treemux/router.go",
        "github.com/guregu/kami/treemux/tree.go",
    ],
    tags = ["automanaged"],
)

go_test(
    name = "github.com/hashicorp/go-multierror_test",
    srcs = [
        "github.com/hashicorp/go-multierror/append_test.go",
        "github.com/hashicorp/go-multierror/flatten_test.go",
        "github.com/hashicorp/go-multierror/format_test.go",
        "github.com/hashicorp/go-multierror/multierror_test.go",
        "github.com/hashicorp/go-multierror/prefix_test.go",
    ],
    library = ":github.com/hashicorp/go-multierror",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/hashicorp/go-multierror",
    srcs = [
        "github.com/hashicorp/go-multierror/append.go",
        "github.com/hashicorp/go-multierror/flatten.go",
        "github.com/hashicorp/go-multierror/format.go",
        "github.com/hashicorp/go-multierror/multierror.go",
        "github.com/hashicorp/go-multierror/prefix.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/hashicorp/errwrap"],
)

go_library(
    name = "github.com/hashicorp/go-multierror/vendor/github.com/hashicorp/errwrap",
    srcs = ["github.com/hashicorp/go-multierror/vendor/github.com/hashicorp/errwrap/errwrap.go"],
    tags = ["automanaged"],
)

go_test(
    name = "github.com/jessevdk/go-flags_test",
    srcs = [
        "github.com/jessevdk/go-flags/arg_test.go",
        "github.com/jessevdk/go-flags/assert_test.go",
        "github.com/jessevdk/go-flags/command_test.go",
        "github.com/jessevdk/go-flags/completion_test.go",
        "github.com/jessevdk/go-flags/convert_test.go",
        "github.com/jessevdk/go-flags/example_test.go",
        "github.com/jessevdk/go-flags/group_test.go",
        "github.com/jessevdk/go-flags/help_test.go",
        "github.com/jessevdk/go-flags/ini_test.go",
        "github.com/jessevdk/go-flags/long_test.go",
        "github.com/jessevdk/go-flags/marshal_test.go",
        "github.com/jessevdk/go-flags/options_test.go",
        "github.com/jessevdk/go-flags/parser_test.go",
        "github.com/jessevdk/go-flags/pointer_test.go",
        "github.com/jessevdk/go-flags/short_test.go",
        "github.com/jessevdk/go-flags/tag_test.go",
        "github.com/jessevdk/go-flags/unknown_test.go",
    ],
    library = ":github.com/jessevdk/go-flags",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/jessevdk/go-flags",
    srcs = [
        "github.com/jessevdk/go-flags/arg.go",
        "github.com/jessevdk/go-flags/closest.go",
        "github.com/jessevdk/go-flags/command.go",
        "github.com/jessevdk/go-flags/completion.go",
        "github.com/jessevdk/go-flags/convert.go",
        "github.com/jessevdk/go-flags/error.go",
        "github.com/jessevdk/go-flags/flags.go",
        "github.com/jessevdk/go-flags/group.go",
        "github.com/jessevdk/go-flags/help.go",
        "github.com/jessevdk/go-flags/ini.go",
        "github.com/jessevdk/go-flags/man.go",
        "github.com/jessevdk/go-flags/multitag.go",
        "github.com/jessevdk/go-flags/option.go",
        "github.com/jessevdk/go-flags/optstyle_other.go",
        "github.com/jessevdk/go-flags/parser.go",
        "github.com/jessevdk/go-flags/termsize.go",
        "github.com/jessevdk/go-flags/termsize_linux.go",
    ],
    tags = ["automanaged"],
)

go_binary(
    name = "github.com/jessevdk/go-flags/examples_bin",
    library = ":github.com/jessevdk/go-flags/examples",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/jessevdk/go-flags/examples",
    srcs = [
        "github.com/jessevdk/go-flags/examples/add.go",
        "github.com/jessevdk/go-flags/examples/main.go",
        "github.com/jessevdk/go-flags/examples/rm.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/jessevdk/go-flags"],
)

go_library(
    name = "github.com/jtolds/gls",
    srcs = [
        "github.com/jtolds/gls/context.go",
        "github.com/jtolds/gls/gen_sym.go",
        "github.com/jtolds/gls/gid.go",
        "github.com/jtolds/gls/id_pool.go",
        "github.com/jtolds/gls/stack_tags.go",
        "github.com/jtolds/gls/stack_tags_main.go",
    ],
    tags = ["automanaged"],
)

go_test(
    name = "github.com/jtolds/gls_xtest",
    srcs = ["github.com/jtolds/gls/context_test.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/jtolds/gls"],
)

go_test(
    name = "github.com/kardianos/osext_test",
    srcs = ["github.com/kardianos/osext/osext_test.go"],
    library = ":github.com/kardianos/osext",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/kardianos/osext",
    srcs = [
        "github.com/kardianos/osext/osext.go",
        "github.com/kardianos/osext/osext_procfs.go",
    ],
    tags = ["automanaged"],
)

go_test(
    name = "github.com/pkg/errors_test",
    srcs = [
        "github.com/pkg/errors/bench_test.go",
        "github.com/pkg/errors/errors_test.go",
        "github.com/pkg/errors/format_test.go",
        "github.com/pkg/errors/stack_test.go",
    ],
    library = ":github.com/pkg/errors",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/pkg/errors",
    srcs = [
        "github.com/pkg/errors/errors.go",
        "github.com/pkg/errors/stack.go",
    ],
    tags = ["automanaged"],
)

go_test(
    name = "github.com/pkg/errors_xtest",
    srcs = ["github.com/pkg/errors/example_test.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/pkg/errors"],
)

go_library(
    name = "github.com/shiena/ansicolor",
    srcs = [
        "github.com/shiena/ansicolor/ansicolor.go",
        "github.com/shiena/ansicolor/ansicolor_ansi.go",
    ],
    tags = ["automanaged"],
)

go_test(
    name = "github.com/shiena/ansicolor_xtest",
    srcs = [
        "github.com/shiena/ansicolor/ansicolor_test.go",
        "github.com/shiena/ansicolor/example_test.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/shiena/ansicolor"],
)

go_binary(
    name = "github.com/shiena/ansicolor/ansicolor_bin",
    library = ":github.com/shiena/ansicolor/ansicolor",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/shiena/ansicolor/ansicolor",
    srcs = ["github.com/shiena/ansicolor/ansicolor/main.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/shiena/ansicolor"],
)

go_test(
    name = "github.com/smartystreets/assertions_test",
    srcs = [
        "github.com/smartystreets/assertions/collections_test.go",
        "github.com/smartystreets/assertions/doc_test.go",
        "github.com/smartystreets/assertions/equality_test.go",
        "github.com/smartystreets/assertions/panic_test.go",
        "github.com/smartystreets/assertions/quantity_test.go",
        "github.com/smartystreets/assertions/serializer_test.go",
        "github.com/smartystreets/assertions/strings_test.go",
        "github.com/smartystreets/assertions/time_test.go",
        "github.com/smartystreets/assertions/type_test.go",
        "github.com/smartystreets/assertions/utilities_for_test.go",
    ],
    library = ":github.com/smartystreets/assertions",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/smartystreets/assertions",
    srcs = [
        "github.com/smartystreets/assertions/collections.go",
        "github.com/smartystreets/assertions/doc.go",
        "github.com/smartystreets/assertions/equality.go",
        "github.com/smartystreets/assertions/filter.go",
        "github.com/smartystreets/assertions/messages.go",
        "github.com/smartystreets/assertions/panic.go",
        "github.com/smartystreets/assertions/quantity.go",
        "github.com/smartystreets/assertions/serializer.go",
        "github.com/smartystreets/assertions/strings.go",
        "github.com/smartystreets/assertions/time.go",
        "github.com/smartystreets/assertions/type.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/smartystreets/assertions/internal/go-render/render",
        "//vendor:github.com/smartystreets/assertions/internal/oglematchers",
    ],
)

go_library(
    name = "github.com/smartystreets/assertions/internal/go-render/render",
    srcs = ["github.com/smartystreets/assertions/internal/go-render/render/render.go"],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/smartystreets/assertions/internal/oglematchers",
    srcs = [
        "github.com/smartystreets/assertions/internal/oglematchers/any_of.go",
        "github.com/smartystreets/assertions/internal/oglematchers/contains.go",
        "github.com/smartystreets/assertions/internal/oglematchers/deep_equals.go",
        "github.com/smartystreets/assertions/internal/oglematchers/equals.go",
        "github.com/smartystreets/assertions/internal/oglematchers/greater_or_equal.go",
        "github.com/smartystreets/assertions/internal/oglematchers/greater_than.go",
        "github.com/smartystreets/assertions/internal/oglematchers/less_or_equal.go",
        "github.com/smartystreets/assertions/internal/oglematchers/less_than.go",
        "github.com/smartystreets/assertions/internal/oglematchers/matcher.go",
        "github.com/smartystreets/assertions/internal/oglematchers/not.go",
        "github.com/smartystreets/assertions/internal/oglematchers/transform_description.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/smartystreets/assertions/should",
    srcs = ["github.com/smartystreets/assertions/should/should.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/smartystreets/assertions"],
)

go_test(
    name = "github.com/smartystreets/goconvey/convey_test",
    srcs = [
        "github.com/smartystreets/goconvey/convey/focused_execution_test.go",
        "github.com/smartystreets/goconvey/convey/isolated_execution_test.go",
        "github.com/smartystreets/goconvey/convey/reporting_hooks_test.go",
        "github.com/smartystreets/goconvey/convey/story_conventions_test.go",
    ],
    library = ":github.com/smartystreets/goconvey/convey",
    tags = ["automanaged"],
    deps = ["//vendor:github.com/smartystreets/goconvey/convey/reporting"],
)

go_library(
    name = "github.com/smartystreets/goconvey/convey",
    srcs = [
        "github.com/smartystreets/goconvey/convey/assertions.go",
        "github.com/smartystreets/goconvey/convey/context.go",
        "github.com/smartystreets/goconvey/convey/discovery.go",
        "github.com/smartystreets/goconvey/convey/doc.go",
        "github.com/smartystreets/goconvey/convey/init.go",
        "github.com/smartystreets/goconvey/convey/nilReporter.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/jtolds/gls",
        "//vendor:github.com/smartystreets/assertions",
        "//vendor:github.com/smartystreets/goconvey/convey/reporting",
    ],
)

go_test(
    name = "github.com/smartystreets/goconvey/convey/gotest_test",
    srcs = ["github.com/smartystreets/goconvey/convey/gotest/doc_test.go"],
    library = ":github.com/smartystreets/goconvey/convey/gotest",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/smartystreets/goconvey/convey/gotest",
    srcs = ["github.com/smartystreets/goconvey/convey/gotest/utils.go"],
    tags = ["automanaged"],
)

go_test(
    name = "github.com/smartystreets/goconvey/convey/reporting_test",
    srcs = [
        "github.com/smartystreets/goconvey/convey/reporting/dot_test.go",
        "github.com/smartystreets/goconvey/convey/reporting/gotest_test.go",
        "github.com/smartystreets/goconvey/convey/reporting/printer_test.go",
        "github.com/smartystreets/goconvey/convey/reporting/problems_test.go",
        "github.com/smartystreets/goconvey/convey/reporting/reporter_test.go",
    ],
    library = ":github.com/smartystreets/goconvey/convey/reporting",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/smartystreets/goconvey/convey/reporting",
    srcs = [
        "github.com/smartystreets/goconvey/convey/reporting/console.go",
        "github.com/smartystreets/goconvey/convey/reporting/doc.go",
        "github.com/smartystreets/goconvey/convey/reporting/dot.go",
        "github.com/smartystreets/goconvey/convey/reporting/gotest.go",
        "github.com/smartystreets/goconvey/convey/reporting/init.go",
        "github.com/smartystreets/goconvey/convey/reporting/json.go",
        "github.com/smartystreets/goconvey/convey/reporting/printer.go",
        "github.com/smartystreets/goconvey/convey/reporting/problems.go",
        "github.com/smartystreets/goconvey/convey/reporting/reporter.go",
        "github.com/smartystreets/goconvey/convey/reporting/reports.go",
        "github.com/smartystreets/goconvey/convey/reporting/statistics.go",
        "github.com/smartystreets/goconvey/convey/reporting/story.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/smartystreets/goconvey/convey/gotest"],
)

go_test(
    name = "github.com/stretchr/testify/assert_test",
    srcs = [
        "github.com/stretchr/testify/assert/assertions_test.go",
        "github.com/stretchr/testify/assert/forward_assertions_test.go",
        "github.com/stretchr/testify/assert/http_assertions_test.go",
    ],
    library = ":github.com/stretchr/testify/assert",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/stretchr/testify/assert",
    srcs = [
        "github.com/stretchr/testify/assert/assertion_forward.go",
        "github.com/stretchr/testify/assert/assertions.go",
        "github.com/stretchr/testify/assert/doc.go",
        "github.com/stretchr/testify/assert/errors.go",
        "github.com/stretchr/testify/assert/forward_assertions.go",
        "github.com/stretchr/testify/assert/http_assertions.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/davecgh/go-spew/spew",
        "//vendor:github.com/pmezard/go-difflib/difflib",
    ],
)

go_test(
    name = "github.com/stretchr/testify/require_test",
    srcs = [
        "github.com/stretchr/testify/require/forward_requirements_test.go",
        "github.com/stretchr/testify/require/requirements_test.go",
    ],
    library = ":github.com/stretchr/testify/require",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/stretchr/testify/require",
    srcs = [
        "github.com/stretchr/testify/require/doc.go",
        "github.com/stretchr/testify/require/forward_requirements.go",
        "github.com/stretchr/testify/require/require.go",
        "github.com/stretchr/testify/require/require_forward.go",
        "github.com/stretchr/testify/require/requirements.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/stretchr/testify/assert"],
)

go_library(
    name = "github.com/stretchr/testify/vendor/github.com/davecgh/go-spew/spew",
    srcs = [
        "github.com/stretchr/testify/vendor/github.com/davecgh/go-spew/spew/bypass.go",
        "github.com/stretchr/testify/vendor/github.com/davecgh/go-spew/spew/common.go",
        "github.com/stretchr/testify/vendor/github.com/davecgh/go-spew/spew/config.go",
        "github.com/stretchr/testify/vendor/github.com/davecgh/go-spew/spew/doc.go",
        "github.com/stretchr/testify/vendor/github.com/davecgh/go-spew/spew/dump.go",
        "github.com/stretchr/testify/vendor/github.com/davecgh/go-spew/spew/format.go",
        "github.com/stretchr/testify/vendor/github.com/davecgh/go-spew/spew/spew.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/stretchr/testify/vendor/github.com/pmezard/go-difflib/difflib",
    srcs = ["github.com/stretchr/testify/vendor/github.com/pmezard/go-difflib/difflib/difflib.go"],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/zenazn/goji/bind",
    srcs = [
        "github.com/zenazn/goji/bind/bind.go",
        "github.com/zenazn/goji/bind/einhorn.go",
        "github.com/zenazn/goji/bind/systemd.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/zenazn/goji/graceful",
    srcs = [
        "github.com/zenazn/goji/graceful/clone16.go",
        "github.com/zenazn/goji/graceful/einhorn.go",
        "github.com/zenazn/goji/graceful/graceful.go",
        "github.com/zenazn/goji/graceful/serve13.go",
        "github.com/zenazn/goji/graceful/server.go",
        "github.com/zenazn/goji/graceful/signal.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/zenazn/goji/graceful/listener"],
)

go_test(
    name = "github.com/zenazn/goji/graceful/listener_test",
    srcs = [
        "github.com/zenazn/goji/graceful/listener/conn_test.go",
        "github.com/zenazn/goji/graceful/listener/fake_test.go",
        "github.com/zenazn/goji/graceful/listener/listener_test.go",
        "github.com/zenazn/goji/graceful/listener/race_test.go",
    ],
    library = ":github.com/zenazn/goji/graceful/listener",
    tags = ["automanaged"],
)

go_library(
    name = "github.com/zenazn/goji/graceful/listener",
    srcs = [
        "github.com/zenazn/goji/graceful/listener/conn.go",
        "github.com/zenazn/goji/graceful/listener/listener.go",
        "github.com/zenazn/goji/graceful/listener/shard.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/zenazn/goji/web/mutil",
    srcs = [
        "github.com/zenazn/goji/web/mutil/mutil.go",
        "github.com/zenazn/goji/web/mutil/writer_proxy.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "golang.org/x/net/context",
    srcs = [
        "golang.org/x/net/context/context.go",
        "golang.org/x/net/context/go17.go",
    ],
    tags = ["automanaged"],
)

go_test(
    name = "golang.org/x/net/context_xtest",
    srcs = ["golang.org/x/net/context/withtimeout_test.go"],
    tags = ["automanaged"],
    deps = ["//vendor:golang.org/x/net/context"],
)

go_test(
    name = "golang.org/x/net/context/ctxhttp_test",
    srcs = [
        "golang.org/x/net/context/ctxhttp/ctxhttp_17_test.go",
        "golang.org/x/net/context/ctxhttp/ctxhttp_test.go",
    ],
    library = ":golang.org/x/net/context/ctxhttp",
    tags = ["automanaged"],
    deps = ["//vendor:golang.org/x/net/context"],
)

go_library(
    name = "golang.org/x/net/context/ctxhttp",
    srcs = ["golang.org/x/net/context/ctxhttp/ctxhttp.go"],
    tags = ["automanaged"],
    deps = ["//vendor:golang.org/x/net/context"],
)

go_library(
    name = "golang.org/x/sync/errgroup",
    srcs = ["golang.org/x/sync/errgroup/errgroup.go"],
    tags = ["automanaged"],
    deps = ["//vendor:golang.org/x/net/context"],
)

go_test(
    name = "golang.org/x/sync/errgroup_xtest",
    srcs = [
        "golang.org/x/sync/errgroup/errgroup_example_md5all_test.go",
        "golang.org/x/sync/errgroup/errgroup_test.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/net/context",
        "//vendor:golang.org/x/sync/errgroup",
    ],
)

go_test(
    name = "google.golang.org/appengine_test",
    srcs = [
        "google.golang.org/appengine/appengine_test.go",
        "google.golang.org/appengine/namespace_test.go",
    ],
    library = ":google.golang.org/appengine",
    tags = ["automanaged"],
    deps = ["//vendor:golang.org/x/net/context"],
)

go_library(
    name = "google.golang.org/appengine",
    srcs = [
        "google.golang.org/appengine/appengine.go",
        "google.golang.org/appengine/appengine_vm.go",
        "google.golang.org/appengine/errors.go",
        "google.golang.org/appengine/identity.go",
        "google.golang.org/appengine/namespace.go",
        "google.golang.org/appengine/timeout.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/app_identity",
        "//vendor:google.golang.org/appengine/internal/modules",
    ],
)

go_test(
    name = "google.golang.org/appengine/aetest_test",
    srcs = ["google.golang.org/appengine/aetest/instance_test.go"],
    library = ":google.golang.org/appengine/aetest",
    tags = ["automanaged"],
    deps = [
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/datastore",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/memcache",
        "//vendor:google.golang.org/appengine/user",
    ],
)

go_library(
    name = "google.golang.org/appengine/aetest",
    srcs = [
        "google.golang.org/appengine/aetest/doc.go",
        "google.golang.org/appengine/aetest/instance.go",
        "google.golang.org/appengine/aetest/instance_vm.go",
        "google.golang.org/appengine/aetest/user.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/user",
    ],
)

go_test(
    name = "google.golang.org/appengine/blobstore_test",
    srcs = ["google.golang.org/appengine/blobstore/blobstore_test.go"],
    library = ":google.golang.org/appengine/blobstore",
    tags = ["automanaged"],
    deps = [
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/internal/aetesting",
        "//vendor:google.golang.org/appengine/internal/blobstore",
    ],
)

go_library(
    name = "google.golang.org/appengine/blobstore",
    srcs = [
        "google.golang.org/appengine/blobstore/blobstore.go",
        "google.golang.org/appengine/blobstore/read.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/datastore",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/base",
        "//vendor:google.golang.org/appengine/internal/blobstore",
    ],
)

go_library(
    name = "google.golang.org/appengine/capability",
    srcs = ["google.golang.org/appengine/capability/capability.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/capability",
        "//vendor:google.golang.org/appengine/log",
    ],
)

go_test(
    name = "google.golang.org/appengine/channel_test",
    srcs = ["google.golang.org/appengine/channel/channel_test.go"],
    library = ":google.golang.org/appengine/channel",
    tags = ["automanaged"],
    deps = ["//vendor:google.golang.org/appengine/internal"],
)

go_library(
    name = "google.golang.org/appengine/channel",
    srcs = ["google.golang.org/appengine/channel/channel.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/base",
        "//vendor:google.golang.org/appengine/internal/channel",
    ],
)

go_library(
    name = "google.golang.org/appengine/cloudsql",
    srcs = [
        "google.golang.org/appengine/cloudsql/cloudsql.go",
        "google.golang.org/appengine/cloudsql/cloudsql_vm.go",
    ],
    tags = ["automanaged"],
)

go_binary(
    name = "google.golang.org/appengine/cmd/aebundler_bin",
    library = ":google.golang.org/appengine/cmd/aebundler",
    tags = ["automanaged"],
)

go_library(
    name = "google.golang.org/appengine/cmd/aebundler",
    srcs = ["google.golang.org/appengine/cmd/aebundler/aebundler.go"],
    tags = ["automanaged"],
)

go_binary(
    name = "google.golang.org/appengine/cmd/aedeploy_bin",
    library = ":google.golang.org/appengine/cmd/aedeploy",
    tags = ["automanaged"],
)

go_library(
    name = "google.golang.org/appengine/cmd/aedeploy",
    srcs = ["google.golang.org/appengine/cmd/aedeploy/aedeploy.go"],
    tags = ["automanaged"],
)

go_binary(
    name = "google.golang.org/appengine/cmd/aefix_bin",
    library = ":google.golang.org/appengine/cmd/aefix",
    tags = ["automanaged"],
)

go_test(
    name = "google.golang.org/appengine/cmd/aefix_test",
    srcs = [
        "google.golang.org/appengine/cmd/aefix/ae_test.go",
        "google.golang.org/appengine/cmd/aefix/main_test.go",
    ],
    library = ":google.golang.org/appengine/cmd/aefix",
    tags = ["automanaged"],
)

go_library(
    name = "google.golang.org/appengine/cmd/aefix",
    srcs = [
        "google.golang.org/appengine/cmd/aefix/ae.go",
        "google.golang.org/appengine/cmd/aefix/fix.go",
        "google.golang.org/appengine/cmd/aefix/main.go",
        "google.golang.org/appengine/cmd/aefix/typecheck.go",
    ],
    tags = ["automanaged"],
)

go_test(
    name = "google.golang.org/appengine/datastore_test",
    srcs = [
        "google.golang.org/appengine/datastore/datastore_test.go",
        "google.golang.org/appengine/datastore/key_test.go",
        "google.golang.org/appengine/datastore/load_test.go",
        "google.golang.org/appengine/datastore/prop_test.go",
        "google.golang.org/appengine/datastore/query_test.go",
        "google.golang.org/appengine/datastore/time_test.go",
    ],
    library = ":google.golang.org/appengine/datastore",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/aetesting",
        "//vendor:google.golang.org/appengine/internal/datastore",
    ],
)

go_library(
    name = "google.golang.org/appengine/datastore",
    srcs = [
        "google.golang.org/appengine/datastore/datastore.go",
        "google.golang.org/appengine/datastore/doc.go",
        "google.golang.org/appengine/datastore/key.go",
        "google.golang.org/appengine/datastore/load.go",
        "google.golang.org/appengine/datastore/metadata.go",
        "google.golang.org/appengine/datastore/prop.go",
        "google.golang.org/appengine/datastore/query.go",
        "google.golang.org/appengine/datastore/save.go",
        "google.golang.org/appengine/datastore/transaction.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/datastore",
    ],
)

go_test(
    name = "google.golang.org/appengine/delay_test",
    srcs = ["google.golang.org/appengine/delay/delay_test.go"],
    library = ":google.golang.org/appengine/delay",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/taskqueue",
    ],
)

go_library(
    name = "google.golang.org/appengine/delay",
    srcs = ["google.golang.org/appengine/delay/delay.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/log",
        "//vendor:google.golang.org/appengine/taskqueue",
    ],
)

go_binary(
    name = "google.golang.org/appengine/demos/guestbook_bin",
    library = ":google.golang.org/appengine/demos/guestbook",
    tags = ["automanaged"],
)

go_library(
    name = "google.golang.org/appengine/demos/guestbook",
    srcs = ["google.golang.org/appengine/demos/guestbook/guestbook.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/datastore",
        "//vendor:google.golang.org/appengine/log",
        "//vendor:google.golang.org/appengine/user",
    ],
)

go_binary(
    name = "google.golang.org/appengine/demos/helloworld_bin",
    library = ":google.golang.org/appengine/demos/helloworld",
    tags = ["automanaged"],
)

go_library(
    name = "google.golang.org/appengine/demos/helloworld",
    srcs = ["google.golang.org/appengine/demos/helloworld/helloworld.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/log",
    ],
)

go_library(
    name = "google.golang.org/appengine/file",
    srcs = ["google.golang.org/appengine/file/file.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/app_identity",
    ],
)

go_library(
    name = "google.golang.org/appengine/image",
    srcs = ["google.golang.org/appengine/image/image.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/image",
    ],
)

go_test(
    name = "google.golang.org/appengine/internal_test",
    srcs = [
        "google.golang.org/appengine/internal/api_test.go",
        "google.golang.org/appengine/internal/app_id_test.go",
        "google.golang.org/appengine/internal/internal_vm_test.go",
        "google.golang.org/appengine/internal/net_test.go",
    ],
    library = ":google.golang.org/appengine/internal",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine/internal/base",
        "//vendor:google.golang.org/appengine/internal/remote_api",
    ],
)

go_library(
    name = "google.golang.org/appengine/internal",
    srcs = [
        "google.golang.org/appengine/internal/api.go",
        "google.golang.org/appengine/internal/api_common.go",
        "google.golang.org/appengine/internal/app_id.go",
        "google.golang.org/appengine/internal/identity.go",
        "google.golang.org/appengine/internal/identity_vm.go",
        "google.golang.org/appengine/internal/internal.go",
        "google.golang.org/appengine/internal/main_vm.go",
        "google.golang.org/appengine/internal/metadata.go",
        "google.golang.org/appengine/internal/net.go",
        "google.golang.org/appengine/internal/transaction.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine/internal/base",
        "//vendor:google.golang.org/appengine/internal/datastore",
        "//vendor:google.golang.org/appengine/internal/log",
        "//vendor:google.golang.org/appengine/internal/remote_api",
    ],
)

go_library(
    name = "google.golang.org/appengine/internal/aetesting",
    srcs = ["google.golang.org/appengine/internal/aetesting/fake.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine/internal",
    ],
)

go_library(
    name = "google.golang.org/appengine/internal/app_identity",
    srcs = ["google.golang.org/appengine/internal/app_identity/app_identity_service.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "google.golang.org/appengine/internal/base",
    srcs = ["google.golang.org/appengine/internal/base/api_base.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "google.golang.org/appengine/internal/blobstore",
    srcs = ["google.golang.org/appengine/internal/blobstore/blobstore_service.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "google.golang.org/appengine/internal/capability",
    srcs = ["google.golang.org/appengine/internal/capability/capability_service.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "google.golang.org/appengine/internal/channel",
    srcs = ["google.golang.org/appengine/internal/channel/channel_service.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "google.golang.org/appengine/internal/datastore",
    srcs = ["google.golang.org/appengine/internal/datastore/datastore_v3.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "google.golang.org/appengine/internal/image",
    srcs = ["google.golang.org/appengine/internal/image/images_service.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "google.golang.org/appengine/internal/log",
    srcs = ["google.golang.org/appengine/internal/log/log_service.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "google.golang.org/appengine/internal/mail",
    srcs = ["google.golang.org/appengine/internal/mail/mail_service.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "google.golang.org/appengine/internal/memcache",
    srcs = ["google.golang.org/appengine/internal/memcache/memcache_service.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "google.golang.org/appengine/internal/modules",
    srcs = ["google.golang.org/appengine/internal/modules/modules_service.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "google.golang.org/appengine/internal/remote_api",
    srcs = ["google.golang.org/appengine/internal/remote_api/remote_api.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "google.golang.org/appengine/internal/search",
    srcs = ["google.golang.org/appengine/internal/search/search.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "google.golang.org/appengine/internal/socket",
    srcs = ["google.golang.org/appengine/internal/socket/socket_service.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "google.golang.org/appengine/internal/system",
    srcs = ["google.golang.org/appengine/internal/system/system_service.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "google.golang.org/appengine/internal/taskqueue",
    srcs = ["google.golang.org/appengine/internal/taskqueue/taskqueue_service.pb.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:google.golang.org/appengine/internal/datastore",
    ],
)

go_library(
    name = "google.golang.org/appengine/internal/urlfetch",
    srcs = ["google.golang.org/appengine/internal/urlfetch/urlfetch_service.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "google.golang.org/appengine/internal/user",
    srcs = ["google.golang.org/appengine/internal/user/user_service.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "google.golang.org/appengine/internal/xmpp",
    srcs = ["google.golang.org/appengine/internal/xmpp/xmpp_service.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_test(
    name = "google.golang.org/appengine/log_test",
    srcs = ["google.golang.org/appengine/log/log_test.go"],
    library = ":google.golang.org/appengine/log",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:google.golang.org/appengine/internal/log",
    ],
)

go_library(
    name = "google.golang.org/appengine/log",
    srcs = [
        "google.golang.org/appengine/log/api.go",
        "google.golang.org/appengine/log/log.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/log",
    ],
)

go_test(
    name = "google.golang.org/appengine/mail_test",
    srcs = ["google.golang.org/appengine/mail/mail_test.go"],
    library = ":google.golang.org/appengine/mail",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:google.golang.org/appengine/internal/aetesting",
        "//vendor:google.golang.org/appengine/internal/base",
        "//vendor:google.golang.org/appengine/internal/mail",
    ],
)

go_library(
    name = "google.golang.org/appengine/mail",
    srcs = ["google.golang.org/appengine/mail/mail.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/base",
        "//vendor:google.golang.org/appengine/internal/mail",
    ],
)

go_test(
    name = "google.golang.org/appengine/memcache_test",
    srcs = ["google.golang.org/appengine/memcache/memcache_test.go"],
    library = ":google.golang.org/appengine/memcache",
    tags = ["automanaged"],
    deps = [
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/internal/aetesting",
        "//vendor:google.golang.org/appengine/internal/memcache",
    ],
)

go_library(
    name = "google.golang.org/appengine/memcache",
    srcs = ["google.golang.org/appengine/memcache/memcache.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/memcache",
    ],
)

go_test(
    name = "google.golang.org/appengine/module_test",
    srcs = ["google.golang.org/appengine/module/module_test.go"],
    library = ":google.golang.org/appengine/module",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:google.golang.org/appengine/internal/aetesting",
        "//vendor:google.golang.org/appengine/internal/modules",
    ],
)

go_library(
    name = "google.golang.org/appengine/module",
    srcs = ["google.golang.org/appengine/module/module.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/modules",
    ],
)

go_test(
    name = "google.golang.org/appengine/remote_api_test",
    srcs = ["google.golang.org/appengine/remote_api/client_test.go"],
    library = ":google.golang.org/appengine/remote_api",
    tags = ["automanaged"],
)

go_library(
    name = "google.golang.org/appengine/remote_api",
    srcs = [
        "google.golang.org/appengine/remote_api/client.go",
        "google.golang.org/appengine/remote_api/remote_api.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/remote_api",
        "//vendor:google.golang.org/appengine/log",
        "//vendor:google.golang.org/appengine/user",
    ],
)

go_test(
    name = "google.golang.org/appengine/runtime_test",
    srcs = ["google.golang.org/appengine/runtime/runtime_test.go"],
    library = ":google.golang.org/appengine/runtime",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine/internal/aetesting",
        "//vendor:google.golang.org/appengine/internal/system",
    ],
)

go_library(
    name = "google.golang.org/appengine/runtime",
    srcs = ["google.golang.org/appengine/runtime/runtime.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/system",
    ],
)

go_test(
    name = "google.golang.org/appengine/search_test",
    srcs = [
        "google.golang.org/appengine/search/search_test.go",
        "google.golang.org/appengine/search/struct_test.go",
    ],
    library = ":google.golang.org/appengine/search",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/internal/aetesting",
        "//vendor:google.golang.org/appengine/internal/search",
    ],
)

go_library(
    name = "google.golang.org/appengine/search",
    srcs = [
        "google.golang.org/appengine/search/doc.go",
        "google.golang.org/appengine/search/field.go",
        "google.golang.org/appengine/search/search.go",
        "google.golang.org/appengine/search/struct.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/search",
    ],
)

go_library(
    name = "google.golang.org/appengine/socket",
    srcs = [
        "google.golang.org/appengine/socket/doc.go",
        "google.golang.org/appengine/socket/socket_vm.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:golang.org/x/net/context"],
)

go_test(
    name = "google.golang.org/appengine/taskqueue_test",
    srcs = ["google.golang.org/appengine/taskqueue/taskqueue_test.go"],
    library = ":google.golang.org/appengine/taskqueue",
    tags = ["automanaged"],
    deps = [
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/aetesting",
        "//vendor:google.golang.org/appengine/internal/taskqueue",
    ],
)

go_library(
    name = "google.golang.org/appengine/taskqueue",
    srcs = ["google.golang.org/appengine/taskqueue/taskqueue.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/datastore",
        "//vendor:google.golang.org/appengine/internal/taskqueue",
    ],
)

go_library(
    name = "google.golang.org/appengine/urlfetch",
    srcs = ["google.golang.org/appengine/urlfetch/urlfetch.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/urlfetch",
    ],
)

go_test(
    name = "google.golang.org/appengine/user_test",
    srcs = ["google.golang.org/appengine/user/user_test.go"],
    library = ":google.golang.org/appengine/user",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/aetesting",
        "//vendor:google.golang.org/appengine/internal/user",
    ],
)

go_library(
    name = "google.golang.org/appengine/user",
    srcs = [
        "google.golang.org/appengine/user/oauth.go",
        "google.golang.org/appengine/user/user.go",
        "google.golang.org/appengine/user/user_vm.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/user",
    ],
)

go_test(
    name = "google.golang.org/appengine/xmpp_test",
    srcs = ["google.golang.org/appengine/xmpp/xmpp_test.go"],
    library = ":google.golang.org/appengine/xmpp",
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/internal/aetesting",
        "//vendor:google.golang.org/appengine/internal/xmpp",
    ],
)

go_library(
    name = "google.golang.org/appengine/xmpp",
    srcs = ["google.golang.org/appengine/xmpp/xmpp.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/net/context",
        "//vendor:google.golang.org/appengine",
        "//vendor:google.golang.org/appengine/internal",
        "//vendor:google.golang.org/appengine/internal/xmpp",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/hashicorp/errwrap",
    srcs = ["github.com/hashicorp/errwrap/errwrap.go"],
    tags = ["automanaged"],
)
