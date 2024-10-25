/*eslint-disable */

import * as fieldkind from '@ongoku/app-lib/src/common/fieldkind'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { App, AppReq, ServiceReq, EntityInfoReq, TypeInfoReq, EnumReq, MethodReq } from '@ongoku/app-lib/src/common/app_v3'
import { Name, ITypeAppNamespace, IEnumAppNamespace, ITypeServiceNamespace, IEnumServiceNamespace, ITypeEntityNamespace, IEnumEntityNamespace } from '@ongoku/app-lib/src/common/namespacev2'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BoolCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

const entityInfoReq: EntityInfoReq = {
    namespace: {
        service: 'core',
        entity: 'file',
    },
    associations: [],
    actions: [],
}

const typeReqs: TypeReq[] = [
    {
        namespace: {
            service: 'core',
            entity: 'file',
            types: ['add_request'],
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
                    name: new Name('core__file'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'file',
                        types: ['file'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const objectTypeInfo = appInfo.getTypeInfo({
                service: 'core',
                entity: 'file',
                types: ['file'],
                enum: undefined,
                method: undefined,
            })

            // Return the empty object
            return {
                object: objectTypeInfo.getEmptyObject(appInfo),
            }
        },
    },
    {
        namespace: {
            service: 'core',
            entity: 'file',
            types: ['file'],
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
                    name: new Name('id'),
                    kind: fieldkind.IDKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('file_name'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('string'),
                    kind: fieldkind.StringKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('storage_client'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('core__file__storage_client'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'core',
                        entity: 'file',
                        types: undefined,
                        enum: 'storage_client',
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('storage_client_identifier'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('string'),
                    kind: fieldkind.StringKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('size_bytes'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: true,
                dtype: {
                    name: new Name('number'),
                    kind: fieldkind.NumberKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('mime_type'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: true,
                dtype: {
                    name: new Name('string'),
                    kind: fieldkind.StringKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('file_hash'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: true,
                dtype: {
                    name: new Name('string'),
                    kind: fieldkind.StringKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('created_at'),
                isRepeated: false,
                isOptional: false,
                isMetaField: true,
                excludeFromForm: true,
                dtype: {
                    name: new Name('timestamp'),
                    kind: fieldkind.TimestampKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('updated_at'),
                isRepeated: false,
                isOptional: false,
                isMetaField: true,
                excludeFromForm: true,
                dtype: {
                    name: new Name('timestamp'),
                    kind: fieldkind.TimestampKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('deleted_at'),
                isRepeated: false,
                isOptional: true,
                isMetaField: true,
                excludeFromForm: true,
                dtype: {
                    name: new Name('timestamp'),
                    kind: fieldkind.TimestampKind,
                    namespace: undefined,
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances

            // Return the empty object
            return {
                id: undefined,
                fileName: undefined,
                storageClient: undefined,
                storageClientIdentifier: undefined,
                sizeBytes: undefined,
                mimeType: undefined,
                fileHash: undefined,
                createdAt: undefined,
                updatedAt: undefined,
                deletedAt: undefined,
            }
        },
    },
    {
        namespace: {
            service: 'core',
            entity: 'file',
            types: ['file_field_condition'],
            enum: undefined,
            method: undefined,
        },
        fields: [],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances

            // Return the empty object
            return {}
        },
    },
    {
        namespace: {
            service: 'core',
            entity: 'file',
            types: ['file_filter'],
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
                    name: new Name('id_condition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('file_name'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('string_condition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('storage_client'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('storage_client_condition'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'file',
                        types: ['storage_client_condition'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('storage_client_identifier'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('string_condition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('size_bytes'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('number_condition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('mime_type'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('string_condition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('file_hash'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('string_condition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('created_at'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('timestamp_condition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('updated_at'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('timestamp_condition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('deleted_at'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('timestamp_condition'),
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
                    name: new Name('core__file__file_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'file',
                        types: ['file_filter'],
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
                    name: new Name('core__file__file_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'file',
                        types: ['file_filter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const storageClientTypeInfo = appInfo.getTypeInfo({
                service: 'core',
                entity: 'file',
                types: ['storage_client_condition'],
                enum: undefined,
                method: undefined,
            })
            const andTypeInfo = appInfo.getTypeInfo({
                service: 'core',
                entity: 'file',
                types: ['file_filter'],
                enum: undefined,
                method: undefined,
            })
            const orTypeInfo = appInfo.getTypeInfo({
                service: 'core',
                entity: 'file',
                types: ['file_filter'],
                enum: undefined,
                method: undefined,
            })

            // Return the empty object
            return {
                id: undefined,
                fileName: undefined,
                storageClient: storageClientTypeInfo.getEmptyObject(appInfo),
                storageClientIdentifier: undefined,
                sizeBytes: undefined,
                mimeType: undefined,
                fileHash: undefined,
                createdAt: undefined,
                updatedAt: undefined,
                deletedAt: undefined,
                and: undefined,
                or: undefined,
            }
        },
    },
    {
        namespace: {
            service: 'core',
            entity: 'file',
            types: ['get_request'],
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
                    name: new Name('id'),
                    kind: fieldkind.IDKind,
                    namespace: undefined,
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances

            // Return the empty object
            return {
                id: undefined,
            }
        },
    },
    {
        namespace: {
            service: 'core',
            entity: 'file',
            types: ['list_request'],
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
                    name: new Name('core__file__file_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'file',
                        types: ['file_filter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const filterTypeInfo = appInfo.getTypeInfo({
                service: 'core',
                entity: 'file',
                types: ['file_filter'],
                enum: undefined,
                method: undefined,
            })

            // Return the empty object
            return {
                filter: filterTypeInfo.getEmptyObject(appInfo),
            }
        },
    },
    {
        namespace: {
            service: 'core',
            entity: 'file',
            types: ['list_response'],
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
                    name: new Name('core__file'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'file',
                        types: ['file'],
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
                    name: new Name('number'),
                    kind: fieldkind.NumberKind,
                    namespace: undefined,
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const itemsTypeInfo = appInfo.getTypeInfo({
                service: 'core',
                entity: 'file',
                types: ['file'],
                enum: undefined,
                method: undefined,
            })

            // Return the empty object
            return {
                items: undefined,
                count: undefined,
            }
        },
    },
    {
        namespace: {
            service: 'core',
            entity: 'file',
            types: ['method_name_condition'],
            enum: undefined,
            method: undefined,
        },
        fields: [],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances

            // Return the empty object
            return {}
        },
    },
    {
        namespace: {
            service: 'core',
            entity: 'file',
            types: ['query_by_text_request'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('query_text'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('string'),
                    kind: fieldkind.StringKind,
                    namespace: undefined,
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances

            // Return the empty object
            return {
                queryText: undefined,
            }
        },
    },
    {
        namespace: {
            service: 'core',
            entity: 'file',
            types: ['standard_entity_request'],
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
                    name: new Name('core__file'),
                    kind: fieldkind.ForeignEntityKind,
                    namespace: {
                        service: 'core',
                        entity: 'file',
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances

            // Return the empty object
            return {
                id: undefined,
            }
        },
    },
    {
        namespace: {
            service: 'core',
            entity: 'file',
            types: ['standard_entity_response'],
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
                    name: new Name('core__file'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'file',
                        types: ['file'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const objectTypeInfo = appInfo.getTypeInfo({
                service: 'core',
                entity: 'file',
                types: ['file'],
                enum: undefined,
                method: undefined,
            })

            // Return the empty object
            return {
                object: objectTypeInfo.getEmptyObject(appInfo),
            }
        },
    },
    {
        namespace: {
            service: 'core',
            entity: 'file',
            types: ['storage_client_condition'],
            enum: undefined,
            method: undefined,
        },
        fields: [],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances

            // Return the empty object
            return {}
        },
    },
    {
        namespace: {
            service: 'core',
            entity: 'file',
            types: ['update_request'],
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
                    name: new Name('core__file'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'file',
                        types: ['file'],
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
                    name: new Name('core__file__file_field'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'core',
                        entity: 'file',
                        types: undefined,
                        enum: 'file_field',
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('exclude_fields'),
                isRepeated: true,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('core__file__file_field'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'core',
                        entity: 'file',
                        types: undefined,
                        enum: 'file_field',
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const objectTypeInfo = appInfo.getTypeInfo({
                service: 'core',
                entity: 'file',
                types: ['file'],
                enum: undefined,
                method: undefined,
            })

            // Return the empty object
            return {
                object: objectTypeInfo.getEmptyObject(appInfo),
                fields: undefined,
                excludeFields: undefined,
            }
        },
    },
    {
        namespace: {
            service: 'core',
            entity: 'file',
            types: ['update_response'],
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
                    name: new Name('core__file'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'file',
                        types: ['file'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const objectTypeInfo = appInfo.getTypeInfo({
                service: 'core',
                entity: 'file',
                types: ['file'],
                enum: undefined,
                method: undefined,
            })

            // Return the empty object
            return {
                object: objectTypeInfo.getEmptyObject(appInfo),
            }
        },
    },
]

const enumReqs: EnumReq[] = [
    {
        namespace: {
            service: 'core',
            entity: 'file',
            types: undefined,
            enum: 'file_field',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'ID',
                description: undefined,
                displayValue: 'ID',
            },
            {
                id: 2,
                value: 'FileName',
                description: undefined,
                displayValue: 'File Name',
            },
            {
                id: 3,
                value: 'StorageClient',
                description: undefined,
                displayValue: 'Storage Client',
            },
            {
                id: 4,
                value: 'StorageClientIdentifier',
                description: undefined,
                displayValue: 'Storage Client Identifier',
            },
            {
                id: 5,
                value: 'SizeBytes',
                description: undefined,
                displayValue: 'Size Bytes',
            },
            {
                id: 6,
                value: 'MimeType',
                description: undefined,
                displayValue: 'Mime Type',
            },
            {
                id: 7,
                value: 'FileHash',
                description: undefined,
                displayValue: 'File Hash',
            },
            {
                id: 8,
                value: 'CreatedAt',
                description: undefined,
                displayValue: 'Created At',
            },
            {
                id: 9,
                value: 'UpdatedAt',
                description: undefined,
                displayValue: 'Updated At',
            },
            {
                id: 10,
                value: 'DeletedAt',
                description: undefined,
                displayValue: 'Deleted At',
            },
        ],
    },
    {
        namespace: {
            service: 'core',
            entity: 'file',
            types: undefined,
            enum: 'method_name',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'Add',
                description: undefined,
                displayValue: 'Add',
            },
            {
                id: 2,
                value: 'Update',
                description: undefined,
                displayValue: 'Update',
            },
            {
                id: 3,
                value: 'Get',
                description: undefined,
                displayValue: 'Get',
            },
            {
                id: 4,
                value: 'List',
                description: undefined,
                displayValue: 'List',
            },
            {
                id: 5,
                value: 'QueryByText',
                description: undefined,
                displayValue: 'Query By Text',
            },
            {
                id: 6,
                value: 'HookCreatePre',
                description: undefined,
                displayValue: 'Hook Create Pre',
            },
        ],
    },
    {
        namespace: {
            service: 'core',
            entity: 'file',
            types: undefined,
            enum: 'storage_client',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'LocalFile',
                description: undefined,
                displayValue: 'Local File',
            },
            {
                id: 2,
                value: 'Database',
                description: undefined,
                displayValue: 'Database',
            },
            {
                id: 3,
                value: 'CloudAmazonS3',
                description: undefined,
                displayValue: 'Cloud Amazon S 3',
            },
        ],
    },
]
