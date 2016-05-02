package statik

import (
		"github.com/rakyll/statik/fs"
)

func init() {
	data := "PK\x03\x04\x14\x00\x08\x00\x00\x00\xb9C\xa2H\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0f\x00\x00\x00defs-image.json{\n  \"description\": \"Definitions particular to OpenContainer Image Specification\",\n  \"definitions\": {\n    \"mediaType\": {\n      \"id\": \"https://opencontainers.org/schema/image/mediaType\",\n      \"type\": \"string\",\n      \"pattern\": \"^[a-z]+/[0-9a-zA-Z.+]+$\"\n    },\n    \"digest\": {\n      \"description\": \"the cryptographic checksum digest of the object, in the pattern '<hash>:<hexadecimal digest>'\",\n      \"type\": \"string\",\n      \"pattern\": \"^[a-z0-9]+:[a-fA-F0-9]+$\"\n    },\n    \"descriptor\": {\n      \"id\": \"https://opencontainers.org/schema/image/descriptor\",\n      \"type\": \"object\",\n      \"required\": [\n        \"mediaType\",\n        \"size\",\n        \"digest\"\n      ],\n      \"properties\": {\n        \"mediaType\": {\n          \"description\": \"the mediatype of the referenced object\",\n          \"$ref\": \"#definitions/mediaType\"\n        },\n        \"size\": {\n          \"description\": \"the size in bytes of the referenced object\",\n          \"type\": \"integer\"\n        },\n        \"digest\": {\n          \"$ref\": \"#definitions/digest\"\n        }\n      }\n    },\n    \"manifestDescriptor\": {\n      \"id\": \"https://opencontainers.org/schema/image/manifestDescriptor\",\n      \"type\": \"object\",\n      \"required\": [\n        \"mediaType\",\n        \"size\",\n        \"digest\",\n        \"platform\"\n      ],\n      \"properties\": {\n        \"mediaType\": {\n          \"description\": \"the mediatype of the referenced object\",\n          \"$ref\": \"#definitions/mediaType\"\n        },\n        \"size\": {\n          \"description\": \"the size in bytes of the referenced object\",\n          \"type\": \"integer\"\n        },\n        \"digest\": {\n          \"$ref\": \"#definitions/digest\"\n        },\n        \"platform\": {\n          \"id\": \"https://opencontainers.org/schema/image/platform\",\n          \"type\": \"object\",\n          \"required\": [\n            \"architecture\",\n            \"os\"\n          ],\n          \"properties\": {\n            \"architecture\": {\n              \"id\": \"https://opencontainers.org/schema/image/platform/architecture\",\n              \"type\": \"string\"\n            },\n            \"os\": {\n              \"id\": \"https://opencontainers.org/schema/image/platform/os\",\n              \"type\": \"string\"\n            },\n            \"os.version\": {\n              \"id\": \"https://opencontainers.org/schema/image/platform/os.version\",\n              \"type\": \"string\"\n            },\n            \"os.features\": {\n              \"id\": \"https://opencontainers.org/schema/image/platform/os.features\",\n              \"type\": \"array\",\n              \"items\": {\n                \"type\": \"string\"\n              }\n            },\n            \"variant\": {\n              \"type\": \"string\"\n            },\n            \"features\": {\n              \"type\": \"array\",\n              \"items\": {\n                \"type\": \"string\"\n              },\n              \"additionalProperties\": false\n            }\n          }\n        }\n      }\n    },\n    \"annotations\": {\n      \"id\": \"https://opencontainers.org/schema/image/annotations\",\n      \"oneOf\": [\n        {\n          \"$ref\": \"defs.json#/definitions/mapStringString\"\n        },\n        {\n          \"type\": \"null\"\n        }\n      ]\n    }\n  }\n}\nPK\x07\x08V5\x0d\xc8\x1c\x0c\x00\x00\x1c\x0c\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00h=\x9dH\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00	\x00\x00\x00defs.json{\n  \"description\": \"Definitions used throughout the OpenContainer Specification\",\n  \"definitions\": {\n    \"int8\": {\n      \"type\": \"integer\",\n      \"minimum\": -128,\n      \"maximum\": 127\n    },\n    \"int16\": {\n      \"type\": \"integer\",\n      \"minimum\": -32768,\n      \"maximum\": 32767\n    },\n    \"int32\": {\n      \"type\": \"integer\",\n      \"minimum\": -2147483648,\n      \"maximum\": 2147483647\n    },\n    \"int64\": {\n      \"type\": \"integer\",\n      \"minimum\": -9223372036854776000,\n      \"maximum\": 9223372036854776000\n    },\n    \"uint8\": {\n      \"type\": \"integer\",\n      \"minimum\": 0,\n      \"maximum\": 255\n    },\n    \"uint16\": {\n      \"type\": \"integer\",\n      \"minimum\": 0,\n      \"maximum\": 65535\n    },\n    \"uint32\": {\n      \"type\": \"integer\",\n      \"minimum\": 0,\n      \"maximum\": 4294967295\n    },\n    \"uint64\": {\n      \"type\": \"integer\",\n      \"minimum\": 0,\n      \"maximum\": 18446744073709552000\n    },\n    \"uint16Pointer\": {\n      \"oneOf\": [\n        {\n          \"$ref\": \"#/definitions/uint16\"\n        },\n        {\n          \"type\": \"null\"\n        }\n      ]\n    },\n    \"uint64Pointer\": {\n      \"oneOf\": [\n        {\n          \"$ref\": \"#/definitions/uint64\"\n        },\n        {\n          \"type\": \"null\"\n        }\n      ]\n    },\n    \"stringPointer\": {\n      \"oneOf\": [\n        {\n          \"type\": \"string\"\n        },\n        {\n          \"type\": \"null\"\n        }\n      ]\n    },\n    \"mapStringString\": {\n      \"type\": \"object\",\n      \"patternProperties\": {\n        \".{1,}\": {\n          \"type\": \"string\"\n        }\n      }\n    },\n    \"UID\": {\n      \"$ref\": \"#/definitions/uint32\"\n    },\n    \"GID\": {\n      \"$ref\": \"#/definitions/uint32\"\n    },\n    \"ArrayOfGIDs\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"#/definitions/GID\"\n      }\n    },\n    \"ArrayOfStrings\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"type\": \"string\"\n      }\n    },\n    \"FilePath\": {\n      \"type\": \"string\"\n    },\n    \"Env\": {\n      \"$ref\": \"#/definitions/ArrayOfStrings\"\n    },\n    \"Hook\": {\n      \"properties\": {\n        \"path\": {\n          \"$ref\": \"#/definitions/FilePath\"\n        },\n        \"args\": {\n          \"$ref\": \"#/definitions/ArrayOfStrings\"\n        },\n        \"env\": {\n          \"$ref\": \"#/definitions/Env\"\n        }\n      }\n    },\n    \"ArrayOfHooks\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"#/definitions/Hook\"\n      }\n    },\n    \"IDMapping\": {\n      \"properties\": {\n        \"hostID\": {\n          \"$ref\": \"#/definitions/uint32\"\n        },\n        \"containerID\": {\n          \"$ref\": \"#/definitions/uint32\"\n        },\n        \"size\": {\n          \"$ref\": \"#/definitions/uint32\"\n        }\n      }\n    },\n    \"Mount\": {\n      \"properties\": {\n        \"source\": {\n          \"$ref\": \"#/definitions/FilePath\"\n        },\n        \"destination\": {\n          \"$ref\": \"#/definitions/FilePath\"\n        },\n        \"options\": {\n          \"$ref\": \"#/definitions/ArrayOfStrings\"\n        },\n        \"type\": {\n          \"type\": \"string\"\n        }\n      },\n      \"required\": [\n        \"destination\",\n        \"source\",\n        \"type\"\n      ]\n    }\n  }\n}\nPK\x07\x08e`\x88L\xe4\x0b\x00\x00\xe4\x0b\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xb9C\xa2H\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1a\x00\x00\x00image-manifest-schema.json{\n  \"description\": \"OpenContainer Image Manifest Specification\",\n  \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n  \"id\": \"https://opencontainers.org/schema/image/manifest\",\n  \"type\": \"object\",\n  \"properties\": {\n    \"schemaVersion\": {\n      \"description\": \"This field specifies the image manifest schema version as an integer\",\n      \"id\": \"https://opencontainers.org/schema/image/manifest/schemaVersion\",\n      \"type\": \"integer\"\n    },\n    \"mediaType\": {\n      \"id\": \"https://opencontainers.org/schema/image/manifest/mediaType\",\n      \"$ref\": \"defs-image.json#/definitions/mediaType\"\n    },\n    \"config\": {\n      \"$ref\": \"defs-image.json#/definitions/descriptor\"\n    },\n    \"layers\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"defs-image.json#/definitions/descriptor\"\n      }\n    },\n    \"annotations\": {\n      \"id\": \"https://opencontainers.org/schema/image/manifest-list/annotations\",\n      \"$ref\": \"defs-image.json#/definitions/annotations\"\n    }\n  },\n  \"required\": [\n    \"schemaVersion\",\n    \"mediaType\",\n    \"config\",\n    \"layers\"\n  ]\n}\nPK\x07\x08\xfe\xfb*_(\x04\x00\x00(\x04\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xb9C\xa2H\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x19\x00\x00\x00manifest-list-schema.json{\n  \"description\": \"OpenContainer Image Manifest List Specification\",\n  \"$schema\": \"http://json-schema.org/draft-04/schema#\",\n  \"id\": \"https://opencontainers.org/schema/image/manifest-list\",\n  \"type\": \"object\",\n  \"properties\": {\n    \"schemaVersion\": {\n      \"description\": \"This field specifies the image manifest-list schema version as an integer\",\n      \"id\": \"https://opencontainers.org/schema/image/manifest-list/schemaVersion\",\n      \"type\": \"integer\"\n    },\n    \"mediaType\": {\n      \"id\": \"https://opencontainers.org/schema/image/manifest-list/mediaType\",\n      \"$ref\": \"defs-image.json#/definitions/mediaType\"\n    },\n    \"manifests\": {\n      \"type\": \"array\",\n      \"items\": {\n        \"$ref\": \"defs-image.json#/definitions/manifestDescriptor\"\n      }\n    },\n    \"annotations\": {\n      \"id\": \"https://opencontainers.org/schema/image/manifest-list/annotations\",\n      \"$ref\": \"defs-image.json#/definitions/annotations\"\n    }\n  },\n  \"required\": [\n    \"schemaVersion\",\n    \"mediaType\",\n    \"manifests\"\n  ]\n}\nPK\x07\x08\xb8\xe9\xcd\xae\xf2\x03\x00\x00\xf2\x03\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xc1n\xa6H\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0e\x00\x00\x00media-types.gopackage schema\n\nconst (\n	MediaTypeManifestList               = `application/vnd.oci.image.manifest.list.v1+json`\n	MediaTypeManifest                   = `application/vnd.oci.image.manifest.v1+json`\n	MediaTypeImageSerialization         = `application/vnd.oci.image.serialization.rootfs.tar.gzip`\n	MediaTypeImageSerializationConfig   = `application/vnd.oci.image.serialization.config.v1+json`\n	MediaTypeImageSerializationCombined = `application/vnd.oci.image.serialization.combined.v1+json`\n)\nPK\x07\x08/\xb7?\x0c\xea\x01\x00\x00\xea\x01\x00\x00PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xb9C\xa2HV5\x0d\xc8\x1c\x0c\x00\x00\x1c\x0c\x00\x00\x0f\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x00\x00\x00\x00defs-image.jsonPK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00h=\x9dHe`\x88L\xe4\x0b\x00\x00\xe4\x0b\x00\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81Y\x0c\x00\x00defs.jsonPK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xb9C\xa2H\xfe\xfb*_(\x04\x00\x00(\x04\x00\x00\x1a\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81t\x18\x00\x00image-manifest-schema.jsonPK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xb9C\xa2H\xb8\xe9\xcd\xae\xf2\x03\x00\x00\xf2\x03\x00\x00\x19\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xe4\x1c\x00\x00manifest-list-schema.jsonPK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xc1n\xa6H/\xb7?\x0c\xea\x01\x00\x00\xea\x01\x00\x00\x0e\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x1d!\x00\x00media-types.goPK\x05\x06\x00\x00\x00\x00\x05\x00\x05\x00?\x01\x00\x00C#\x00\x00\x00\x00"
	fs.Register(data)
}
