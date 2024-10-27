/*eslint-disable */

import * as fieldkind from '@ongoku/app-lib/src/common/fieldkind'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { App, AppReq, ServiceReq, EntityInfoReq, TypeInfoReq, EnumReq, MethodReq } from '@ongoku/app-lib/src/common/app_v3'
import { Name, ITypeAppNamespace, IEnumAppNamespace, ITypeServiceNamespace, IEnumServiceNamespace, ITypeEntityNamespace, IEnumEntityNamespace } from '@ongoku/app-lib/src/common/namespacev2'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BoolCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

const entityInfoReq: EntityInfoReq = {
    namespace: {
        service: 'core',
        entity: 'secret_key',
    },
    associations: [
        {
            relationship: 'parent_of',
            type: 'many',
            entityNamespace: {
                service: 'core',
                entity: 'secret_decryptable',
            },
            name: new Name('secrets'),
            otherAssociationName: new Name('secret_key'),
        },
    ],
    actions: [],
}

const typeReqs: TypeReq[] = [
    {
        namespace: {
            service: 'core',
            entity: 'secret_key',
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
                    name: new Name('core__secret_key'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_key',
                        types: ['secret_key'],
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
                entity: 'secret_key',
                types: ['secret_key'],
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
            entity: 'secret_key',
            types: ['format_condition'],
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
            entity: 'secret_key',
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
            entity: 'secret_key',
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
                    name: new Name('core__secret_key__secret_key_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_key',
                        types: ['secret_key_filter'],
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
                entity: 'secret_key',
                types: ['secret_key_filter'],
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
            entity: 'secret_key',
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
                    name: new Name('core__secret_key'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_key',
                        types: ['secret_key'],
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
                entity: 'secret_key',
                types: ['secret_key'],
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
            entity: 'secret_key',
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
            entity: 'secret_key',
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
            entity: 'secret_key',
            types: ['secret_key'],
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
                name: new Name('type'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('core__secret_key__type'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_key',
                        types: undefined,
                        enum: 'type',
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('public_key'),
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
                name: new Name('private_key_format'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: true,
                dtype: {
                    name: new Name('core__secret_key__format'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_key',
                        types: undefined,
                        enum: 'format',
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('public_key_format'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: true,
                dtype: {
                    name: new Name('core__secret_key__format'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_key',
                        types: undefined,
                        enum: 'format',
                        method: undefined,
                    },
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
                type: undefined,
                publicKey: undefined,
                privateKeyFormat: undefined,
                publicKeyFormat: undefined,
                createdAt: undefined,
                updatedAt: undefined,
                deletedAt: undefined,
            }
        },
    },
    {
        namespace: {
            service: 'core',
            entity: 'secret_key',
            types: ['secret_key_field_condition'],
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
            entity: 'secret_key',
            types: ['secret_key_filter'],
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
                name: new Name('type'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('type_condition'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_key',
                        types: ['type_condition'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('public_key'),
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
                name: new Name('private_key_format'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('format_condition'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_key',
                        types: ['format_condition'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('public_key_format'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('format_condition'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_key',
                        types: ['format_condition'],
                        enum: undefined,
                        method: undefined,
                    },
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
                    name: new Name('core__secret_key__secret_key_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_key',
                        types: ['secret_key_filter'],
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
                    name: new Name('core__secret_key__secret_key_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_key',
                        types: ['secret_key_filter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const typeTypeInfo = appInfo.getTypeInfo({
                service: 'core',
                entity: 'secret_key',
                types: ['type_condition'],
                enum: undefined,
                method: undefined,
            })
            const privateKeyFormatTypeInfo = appInfo.getTypeInfo({
                service: 'core',
                entity: 'secret_key',
                types: ['format_condition'],
                enum: undefined,
                method: undefined,
            })
            const publicKeyFormatTypeInfo = appInfo.getTypeInfo({
                service: 'core',
                entity: 'secret_key',
                types: ['format_condition'],
                enum: undefined,
                method: undefined,
            })
            const andTypeInfo = appInfo.getTypeInfo({
                service: 'core',
                entity: 'secret_key',
                types: ['secret_key_filter'],
                enum: undefined,
                method: undefined,
            })
            const orTypeInfo = appInfo.getTypeInfo({
                service: 'core',
                entity: 'secret_key',
                types: ['secret_key_filter'],
                enum: undefined,
                method: undefined,
            })

            // Return the empty object
            return {
                id: undefined,
                type: typeTypeInfo.getEmptyObject(appInfo),
                publicKey: undefined,
                privateKeyFormat: privateKeyFormatTypeInfo.getEmptyObject(appInfo),
                publicKeyFormat: publicKeyFormatTypeInfo.getEmptyObject(appInfo),
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
            entity: 'secret_key',
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
                    name: new Name('core__secret_key'),
                    kind: fieldkind.ForeignEntityKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_key',
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
            entity: 'secret_key',
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
                    name: new Name('core__secret_key'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_key',
                        types: ['secret_key'],
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
                entity: 'secret_key',
                types: ['secret_key'],
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
            entity: 'secret_key',
            types: ['type_condition'],
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
            entity: 'secret_key',
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
                    name: new Name('core__secret_key'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_key',
                        types: ['secret_key'],
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
                    name: new Name('core__secret_key__secret_key_field'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_key',
                        types: undefined,
                        enum: 'secret_key_field',
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
                    name: new Name('core__secret_key__secret_key_field'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_key',
                        types: undefined,
                        enum: 'secret_key_field',
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const objectTypeInfo = appInfo.getTypeInfo({
                service: 'core',
                entity: 'secret_key',
                types: ['secret_key'],
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
            entity: 'secret_key',
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
                    name: new Name('core__secret_key'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_key',
                        types: ['secret_key'],
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
                entity: 'secret_key',
                types: ['secret_key'],
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
            entity: 'secret_key',
            types: undefined,
            enum: 'format',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'PEM',
                description: undefined,
                displayValue: 'PEM',
            },
            {
                id: 2,
                value: 'OpenSSH',
                description: undefined,
                displayValue: 'Open SSH',
            },
        ],
    },
    {
        namespace: {
            service: 'core',
            entity: 'secret_key',
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
                value: 'HookInit',
                description: undefined,
                displayValue: 'Hook Init',
            },
            {
                id: 7,
                value: 'HookCreatePre',
                description: undefined,
                displayValue: 'Hook Create Pre',
            },
        ],
    },
    {
        namespace: {
            service: 'core',
            entity: 'secret_key',
            types: undefined,
            enum: 'secret_key_field',
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
                value: 'Type',
                description: undefined,
                displayValue: 'Type',
            },
            {
                id: 3,
                value: 'PublicKey',
                description: undefined,
                displayValue: 'Public Key',
            },
            {
                id: 4,
                value: 'PrivateKeyFormat',
                description: undefined,
                displayValue: 'Private Key Format',
            },
            {
                id: 5,
                value: 'PublicKeyFormat',
                description: undefined,
                displayValue: 'Public Key Format',
            },
            {
                id: 6,
                value: 'CreatedAt',
                description: undefined,
                displayValue: 'Created At',
            },
            {
                id: 7,
                value: 'UpdatedAt',
                description: undefined,
                displayValue: 'Updated At',
            },
            {
                id: 8,
                value: 'DeletedAt',
                description: undefined,
                displayValue: 'Deleted At',
            },
        ],
    },
    {
        namespace: {
            service: 'core',
            entity: 'secret_key',
            types: undefined,
            enum: 'type',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'Ed25519',
                description: undefined,
                displayValue: 'Ed 25519',
            },
            {
                id: 2,
                value: 'RSA',
                description: undefined,
                displayValue: 'RSA',
            },
        ],
    },
]
