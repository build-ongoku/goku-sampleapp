/*eslint-disable */

import * as fieldkind from '@ongoku/app-lib/src/common/fieldkind'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { AppReq, ServiceReq, EntityInfoReq, TypeInfoReq, EnumReq, MethodReq } from '@ongoku/app-lib/src/common/app_v3'
import { Name, ITypeAppNamespace, IEnumAppNamespace, ITypeServiceNamespace, IEnumServiceNamespace, ITypeEntityNamespace, IEnumEntityNamespace } from '@ongoku/app-lib/src/common/namespacev2'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BooleanCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

const entityInfoReq: EntityInfoReq = {
    namespace: {
        service: 'Core',
        entity: 'File',
    },
    associations: [],
    actions: [],
}

const typeReqs: TypeReq[] = [
    {
        namespace: {
            service: 'Core',
            entity: 'File',
            types: ['File'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('id'),
                isRepeated: false,
                isOptional: false,
                isMetaField: true,
                excludeFromForm: true,
                dtype: {
                    name: new Name('ID'),
                    kind: fieldkind.IDKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('fileName'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('String'),
                    kind: fieldkind.StringKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('storageClient'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_File_StorageClient'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'Core',
                        entity: 'File',
                        types: undefined,
                        enum: 'StorageClient',
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('storageClientIdentifier'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('String'),
                    kind: fieldkind.StringKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('sizeBytes'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: true,
                dtype: {
                    name: new Name('Number'),
                    kind: fieldkind.NumberKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('mimeType'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: true,
                dtype: {
                    name: new Name('String'),
                    kind: fieldkind.StringKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('fileHash'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: true,
                dtype: {
                    name: new Name('String'),
                    kind: fieldkind.StringKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('createdAt'),
                isRepeated: false,
                isOptional: false,
                isMetaField: true,
                excludeFromForm: true,
                dtype: {
                    name: new Name('Timestamp'),
                    kind: fieldkind.TimestampKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('updatedAt'),
                isRepeated: false,
                isOptional: false,
                isMetaField: true,
                excludeFromForm: true,
                dtype: {
                    name: new Name('Timestamp'),
                    kind: fieldkind.TimestampKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('deletedAt'),
                isRepeated: false,
                isOptional: true,
                isMetaField: true,
                excludeFromForm: true,
                dtype: {
                    name: new Name('Timestamp'),
                    kind: fieldkind.TimestampKind,
                    namespace: undefined,
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'File',
            types: ['FileFieldCondition'],
            enum: undefined,
            method: undefined,
        },
        fields: [],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'File',
            types: ['FileFilter'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('id'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('IDCondition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('fileName'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('StringCondition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('storageClient'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('StorageClientCondition'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Core',
                        entity: 'File',
                        types: ['StorageClientCondition'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('storageClientIdentifier'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('StringCondition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('sizeBytes'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('NumberCondition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('mimeType'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('StringCondition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('fileHash'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('StringCondition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('createdAt'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('TimestampCondition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('updatedAt'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('TimestampCondition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('deletedAt'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('TimestampCondition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('and'),
                isRepeated: true,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_File_FileFilter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Core',
                        entity: 'File',
                        types: ['FileFilter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('or'),
                isRepeated: true,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_File_FileFilter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Core',
                        entity: 'File',
                        types: ['FileFilter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'File',
            types: ['StandardEntityRequest'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('id'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_File'),
                    kind: fieldkind.ForeignEntityKind,
                    namespace: {
                        service: 'Core',
                        entity: 'File',
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'File',
            types: ['StandardEntityResponse'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('object'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_File'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Core',
                        entity: 'File',
                        types: ['File'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'File',
            types: ['StorageClientCondition'],
            enum: undefined,
            method: undefined,
        },
        fields: [],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'File',
            types: ['AddRequest'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('object'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_File'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Core',
                        entity: 'File',
                        types: ['File'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'File',
            types: ['UpdateRequest'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('object'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_File'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Core',
                        entity: 'File',
                        types: ['File'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('fields'),
                isRepeated: true,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_File_FileField'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'Core',
                        entity: 'File',
                        types: undefined,
                        enum: 'FileField',
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('excludeFields'),
                isRepeated: true,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_File_FileField'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'Core',
                        entity: 'File',
                        types: undefined,
                        enum: 'FileField',
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'File',
            types: ['UpdateResponse'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('object'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_File'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Core',
                        entity: 'File',
                        types: ['File'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'File',
            types: ['GetRequest'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('id'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('ID'),
                    kind: fieldkind.IDKind,
                    namespace: undefined,
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'File',
            types: ['ListRequest'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('filter'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_File_FileFilter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Core',
                        entity: 'File',
                        types: ['FileFilter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'File',
            types: ['ListResponse'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('items'),
                isRepeated: true,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_File'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Core',
                        entity: 'File',
                        types: ['File'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('count'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Number'),
                    kind: fieldkind.NumberKind,
                    namespace: undefined,
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'File',
            types: ['QueryByTextRequest'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('queryText'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('String'),
                    kind: fieldkind.StringKind,
                    namespace: undefined,
                },
            },
        ],
    },
]

const enumReqs: EnumReq[] = [
    {
        namespace: {
            service: 'Core',
            entity: 'File',
            types: undefined,
            enum: 'FileField',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'ID',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 2,
                value: 'FileName',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 3,
                value: 'StorageClient',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 4,
                value: 'StorageClientIdentifier',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 5,
                value: 'SizeBytes',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 6,
                value: 'MimeType',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 7,
                value: 'FileHash',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 8,
                value: 'CreatedAt',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 9,
                value: 'UpdatedAt',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 10,
                value: 'DeletedAt',
                description: undefined,
                displayValue: undefined,
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'File',
            types: undefined,
            enum: 'StorageClient',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'LocalFile',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 2,
                value: 'Database',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 3,
                value: 'CloudAmazonS3',
                description: undefined,
                displayValue: undefined,
            },
        ],
    },
]
