/*eslint-disable */

import * as fieldkind from '@ongoku/app-lib/src/common/fieldkind'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { App, AppReq, ServiceReq, EntityInfoReq, TypeInfoReq, EnumReq, MethodReq } from '@ongoku/app-lib/src/common/app_v3'
import { Name, ITypeAppNamespace, IEnumAppNamespace, ITypeServiceNamespace, IEnumServiceNamespace, ITypeEntityNamespace, IEnumEntityNamespace } from '@ongoku/app-lib/src/common/namespacev2'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BoolCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

const entityInfoReq: EntityInfoReq = {
    namespace: {
        service: 'auth',
        entity: 'secret',
    },
    associations: [],
    actions: [],
}

const typeReqs: TypeReq[] = [
    {
        namespace: {
            service: 'auth',
            entity: 'secret',
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
                    name: new Name('auth__secret'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'auth',
                        entity: 'secret',
                        types: ['secret'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const objectTypeInfo = appInfo.getTypeInfo({
                service: 'auth',
                entity: 'secret',
                types: ['secret'],
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
            service: 'auth',
            entity: 'secret',
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
            service: 'auth',
            entity: 'secret',
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
                    name: new Name('auth__secret__secret_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'auth',
                        entity: 'secret',
                        types: ['secret_filter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const filterTypeInfo = appInfo.getTypeInfo({
                service: 'auth',
                entity: 'secret',
                types: ['secret_filter'],
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
            service: 'auth',
            entity: 'secret',
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
                    name: new Name('auth__secret'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'auth',
                        entity: 'secret',
                        types: ['secret'],
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
                service: 'auth',
                entity: 'secret',
                types: ['secret'],
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
            service: 'auth',
            entity: 'secret',
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
            service: 'auth',
            entity: 'secret',
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
            service: 'auth',
            entity: 'secret',
            types: ['secret'],
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
                name: new Name('user_id'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('user__user'),
                    kind: fieldkind.ForeignEntityKind,
                    namespace: {
                        service: 'user',
                        entity: 'user',
                    },
                },
            },
            {
                name: new Name('type'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('auth__secret__type'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'auth',
                        entity: 'secret',
                        types: undefined,
                        enum: 'type',
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('secret'),
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
                name: new Name('expires_at'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('timestamp'),
                    kind: fieldkind.TimestampKind,
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
                userID: undefined,
                type: undefined,
                secret: undefined,
                expiresAt: undefined,
                createdAt: undefined,
                updatedAt: undefined,
                deletedAt: undefined,
            }
        },
    },
    {
        namespace: {
            service: 'auth',
            entity: 'secret',
            types: ['secret_field_condition'],
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
            service: 'auth',
            entity: 'secret',
            types: ['secret_filter'],
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
                name: new Name('user_id'),
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
                        service: 'auth',
                        entity: 'secret',
                        types: ['type_condition'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('secret'),
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
                name: new Name('expires_at'),
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
                    name: new Name('auth__secret__secret_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'auth',
                        entity: 'secret',
                        types: ['secret_filter'],
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
                    name: new Name('auth__secret__secret_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'auth',
                        entity: 'secret',
                        types: ['secret_filter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const typeTypeInfo = appInfo.getTypeInfo({
                service: 'auth',
                entity: 'secret',
                types: ['type_condition'],
                enum: undefined,
                method: undefined,
            })
            const andTypeInfo = appInfo.getTypeInfo({
                service: 'auth',
                entity: 'secret',
                types: ['secret_filter'],
                enum: undefined,
                method: undefined,
            })
            const orTypeInfo = appInfo.getTypeInfo({
                service: 'auth',
                entity: 'secret',
                types: ['secret_filter'],
                enum: undefined,
                method: undefined,
            })

            // Return the empty object
            return {
                id: undefined,
                userID: undefined,
                type: typeTypeInfo.getEmptyObject(appInfo),
                secret: undefined,
                expiresAt: undefined,
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
            service: 'auth',
            entity: 'secret',
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
                    name: new Name('auth__secret'),
                    kind: fieldkind.ForeignEntityKind,
                    namespace: {
                        service: 'auth',
                        entity: 'secret',
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
            service: 'auth',
            entity: 'secret',
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
                    name: new Name('auth__secret'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'auth',
                        entity: 'secret',
                        types: ['secret'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const objectTypeInfo = appInfo.getTypeInfo({
                service: 'auth',
                entity: 'secret',
                types: ['secret'],
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
            service: 'auth',
            entity: 'secret',
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
            service: 'auth',
            entity: 'secret',
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
                    name: new Name('auth__secret'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'auth',
                        entity: 'secret',
                        types: ['secret'],
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
                    name: new Name('auth__secret__secret_field'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'auth',
                        entity: 'secret',
                        types: undefined,
                        enum: 'secret_field',
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
                    name: new Name('auth__secret__secret_field'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'auth',
                        entity: 'secret',
                        types: undefined,
                        enum: 'secret_field',
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const objectTypeInfo = appInfo.getTypeInfo({
                service: 'auth',
                entity: 'secret',
                types: ['secret'],
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
            service: 'auth',
            entity: 'secret',
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
                    name: new Name('auth__secret'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'auth',
                        entity: 'secret',
                        types: ['secret'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const objectTypeInfo = appInfo.getTypeInfo({
                service: 'auth',
                entity: 'secret',
                types: ['secret'],
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
            service: 'auth',
            entity: 'secret',
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
        ],
    },
    {
        namespace: {
            service: 'auth',
            entity: 'secret',
            types: undefined,
            enum: 'secret_field',
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
                value: 'UserID',
                description: undefined,
                displayValue: 'User ID',
            },
            {
                id: 3,
                value: 'Type',
                description: undefined,
                displayValue: 'Type',
            },
            {
                id: 4,
                value: 'Secret',
                description: undefined,
                displayValue: 'Secret',
            },
            {
                id: 5,
                value: 'ExpiresAt',
                description: undefined,
                displayValue: 'Expires At',
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
            service: 'auth',
            entity: 'secret',
            types: undefined,
            enum: 'type',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'Password',
                description: undefined,
                displayValue: 'Password',
            },
            {
                id: 2,
                value: 'GithubToken',
                description: undefined,
                displayValue: 'Github Token',
            },
        ],
    },
]
