package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_binary",
    "go_library",
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

go_library(
    name = "github.com/aws/aws-sdk-go/private/protocol",
    srcs = [
        "github.com/aws/aws-sdk-go/private/protocol/idempotency.go",
        "github.com/aws/aws-sdk-go/private/protocol/unmarshal.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/aws/aws-sdk-go/aws/request"],
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

go_library(
    name = "github.com/aws/aws-sdk-go/private/protocol/json/jsonutil",
    srcs = [
        "github.com/aws/aws-sdk-go/private/protocol/json/jsonutil/build.go",
        "github.com/aws/aws-sdk-go/private/protocol/json/jsonutil/unmarshal.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/aws/aws-sdk-go/private/protocol"],
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

go_library(
    name = "github.com/aws/aws-sdk-go/private/protocol/query/queryutil",
    srcs = ["github.com/aws/aws-sdk-go/private/protocol/query/queryutil/queryutil.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/aws/aws-sdk-go/private/protocol"],
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

go_library(
    name = "github.com/dimfeld/httppath",
    srcs = ["github.com/dimfeld/httppath/path.go"],
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

go_library(
    name = "github.com/golang/protobuf/ptypes/any",
    srcs = ["github.com/golang/protobuf/ptypes/any/any.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "github.com/google/shlex",
    srcs = ["github.com/google/shlex/shlex.go"],
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
    name = "github.com/kardianos/osext",
    srcs = [
        "github.com/kardianos/osext/osext.go",
        "github.com/kardianos/osext/osext_procfs.go",
    ],
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

go_library(
    name = "github.com/go-ini/ini",
    srcs = [
        "github.com/go-ini/ini/error.go",
        "github.com/go-ini/ini/ini.go",
        "github.com/go-ini/ini/key.go",
        "github.com/go-ini/ini/parser.go",
        "github.com/go-ini/ini/section.go",
        "github.com/go-ini/ini/struct.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/jmespath/go-jmespath",
    srcs = [
        "github.com/jmespath/go-jmespath/api.go",
        "github.com/jmespath/go-jmespath/astnodetype_string.go",
        "github.com/jmespath/go-jmespath/functions.go",
        "github.com/jmespath/go-jmespath/interpreter.go",
        "github.com/jmespath/go-jmespath/lexer.go",
        "github.com/jmespath/go-jmespath/parser.go",
        "github.com/jmespath/go-jmespath/toktype_string.go",
        "github.com/jmespath/go-jmespath/util.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/jmespath/go-jmespath/fuzz",
    srcs = ["github.com/jmespath/go-jmespath/fuzz/jmespath.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/jmespath/go-jmespath"],
)

go_library(
    name = "golang.org/x/net/context",
    srcs = [
        "golang.org/x/net/context/context.go",
        "golang.org/x/net/context/go17.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "golang.org/x/net/context/ctxhttp",
    srcs = ["golang.org/x/net/context/ctxhttp/ctxhttp.go"],
    tags = ["automanaged"],
    deps = ["//vendor:golang.org/x/net/context"],
)

go_library(
    name = "github.com/julienschmidt/httprouter",
    srcs = [
        "github.com/julienschmidt/httprouter/path.go",
        "github.com/julienschmidt/httprouter/router.go",
        "github.com/julienschmidt/httprouter/tree.go",
    ],
    tags = ["automanaged"],
)
