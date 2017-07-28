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
        "github.com/aws/aws-sdk-go/aws/context.go",
        "github.com/aws/aws-sdk-go/aws/context_1_7.go",
        "github.com/aws/aws-sdk-go/aws/convert_types.go",
        "github.com/aws/aws-sdk-go/aws/doc.go",
        "github.com/aws/aws-sdk-go/aws/errors.go",
        "github.com/aws/aws-sdk-go/aws/jsonvalue.go",
        "github.com/aws/aws-sdk-go/aws/logger.go",
        "github.com/aws/aws-sdk-go/aws/types.go",
        "github.com/aws/aws-sdk-go/aws/url_1_7.go",
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
        "github.com/aws/aws-sdk-go/aws/client/logger.go",
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
        "//vendor:github.com/aws/aws-sdk-go/internal/shareddefaults",
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
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor:github.com/aws/aws-sdk-go/service/sts",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/aws/defaults",
    srcs = [
        "github.com/aws/aws-sdk-go/aws/defaults/defaults.go",
        "github.com/aws/aws-sdk-go/aws/defaults/shared_config.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/aws/corehandlers",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials/endpointcreds",
        "//vendor:github.com/aws/aws-sdk-go/aws/ec2metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/endpoints",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/internal/shareddefaults",
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
        "github.com/aws/aws-sdk-go/aws/request/connection_reset_error.go",
        "github.com/aws/aws-sdk-go/aws/request/handlers.go",
        "github.com/aws/aws-sdk-go/aws/request/http_request.go",
        "github.com/aws/aws-sdk-go/aws/request/offset_reader.go",
        "github.com/aws/aws-sdk-go/aws/request/request.go",
        "github.com/aws/aws-sdk-go/aws/request/request_1_7.go",
        "github.com/aws/aws-sdk-go/aws/request/request_context.go",
        "github.com/aws/aws-sdk-go/aws/request/request_pagination.go",
        "github.com/aws/aws-sdk-go/aws/request/retryer.go",
        "github.com/aws/aws-sdk-go/aws/request/timeout_read_closer.go",
        "github.com/aws/aws-sdk-go/aws/request/validation.go",
        "github.com/aws/aws-sdk-go/aws/request/waiter.go",
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
        "github.com/aws/aws-sdk-go/aws/signer/v4/options.go",
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
    name = "github.com/aws/aws-sdk-go/service/kms",
    srcs = [
        "github.com/aws/aws-sdk-go/service/kms/api.go",
        "github.com/aws/aws-sdk-go/service/kms/doc.go",
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
        "//vendor:github.com/aws/aws-sdk-go/aws",
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
        "github.com/aws/aws-sdk-go/service/s3/doc.go",
        "github.com/aws/aws-sdk-go/service/s3/doc_custom.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/s3/s3crypto",
    srcs = [
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/aes_cbc.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/aes_cbc_content_cipher.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/aes_cbc_padder.go",
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
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/padder.go",
        "github.com/aws/aws-sdk-go/service/s3/s3crypto/pkcs7_padder.go",
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
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/s3",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/s3/s3manager",
    srcs = [
        "github.com/aws/aws-sdk-go/service/s3/s3manager/batch.go",
        "github.com/aws/aws-sdk-go/service/s3/s3manager/bucket_region.go",
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
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
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
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/service/s3",
        "//vendor:github.com/aws/aws-sdk-go/service/s3/s3manager",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/sts",
    srcs = [
        "github.com/aws/aws-sdk-go/service/sts/api.go",
        "github.com/aws/aws-sdk-go/service/sts/customizations.go",
        "github.com/aws/aws-sdk-go/service/sts/doc.go",
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
        "//vendor:github.com/aws/aws-sdk-go/aws",
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

go_library(
    name = "github.com/beorn7/perks/quantile",
    srcs = ["github.com/beorn7/perks/quantile/stream.go"],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/matttproud/golang_protobuf_extensions/pbutil",
    srcs = [
        "github.com/matttproud/golang_protobuf_extensions/pbutil/decode.go",
        "github.com/matttproud/golang_protobuf_extensions/pbutil/doc.go",
        "github.com/matttproud/golang_protobuf_extensions/pbutil/encode.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "github.com/prometheus/client_golang/prometheus",
    srcs = [
        "github.com/prometheus/client_golang/prometheus/collector.go",
        "github.com/prometheus/client_golang/prometheus/counter.go",
        "github.com/prometheus/client_golang/prometheus/desc.go",
        "github.com/prometheus/client_golang/prometheus/doc.go",
        "github.com/prometheus/client_golang/prometheus/expvar_collector.go",
        "github.com/prometheus/client_golang/prometheus/fnv.go",
        "github.com/prometheus/client_golang/prometheus/gauge.go",
        "github.com/prometheus/client_golang/prometheus/go_collector.go",
        "github.com/prometheus/client_golang/prometheus/histogram.go",
        "github.com/prometheus/client_golang/prometheus/http.go",
        "github.com/prometheus/client_golang/prometheus/metric.go",
        "github.com/prometheus/client_golang/prometheus/observer.go",
        "github.com/prometheus/client_golang/prometheus/process_collector.go",
        "github.com/prometheus/client_golang/prometheus/registry.go",
        "github.com/prometheus/client_golang/prometheus/summary.go",
        "github.com/prometheus/client_golang/prometheus/timer.go",
        "github.com/prometheus/client_golang/prometheus/untyped.go",
        "github.com/prometheus/client_golang/prometheus/value.go",
        "github.com/prometheus/client_golang/prometheus/vec.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/beorn7/perks/quantile",
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:github.com/prometheus/client_model/go",
        "//vendor:github.com/prometheus/common/expfmt",
        "//vendor:github.com/prometheus/common/model",
        "//vendor:github.com/prometheus/procfs",
    ],
)

go_library(
    name = "github.com/prometheus/client_golang/prometheus/graphite",
    srcs = ["github.com/prometheus/client_golang/prometheus/graphite/bridge.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/prometheus/client_golang/prometheus",
        "//vendor:github.com/prometheus/client_model/go",
        "//vendor:github.com/prometheus/common/expfmt",
        "//vendor:github.com/prometheus/common/model",
        "//vendor:golang.org/x/net/context",
    ],
)

go_library(
    name = "github.com/prometheus/client_golang/prometheus/promhttp",
    srcs = [
        "github.com/prometheus/client_golang/prometheus/promhttp/delegator.go",
        "github.com/prometheus/client_golang/prometheus/promhttp/delegator_pre_1_8.go",
        "github.com/prometheus/client_golang/prometheus/promhttp/http.go",
        "github.com/prometheus/client_golang/prometheus/promhttp/instrument_client.go",
        "github.com/prometheus/client_golang/prometheus/promhttp/instrument_server.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/prometheus/client_golang/prometheus",
        "//vendor:github.com/prometheus/client_model/go",
        "//vendor:github.com/prometheus/common/expfmt",
    ],
)

go_library(
    name = "github.com/prometheus/client_golang/prometheus/push",
    srcs = ["github.com/prometheus/client_golang/prometheus/push/push.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/prometheus/client_golang/prometheus",
        "//vendor:github.com/prometheus/common/expfmt",
        "//vendor:github.com/prometheus/common/model",
    ],
)

go_library(
    name = "github.com/prometheus/client_model/go",
    srcs = ["github.com/prometheus/client_model/go/metrics.pb.go"],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/golang/protobuf/proto"],
)

go_library(
    name = "github.com/prometheus/common/expfmt",
    srcs = [
        "github.com/prometheus/common/expfmt/decode.go",
        "github.com/prometheus/common/expfmt/encode.go",
        "github.com/prometheus/common/expfmt/expfmt.go",
        "github.com/prometheus/common/expfmt/text_create.go",
        "github.com/prometheus/common/expfmt/text_parse.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/golang/protobuf/proto",
        "//vendor:github.com/matttproud/golang_protobuf_extensions/pbutil",
        "//vendor:github.com/prometheus/client_model/go",
        "//vendor:github.com/prometheus/common/internal/bitbucket.org/ww/goautoneg",
        "//vendor:github.com/prometheus/common/model",
    ],
)

go_library(
    name = "github.com/prometheus/common/internal/bitbucket.org/ww/goautoneg",
    srcs = ["github.com/prometheus/common/internal/bitbucket.org/ww/goautoneg/autoneg.go"],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/prometheus/common/model",
    srcs = [
        "github.com/prometheus/common/model/alert.go",
        "github.com/prometheus/common/model/fingerprinting.go",
        "github.com/prometheus/common/model/fnv.go",
        "github.com/prometheus/common/model/labels.go",
        "github.com/prometheus/common/model/labelset.go",
        "github.com/prometheus/common/model/metric.go",
        "github.com/prometheus/common/model/model.go",
        "github.com/prometheus/common/model/signature.go",
        "github.com/prometheus/common/model/silence.go",
        "github.com/prometheus/common/model/time.go",
        "github.com/prometheus/common/model/value.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/prometheus/procfs",
    srcs = [
        "github.com/prometheus/procfs/buddyinfo.go",
        "github.com/prometheus/procfs/doc.go",
        "github.com/prometheus/procfs/fs.go",
        "github.com/prometheus/procfs/ipvs.go",
        "github.com/prometheus/procfs/mdstat.go",
        "github.com/prometheus/procfs/mountstats.go",
        "github.com/prometheus/procfs/proc.go",
        "github.com/prometheus/procfs/proc_io.go",
        "github.com/prometheus/procfs/proc_limits.go",
        "github.com/prometheus/procfs/proc_stat.go",
        "github.com/prometheus/procfs/stat.go",
        "github.com/prometheus/procfs/xfrm.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/prometheus/procfs/xfs"],
)

go_library(
    name = "github.com/prometheus/procfs/bcache",
    srcs = [
        "github.com/prometheus/procfs/bcache/bcache.go",
        "github.com/prometheus/procfs/bcache/get.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/prometheus/procfs/sysfs",
    srcs = [
        "github.com/prometheus/procfs/sysfs/doc.go",
        "github.com/prometheus/procfs/sysfs/fs.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/prometheus/procfs/bcache",
        "//vendor:github.com/prometheus/procfs/xfs",
    ],
)

go_library(
    name = "github.com/prometheus/procfs/xfs",
    srcs = [
        "github.com/prometheus/procfs/xfs/parse.go",
        "github.com/prometheus/procfs/xfs/xfs.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/aws/aws-sdk-go",
    srcs = ["github.com/aws/aws-sdk-go/doc.go"],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/aws/aws-sdk-go/aws/credentials/plugincreds",
    srcs = ["github.com/aws/aws-sdk-go/aws/credentials/plugincreds/doc_1_7.go"],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/aws/aws-sdk-go/awsmigrate/awsmigrate-renamer/vendor/golang.org/x/tools/go/ast/astutil",
    srcs = [
        "github.com/aws/aws-sdk-go/awsmigrate/awsmigrate-renamer/vendor/golang.org/x/tools/go/ast/astutil/enclosing.go",
        "github.com/aws/aws-sdk-go/awsmigrate/awsmigrate-renamer/vendor/golang.org/x/tools/go/ast/astutil/imports.go",
        "github.com/aws/aws-sdk-go/awsmigrate/awsmigrate-renamer/vendor/golang.org/x/tools/go/ast/astutil/util.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/aws/aws-sdk-go/awsmigrate/awsmigrate-renamer/vendor/golang.org/x/tools/go/buildutil",
    srcs = [
        "github.com/aws/aws-sdk-go/awsmigrate/awsmigrate-renamer/vendor/golang.org/x/tools/go/buildutil/allpackages.go",
        "github.com/aws/aws-sdk-go/awsmigrate/awsmigrate-renamer/vendor/golang.org/x/tools/go/buildutil/fakecontext.go",
        "github.com/aws/aws-sdk-go/awsmigrate/awsmigrate-renamer/vendor/golang.org/x/tools/go/buildutil/tags.go",
        "github.com/aws/aws-sdk-go/awsmigrate/awsmigrate-renamer/vendor/golang.org/x/tools/go/buildutil/util.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/aws/aws-sdk-go/awsmigrate/awsmigrate-renamer/vendor/golang.org/x/tools/go/loader",
    srcs = [
        "github.com/aws/aws-sdk-go/awsmigrate/awsmigrate-renamer/vendor/golang.org/x/tools/go/loader/cgo.go",
        "github.com/aws/aws-sdk-go/awsmigrate/awsmigrate-renamer/vendor/golang.org/x/tools/go/loader/cgo_pkgconfig.go",
        "github.com/aws/aws-sdk-go/awsmigrate/awsmigrate-renamer/vendor/golang.org/x/tools/go/loader/doc.go",
        "github.com/aws/aws-sdk-go/awsmigrate/awsmigrate-renamer/vendor/golang.org/x/tools/go/loader/go16.go",
        "github.com/aws/aws-sdk-go/awsmigrate/awsmigrate-renamer/vendor/golang.org/x/tools/go/loader/loader.go",
        "github.com/aws/aws-sdk-go/awsmigrate/awsmigrate-renamer/vendor/golang.org/x/tools/go/loader/util.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/tools/go/ast/astutil",
        "//vendor:golang.org/x/tools/go/buildutil",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/awstesting",
    srcs = [
        "github.com/aws/aws-sdk-go/awstesting/assert.go",
        "github.com/aws/aws-sdk-go/awstesting/client.go",
        "github.com/aws/aws-sdk-go/awstesting/custom_ca_bundle.go",
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
    name = "github.com/aws/aws-sdk-go/internal/shareddefaults",
    srcs = ["github.com/aws/aws-sdk-go/internal/shareddefaults/shared_config.go"],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/aws/aws-sdk-go/models/endpoints",
    srcs = ["github.com/aws/aws-sdk-go/models/endpoints/generate.go"],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/aws/aws-sdk-go/private/model/api",
    srcs = ["github.com/aws/aws-sdk-go/private/model/api/list_of_shame.go"],
    tags = ["automanaged"],
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
    name = "github.com/aws/aws-sdk-go/service",
    srcs = ["github.com/aws/aws-sdk-go/service/generate.go"],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/acm",
    srcs = [
        "github.com/aws/aws-sdk-go/service/acm/api.go",
        "github.com/aws/aws-sdk-go/service/acm/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/acm/acmiface",
    srcs = ["github.com/aws/aws-sdk-go/service/acm/acmiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/acm",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/apigateway",
    srcs = [
        "github.com/aws/aws-sdk-go/service/apigateway/api.go",
        "github.com/aws/aws-sdk-go/service/apigateway/customization.go",
        "github.com/aws/aws-sdk-go/service/apigateway/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/apigateway/apigatewayiface",
    srcs = ["github.com/aws/aws-sdk-go/service/apigateway/apigatewayiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/apigateway",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/applicationautoscaling",
    srcs = [
        "github.com/aws/aws-sdk-go/service/applicationautoscaling/api.go",
        "github.com/aws/aws-sdk-go/service/applicationautoscaling/doc.go",
        "github.com/aws/aws-sdk-go/service/applicationautoscaling/errors.go",
        "github.com/aws/aws-sdk-go/service/applicationautoscaling/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/applicationautoscaling/applicationautoscalingiface",
    srcs = ["github.com/aws/aws-sdk-go/service/applicationautoscaling/applicationautoscalingiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/applicationautoscaling",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/applicationdiscoveryservice",
    srcs = [
        "github.com/aws/aws-sdk-go/service/applicationdiscoveryservice/api.go",
        "github.com/aws/aws-sdk-go/service/applicationdiscoveryservice/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/applicationdiscoveryservice/applicationdiscoveryserviceiface",
    srcs = ["github.com/aws/aws-sdk-go/service/applicationdiscoveryservice/applicationdiscoveryserviceiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/applicationdiscoveryservice",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/appstream",
    srcs = [
        "github.com/aws/aws-sdk-go/service/appstream/api.go",
        "github.com/aws/aws-sdk-go/service/appstream/doc.go",
        "github.com/aws/aws-sdk-go/service/appstream/errors.go",
        "github.com/aws/aws-sdk-go/service/appstream/service.go",
        "github.com/aws/aws-sdk-go/service/appstream/waiters.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/appstream/appstreamiface",
    srcs = ["github.com/aws/aws-sdk-go/service/appstream/appstreamiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/appstream",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/athena",
    srcs = [
        "github.com/aws/aws-sdk-go/service/athena/api.go",
        "github.com/aws/aws-sdk-go/service/athena/doc.go",
        "github.com/aws/aws-sdk-go/service/athena/errors.go",
        "github.com/aws/aws-sdk-go/service/athena/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/athena/athenaiface",
    srcs = ["github.com/aws/aws-sdk-go/service/athena/athenaiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/athena",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/autoscaling",
    srcs = [
        "github.com/aws/aws-sdk-go/service/autoscaling/api.go",
        "github.com/aws/aws-sdk-go/service/autoscaling/doc.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface",
    srcs = ["github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/autoscaling",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/batch",
    srcs = [
        "github.com/aws/aws-sdk-go/service/batch/api.go",
        "github.com/aws/aws-sdk-go/service/batch/doc.go",
        "github.com/aws/aws-sdk-go/service/batch/errors.go",
        "github.com/aws/aws-sdk-go/service/batch/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/batch/batchiface",
    srcs = ["github.com/aws/aws-sdk-go/service/batch/batchiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/batch",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/budgets",
    srcs = [
        "github.com/aws/aws-sdk-go/service/budgets/api.go",
        "github.com/aws/aws-sdk-go/service/budgets/doc.go",
        "github.com/aws/aws-sdk-go/service/budgets/errors.go",
        "github.com/aws/aws-sdk-go/service/budgets/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/budgets/budgetsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/budgets/budgetsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/budgets",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/clouddirectory",
    srcs = [
        "github.com/aws/aws-sdk-go/service/clouddirectory/api.go",
        "github.com/aws/aws-sdk-go/service/clouddirectory/doc.go",
        "github.com/aws/aws-sdk-go/service/clouddirectory/errors.go",
        "github.com/aws/aws-sdk-go/service/clouddirectory/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/clouddirectory/clouddirectoryiface",
    srcs = ["github.com/aws/aws-sdk-go/service/clouddirectory/clouddirectoryiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/clouddirectory",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudformation",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cloudformation/api.go",
        "github.com/aws/aws-sdk-go/service/cloudformation/doc.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudformation",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudfront",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cloudfront/api.go",
        "github.com/aws/aws-sdk-go/service/cloudfront/doc.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudfront/cloudfrontiface",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudfront/cloudfrontiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudfront",
    ],
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
        "github.com/aws/aws-sdk-go/service/cloudhsm/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudhsm/cloudhsmiface",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudhsm/cloudhsmiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudhsm",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudsearch",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cloudsearch/api.go",
        "github.com/aws/aws-sdk-go/service/cloudsearch/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudsearch/cloudsearchiface",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudsearch/cloudsearchiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudsearch",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudsearchdomain",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cloudsearchdomain/api.go",
        "github.com/aws/aws-sdk-go/service/cloudsearchdomain/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudsearchdomain/cloudsearchdomainiface",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudsearchdomain/cloudsearchdomainiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudsearchdomain",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudtrail",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cloudtrail/api.go",
        "github.com/aws/aws-sdk-go/service/cloudtrail/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudtrail/cloudtrailiface",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudtrail/cloudtrailiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudtrail",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudwatch",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cloudwatch/api.go",
        "github.com/aws/aws-sdk-go/service/cloudwatch/doc.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudwatch",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudwatchevents",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cloudwatchevents/api.go",
        "github.com/aws/aws-sdk-go/service/cloudwatchevents/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudwatchevents/cloudwatcheventsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudwatchevents/cloudwatcheventsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudwatchevents",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudwatchlogs",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cloudwatchlogs/api.go",
        "github.com/aws/aws-sdk-go/service/cloudwatchlogs/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cloudwatchlogs",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/codebuild",
    srcs = [
        "github.com/aws/aws-sdk-go/service/codebuild/api.go",
        "github.com/aws/aws-sdk-go/service/codebuild/doc.go",
        "github.com/aws/aws-sdk-go/service/codebuild/errors.go",
        "github.com/aws/aws-sdk-go/service/codebuild/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/codebuild/codebuildiface",
    srcs = ["github.com/aws/aws-sdk-go/service/codebuild/codebuildiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/codebuild",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/codecommit",
    srcs = [
        "github.com/aws/aws-sdk-go/service/codecommit/api.go",
        "github.com/aws/aws-sdk-go/service/codecommit/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/codecommit/codecommitiface",
    srcs = ["github.com/aws/aws-sdk-go/service/codecommit/codecommitiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/codecommit",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/codedeploy",
    srcs = [
        "github.com/aws/aws-sdk-go/service/codedeploy/api.go",
        "github.com/aws/aws-sdk-go/service/codedeploy/doc.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/codedeploy/codedeployiface",
    srcs = ["github.com/aws/aws-sdk-go/service/codedeploy/codedeployiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/codedeploy",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/codepipeline",
    srcs = [
        "github.com/aws/aws-sdk-go/service/codepipeline/api.go",
        "github.com/aws/aws-sdk-go/service/codepipeline/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/codepipeline/codepipelineiface",
    srcs = ["github.com/aws/aws-sdk-go/service/codepipeline/codepipelineiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/codepipeline",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/codestar",
    srcs = [
        "github.com/aws/aws-sdk-go/service/codestar/api.go",
        "github.com/aws/aws-sdk-go/service/codestar/doc.go",
        "github.com/aws/aws-sdk-go/service/codestar/errors.go",
        "github.com/aws/aws-sdk-go/service/codestar/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/codestar/codestariface",
    srcs = ["github.com/aws/aws-sdk-go/service/codestar/codestariface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/codestar",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cognitoidentity",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cognitoidentity/api.go",
        "github.com/aws/aws-sdk-go/service/cognitoidentity/customizations.go",
        "github.com/aws/aws-sdk-go/service/cognitoidentity/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/cognitoidentity/cognitoidentityiface",
    srcs = ["github.com/aws/aws-sdk-go/service/cognitoidentity/cognitoidentityiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cognitoidentity",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cognitoidentityprovider",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cognitoidentityprovider/api.go",
        "github.com/aws/aws-sdk-go/service/cognitoidentityprovider/doc.go",
        "github.com/aws/aws-sdk-go/service/cognitoidentityprovider/errors.go",
        "github.com/aws/aws-sdk-go/service/cognitoidentityprovider/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/jsonrpc",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface",
    srcs = ["github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cognitoidentityprovider",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/cognitosync",
    srcs = [
        "github.com/aws/aws-sdk-go/service/cognitosync/api.go",
        "github.com/aws/aws-sdk-go/service/cognitosync/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/cognitosync/cognitosynciface",
    srcs = ["github.com/aws/aws-sdk-go/service/cognitosync/cognitosynciface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/cognitosync",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/configservice",
    srcs = [
        "github.com/aws/aws-sdk-go/service/configservice/api.go",
        "github.com/aws/aws-sdk-go/service/configservice/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/configservice/configserviceiface",
    srcs = ["github.com/aws/aws-sdk-go/service/configservice/configserviceiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/configservice",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/costandusagereportservice",
    srcs = [
        "github.com/aws/aws-sdk-go/service/costandusagereportservice/api.go",
        "github.com/aws/aws-sdk-go/service/costandusagereportservice/doc.go",
        "github.com/aws/aws-sdk-go/service/costandusagereportservice/errors.go",
        "github.com/aws/aws-sdk-go/service/costandusagereportservice/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/costandusagereportservice/costandusagereportserviceiface",
    srcs = ["github.com/aws/aws-sdk-go/service/costandusagereportservice/costandusagereportserviceiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/costandusagereportservice",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/databasemigrationservice",
    srcs = [
        "github.com/aws/aws-sdk-go/service/databasemigrationservice/api.go",
        "github.com/aws/aws-sdk-go/service/databasemigrationservice/doc.go",
        "github.com/aws/aws-sdk-go/service/databasemigrationservice/errors.go",
        "github.com/aws/aws-sdk-go/service/databasemigrationservice/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/databasemigrationservice/databasemigrationserviceiface",
    srcs = ["github.com/aws/aws-sdk-go/service/databasemigrationservice/databasemigrationserviceiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/databasemigrationservice",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/datapipeline",
    srcs = [
        "github.com/aws/aws-sdk-go/service/datapipeline/api.go",
        "github.com/aws/aws-sdk-go/service/datapipeline/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/datapipeline/datapipelineiface",
    srcs = ["github.com/aws/aws-sdk-go/service/datapipeline/datapipelineiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/datapipeline",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/dax",
    srcs = [
        "github.com/aws/aws-sdk-go/service/dax/api.go",
        "github.com/aws/aws-sdk-go/service/dax/doc.go",
        "github.com/aws/aws-sdk-go/service/dax/errors.go",
        "github.com/aws/aws-sdk-go/service/dax/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/dax/daxiface",
    srcs = ["github.com/aws/aws-sdk-go/service/dax/daxiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/dax",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/devicefarm",
    srcs = [
        "github.com/aws/aws-sdk-go/service/devicefarm/api.go",
        "github.com/aws/aws-sdk-go/service/devicefarm/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/devicefarm/devicefarmiface",
    srcs = ["github.com/aws/aws-sdk-go/service/devicefarm/devicefarmiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/devicefarm",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/directconnect",
    srcs = [
        "github.com/aws/aws-sdk-go/service/directconnect/api.go",
        "github.com/aws/aws-sdk-go/service/directconnect/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/directconnect/directconnectiface",
    srcs = ["github.com/aws/aws-sdk-go/service/directconnect/directconnectiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/directconnect",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/directoryservice",
    srcs = [
        "github.com/aws/aws-sdk-go/service/directoryservice/api.go",
        "github.com/aws/aws-sdk-go/service/directoryservice/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/directoryservice/directoryserviceiface",
    srcs = ["github.com/aws/aws-sdk-go/service/directoryservice/directoryserviceiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/directoryservice",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/dynamodb",
    srcs = [
        "github.com/aws/aws-sdk-go/service/dynamodb/api.go",
        "github.com/aws/aws-sdk-go/service/dynamodb/customizations.go",
        "github.com/aws/aws-sdk-go/service/dynamodb/doc.go",
        "github.com/aws/aws-sdk-go/service/dynamodb/doc_custom.go",
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
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awserr",
        "//vendor:github.com/aws/aws-sdk-go/service/dynamodb",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface",
    srcs = ["github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/dynamodb",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/dynamodbstreams",
    srcs = [
        "github.com/aws/aws-sdk-go/service/dynamodbstreams/api.go",
        "github.com/aws/aws-sdk-go/service/dynamodbstreams/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/dynamodbstreams/dynamodbstreamsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/dynamodbstreams/dynamodbstreamsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/dynamodbstreams",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/ec2",
    srcs = [
        "github.com/aws/aws-sdk-go/service/ec2/api.go",
        "github.com/aws/aws-sdk-go/service/ec2/customizations.go",
        "github.com/aws/aws-sdk-go/service/ec2/doc.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/ec2/ec2iface",
    srcs = ["github.com/aws/aws-sdk-go/service/ec2/ec2iface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/ec2",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/ecr",
    srcs = [
        "github.com/aws/aws-sdk-go/service/ecr/api.go",
        "github.com/aws/aws-sdk-go/service/ecr/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/ecr/ecriface",
    srcs = ["github.com/aws/aws-sdk-go/service/ecr/ecriface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/ecr",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/ecs",
    srcs = [
        "github.com/aws/aws-sdk-go/service/ecs/api.go",
        "github.com/aws/aws-sdk-go/service/ecs/doc.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/ecs/ecsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/ecs/ecsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/ecs",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/efs",
    srcs = [
        "github.com/aws/aws-sdk-go/service/efs/api.go",
        "github.com/aws/aws-sdk-go/service/efs/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/efs/efsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/efs/efsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/efs",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/elasticache",
    srcs = [
        "github.com/aws/aws-sdk-go/service/elasticache/api.go",
        "github.com/aws/aws-sdk-go/service/elasticache/doc.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/elasticache/elasticacheiface",
    srcs = ["github.com/aws/aws-sdk-go/service/elasticache/elasticacheiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/elasticache",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/elasticbeanstalk",
    srcs = [
        "github.com/aws/aws-sdk-go/service/elasticbeanstalk/api.go",
        "github.com/aws/aws-sdk-go/service/elasticbeanstalk/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/elasticbeanstalk/elasticbeanstalkiface",
    srcs = ["github.com/aws/aws-sdk-go/service/elasticbeanstalk/elasticbeanstalkiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/elasticbeanstalk",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/elasticsearchservice",
    srcs = [
        "github.com/aws/aws-sdk-go/service/elasticsearchservice/api.go",
        "github.com/aws/aws-sdk-go/service/elasticsearchservice/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/elasticsearchservice/elasticsearchserviceiface",
    srcs = ["github.com/aws/aws-sdk-go/service/elasticsearchservice/elasticsearchserviceiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/elasticsearchservice",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/elastictranscoder",
    srcs = [
        "github.com/aws/aws-sdk-go/service/elastictranscoder/api.go",
        "github.com/aws/aws-sdk-go/service/elastictranscoder/doc.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/elastictranscoder/elastictranscoderiface",
    srcs = ["github.com/aws/aws-sdk-go/service/elastictranscoder/elastictranscoderiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/elastictranscoder",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/elb",
    srcs = [
        "github.com/aws/aws-sdk-go/service/elb/api.go",
        "github.com/aws/aws-sdk-go/service/elb/doc.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/elb/elbiface",
    srcs = ["github.com/aws/aws-sdk-go/service/elb/elbiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/elb",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/elbv2",
    srcs = [
        "github.com/aws/aws-sdk-go/service/elbv2/api.go",
        "github.com/aws/aws-sdk-go/service/elbv2/doc.go",
        "github.com/aws/aws-sdk-go/service/elbv2/errors.go",
        "github.com/aws/aws-sdk-go/service/elbv2/service.go",
        "github.com/aws/aws-sdk-go/service/elbv2/waiters.go",
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
    name = "github.com/aws/aws-sdk-go/service/elbv2/elbv2iface",
    srcs = ["github.com/aws/aws-sdk-go/service/elbv2/elbv2iface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/elbv2",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/emr",
    srcs = [
        "github.com/aws/aws-sdk-go/service/emr/api.go",
        "github.com/aws/aws-sdk-go/service/emr/doc.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/emr/emriface",
    srcs = ["github.com/aws/aws-sdk-go/service/emr/emriface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/emr",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/firehose",
    srcs = [
        "github.com/aws/aws-sdk-go/service/firehose/api.go",
        "github.com/aws/aws-sdk-go/service/firehose/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/firehose/firehoseiface",
    srcs = ["github.com/aws/aws-sdk-go/service/firehose/firehoseiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/firehose",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/gamelift",
    srcs = [
        "github.com/aws/aws-sdk-go/service/gamelift/api.go",
        "github.com/aws/aws-sdk-go/service/gamelift/doc.go",
        "github.com/aws/aws-sdk-go/service/gamelift/errors.go",
        "github.com/aws/aws-sdk-go/service/gamelift/service.go",
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
    name = "github.com/aws/aws-sdk-go/service/gamelift/gameliftiface",
    srcs = ["github.com/aws/aws-sdk-go/service/gamelift/gameliftiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/gamelift",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/glacier",
    srcs = [
        "github.com/aws/aws-sdk-go/service/glacier/api.go",
        "github.com/aws/aws-sdk-go/service/glacier/customizations.go",
        "github.com/aws/aws-sdk-go/service/glacier/doc.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/glacier/glacieriface",
    srcs = ["github.com/aws/aws-sdk-go/service/glacier/glacieriface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/glacier",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/greengrass",
    srcs = [
        "github.com/aws/aws-sdk-go/service/greengrass/api.go",
        "github.com/aws/aws-sdk-go/service/greengrass/doc.go",
        "github.com/aws/aws-sdk-go/service/greengrass/errors.go",
        "github.com/aws/aws-sdk-go/service/greengrass/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/greengrass/greengrassiface",
    srcs = ["github.com/aws/aws-sdk-go/service/greengrass/greengrassiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/greengrass",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/health",
    srcs = [
        "github.com/aws/aws-sdk-go/service/health/api.go",
        "github.com/aws/aws-sdk-go/service/health/doc.go",
        "github.com/aws/aws-sdk-go/service/health/errors.go",
        "github.com/aws/aws-sdk-go/service/health/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/health/healthiface",
    srcs = ["github.com/aws/aws-sdk-go/service/health/healthiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/health",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/iam",
    srcs = [
        "github.com/aws/aws-sdk-go/service/iam/api.go",
        "github.com/aws/aws-sdk-go/service/iam/doc.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/iam/iamiface",
    srcs = ["github.com/aws/aws-sdk-go/service/iam/iamiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/iam",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/inspector",
    srcs = [
        "github.com/aws/aws-sdk-go/service/inspector/api.go",
        "github.com/aws/aws-sdk-go/service/inspector/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/inspector/inspectoriface",
    srcs = ["github.com/aws/aws-sdk-go/service/inspector/inspectoriface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/inspector",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/iot",
    srcs = [
        "github.com/aws/aws-sdk-go/service/iot/api.go",
        "github.com/aws/aws-sdk-go/service/iot/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/iot/iotiface",
    srcs = ["github.com/aws/aws-sdk-go/service/iot/iotiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/iot",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/iotdataplane",
    srcs = [
        "github.com/aws/aws-sdk-go/service/iotdataplane/api.go",
        "github.com/aws/aws-sdk-go/service/iotdataplane/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/iotdataplane/iotdataplaneiface",
    srcs = ["github.com/aws/aws-sdk-go/service/iotdataplane/iotdataplaneiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/iotdataplane",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/kinesis",
    srcs = [
        "github.com/aws/aws-sdk-go/service/kinesis/api.go",
        "github.com/aws/aws-sdk-go/service/kinesis/customizations.go",
        "github.com/aws/aws-sdk-go/service/kinesis/doc.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/kinesis/kinesisiface",
    srcs = ["github.com/aws/aws-sdk-go/service/kinesis/kinesisiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/kinesis",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/kinesisanalytics",
    srcs = [
        "github.com/aws/aws-sdk-go/service/kinesisanalytics/api.go",
        "github.com/aws/aws-sdk-go/service/kinesisanalytics/doc.go",
        "github.com/aws/aws-sdk-go/service/kinesisanalytics/errors.go",
        "github.com/aws/aws-sdk-go/service/kinesisanalytics/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/kinesisanalytics/kinesisanalyticsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/kinesisanalytics/kinesisanalyticsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/kinesisanalytics",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/lambda",
    srcs = [
        "github.com/aws/aws-sdk-go/service/lambda/api.go",
        "github.com/aws/aws-sdk-go/service/lambda/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/lambda/lambdaiface",
    srcs = ["github.com/aws/aws-sdk-go/service/lambda/lambdaiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/lambda",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/lexmodelbuildingservice",
    srcs = [
        "github.com/aws/aws-sdk-go/service/lexmodelbuildingservice/api.go",
        "github.com/aws/aws-sdk-go/service/lexmodelbuildingservice/doc.go",
        "github.com/aws/aws-sdk-go/service/lexmodelbuildingservice/errors.go",
        "github.com/aws/aws-sdk-go/service/lexmodelbuildingservice/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/lexmodelbuildingservice/lexmodelbuildingserviceiface",
    srcs = ["github.com/aws/aws-sdk-go/service/lexmodelbuildingservice/lexmodelbuildingserviceiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/lexmodelbuildingservice",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/lexruntimeservice",
    srcs = [
        "github.com/aws/aws-sdk-go/service/lexruntimeservice/api.go",
        "github.com/aws/aws-sdk-go/service/lexruntimeservice/doc.go",
        "github.com/aws/aws-sdk-go/service/lexruntimeservice/errors.go",
        "github.com/aws/aws-sdk-go/service/lexruntimeservice/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/lexruntimeservice/lexruntimeserviceiface",
    srcs = ["github.com/aws/aws-sdk-go/service/lexruntimeservice/lexruntimeserviceiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/lexruntimeservice",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/lightsail",
    srcs = [
        "github.com/aws/aws-sdk-go/service/lightsail/api.go",
        "github.com/aws/aws-sdk-go/service/lightsail/doc.go",
        "github.com/aws/aws-sdk-go/service/lightsail/errors.go",
        "github.com/aws/aws-sdk-go/service/lightsail/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/lightsail/lightsailiface",
    srcs = ["github.com/aws/aws-sdk-go/service/lightsail/lightsailiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/lightsail",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/machinelearning",
    srcs = [
        "github.com/aws/aws-sdk-go/service/machinelearning/api.go",
        "github.com/aws/aws-sdk-go/service/machinelearning/customizations.go",
        "github.com/aws/aws-sdk-go/service/machinelearning/doc.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/machinelearning/machinelearningiface",
    srcs = ["github.com/aws/aws-sdk-go/service/machinelearning/machinelearningiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/machinelearning",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics",
    srcs = [
        "github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics/api.go",
        "github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics/marketplacecommerceanalyticsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics/marketplacecommerceanalyticsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/marketplaceentitlementservice",
    srcs = [
        "github.com/aws/aws-sdk-go/service/marketplaceentitlementservice/api.go",
        "github.com/aws/aws-sdk-go/service/marketplaceentitlementservice/doc.go",
        "github.com/aws/aws-sdk-go/service/marketplaceentitlementservice/errors.go",
        "github.com/aws/aws-sdk-go/service/marketplaceentitlementservice/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/marketplaceentitlementservice/marketplaceentitlementserviceiface",
    srcs = ["github.com/aws/aws-sdk-go/service/marketplaceentitlementservice/marketplaceentitlementserviceiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/marketplaceentitlementservice",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/marketplacemetering",
    srcs = [
        "github.com/aws/aws-sdk-go/service/marketplacemetering/api.go",
        "github.com/aws/aws-sdk-go/service/marketplacemetering/doc.go",
        "github.com/aws/aws-sdk-go/service/marketplacemetering/errors.go",
        "github.com/aws/aws-sdk-go/service/marketplacemetering/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/marketplacemetering/marketplacemeteringiface",
    srcs = ["github.com/aws/aws-sdk-go/service/marketplacemetering/marketplacemeteringiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/marketplacemetering",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/mobileanalytics",
    srcs = [
        "github.com/aws/aws-sdk-go/service/mobileanalytics/api.go",
        "github.com/aws/aws-sdk-go/service/mobileanalytics/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/mobileanalytics/mobileanalyticsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/mobileanalytics/mobileanalyticsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/mobileanalytics",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/mturk",
    srcs = [
        "github.com/aws/aws-sdk-go/service/mturk/api.go",
        "github.com/aws/aws-sdk-go/service/mturk/doc.go",
        "github.com/aws/aws-sdk-go/service/mturk/errors.go",
        "github.com/aws/aws-sdk-go/service/mturk/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/mturk/mturkiface",
    srcs = ["github.com/aws/aws-sdk-go/service/mturk/mturkiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/mturk",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/opsworks",
    srcs = [
        "github.com/aws/aws-sdk-go/service/opsworks/api.go",
        "github.com/aws/aws-sdk-go/service/opsworks/doc.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/opsworks/opsworksiface",
    srcs = ["github.com/aws/aws-sdk-go/service/opsworks/opsworksiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/opsworks",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/opsworkscm",
    srcs = [
        "github.com/aws/aws-sdk-go/service/opsworkscm/api.go",
        "github.com/aws/aws-sdk-go/service/opsworkscm/doc.go",
        "github.com/aws/aws-sdk-go/service/opsworkscm/errors.go",
        "github.com/aws/aws-sdk-go/service/opsworkscm/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/opsworkscm/opsworkscmiface",
    srcs = ["github.com/aws/aws-sdk-go/service/opsworkscm/opsworkscmiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/opsworkscm",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/organizations",
    srcs = [
        "github.com/aws/aws-sdk-go/service/organizations/api.go",
        "github.com/aws/aws-sdk-go/service/organizations/doc.go",
        "github.com/aws/aws-sdk-go/service/organizations/errors.go",
        "github.com/aws/aws-sdk-go/service/organizations/service.go",
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
    name = "github.com/aws/aws-sdk-go/service/organizations/organizationsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/organizations/organizationsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/organizations",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/pinpoint",
    srcs = [
        "github.com/aws/aws-sdk-go/service/pinpoint/api.go",
        "github.com/aws/aws-sdk-go/service/pinpoint/doc.go",
        "github.com/aws/aws-sdk-go/service/pinpoint/errors.go",
        "github.com/aws/aws-sdk-go/service/pinpoint/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/pinpoint/pinpointiface",
    srcs = ["github.com/aws/aws-sdk-go/service/pinpoint/pinpointiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/pinpoint",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/polly",
    srcs = [
        "github.com/aws/aws-sdk-go/service/polly/api.go",
        "github.com/aws/aws-sdk-go/service/polly/customizations.go",
        "github.com/aws/aws-sdk-go/service/polly/doc.go",
        "github.com/aws/aws-sdk-go/service/polly/errors.go",
        "github.com/aws/aws-sdk-go/service/polly/service.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/awsutil",
        "//vendor:github.com/aws/aws-sdk-go/aws/client",
        "//vendor:github.com/aws/aws-sdk-go/aws/client/metadata",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/rest",
        "//vendor:github.com/aws/aws-sdk-go/private/protocol/restjson",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/polly/pollyiface",
    srcs = ["github.com/aws/aws-sdk-go/service/polly/pollyiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/polly",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/rds",
    srcs = [
        "github.com/aws/aws-sdk-go/service/rds/api.go",
        "github.com/aws/aws-sdk-go/service/rds/customizations.go",
        "github.com/aws/aws-sdk-go/service/rds/doc.go",
        "github.com/aws/aws-sdk-go/service/rds/doc_custom.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/rds/rdsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/rds/rdsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/rds",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/rds/rdsutils",
    srcs = ["github.com/aws/aws-sdk-go/service/rds/rdsutils/connect.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws/credentials",
        "//vendor:github.com/aws/aws-sdk-go/aws/signer/v4",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/redshift",
    srcs = [
        "github.com/aws/aws-sdk-go/service/redshift/api.go",
        "github.com/aws/aws-sdk-go/service/redshift/doc.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/redshift/redshiftiface",
    srcs = ["github.com/aws/aws-sdk-go/service/redshift/redshiftiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/redshift",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/rekognition",
    srcs = [
        "github.com/aws/aws-sdk-go/service/rekognition/api.go",
        "github.com/aws/aws-sdk-go/service/rekognition/doc.go",
        "github.com/aws/aws-sdk-go/service/rekognition/errors.go",
        "github.com/aws/aws-sdk-go/service/rekognition/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/rekognition/rekognitioniface",
    srcs = ["github.com/aws/aws-sdk-go/service/rekognition/rekognitioniface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/rekognition",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi",
    srcs = [
        "github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi/api.go",
        "github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi/doc.go",
        "github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi/errors.go",
        "github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi/resourcegroupstaggingapiiface",
    srcs = ["github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi/resourcegroupstaggingapiiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/route53",
    srcs = [
        "github.com/aws/aws-sdk-go/service/route53/api.go",
        "github.com/aws/aws-sdk-go/service/route53/customizations.go",
        "github.com/aws/aws-sdk-go/service/route53/doc.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/route53/route53iface",
    srcs = ["github.com/aws/aws-sdk-go/service/route53/route53iface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/route53",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/route53domains",
    srcs = [
        "github.com/aws/aws-sdk-go/service/route53domains/api.go",
        "github.com/aws/aws-sdk-go/service/route53domains/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/route53domains/route53domainsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/route53domains/route53domainsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/route53domains",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/servicecatalog",
    srcs = [
        "github.com/aws/aws-sdk-go/service/servicecatalog/api.go",
        "github.com/aws/aws-sdk-go/service/servicecatalog/doc.go",
        "github.com/aws/aws-sdk-go/service/servicecatalog/errors.go",
        "github.com/aws/aws-sdk-go/service/servicecatalog/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/servicecatalog/servicecatalogiface",
    srcs = ["github.com/aws/aws-sdk-go/service/servicecatalog/servicecatalogiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/servicecatalog",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/ses",
    srcs = [
        "github.com/aws/aws-sdk-go/service/ses/api.go",
        "github.com/aws/aws-sdk-go/service/ses/doc.go",
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
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/ses/sesiface",
    srcs = ["github.com/aws/aws-sdk-go/service/ses/sesiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/ses",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/sfn",
    srcs = [
        "github.com/aws/aws-sdk-go/service/sfn/api.go",
        "github.com/aws/aws-sdk-go/service/sfn/doc.go",
        "github.com/aws/aws-sdk-go/service/sfn/errors.go",
        "github.com/aws/aws-sdk-go/service/sfn/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/sfn/sfniface",
    srcs = ["github.com/aws/aws-sdk-go/service/sfn/sfniface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/sfn",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/shield",
    srcs = [
        "github.com/aws/aws-sdk-go/service/shield/api.go",
        "github.com/aws/aws-sdk-go/service/shield/doc.go",
        "github.com/aws/aws-sdk-go/service/shield/errors.go",
        "github.com/aws/aws-sdk-go/service/shield/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/shield/shieldiface",
    srcs = ["github.com/aws/aws-sdk-go/service/shield/shieldiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/shield",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/simpledb",
    srcs = [
        "github.com/aws/aws-sdk-go/service/simpledb/api.go",
        "github.com/aws/aws-sdk-go/service/simpledb/customizations.go",
        "github.com/aws/aws-sdk-go/service/simpledb/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/simpledb/simpledbiface",
    srcs = ["github.com/aws/aws-sdk-go/service/simpledb/simpledbiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/simpledb",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/sms",
    srcs = [
        "github.com/aws/aws-sdk-go/service/sms/api.go",
        "github.com/aws/aws-sdk-go/service/sms/doc.go",
        "github.com/aws/aws-sdk-go/service/sms/errors.go",
        "github.com/aws/aws-sdk-go/service/sms/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/sms/smsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/sms/smsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/sms",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/snowball",
    srcs = [
        "github.com/aws/aws-sdk-go/service/snowball/api.go",
        "github.com/aws/aws-sdk-go/service/snowball/doc.go",
        "github.com/aws/aws-sdk-go/service/snowball/errors.go",
        "github.com/aws/aws-sdk-go/service/snowball/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/snowball/snowballiface",
    srcs = ["github.com/aws/aws-sdk-go/service/snowball/snowballiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/snowball",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/sns",
    srcs = [
        "github.com/aws/aws-sdk-go/service/sns/api.go",
        "github.com/aws/aws-sdk-go/service/sns/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/sns/snsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/sns/snsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
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
        "github.com/aws/aws-sdk-go/service/sqs/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/sqs/sqsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/sqs/sqsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/sqs",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/ssm",
    srcs = [
        "github.com/aws/aws-sdk-go/service/ssm/api.go",
        "github.com/aws/aws-sdk-go/service/ssm/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/ssm/ssmiface",
    srcs = ["github.com/aws/aws-sdk-go/service/ssm/ssmiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/ssm",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/storagegateway",
    srcs = [
        "github.com/aws/aws-sdk-go/service/storagegateway/api.go",
        "github.com/aws/aws-sdk-go/service/storagegateway/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/storagegateway/storagegatewayiface",
    srcs = ["github.com/aws/aws-sdk-go/service/storagegateway/storagegatewayiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/storagegateway",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/support",
    srcs = [
        "github.com/aws/aws-sdk-go/service/support/api.go",
        "github.com/aws/aws-sdk-go/service/support/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/support/supportiface",
    srcs = ["github.com/aws/aws-sdk-go/service/support/supportiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/support",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/swf",
    srcs = [
        "github.com/aws/aws-sdk-go/service/swf/api.go",
        "github.com/aws/aws-sdk-go/service/swf/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/swf/swfiface",
    srcs = ["github.com/aws/aws-sdk-go/service/swf/swfiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/swf",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/waf",
    srcs = [
        "github.com/aws/aws-sdk-go/service/waf/api.go",
        "github.com/aws/aws-sdk-go/service/waf/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/waf/wafiface",
    srcs = ["github.com/aws/aws-sdk-go/service/waf/wafiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/waf",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/wafregional",
    srcs = [
        "github.com/aws/aws-sdk-go/service/wafregional/api.go",
        "github.com/aws/aws-sdk-go/service/wafregional/doc.go",
        "github.com/aws/aws-sdk-go/service/wafregional/errors.go",
        "github.com/aws/aws-sdk-go/service/wafregional/service.go",
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
        "//vendor:github.com/aws/aws-sdk-go/service/waf",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/wafregional/wafregionaliface",
    srcs = ["github.com/aws/aws-sdk-go/service/wafregional/wafregionaliface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/waf",
        "//vendor:github.com/aws/aws-sdk-go/service/wafregional",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/workdocs",
    srcs = [
        "github.com/aws/aws-sdk-go/service/workdocs/api.go",
        "github.com/aws/aws-sdk-go/service/workdocs/doc.go",
        "github.com/aws/aws-sdk-go/service/workdocs/errors.go",
        "github.com/aws/aws-sdk-go/service/workdocs/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/workdocs/workdocsiface",
    srcs = ["github.com/aws/aws-sdk-go/service/workdocs/workdocsiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/workdocs",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/workspaces",
    srcs = [
        "github.com/aws/aws-sdk-go/service/workspaces/api.go",
        "github.com/aws/aws-sdk-go/service/workspaces/doc.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/workspaces/workspacesiface",
    srcs = ["github.com/aws/aws-sdk-go/service/workspaces/workspacesiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/workspaces",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/service/xray",
    srcs = [
        "github.com/aws/aws-sdk-go/service/xray/api.go",
        "github.com/aws/aws-sdk-go/service/xray/doc.go",
        "github.com/aws/aws-sdk-go/service/xray/errors.go",
        "github.com/aws/aws-sdk-go/service/xray/service.go",
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

go_library(
    name = "github.com/aws/aws-sdk-go/service/xray/xrayiface",
    srcs = ["github.com/aws/aws-sdk-go/service/xray/xrayiface/interface.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:github.com/aws/aws-sdk-go/aws",
        "//vendor:github.com/aws/aws-sdk-go/aws/request",
        "//vendor:github.com/aws/aws-sdk-go/service/xray",
    ],
)

go_library(
    name = "github.com/aws/aws-sdk-go/vendor/github.com/go-ini/ini",
    srcs = [
        "github.com/aws/aws-sdk-go/vendor/github.com/go-ini/ini/error.go",
        "github.com/aws/aws-sdk-go/vendor/github.com/go-ini/ini/ini.go",
        "github.com/aws/aws-sdk-go/vendor/github.com/go-ini/ini/key.go",
        "github.com/aws/aws-sdk-go/vendor/github.com/go-ini/ini/parser.go",
        "github.com/aws/aws-sdk-go/vendor/github.com/go-ini/ini/section.go",
        "github.com/aws/aws-sdk-go/vendor/github.com/go-ini/ini/struct.go",
    ],
    tags = ["automanaged"],
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

go_library(
    name = "github.com/go-sql-driver/mysql",
    srcs = [
        "github.com/go-sql-driver/mysql/buffer.go",
        "github.com/go-sql-driver/mysql/collations.go",
        "github.com/go-sql-driver/mysql/connection.go",
        "github.com/go-sql-driver/mysql/const.go",
        "github.com/go-sql-driver/mysql/driver.go",
        "github.com/go-sql-driver/mysql/dsn.go",
        "github.com/go-sql-driver/mysql/errors.go",
        "github.com/go-sql-driver/mysql/infile.go",
        "github.com/go-sql-driver/mysql/packets.go",
        "github.com/go-sql-driver/mysql/result.go",
        "github.com/go-sql-driver/mysql/rows.go",
        "github.com/go-sql-driver/mysql/statement.go",
        "github.com/go-sql-driver/mysql/transaction.go",
        "github.com/go-sql-driver/mysql/utils.go",
        "github.com/go-sql-driver/mysql/utils_go17.go",
    ],
    tags = ["automanaged"],
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
    name = "github.com/shiena/ansicolor",
    srcs = [
        "github.com/shiena/ansicolor/ansicolor.go",
        "github.com/shiena/ansicolor/ansicolor_ansi.go",
    ],
    tags = ["automanaged"],
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

go_library(
    name = "github.com/stretchr/testify/assert",
    srcs = [
        "github.com/stretchr/testify/assert/assertion_format.go",
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
    name = "golang.org/x/net/html",
    srcs = [
        "golang.org/x/net/html/const.go",
        "golang.org/x/net/html/doc.go",
        "golang.org/x/net/html/doctype.go",
        "golang.org/x/net/html/entity.go",
        "golang.org/x/net/html/escape.go",
        "golang.org/x/net/html/foreign.go",
        "golang.org/x/net/html/node.go",
        "golang.org/x/net/html/parse.go",
        "golang.org/x/net/html/render.go",
        "golang.org/x/net/html/token.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:golang.org/x/net/html/atom"],
)

go_library(
    name = "golang.org/x/net/html/atom",
    srcs = [
        "golang.org/x/net/html/atom/atom.go",
        "golang.org/x/net/html/atom/table.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "golang.org/x/net/html/charset",
    srcs = ["golang.org/x/net/html/charset/charset.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/net/html",
        "//vendor:golang.org/x/text/encoding",
        "//vendor:golang.org/x/text/encoding/charmap",
        "//vendor:golang.org/x/text/encoding/htmlindex",
        "//vendor:golang.org/x/text/transform",
    ],
)

go_library(
    name = "golang.org/x/text/encoding",
    srcs = ["golang.org/x/text/encoding/encoding.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/text/encoding/internal/identifier",
        "//vendor:golang.org/x/text/transform",
    ],
)

go_library(
    name = "golang.org/x/text/encoding/charmap",
    srcs = [
        "golang.org/x/text/encoding/charmap/charmap.go",
        "golang.org/x/text/encoding/charmap/tables.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/text/encoding",
        "//vendor:golang.org/x/text/encoding/internal",
        "//vendor:golang.org/x/text/encoding/internal/identifier",
        "//vendor:golang.org/x/text/transform",
    ],
)

go_library(
    name = "golang.org/x/text/encoding/htmlindex",
    srcs = [
        "golang.org/x/text/encoding/htmlindex/htmlindex.go",
        "golang.org/x/text/encoding/htmlindex/map.go",
        "golang.org/x/text/encoding/htmlindex/tables.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/text/encoding",
        "//vendor:golang.org/x/text/encoding/charmap",
        "//vendor:golang.org/x/text/encoding/internal/identifier",
        "//vendor:golang.org/x/text/encoding/japanese",
        "//vendor:golang.org/x/text/encoding/korean",
        "//vendor:golang.org/x/text/encoding/simplifiedchinese",
        "//vendor:golang.org/x/text/encoding/traditionalchinese",
        "//vendor:golang.org/x/text/encoding/unicode",
        "//vendor:golang.org/x/text/language",
    ],
)

go_library(
    name = "golang.org/x/text/encoding/ianaindex",
    srcs = [
        "golang.org/x/text/encoding/ianaindex/ianaindex.go",
        "golang.org/x/text/encoding/ianaindex/tables.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/text/encoding",
        "//vendor:golang.org/x/text/encoding/charmap",
        "//vendor:golang.org/x/text/encoding/internal/identifier",
        "//vendor:golang.org/x/text/encoding/japanese",
        "//vendor:golang.org/x/text/encoding/korean",
        "//vendor:golang.org/x/text/encoding/simplifiedchinese",
        "//vendor:golang.org/x/text/encoding/traditionalchinese",
        "//vendor:golang.org/x/text/encoding/unicode",
    ],
)

go_library(
    name = "golang.org/x/text/encoding/internal",
    srcs = ["golang.org/x/text/encoding/internal/internal.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/text/encoding",
        "//vendor:golang.org/x/text/encoding/internal/identifier",
        "//vendor:golang.org/x/text/transform",
    ],
)

go_library(
    name = "golang.org/x/text/encoding/internal/enctest",
    srcs = ["golang.org/x/text/encoding/internal/enctest/enctest.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/text/encoding",
        "//vendor:golang.org/x/text/encoding/internal/identifier",
        "//vendor:golang.org/x/text/transform",
    ],
)

go_library(
    name = "golang.org/x/text/encoding/internal/identifier",
    srcs = [
        "golang.org/x/text/encoding/internal/identifier/identifier.go",
        "golang.org/x/text/encoding/internal/identifier/mib.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "golang.org/x/text/encoding/japanese",
    srcs = [
        "golang.org/x/text/encoding/japanese/all.go",
        "golang.org/x/text/encoding/japanese/eucjp.go",
        "golang.org/x/text/encoding/japanese/iso2022jp.go",
        "golang.org/x/text/encoding/japanese/shiftjis.go",
        "golang.org/x/text/encoding/japanese/tables.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/text/encoding",
        "//vendor:golang.org/x/text/encoding/internal",
        "//vendor:golang.org/x/text/encoding/internal/identifier",
        "//vendor:golang.org/x/text/transform",
    ],
)

go_library(
    name = "golang.org/x/text/encoding/korean",
    srcs = [
        "golang.org/x/text/encoding/korean/euckr.go",
        "golang.org/x/text/encoding/korean/tables.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/text/encoding",
        "//vendor:golang.org/x/text/encoding/internal",
        "//vendor:golang.org/x/text/encoding/internal/identifier",
        "//vendor:golang.org/x/text/transform",
    ],
)

go_library(
    name = "golang.org/x/text/encoding/simplifiedchinese",
    srcs = [
        "golang.org/x/text/encoding/simplifiedchinese/all.go",
        "golang.org/x/text/encoding/simplifiedchinese/gbk.go",
        "golang.org/x/text/encoding/simplifiedchinese/hzgb2312.go",
        "golang.org/x/text/encoding/simplifiedchinese/tables.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/text/encoding",
        "//vendor:golang.org/x/text/encoding/internal",
        "//vendor:golang.org/x/text/encoding/internal/identifier",
        "//vendor:golang.org/x/text/transform",
    ],
)

go_library(
    name = "golang.org/x/text/encoding/traditionalchinese",
    srcs = [
        "golang.org/x/text/encoding/traditionalchinese/big5.go",
        "golang.org/x/text/encoding/traditionalchinese/tables.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/text/encoding",
        "//vendor:golang.org/x/text/encoding/internal",
        "//vendor:golang.org/x/text/encoding/internal/identifier",
        "//vendor:golang.org/x/text/transform",
    ],
)

go_library(
    name = "golang.org/x/text/encoding/unicode",
    srcs = [
        "golang.org/x/text/encoding/unicode/override.go",
        "golang.org/x/text/encoding/unicode/unicode.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/text/encoding",
        "//vendor:golang.org/x/text/encoding/internal",
        "//vendor:golang.org/x/text/encoding/internal/identifier",
        "//vendor:golang.org/x/text/internal/utf8internal",
        "//vendor:golang.org/x/text/runes",
        "//vendor:golang.org/x/text/transform",
    ],
)

go_library(
    name = "golang.org/x/text/encoding/unicode/utf32",
    srcs = ["golang.org/x/text/encoding/unicode/utf32/utf32.go"],
    tags = ["automanaged"],
    deps = [
        "//vendor:golang.org/x/text/encoding",
        "//vendor:golang.org/x/text/encoding/internal/identifier",
        "//vendor:golang.org/x/text/transform",
    ],
)

go_library(
    name = "golang.org/x/text/internal/gen",
    srcs = [
        "golang.org/x/text/internal/gen/code.go",
        "golang.org/x/text/internal/gen/gen.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:golang.org/x/text/unicode/cldr"],
)

go_library(
    name = "golang.org/x/text/internal/tag",
    srcs = ["golang.org/x/text/internal/tag/tag.go"],
    tags = ["automanaged"],
)

go_library(
    name = "golang.org/x/text/internal/utf8internal",
    srcs = ["golang.org/x/text/internal/utf8internal/utf8internal.go"],
    tags = ["automanaged"],
)

go_library(
    name = "golang.org/x/text/language",
    srcs = [
        "golang.org/x/text/language/common.go",
        "golang.org/x/text/language/coverage.go",
        "golang.org/x/text/language/go1_2.go",
        "golang.org/x/text/language/index.go",
        "golang.org/x/text/language/language.go",
        "golang.org/x/text/language/lookup.go",
        "golang.org/x/text/language/match.go",
        "golang.org/x/text/language/parse.go",
        "golang.org/x/text/language/tables.go",
        "golang.org/x/text/language/tags.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:golang.org/x/text/internal/tag"],
)

go_library(
    name = "golang.org/x/text/language/display",
    srcs = [
        "golang.org/x/text/language/display/dict.go",
        "golang.org/x/text/language/display/display.go",
        "golang.org/x/text/language/display/lookup.go",
        "golang.org/x/text/language/display/tables.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:golang.org/x/text/language"],
)

go_library(
    name = "golang.org/x/text/runes",
    srcs = [
        "golang.org/x/text/runes/cond.go",
        "golang.org/x/text/runes/runes.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:golang.org/x/text/transform"],
)

go_library(
    name = "golang.org/x/text/transform",
    srcs = ["golang.org/x/text/transform/transform.go"],
    tags = ["automanaged"],
)

go_library(
    name = "golang.org/x/text/unicode/cldr",
    srcs = [
        "golang.org/x/text/unicode/cldr/base.go",
        "golang.org/x/text/unicode/cldr/cldr.go",
        "golang.org/x/text/unicode/cldr/collate.go",
        "golang.org/x/text/unicode/cldr/decode.go",
        "golang.org/x/text/unicode/cldr/resolve.go",
        "golang.org/x/text/unicode/cldr/slice.go",
        "golang.org/x/text/unicode/cldr/xml.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/lib/pq",
    srcs = [
        "github.com/lib/pq/array.go",
        "github.com/lib/pq/buf.go",
        "github.com/lib/pq/conn.go",
        "github.com/lib/pq/copy.go",
        "github.com/lib/pq/doc.go",
        "github.com/lib/pq/encode.go",
        "github.com/lib/pq/error.go",
        "github.com/lib/pq/notify.go",
        "github.com/lib/pq/ssl.go",
        "github.com/lib/pq/ssl_go1.7.go",
        "github.com/lib/pq/ssl_permissions.go",
        "github.com/lib/pq/url.go",
        "github.com/lib/pq/user_posix.go",
        "github.com/lib/pq/uuid.go",
    ],
    tags = ["automanaged"],
    deps = ["//vendor:github.com/lib/pq/oid"],
)

go_library(
    name = "github.com/lib/pq/hstore",
    srcs = ["github.com/lib/pq/hstore/hstore.go"],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/lib/pq/listen_example",
    srcs = ["github.com/lib/pq/listen_example/doc.go"],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/lib/pq/oid",
    srcs = [
        "github.com/lib/pq/oid/doc.go",
        "github.com/lib/pq/oid/types.go",
    ],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/jmoiron/sqlx",
    srcs = glob(["github.com/jmoiron/sqlx/*.go"]),
    deps = ["//vendor:github.com/jmoiron/sqlx/reflectx"],
)

go_library(
    name = "github.com/jmoiron/sqlx/reflectx",
    srcs = ["github.com/jmoiron/sqlx/reflectx/reflect.go"],
    tags = ["automanaged"],
)

go_library(
    name = "github.com/jmoiron/sqlx/types",
    srcs = ["github.com/jmoiron/sqlx/types/types.go"],
    tags = ["automanaged"],
)
