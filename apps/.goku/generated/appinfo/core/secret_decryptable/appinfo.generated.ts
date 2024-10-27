/*eslint-disable */

import * as fieldkind from '@ongoku/app-lib/src/common/fieldkind'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { App, AppReq, ServiceReq, EntityInfoReq, TypeInfoReq, EnumReq, MethodReq } from '@ongoku/app-lib/src/common/app_v3'
import { Name, ITypeAppNamespace, IEnumAppNamespace, ITypeServiceNamespace, IEnumServiceNamespace, ITypeEntityNamespace, IEnumEntityNamespace } from '@ongoku/app-lib/src/common/namespacev2'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BoolCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

const entityInfoReq: EntityInfoReq = {
    namespace: {
        service: 'core',
        entity: 'secret_decryptable',
    },
    associations: [
        {
            relationship: 'child_of',
            type: 'one',
            entityNamespace: {
                service: 'core',
                entity: 'secret_key',
            },
            name: new Name('secret_key'),
            otherAssociationName: new Name('secrets'),
        },
    ],
    actions: [
        {
            name: new Name('decrypt'),
            methodNamespace: {
                service: 'core',
                entity: 'secret_decryptable',
                types: undefined,
                enum: undefined,
                method: 'action_decrypt',
            },
        },
    ],
}

const typeReqs: TypeReq[] = [
    {
        namespace: {
            service: 'core',
            entity: 'secret_decryptable',
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
                    name: new Name('core__secret_decryptable'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_decryptable',
                        types: ['secret_decryptable'],
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
                entity: 'secret_decryptable',
                types: ['secret_decryptable'],
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
            entity: 'secret_decryptable',
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
            entity: 'secret_decryptable',
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
                    name: new Name('core__secret_decryptable__secret_decryptable_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_decryptable',
                        types: ['secret_decryptable_filter'],
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
                entity: 'secret_decryptable',
                types: ['secret_decryptable_filter'],
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
            entity: 'secret_decryptable',
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
                    name: new Name('core__secret_decryptable'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_decryptable',
                        types: ['secret_decryptable'],
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
                entity: 'secret_decryptable',
                types: ['secret_decryptable'],
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
            entity: 'secret_decryptable',
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
            entity: 'secret_decryptable',
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
            entity: 'secret_decryptable',
            types: ['secret_decryptable'],
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
                name: new Name('raw_value'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('string'),
                    kind: fieldkind.StringKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('encrypted_value'),
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
                name: new Name('salt'),
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
            {
                name: new Name('secret_key_id'),
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
                rawValue: undefined,
                encryptedValue: undefined,
                salt: undefined,
                createdAt: undefined,
                updatedAt: undefined,
                deletedAt: undefined,
                secretKeyID: undefined,
            }
        },
    },
    {
        namespace: {
            service: 'core',
            entity: 'secret_decryptable',
            types: ['secret_decryptable_action_condition'],
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
            entity: 'secret_decryptable',
            types: ['secret_decryptable_field_condition'],
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
            entity: 'secret_decryptable',
            types: ['secret_decryptable_filter'],
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
                name: new Name('raw_value'),
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
                name: new Name('encrypted_value'),
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
                name: new Name('salt'),
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
                name: new Name('secret_key_id'),
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
                name: new Name('and'),
                isRepeated: true,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('core__secret_decryptable__secret_decryptable_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_decryptable',
                        types: ['secret_decryptable_filter'],
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
                    name: new Name('core__secret_decryptable__secret_decryptable_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_decryptable',
                        types: ['secret_decryptable_filter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const andTypeInfo = appInfo.getTypeInfo({
                service: 'core',
                entity: 'secret_decryptable',
                types: ['secret_decryptable_filter'],
                enum: undefined,
                method: undefined,
            })
            const orTypeInfo = appInfo.getTypeInfo({
                service: 'core',
                entity: 'secret_decryptable',
                types: ['secret_decryptable_filter'],
                enum: undefined,
                method: undefined,
            })

            // Return the empty object
            return {
                id: undefined,
                rawValue: undefined,
                encryptedValue: undefined,
                salt: undefined,
                createdAt: undefined,
                updatedAt: undefined,
                deletedAt: undefined,
                secretKeyID: undefined,
                and: undefined,
                or: undefined,
            }
        },
    },
    {
        namespace: {
            service: 'core',
            entity: 'secret_decryptable',
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
                    name: new Name('core__secret_decryptable'),
                    kind: fieldkind.ForeignEntityKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_decryptable',
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
            entity: 'secret_decryptable',
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
                    name: new Name('core__secret_decryptable'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_decryptable',
                        types: ['secret_decryptable'],
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
                entity: 'secret_decryptable',
                types: ['secret_decryptable'],
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
            entity: 'secret_decryptable',
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
                    name: new Name('core__secret_decryptable'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_decryptable',
                        types: ['secret_decryptable'],
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
                    name: new Name('core__secret_decryptable__secret_decryptable_field'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_decryptable',
                        types: undefined,
                        enum: 'secret_decryptable_field',
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
                    name: new Name('core__secret_decryptable__secret_decryptable_field'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_decryptable',
                        types: undefined,
                        enum: 'secret_decryptable_field',
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const objectTypeInfo = appInfo.getTypeInfo({
                service: 'core',
                entity: 'secret_decryptable',
                types: ['secret_decryptable'],
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
            entity: 'secret_decryptable',
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
                    name: new Name('core__secret_decryptable'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'core',
                        entity: 'secret_decryptable',
                        types: ['secret_decryptable'],
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
                entity: 'secret_decryptable',
                types: ['secret_decryptable'],
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
            entity: 'secret_decryptable',
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
            {
                id: 7,
                value: 'ActionDecrypt',
                description: undefined,
                displayValue: 'Action Decrypt',
            },
        ],
    },
    {
        namespace: {
            service: 'core',
            entity: 'secret_decryptable',
            types: undefined,
            enum: 'secret_decryptable_action',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'Decrypt',
                description: 'Decrypt the secret',
                displayValue: 'Decrypt',
            },
        ],
    },
    {
        namespace: {
            service: 'core',
            entity: 'secret_decryptable',
            types: undefined,
            enum: 'secret_decryptable_field',
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
                value: 'RawValue',
                description: undefined,
                displayValue: 'Raw Value',
            },
            {
                id: 3,
                value: 'EncryptedValue',
                description: undefined,
                displayValue: 'Encrypted Value',
            },
            {
                id: 4,
                value: 'Salt',
                description: undefined,
                displayValue: 'Salt',
            },
            {
                id: 5,
                value: 'CreatedAt',
                description: undefined,
                displayValue: 'Created At',
            },
            {
                id: 6,
                value: 'UpdatedAt',
                description: undefined,
                displayValue: 'Updated At',
            },
            {
                id: 7,
                value: 'DeletedAt',
                description: undefined,
                displayValue: 'Deleted At',
            },
            {
                id: 8,
                value: 'SecretKeyID',
                description: undefined,
                displayValue: 'Secret Key ID',
            },
        ],
    },
]
