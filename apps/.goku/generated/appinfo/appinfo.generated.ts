/*eslint-disable */

import * as fieldkind from '@ongoku/app-lib/src/common/fieldkind'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { App, AppReq, ServiceReq, EntityInfoReq, TypeInfoReq, EnumReq, MethodReq } from '@ongoku/app-lib/src/common/app_v3'
import { Name, ITypeAppNamespace, IEnumAppNamespace, ITypeServiceNamespace, IEnumServiceNamespace, ITypeEntityNamespace, IEnumEntityNamespace } from '@ongoku/app-lib/src/common/namespacev2'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BoolCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

export const appReq: AppReq = {
    name: new Name('sampleapp'),
    services: [
        {
            namespace: {
                service: 'auth',
                entity: undefined,
            },
        },
        {
            namespace: {
                service: 'core',
                entity: undefined,
            },
        },
        {
            namespace: {
                service: 'user',
                entity: undefined,
            },
        },
    ],
    entityInfos: [
        {
            namespace: {
                service: 'core',
                entity: 'file',
            },
            associations: [],
            actions: [],
        },
        {
            namespace: {
                service: 'auth',
                entity: 'secret',
            },
            associations: [],
            actions: [],
        },
        {
            namespace: {
                service: 'core',
                entity: 'task',
            },
            associations: [
                {
                    relationship: 'parent_of',
                    type: 'many',
                    entityNamespace: {
                        service: 'core',
                        entity: 'task_run',
                    },
                    name: new Name('task_runs'),
                    otherAssociationName: new Name('task'),
                },
            ],
            actions: [
                {
                    name: new Name('run'),
                    methodNamespace: {
                        service: 'core',
                        entity: 'task',
                        types: undefined,
                        enum: undefined,
                        method: 'action_run',
                    },
                },
            ],
        },
        {
            namespace: {
                service: 'core',
                entity: 'task_run',
            },
            associations: [
                {
                    relationship: 'child_of',
                    type: 'one',
                    entityNamespace: {
                        service: 'core',
                        entity: 'task',
                    },
                    name: new Name('task'),
                },
            ],
            actions: [
                {
                    name: new Name('run'),
                    methodNamespace: {
                        service: 'core',
                        entity: 'task_run',
                        types: undefined,
                        enum: undefined,
                        method: 'action_run',
                    },
                },
            ],
        },
        {
            namespace: {
                service: 'user',
                entity: 'team',
            },
            associations: [
                {
                    relationship: 'parent_of',
                    type: 'many',
                    entityNamespace: {
                        service: 'user',
                        entity: 'user',
                    },
                    name: new Name('users'),
                    otherAssociationName: new Name('teams'),
                },
            ],
            actions: [],
        },
        {
            namespace: {
                service: 'user',
                entity: 'user',
            },
            associations: [
                {
                    relationship: 'child_of',
                    type: 'many',
                    entityNamespace: {
                        service: 'user',
                        entity: 'team',
                    },
                    name: new Name('teams'),
                    otherAssociationName: new Name('users'),
                },
            ],
            actions: [],
        },
    ],
    typeInfos: [
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
                service: 'user',
                entity: 'user',
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
                        name: new Name('user__user'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'user',
                            entity: 'user',
                            types: ['user'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const objectTypeInfo = appInfo.getTypeInfo({
                    service: 'user',
                    entity: 'user',
                    types: ['user'],
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
                service: 'user',
                entity: 'team',
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
                        name: new Name('user__team'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'user',
                            entity: 'team',
                            types: ['team'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const objectTypeInfo = appInfo.getTypeInfo({
                    service: 'user',
                    entity: 'team',
                    types: ['team'],
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
                entity: 'task_run',
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
                        name: new Name('core__task_run'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'core',
                            entity: 'task_run',
                            types: ['task_run'],
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
                    entity: 'task_run',
                    types: ['task_run'],
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
                entity: 'task',
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
                        name: new Name('core__task'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'core',
                            entity: 'task',
                            types: ['task'],
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
                    entity: 'task',
                    types: ['task'],
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
                service: undefined,
                entity: undefined,
                types: ['address'],
                enum: undefined,
                method: undefined,
            },
            fields: [
                {
                    name: new Name('line_1'),
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
                    name: new Name('line_2'),
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
                    name: new Name('city'),
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
                    name: new Name('state'),
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
                    name: new Name('postal_code'),
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
                    name: new Name('country'),
                    isRepeated: false,
                    isOptional: false,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('country'),
                        kind: fieldkind.EnumKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: undefined,
                            enum: 'country',
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances

                // Return the empty object
                return {
                    line1: undefined,
                    line2: undefined,
                    city: undefined,
                    state: undefined,
                    postalCode: undefined,
                    country: undefined,
                }
            },
        },
        {
            namespace: {
                service: undefined,
                entity: undefined,
                types: ['address_field_condition'],
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
                service: undefined,
                entity: undefined,
                types: ['address_filter'],
                enum: undefined,
                method: undefined,
            },
            fields: [
                {
                    name: new Name('parent_id'),
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
                    name: new Name('line_1'),
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
                    name: new Name('line_2'),
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
                    name: new Name('city'),
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
                    name: new Name('state'),
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
                    name: new Name('postal_code'),
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
                    name: new Name('country'),
                    isRepeated: false,
                    isOptional: true,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('country_condition'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: ['country_condition'],
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
                        name: new Name('address_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: ['address_filter'],
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
                        name: new Name('address_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: ['address_filter'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const countryTypeInfo = appInfo.getTypeInfo({
                    service: undefined,
                    entity: undefined,
                    types: ['country_condition'],
                    enum: undefined,
                    method: undefined,
                })
                const andTypeInfo = appInfo.getTypeInfo({
                    service: undefined,
                    entity: undefined,
                    types: ['address_filter'],
                    enum: undefined,
                    method: undefined,
                })
                const orTypeInfo = appInfo.getTypeInfo({
                    service: undefined,
                    entity: undefined,
                    types: ['address_filter'],
                    enum: undefined,
                    method: undefined,
                })

                // Return the empty object
                return {
                    parentID: undefined,
                    id: undefined,
                    line1: undefined,
                    line2: undefined,
                    city: undefined,
                    state: undefined,
                    postalCode: undefined,
                    country: countryTypeInfo.getEmptyObject(appInfo),
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
                entity: undefined,
                types: ['authenticate_response'],
                enum: undefined,
                method: undefined,
            },
            fields: [
                {
                    name: new Name('token'),
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
                    token: undefined,
                }
            },
        },
        {
            namespace: {
                service: 'auth',
                entity: undefined,
                types: ['authenticate_token_request'],
                enum: undefined,
                method: undefined,
            },
            fields: [
                {
                    name: new Name('token'),
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
                    token: undefined,
                }
            },
        },
        {
            namespace: {
                service: undefined,
                entity: undefined,
                types: ['contact'],
                enum: undefined,
                method: undefined,
            },
            fields: [
                {
                    name: new Name('name'),
                    isRepeated: false,
                    isOptional: false,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('person_name'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: ['person_name'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
                {
                    name: new Name('email'),
                    isRepeated: false,
                    isOptional: false,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('email'),
                        kind: fieldkind.EmailKind,
                        namespace: undefined,
                    },
                },
                {
                    name: new Name('address'),
                    isRepeated: false,
                    isOptional: false,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('address'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: ['address'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const nameTypeInfo = appInfo.getTypeInfo({
                    service: undefined,
                    entity: undefined,
                    types: ['person_name'],
                    enum: undefined,
                    method: undefined,
                })
                const addressTypeInfo = appInfo.getTypeInfo({
                    service: undefined,
                    entity: undefined,
                    types: ['address'],
                    enum: undefined,
                    method: undefined,
                })

                // Return the empty object
                return {
                    name: nameTypeInfo.getEmptyObject(appInfo),
                    email: undefined,
                    address: addressTypeInfo.getEmptyObject(appInfo),
                }
            },
        },
        {
            namespace: {
                service: undefined,
                entity: undefined,
                types: ['contact_field_condition'],
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
                service: undefined,
                entity: undefined,
                types: ['contact_filter'],
                enum: undefined,
                method: undefined,
            },
            fields: [
                {
                    name: new Name('parent_id'),
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
                    name: new Name('name'),
                    isRepeated: false,
                    isOptional: true,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('person_name_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: ['person_name_filter'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
                {
                    name: new Name('email'),
                    isRepeated: false,
                    isOptional: true,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('email_condition'),
                        kind: fieldkind.ConditionKind,
                        namespace: undefined,
                    },
                },
                {
                    name: new Name('address'),
                    isRepeated: false,
                    isOptional: true,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('address_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: ['address_filter'],
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
                        name: new Name('contact_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: ['contact_filter'],
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
                        name: new Name('contact_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: ['contact_filter'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const nameTypeInfo = appInfo.getTypeInfo({
                    service: undefined,
                    entity: undefined,
                    types: ['person_name_filter'],
                    enum: undefined,
                    method: undefined,
                })
                const addressTypeInfo = appInfo.getTypeInfo({
                    service: undefined,
                    entity: undefined,
                    types: ['address_filter'],
                    enum: undefined,
                    method: undefined,
                })
                const andTypeInfo = appInfo.getTypeInfo({
                    service: undefined,
                    entity: undefined,
                    types: ['contact_filter'],
                    enum: undefined,
                    method: undefined,
                })
                const orTypeInfo = appInfo.getTypeInfo({
                    service: undefined,
                    entity: undefined,
                    types: ['contact_filter'],
                    enum: undefined,
                    method: undefined,
                })

                // Return the empty object
                return {
                    parentID: undefined,
                    id: undefined,
                    name: nameTypeInfo.getEmptyObject(appInfo),
                    email: undefined,
                    address: addressTypeInfo.getEmptyObject(appInfo),
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
                service: undefined,
                entity: undefined,
                types: ['country_condition'],
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
                entity: undefined,
                types: ['cron_expression'],
                enum: undefined,
                method: undefined,
            },
            fields: [
                {
                    name: new Name('second'),
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
                    name: new Name('minute'),
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
                    name: new Name('hour'),
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
                    name: new Name('day_of_month'),
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
                    name: new Name('month'),
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
                    name: new Name('day_of_week'),
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
                    second: undefined,
                    minute: undefined,
                    hour: undefined,
                    dayOfMonth: undefined,
                    month: undefined,
                    dayOfWeek: undefined,
                }
            },
        },
        {
            namespace: {
                service: 'core',
                entity: undefined,
                types: ['cron_expression_field_condition'],
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
                entity: undefined,
                types: ['cron_expression_filter'],
                enum: undefined,
                method: undefined,
            },
            fields: [
                {
                    name: new Name('parent_id'),
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
                    name: new Name('second'),
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
                    name: new Name('minute'),
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
                    name: new Name('hour'),
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
                    name: new Name('day_of_month'),
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
                    name: new Name('month'),
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
                    name: new Name('day_of_week'),
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
                        name: new Name('core__cron_expression_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'core',
                            entity: undefined,
                            types: ['cron_expression_filter'],
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
                        name: new Name('core__cron_expression_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'core',
                            entity: undefined,
                            types: ['cron_expression_filter'],
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
                    entity: undefined,
                    types: ['cron_expression_filter'],
                    enum: undefined,
                    method: undefined,
                })
                const orTypeInfo = appInfo.getTypeInfo({
                    service: 'core',
                    entity: undefined,
                    types: ['cron_expression_filter'],
                    enum: undefined,
                    method: undefined,
                })

                // Return the empty object
                return {
                    parentID: undefined,
                    id: undefined,
                    second: undefined,
                    minute: undefined,
                    hour: undefined,
                    dayOfMonth: undefined,
                    month: undefined,
                    dayOfWeek: undefined,
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
                entity: undefined,
                types: ['entity_name_condition'],
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
                service: undefined,
                entity: undefined,
                types: ['entity_name_condition'],
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
                entity: undefined,
                types: ['entity_name_condition'],
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
                service: 'user',
                entity: undefined,
                types: ['entity_name_condition'],
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
                service: undefined,
                entity: undefined,
                types: ['gender_condition'],
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
                service: 'user',
                entity: 'user',
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
                service: 'user',
                entity: 'team',
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
                entity: 'task_run',
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
                entity: 'task',
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
                service: 'user',
                entity: 'team',
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
                        name: new Name('user__team__team_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'user',
                            entity: 'team',
                            types: ['team_filter'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const filterTypeInfo = appInfo.getTypeInfo({
                    service: 'user',
                    entity: 'team',
                    types: ['team_filter'],
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
                service: 'user',
                entity: 'user',
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
                        name: new Name('user__user__user_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'user',
                            entity: 'user',
                            types: ['user_filter'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const filterTypeInfo = appInfo.getTypeInfo({
                    service: 'user',
                    entity: 'user',
                    types: ['user_filter'],
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
                entity: 'task',
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
                        name: new Name('core__task__task_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'core',
                            entity: 'task',
                            types: ['task_filter'],
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
                    entity: 'task',
                    types: ['task_filter'],
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
                entity: 'task_run',
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
                        name: new Name('core__task_run__task_run_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'core',
                            entity: 'task_run',
                            types: ['task_run_filter'],
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
                    entity: 'task_run',
                    types: ['task_run_filter'],
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
                entity: 'task_run',
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
                        name: new Name('core__task_run'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'core',
                            entity: 'task_run',
                            types: ['task_run'],
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
                    entity: 'task_run',
                    types: ['task_run'],
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
                entity: 'task',
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
                        name: new Name('core__task'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'core',
                            entity: 'task',
                            types: ['task'],
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
                    entity: 'task',
                    types: ['task'],
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
                service: 'user',
                entity: 'user',
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
                        name: new Name('user__user'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'user',
                            entity: 'user',
                            types: ['user'],
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
                    service: 'user',
                    entity: 'user',
                    types: ['user'],
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
                service: 'user',
                entity: 'team',
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
                        name: new Name('user__team'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'user',
                            entity: 'team',
                            types: ['team'],
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
                    service: 'user',
                    entity: 'team',
                    types: ['team'],
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
                entity: undefined,
                types: ['login_request'],
                enum: undefined,
                method: undefined,
            },
            fields: [
                {
                    name: new Name('email'),
                    isRepeated: false,
                    isOptional: false,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('email'),
                        kind: fieldkind.EmailKind,
                        namespace: undefined,
                    },
                },
                {
                    name: new Name('password'),
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
                    email: undefined,
                    password: undefined,
                }
            },
        },
        {
            namespace: {
                service: undefined,
                entity: undefined,
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
                entity: 'task',
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
                entity: undefined,
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
                entity: undefined,
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
                entity: 'task_run',
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
                service: 'user',
                entity: 'user',
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
                service: 'user',
                entity: 'team',
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
                service: undefined,
                entity: undefined,
                types: ['person_name'],
                enum: undefined,
                method: undefined,
            },
            fields: [
                {
                    name: new Name('first_name'),
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
                    name: new Name('middle_initial'),
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
                    name: new Name('last_name'),
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
                    firstName: undefined,
                    middleInitial: undefined,
                    lastName: undefined,
                }
            },
        },
        {
            namespace: {
                service: undefined,
                entity: undefined,
                types: ['person_name_field_condition'],
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
                service: undefined,
                entity: undefined,
                types: ['person_name_filter'],
                enum: undefined,
                method: undefined,
            },
            fields: [
                {
                    name: new Name('parent_id'),
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
                    name: new Name('first_name'),
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
                    name: new Name('middle_initial'),
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
                    name: new Name('last_name'),
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
                        name: new Name('person_name_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: ['person_name_filter'],
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
                        name: new Name('person_name_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: ['person_name_filter'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const andTypeInfo = appInfo.getTypeInfo({
                    service: undefined,
                    entity: undefined,
                    types: ['person_name_filter'],
                    enum: undefined,
                    method: undefined,
                })
                const orTypeInfo = appInfo.getTypeInfo({
                    service: undefined,
                    entity: undefined,
                    types: ['person_name_filter'],
                    enum: undefined,
                    method: undefined,
                })

                // Return the empty object
                return {
                    parentID: undefined,
                    id: undefined,
                    firstName: undefined,
                    middleInitial: undefined,
                    lastName: undefined,
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
                service: undefined,
                entity: undefined,
                types: ['phone_number'],
                enum: undefined,
                method: undefined,
            },
            fields: [
                {
                    name: new Name('country_code'),
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
                {
                    name: new Name('number'),
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
                    name: new Name('extension'),
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
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances

                // Return the empty object
                return {
                    countryCode: undefined,
                    number: undefined,
                    extension: undefined,
                }
            },
        },
        {
            namespace: {
                service: undefined,
                entity: undefined,
                types: ['phone_number_field_condition'],
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
                service: undefined,
                entity: undefined,
                types: ['phone_number_filter'],
                enum: undefined,
                method: undefined,
            },
            fields: [
                {
                    name: new Name('parent_id'),
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
                    name: new Name('country_code'),
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
                    name: new Name('number'),
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
                    name: new Name('extension'),
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
                        name: new Name('phone_number_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: ['phone_number_filter'],
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
                        name: new Name('phone_number_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: ['phone_number_filter'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const andTypeInfo = appInfo.getTypeInfo({
                    service: undefined,
                    entity: undefined,
                    types: ['phone_number_filter'],
                    enum: undefined,
                    method: undefined,
                })
                const orTypeInfo = appInfo.getTypeInfo({
                    service: undefined,
                    entity: undefined,
                    types: ['phone_number_filter'],
                    enum: undefined,
                    method: undefined,
                })

                // Return the empty object
                return {
                    parentID: undefined,
                    id: undefined,
                    countryCode: undefined,
                    number: undefined,
                    extension: undefined,
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
                entity: 'task',
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
                entity: 'task_run',
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
                service: 'user',
                entity: 'team',
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
                service: 'user',
                entity: 'user',
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
                entity: undefined,
                types: ['register_user_request'],
                enum: undefined,
                method: undefined,
            },
            fields: [
                {
                    name: new Name('email'),
                    isRepeated: false,
                    isOptional: false,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('email'),
                        kind: fieldkind.EmailKind,
                        namespace: undefined,
                    },
                },
                {
                    name: new Name('name'),
                    isRepeated: false,
                    isOptional: false,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('person_name'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: ['person_name'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
                {
                    name: new Name('password'),
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
                    name: new Name('team_id'),
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
                const nameTypeInfo = appInfo.getTypeInfo({
                    service: undefined,
                    entity: undefined,
                    types: ['person_name'],
                    enum: undefined,
                    method: undefined,
                })

                // Return the empty object
                return {
                    email: undefined,
                    name: nameTypeInfo.getEmptyObject(appInfo),
                    password: undefined,
                    teamID: undefined,
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
                service: undefined,
                entity: undefined,
                types: ['service_name_condition'],
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
                service: 'core',
                entity: 'task_run',
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
                        name: new Name('core__task_run'),
                        kind: fieldkind.ForeignEntityKind,
                        namespace: {
                            service: 'core',
                            entity: 'task_run',
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
                entity: 'task',
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
                        name: new Name('core__task'),
                        kind: fieldkind.ForeignEntityKind,
                        namespace: {
                            service: 'core',
                            entity: 'task',
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
                service: 'user',
                entity: 'user',
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
                        name: new Name('user__user'),
                        kind: fieldkind.ForeignEntityKind,
                        namespace: {
                            service: 'user',
                            entity: 'user',
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
                service: 'user',
                entity: 'team',
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
                        name: new Name('user__team'),
                        kind: fieldkind.ForeignEntityKind,
                        namespace: {
                            service: 'user',
                            entity: 'team',
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
                entity: 'task',
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
                        name: new Name('core__task'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'core',
                            entity: 'task',
                            types: ['task'],
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
                    entity: 'task',
                    types: ['task'],
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
                entity: 'task_run',
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
                        name: new Name('core__task_run'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'core',
                            entity: 'task_run',
                            types: ['task_run'],
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
                    entity: 'task_run',
                    types: ['task_run'],
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
                service: 'user',
                entity: 'user',
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
                        name: new Name('user__user'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'user',
                            entity: 'user',
                            types: ['user'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const objectTypeInfo = appInfo.getTypeInfo({
                    service: 'user',
                    entity: 'user',
                    types: ['user'],
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
                service: 'user',
                entity: 'team',
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
                        name: new Name('user__team'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'user',
                            entity: 'team',
                            types: ['team'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const objectTypeInfo = appInfo.getTypeInfo({
                    service: 'user',
                    entity: 'team',
                    types: ['team'],
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
                entity: 'task_run',
                types: ['status_condition'],
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
                entity: 'task',
                types: ['task'],
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
                    name: new Name('name'),
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
                    name: new Name('description'),
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
                    name: new Name('method'),
                    isRepeated: false,
                    isOptional: false,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('method_name'),
                        kind: fieldkind.EnumKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: undefined,
                            enum: 'method_name',
                            method: undefined,
                        },
                    },
                },
                {
                    name: new Name('method_request_data'),
                    isRepeated: false,
                    isOptional: false,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('generic_data'),
                        kind: fieldkind.GenericDataKind,
                        namespace: undefined,
                    },
                },
                {
                    name: new Name('enabled'),
                    isRepeated: false,
                    isOptional: false,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('bool'),
                        kind: fieldkind.BoolKind,
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
                    name: undefined,
                    description: undefined,
                    method: undefined,
                    methodRequestData: undefined,
                    enabled: undefined,
                    createdAt: undefined,
                    updatedAt: undefined,
                    deletedAt: undefined,
                }
            },
        },
        {
            namespace: {
                service: 'core',
                entity: 'task',
                types: ['task_action_condition'],
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
                entity: 'task',
                types: ['task_field_condition'],
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
                entity: 'task',
                types: ['task_filter'],
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
                    name: new Name('name'),
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
                    name: new Name('description'),
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
                    name: new Name('method'),
                    isRepeated: false,
                    isOptional: true,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('method_name_condition'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: ['method_name_condition'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
                {
                    name: new Name('method_request_data'),
                    isRepeated: false,
                    isOptional: true,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('generic_data_condition'),
                        kind: fieldkind.ConditionKind,
                        namespace: undefined,
                    },
                },
                {
                    name: new Name('enabled'),
                    isRepeated: false,
                    isOptional: true,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('bool_condition'),
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
                        name: new Name('core__task__task_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'core',
                            entity: 'task',
                            types: ['task_filter'],
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
                        name: new Name('core__task__task_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'core',
                            entity: 'task',
                            types: ['task_filter'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const methodTypeInfo = appInfo.getTypeInfo({
                    service: undefined,
                    entity: undefined,
                    types: ['method_name_condition'],
                    enum: undefined,
                    method: undefined,
                })
                const andTypeInfo = appInfo.getTypeInfo({
                    service: 'core',
                    entity: 'task',
                    types: ['task_filter'],
                    enum: undefined,
                    method: undefined,
                })
                const orTypeInfo = appInfo.getTypeInfo({
                    service: 'core',
                    entity: 'task',
                    types: ['task_filter'],
                    enum: undefined,
                    method: undefined,
                })

                // Return the empty object
                return {
                    id: undefined,
                    name: undefined,
                    description: undefined,
                    method: methodTypeInfo.getEmptyObject(appInfo),
                    methodRequestData: undefined,
                    enabled: undefined,
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
                entity: 'task_run',
                types: ['task_run'],
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
                    name: new Name('started_at'),
                    isRepeated: false,
                    isOptional: false,
                    isMetaField: false,
                    excludeFromForm: true,
                    dtype: {
                        name: new Name('timestamp'),
                        kind: fieldkind.TimestampKind,
                        namespace: undefined,
                    },
                },
                {
                    name: new Name('completed_at'),
                    isRepeated: false,
                    isOptional: false,
                    isMetaField: false,
                    excludeFromForm: true,
                    dtype: {
                        name: new Name('timestamp'),
                        kind: fieldkind.TimestampKind,
                        namespace: undefined,
                    },
                },
                {
                    name: new Name('status'),
                    isRepeated: false,
                    isOptional: false,
                    isMetaField: false,
                    excludeFromForm: true,
                    dtype: {
                        name: new Name('core__task_run__status'),
                        kind: fieldkind.EnumKind,
                        namespace: {
                            service: 'core',
                            entity: 'task_run',
                            types: undefined,
                            enum: 'status',
                            method: undefined,
                        },
                    },
                },
                {
                    name: new Name('method_request_data'),
                    isRepeated: false,
                    isOptional: false,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('generic_data'),
                        kind: fieldkind.GenericDataKind,
                        namespace: undefined,
                    },
                },
                {
                    name: new Name('method_response_data'),
                    isRepeated: false,
                    isOptional: false,
                    isMetaField: false,
                    excludeFromForm: true,
                    dtype: {
                        name: new Name('generic_data'),
                        kind: fieldkind.GenericDataKind,
                        namespace: undefined,
                    },
                },
                {
                    name: new Name('triggered_by'),
                    isRepeated: false,
                    isOptional: false,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('core__task_run__triggered_by'),
                        kind: fieldkind.EnumKind,
                        namespace: {
                            service: 'core',
                            entity: 'task_run',
                            types: undefined,
                            enum: 'triggered_by',
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
                {
                    name: new Name('task_id'),
                    isRepeated: false,
                    isOptional: false,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('core__task'),
                        kind: fieldkind.ForeignEntityKind,
                        namespace: {
                            service: 'core',
                            entity: 'task',
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances

                // Return the empty object
                return {
                    id: undefined,
                    startedAt: undefined,
                    completedAt: undefined,
                    status: undefined,
                    methodRequestData: undefined,
                    methodResponseData: undefined,
                    triggeredBy: undefined,
                    createdAt: undefined,
                    updatedAt: undefined,
                    deletedAt: undefined,
                    taskID: undefined,
                }
            },
        },
        {
            namespace: {
                service: 'core',
                entity: 'task_run',
                types: ['task_run_action_condition'],
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
                entity: 'task_run',
                types: ['task_run_field_condition'],
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
                entity: 'task_run',
                types: ['task_run_filter'],
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
                    name: new Name('started_at'),
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
                    name: new Name('completed_at'),
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
                    name: new Name('status'),
                    isRepeated: false,
                    isOptional: true,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('status_condition'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'core',
                            entity: 'task_run',
                            types: ['status_condition'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
                {
                    name: new Name('method_request_data'),
                    isRepeated: false,
                    isOptional: true,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('generic_data_condition'),
                        kind: fieldkind.ConditionKind,
                        namespace: undefined,
                    },
                },
                {
                    name: new Name('method_response_data'),
                    isRepeated: false,
                    isOptional: true,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('generic_data_condition'),
                        kind: fieldkind.ConditionKind,
                        namespace: undefined,
                    },
                },
                {
                    name: new Name('triggered_by'),
                    isRepeated: false,
                    isOptional: true,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('triggered_by_condition'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'core',
                            entity: 'task_run',
                            types: ['triggered_by_condition'],
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
                    name: new Name('task_id'),
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
                        name: new Name('core__task_run__task_run_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'core',
                            entity: 'task_run',
                            types: ['task_run_filter'],
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
                        name: new Name('core__task_run__task_run_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'core',
                            entity: 'task_run',
                            types: ['task_run_filter'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const statusTypeInfo = appInfo.getTypeInfo({
                    service: 'core',
                    entity: 'task_run',
                    types: ['status_condition'],
                    enum: undefined,
                    method: undefined,
                })
                const triggeredByTypeInfo = appInfo.getTypeInfo({
                    service: 'core',
                    entity: 'task_run',
                    types: ['triggered_by_condition'],
                    enum: undefined,
                    method: undefined,
                })
                const andTypeInfo = appInfo.getTypeInfo({
                    service: 'core',
                    entity: 'task_run',
                    types: ['task_run_filter'],
                    enum: undefined,
                    method: undefined,
                })
                const orTypeInfo = appInfo.getTypeInfo({
                    service: 'core',
                    entity: 'task_run',
                    types: ['task_run_filter'],
                    enum: undefined,
                    method: undefined,
                })

                // Return the empty object
                return {
                    id: undefined,
                    startedAt: undefined,
                    completedAt: undefined,
                    status: statusTypeInfo.getEmptyObject(appInfo),
                    methodRequestData: undefined,
                    methodResponseData: undefined,
                    triggeredBy: triggeredByTypeInfo.getEmptyObject(appInfo),
                    createdAt: undefined,
                    updatedAt: undefined,
                    deletedAt: undefined,
                    taskID: undefined,
                    and: undefined,
                    or: undefined,
                }
            },
        },
        {
            namespace: {
                service: 'user',
                entity: 'team',
                types: ['team'],
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
                    name: new Name('name'),
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
                    name: undefined,
                    createdAt: undefined,
                    updatedAt: undefined,
                    deletedAt: undefined,
                }
            },
        },
        {
            namespace: {
                service: 'user',
                entity: 'team',
                types: ['team_field_condition'],
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
                service: 'user',
                entity: 'team',
                types: ['team_filter'],
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
                    name: new Name('name'),
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
                        name: new Name('user__team__team_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'user',
                            entity: 'team',
                            types: ['team_filter'],
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
                        name: new Name('user__team__team_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'user',
                            entity: 'team',
                            types: ['team_filter'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const andTypeInfo = appInfo.getTypeInfo({
                    service: 'user',
                    entity: 'team',
                    types: ['team_filter'],
                    enum: undefined,
                    method: undefined,
                })
                const orTypeInfo = appInfo.getTypeInfo({
                    service: 'user',
                    entity: 'team',
                    types: ['team_filter'],
                    enum: undefined,
                    method: undefined,
                })

                // Return the empty object
                return {
                    id: undefined,
                    name: undefined,
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
                entity: 'task_run',
                types: ['triggered_by_condition'],
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
                service: 'user',
                entity: 'team',
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
                        name: new Name('user__team'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'user',
                            entity: 'team',
                            types: ['team'],
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
                        name: new Name('user__team__team_field'),
                        kind: fieldkind.EnumKind,
                        namespace: {
                            service: 'user',
                            entity: 'team',
                            types: undefined,
                            enum: 'team_field',
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
                        name: new Name('user__team__team_field'),
                        kind: fieldkind.EnumKind,
                        namespace: {
                            service: 'user',
                            entity: 'team',
                            types: undefined,
                            enum: 'team_field',
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const objectTypeInfo = appInfo.getTypeInfo({
                    service: 'user',
                    entity: 'team',
                    types: ['team'],
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
                service: 'core',
                entity: 'task',
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
                        name: new Name('core__task'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'core',
                            entity: 'task',
                            types: ['task'],
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
                        name: new Name('core__task__task_field'),
                        kind: fieldkind.EnumKind,
                        namespace: {
                            service: 'core',
                            entity: 'task',
                            types: undefined,
                            enum: 'task_field',
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
                        name: new Name('core__task__task_field'),
                        kind: fieldkind.EnumKind,
                        namespace: {
                            service: 'core',
                            entity: 'task',
                            types: undefined,
                            enum: 'task_field',
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const objectTypeInfo = appInfo.getTypeInfo({
                    service: 'core',
                    entity: 'task',
                    types: ['task'],
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
                entity: 'task_run',
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
                        name: new Name('core__task_run'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'core',
                            entity: 'task_run',
                            types: ['task_run'],
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
                        name: new Name('core__task_run__task_run_field'),
                        kind: fieldkind.EnumKind,
                        namespace: {
                            service: 'core',
                            entity: 'task_run',
                            types: undefined,
                            enum: 'task_run_field',
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
                        name: new Name('core__task_run__task_run_field'),
                        kind: fieldkind.EnumKind,
                        namespace: {
                            service: 'core',
                            entity: 'task_run',
                            types: undefined,
                            enum: 'task_run_field',
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const objectTypeInfo = appInfo.getTypeInfo({
                    service: 'core',
                    entity: 'task_run',
                    types: ['task_run'],
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
                service: 'user',
                entity: 'user',
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
                        name: new Name('user__user'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'user',
                            entity: 'user',
                            types: ['user'],
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
                        name: new Name('user__user__user_field'),
                        kind: fieldkind.EnumKind,
                        namespace: {
                            service: 'user',
                            entity: 'user',
                            types: undefined,
                            enum: 'user_field',
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
                        name: new Name('user__user__user_field'),
                        kind: fieldkind.EnumKind,
                        namespace: {
                            service: 'user',
                            entity: 'user',
                            types: undefined,
                            enum: 'user_field',
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const objectTypeInfo = appInfo.getTypeInfo({
                    service: 'user',
                    entity: 'user',
                    types: ['user'],
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
        {
            namespace: {
                service: 'user',
                entity: 'team',
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
                        name: new Name('user__team'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'user',
                            entity: 'team',
                            types: ['team'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const objectTypeInfo = appInfo.getTypeInfo({
                    service: 'user',
                    entity: 'team',
                    types: ['team'],
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
                entity: 'task',
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
                        name: new Name('core__task'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'core',
                            entity: 'task',
                            types: ['task'],
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
                    entity: 'task',
                    types: ['task'],
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
                entity: 'task_run',
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
                        name: new Name('core__task_run'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'core',
                            entity: 'task_run',
                            types: ['task_run'],
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
                    entity: 'task_run',
                    types: ['task_run'],
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
                service: 'user',
                entity: 'user',
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
                        name: new Name('user__user'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'user',
                            entity: 'user',
                            types: ['user'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const objectTypeInfo = appInfo.getTypeInfo({
                    service: 'user',
                    entity: 'user',
                    types: ['user'],
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
        {
            namespace: {
                service: 'user',
                entity: 'user',
                types: ['user'],
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
                    name: new Name('name'),
                    isRepeated: false,
                    isOptional: false,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('person_name_[meta_fields]'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: ['person_name'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
                {
                    name: new Name('email'),
                    isRepeated: false,
                    isOptional: false,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('email'),
                        kind: fieldkind.EmailKind,
                        namespace: undefined,
                    },
                },
                {
                    name: new Name('addresses'),
                    isRepeated: true,
                    isOptional: true,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('address_[meta_fields]'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: ['address'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
                {
                    name: new Name('past_team_ids'),
                    isRepeated: true,
                    isOptional: true,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('user__team'),
                        kind: fieldkind.ForeignEntityKind,
                        namespace: {
                            service: 'user',
                            entity: 'team',
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
                {
                    name: new Name('teams_ids'),
                    isRepeated: true,
                    isOptional: false,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('user__team'),
                        kind: fieldkind.ForeignEntityKind,
                        namespace: {
                            service: 'user',
                            entity: 'team',
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const nameTypeInfo = appInfo.getTypeInfo({
                    service: undefined,
                    entity: undefined,
                    types: ['person_name'],
                    enum: undefined,
                    method: undefined,
                })
                const addressesTypeInfo = appInfo.getTypeInfo({
                    service: undefined,
                    entity: undefined,
                    types: ['address'],
                    enum: undefined,
                    method: undefined,
                })

                // Return the empty object
                return {
                    id: undefined,
                    name: nameTypeInfo.getEmptyObject(appInfo),
                    email: undefined,
                    addresses: undefined,
                    pastTeamIDs: undefined,
                    createdAt: undefined,
                    updatedAt: undefined,
                    deletedAt: undefined,
                    teamsIDs: undefined,
                }
            },
        },
        {
            namespace: {
                service: 'user',
                entity: 'user',
                types: ['user_field_condition'],
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
                service: 'user',
                entity: 'user',
                types: ['user_filter'],
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
                    name: new Name('name'),
                    isRepeated: false,
                    isOptional: true,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('person_name_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: ['person_name_filter'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
                {
                    name: new Name('email'),
                    isRepeated: false,
                    isOptional: true,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('email_condition'),
                        kind: fieldkind.ConditionKind,
                        namespace: undefined,
                    },
                },
                {
                    name: new Name('having_addresses'),
                    isRepeated: false,
                    isOptional: true,
                    isMetaField: false,
                    excludeFromForm: undefined,
                    dtype: {
                        name: new Name('address_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: undefined,
                            entity: undefined,
                            types: ['address_filter'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
                {
                    name: new Name('having_past_team_ids'),
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
                    name: new Name('having_teams_ids'),
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
                        name: new Name('user__user__user_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'user',
                            entity: 'user',
                            types: ['user_filter'],
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
                        name: new Name('user__user__user_filter'),
                        kind: fieldkind.NestedKind,
                        namespace: {
                            service: 'user',
                            entity: 'user',
                            types: ['user_filter'],
                            enum: undefined,
                            method: undefined,
                        },
                    },
                },
            ],
            getEmptyObjectFunc: (appInfo: App) => {
                // All nested types should be initialized with their empty instances
                const nameTypeInfo = appInfo.getTypeInfo({
                    service: undefined,
                    entity: undefined,
                    types: ['person_name_filter'],
                    enum: undefined,
                    method: undefined,
                })
                const havingAddressesTypeInfo = appInfo.getTypeInfo({
                    service: undefined,
                    entity: undefined,
                    types: ['address_filter'],
                    enum: undefined,
                    method: undefined,
                })
                const andTypeInfo = appInfo.getTypeInfo({
                    service: 'user',
                    entity: 'user',
                    types: ['user_filter'],
                    enum: undefined,
                    method: undefined,
                })
                const orTypeInfo = appInfo.getTypeInfo({
                    service: 'user',
                    entity: 'user',
                    types: ['user_filter'],
                    enum: undefined,
                    method: undefined,
                })

                // Return the empty object
                return {
                    id: undefined,
                    name: nameTypeInfo.getEmptyObject(appInfo),
                    email: undefined,
                    havingAddresses: havingAddressesTypeInfo.getEmptyObject(appInfo),
                    havingPastTeamIDs: undefined,
                    createdAt: undefined,
                    updatedAt: undefined,
                    deletedAt: undefined,
                    havingTeamsIDs: undefined,
                    and: undefined,
                    or: undefined,
                }
            },
        },
    ],
    enums: [
        {
            namespace: {
                service: undefined,
                entity: undefined,
                types: undefined,
                enum: 'address_field',
                method: undefined,
            },
            values: [
                {
                    id: 1,
                    value: 'ParentID',
                    description: undefined,
                    displayValue: 'Parent ID',
                },
                {
                    id: 2,
                    value: 'ID',
                    description: undefined,
                    displayValue: 'ID',
                },
                {
                    id: 3,
                    value: 'Line1',
                    description: undefined,
                    displayValue: 'Line 1',
                },
                {
                    id: 4,
                    value: 'Line2',
                    description: undefined,
                    displayValue: 'Line 2',
                },
                {
                    id: 5,
                    value: 'City',
                    description: undefined,
                    displayValue: 'City',
                },
                {
                    id: 6,
                    value: 'State',
                    description: undefined,
                    displayValue: 'State',
                },
                {
                    id: 7,
                    value: 'PostalCode',
                    description: undefined,
                    displayValue: 'Postal Code',
                },
                {
                    id: 8,
                    value: 'Country',
                    description: undefined,
                    displayValue: 'Country',
                },
                {
                    id: 9,
                    value: 'CreatedAt',
                    description: undefined,
                    displayValue: 'Created At',
                },
                {
                    id: 10,
                    value: 'UpdatedAt',
                    description: undefined,
                    displayValue: 'Updated At',
                },
                {
                    id: 11,
                    value: 'DeletedAt',
                    description: undefined,
                    displayValue: 'Deleted At',
                },
            ],
        },
        {
            namespace: {
                service: undefined,
                entity: undefined,
                types: undefined,
                enum: 'contact_field',
                method: undefined,
            },
            values: [
                {
                    id: 1,
                    value: 'ParentID',
                    description: undefined,
                    displayValue: 'Parent ID',
                },
                {
                    id: 2,
                    value: 'ID',
                    description: undefined,
                    displayValue: 'ID',
                },
                {
                    id: 3,
                    value: 'Name',
                    description: undefined,
                    displayValue: 'Name',
                },
                {
                    id: 4,
                    value: 'Name_ParentID',
                    description: undefined,
                    displayValue: 'Name -  Parent ID',
                },
                {
                    id: 5,
                    value: 'Name_ID',
                    description: undefined,
                    displayValue: 'Name -  ID',
                },
                {
                    id: 6,
                    value: 'Name_FirstName',
                    description: undefined,
                    displayValue: 'Name -  First Name',
                },
                {
                    id: 7,
                    value: 'Name_MiddleInitial',
                    description: undefined,
                    displayValue: 'Name -  Middle Initial',
                },
                {
                    id: 8,
                    value: 'Name_LastName',
                    description: undefined,
                    displayValue: 'Name -  Last Name',
                },
                {
                    id: 9,
                    value: 'Name_CreatedAt',
                    description: undefined,
                    displayValue: 'Name -  Created At',
                },
                {
                    id: 10,
                    value: 'Name_UpdatedAt',
                    description: undefined,
                    displayValue: 'Name -  Updated At',
                },
                {
                    id: 11,
                    value: 'Name_DeletedAt',
                    description: undefined,
                    displayValue: 'Name -  Deleted At',
                },
                {
                    id: 12,
                    value: 'Email',
                    description: undefined,
                    displayValue: 'Email',
                },
                {
                    id: 13,
                    value: 'Address',
                    description: undefined,
                    displayValue: 'Address',
                },
                {
                    id: 14,
                    value: 'Address_ParentID',
                    description: undefined,
                    displayValue: 'Address -  Parent ID',
                },
                {
                    id: 15,
                    value: 'Address_ID',
                    description: undefined,
                    displayValue: 'Address -  ID',
                },
                {
                    id: 16,
                    value: 'Address_Line1',
                    description: undefined,
                    displayValue: 'Address -  Line 1',
                },
                {
                    id: 17,
                    value: 'Address_Line2',
                    description: undefined,
                    displayValue: 'Address -  Line 2',
                },
                {
                    id: 18,
                    value: 'Address_City',
                    description: undefined,
                    displayValue: 'Address -  City',
                },
                {
                    id: 19,
                    value: 'Address_State',
                    description: undefined,
                    displayValue: 'Address -  State',
                },
                {
                    id: 20,
                    value: 'Address_PostalCode',
                    description: undefined,
                    displayValue: 'Address -  Postal Code',
                },
                {
                    id: 21,
                    value: 'Address_Country',
                    description: undefined,
                    displayValue: 'Address -  Country',
                },
                {
                    id: 22,
                    value: 'Address_CreatedAt',
                    description: undefined,
                    displayValue: 'Address -  Created At',
                },
                {
                    id: 23,
                    value: 'Address_UpdatedAt',
                    description: undefined,
                    displayValue: 'Address -  Updated At',
                },
                {
                    id: 24,
                    value: 'Address_DeletedAt',
                    description: undefined,
                    displayValue: 'Address -  Deleted At',
                },
                {
                    id: 25,
                    value: 'CreatedAt',
                    description: undefined,
                    displayValue: 'Created At',
                },
                {
                    id: 26,
                    value: 'UpdatedAt',
                    description: undefined,
                    displayValue: 'Updated At',
                },
                {
                    id: 27,
                    value: 'DeletedAt',
                    description: undefined,
                    displayValue: 'Deleted At',
                },
            ],
        },
        {
            namespace: {
                service: undefined,
                entity: undefined,
                types: undefined,
                enum: 'country',
                method: undefined,
            },
            values: [
                {
                    id: 1,
                    value: 'USA',
                    description: undefined,
                    displayValue: 'USA',
                },
                {
                    id: 2,
                    value: 'Canada',
                    description: undefined,
                    displayValue: 'Canada',
                },
                {
                    id: 3,
                    value: 'Mexico',
                    description: undefined,
                    displayValue: 'Mexico',
                },
            ],
        },
        {
            namespace: {
                service: 'core',
                entity: undefined,
                types: undefined,
                enum: 'cron_expression_field',
                method: undefined,
            },
            values: [
                {
                    id: 1,
                    value: 'ParentID',
                    description: undefined,
                    displayValue: 'Parent ID',
                },
                {
                    id: 2,
                    value: 'ID',
                    description: undefined,
                    displayValue: 'ID',
                },
                {
                    id: 3,
                    value: 'Second',
                    description: undefined,
                    displayValue: 'Second',
                },
                {
                    id: 4,
                    value: 'Minute',
                    description: undefined,
                    displayValue: 'Minute',
                },
                {
                    id: 5,
                    value: 'Hour',
                    description: undefined,
                    displayValue: 'Hour',
                },
                {
                    id: 6,
                    value: 'DayOfMonth',
                    description: undefined,
                    displayValue: 'Day Of Month',
                },
                {
                    id: 7,
                    value: 'Month',
                    description: undefined,
                    displayValue: 'Month',
                },
                {
                    id: 8,
                    value: 'DayOfWeek',
                    description: undefined,
                    displayValue: 'Day Of Week',
                },
                {
                    id: 9,
                    value: 'CreatedAt',
                    description: undefined,
                    displayValue: 'Created At',
                },
                {
                    id: 10,
                    value: 'UpdatedAt',
                    description: undefined,
                    displayValue: 'Updated At',
                },
                {
                    id: 11,
                    value: 'DeletedAt',
                    description: undefined,
                    displayValue: 'Deleted At',
                },
            ],
        },
        {
            namespace: {
                service: 'core',
                entity: undefined,
                types: undefined,
                enum: 'entity_name',
                method: undefined,
            },
            values: [
                {
                    id: 1,
                    value: 'File',
                    description: undefined,
                    displayValue: 'File',
                },
                {
                    id: 2,
                    value: 'Task',
                    description: undefined,
                    displayValue: 'Task',
                },
                {
                    id: 3,
                    value: 'TaskRun',
                    description: undefined,
                    displayValue: 'Task Run',
                },
            ],
        },
        {
            namespace: {
                service: undefined,
                entity: undefined,
                types: undefined,
                enum: 'entity_name',
                method: undefined,
            },
            values: [
                {
                    id: 1,
                    value: 'Auth_Secret',
                    description: undefined,
                    displayValue: 'Auth -  Secret',
                },
                {
                    id: 2,
                    value: 'Core_File',
                    description: undefined,
                    displayValue: 'Core -  File',
                },
                {
                    id: 3,
                    value: 'Core_Task',
                    description: undefined,
                    displayValue: 'Core -  Task',
                },
                {
                    id: 4,
                    value: 'Core_TaskRun',
                    description: undefined,
                    displayValue: 'Core -  Task Run',
                },
                {
                    id: 5,
                    value: 'User_User',
                    description: undefined,
                    displayValue: 'User -  User',
                },
                {
                    id: 6,
                    value: 'User_Team',
                    description: undefined,
                    displayValue: 'User -  Team',
                },
            ],
        },
        {
            namespace: {
                service: 'user',
                entity: undefined,
                types: undefined,
                enum: 'entity_name',
                method: undefined,
            },
            values: [
                {
                    id: 1,
                    value: 'User',
                    description: undefined,
                    displayValue: 'User',
                },
                {
                    id: 2,
                    value: 'Team',
                    description: undefined,
                    displayValue: 'Team',
                },
            ],
        },
        {
            namespace: {
                service: 'auth',
                entity: undefined,
                types: undefined,
                enum: 'entity_name',
                method: undefined,
            },
            values: [
                {
                    id: 1,
                    value: 'Secret',
                    description: undefined,
                    displayValue: 'Secret',
                },
            ],
        },
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
                service: undefined,
                entity: undefined,
                types: undefined,
                enum: 'gender',
                method: undefined,
            },
            values: [
                {
                    id: 1,
                    value: 'Male',
                    description: undefined,
                    displayValue: 'Male',
                },
                {
                    id: 2,
                    value: 'Female',
                    description: undefined,
                    displayValue: 'Female',
                },
                {
                    id: 3,
                    value: 'Other',
                    description: undefined,
                    displayValue: 'Other',
                },
            ],
        },
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
                service: 'user',
                entity: 'team',
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
                service: undefined,
                entity: undefined,
                types: undefined,
                enum: 'method_name',
                method: undefined,
            },
            values: [
                {
                    id: 1,
                    value: 'Auth_LoginUser',
                    description: undefined,
                    displayValue: 'Auth -  Login User',
                },
                {
                    id: 2,
                    value: 'Auth_RegisterUser',
                    description: undefined,
                    displayValue: 'Auth -  Register User',
                },
                {
                    id: 3,
                    value: 'Auth_AuthenticateToken',
                    description: undefined,
                    displayValue: 'Auth -  Authenticate Token',
                },
                {
                    id: 4,
                    value: 'Core_FileUpload',
                    description: undefined,
                    displayValue: 'Core -  File Upload',
                },
                {
                    id: 5,
                    value: 'Auth_Secret_Add',
                    description: undefined,
                    displayValue: 'Auth -  Secret -  Add',
                },
                {
                    id: 6,
                    value: 'Auth_Secret_Update',
                    description: undefined,
                    displayValue: 'Auth -  Secret -  Update',
                },
                {
                    id: 7,
                    value: 'Auth_Secret_Get',
                    description: undefined,
                    displayValue: 'Auth -  Secret -  Get',
                },
                {
                    id: 8,
                    value: 'Auth_Secret_List',
                    description: undefined,
                    displayValue: 'Auth -  Secret -  List',
                },
                {
                    id: 9,
                    value: 'Auth_Secret_QueryByText',
                    description: undefined,
                    displayValue: 'Auth -  Secret -  Query By Text',
                },
                {
                    id: 10,
                    value: 'Core_File_Add',
                    description: undefined,
                    displayValue: 'Core -  File -  Add',
                },
                {
                    id: 11,
                    value: 'Core_File_Update',
                    description: undefined,
                    displayValue: 'Core -  File -  Update',
                },
                {
                    id: 12,
                    value: 'Core_File_Get',
                    description: undefined,
                    displayValue: 'Core -  File -  Get',
                },
                {
                    id: 13,
                    value: 'Core_File_List',
                    description: undefined,
                    displayValue: 'Core -  File -  List',
                },
                {
                    id: 14,
                    value: 'Core_File_QueryByText',
                    description: undefined,
                    displayValue: 'Core -  File -  Query By Text',
                },
                {
                    id: 15,
                    value: 'Core_Task_Add',
                    description: undefined,
                    displayValue: 'Core -  Task -  Add',
                },
                {
                    id: 16,
                    value: 'Core_Task_Update',
                    description: undefined,
                    displayValue: 'Core -  Task -  Update',
                },
                {
                    id: 17,
                    value: 'Core_Task_Get',
                    description: undefined,
                    displayValue: 'Core -  Task -  Get',
                },
                {
                    id: 18,
                    value: 'Core_Task_List',
                    description: undefined,
                    displayValue: 'Core -  Task -  List',
                },
                {
                    id: 19,
                    value: 'Core_Task_QueryByText',
                    description: undefined,
                    displayValue: 'Core -  Task -  Query By Text',
                },
                {
                    id: 20,
                    value: 'Core_TaskRun_Add',
                    description: undefined,
                    displayValue: 'Core -  Task Run -  Add',
                },
                {
                    id: 21,
                    value: 'Core_TaskRun_Update',
                    description: undefined,
                    displayValue: 'Core -  Task Run -  Update',
                },
                {
                    id: 22,
                    value: 'Core_TaskRun_Get',
                    description: undefined,
                    displayValue: 'Core -  Task Run -  Get',
                },
                {
                    id: 23,
                    value: 'Core_TaskRun_List',
                    description: undefined,
                    displayValue: 'Core -  Task Run -  List',
                },
                {
                    id: 24,
                    value: 'Core_TaskRun_QueryByText',
                    description: undefined,
                    displayValue: 'Core -  Task Run -  Query By Text',
                },
                {
                    id: 25,
                    value: 'User_User_Add',
                    description: undefined,
                    displayValue: 'User -  User -  Add',
                },
                {
                    id: 26,
                    value: 'User_User_Update',
                    description: undefined,
                    displayValue: 'User -  User -  Update',
                },
                {
                    id: 27,
                    value: 'User_User_Get',
                    description: undefined,
                    displayValue: 'User -  User -  Get',
                },
                {
                    id: 28,
                    value: 'User_User_List',
                    description: undefined,
                    displayValue: 'User -  User -  List',
                },
                {
                    id: 29,
                    value: 'User_User_QueryByText',
                    description: undefined,
                    displayValue: 'User -  User -  Query By Text',
                },
                {
                    id: 30,
                    value: 'User_Team_Add',
                    description: undefined,
                    displayValue: 'User -  Team -  Add',
                },
                {
                    id: 31,
                    value: 'User_Team_Update',
                    description: undefined,
                    displayValue: 'User -  Team -  Update',
                },
                {
                    id: 32,
                    value: 'User_Team_Get',
                    description: undefined,
                    displayValue: 'User -  Team -  Get',
                },
                {
                    id: 33,
                    value: 'User_Team_List',
                    description: undefined,
                    displayValue: 'User -  Team -  List',
                },
                {
                    id: 34,
                    value: 'User_Team_QueryByText',
                    description: undefined,
                    displayValue: 'User -  Team -  Query By Text',
                },
                {
                    id: 35,
                    value: 'Core_File_HookCreatePre',
                    description: undefined,
                    displayValue: 'Core -  File -  Hook Create Pre',
                },
                {
                    id: 36,
                    value: 'Core_Task_HookSavePre',
                    description: undefined,
                    displayValue: 'Core -  Task -  Hook Save Pre',
                },
                {
                    id: 37,
                    value: 'Core_TaskRun_HookCreatePre',
                    description: undefined,
                    displayValue: 'Core -  Task Run -  Hook Create Pre',
                },
                {
                    id: 38,
                    value: 'User_User_HookInit',
                    description: undefined,
                    displayValue: 'User -  User -  Hook Init',
                },
                {
                    id: 39,
                    value: 'Core_Task_ActionRun',
                    description: undefined,
                    displayValue: 'Core -  Task -  Action Run',
                },
                {
                    id: 40,
                    value: 'Core_TaskRun_ActionRun',
                    description: undefined,
                    displayValue: 'Core -  Task Run -  Action Run',
                },
            ],
        },
        {
            namespace: {
                service: 'user',
                entity: 'user',
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
            ],
        },
        {
            namespace: {
                service: 'core',
                entity: 'task_run',
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
                    value: 'ActionRun',
                    description: undefined,
                    displayValue: 'Action Run',
                },
            ],
        },
        {
            namespace: {
                service: 'core',
                entity: 'task',
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
                    value: 'HookSavePre',
                    description: undefined,
                    displayValue: 'Hook Save Pre',
                },
                {
                    id: 7,
                    value: 'ActionRun',
                    description: undefined,
                    displayValue: 'Action Run',
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
                entity: undefined,
                types: undefined,
                enum: 'method_name',
                method: undefined,
            },
            values: [
                {
                    id: 1,
                    value: 'FileUpload',
                    description: undefined,
                    displayValue: 'File Upload',
                },
            ],
        },
        {
            namespace: {
                service: 'auth',
                entity: undefined,
                types: undefined,
                enum: 'method_name',
                method: undefined,
            },
            values: [
                {
                    id: 1,
                    value: 'LoginUser',
                    description: undefined,
                    displayValue: 'Login User',
                },
                {
                    id: 2,
                    value: 'RegisterUser',
                    description: undefined,
                    displayValue: 'Register User',
                },
                {
                    id: 3,
                    value: 'AuthenticateToken',
                    description: undefined,
                    displayValue: 'Authenticate Token',
                },
            ],
        },
        {
            namespace: {
                service: undefined,
                entity: undefined,
                types: undefined,
                enum: 'person_name_field',
                method: undefined,
            },
            values: [
                {
                    id: 1,
                    value: 'ParentID',
                    description: undefined,
                    displayValue: 'Parent ID',
                },
                {
                    id: 2,
                    value: 'ID',
                    description: undefined,
                    displayValue: 'ID',
                },
                {
                    id: 3,
                    value: 'FirstName',
                    description: undefined,
                    displayValue: 'First Name',
                },
                {
                    id: 4,
                    value: 'MiddleInitial',
                    description: undefined,
                    displayValue: 'Middle Initial',
                },
                {
                    id: 5,
                    value: 'LastName',
                    description: undefined,
                    displayValue: 'Last Name',
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
                service: undefined,
                entity: undefined,
                types: undefined,
                enum: 'phone_number_field',
                method: undefined,
            },
            values: [
                {
                    id: 1,
                    value: 'ParentID',
                    description: undefined,
                    displayValue: 'Parent ID',
                },
                {
                    id: 2,
                    value: 'ID',
                    description: undefined,
                    displayValue: 'ID',
                },
                {
                    id: 3,
                    value: 'CountryCode',
                    description: undefined,
                    displayValue: 'Country Code',
                },
                {
                    id: 4,
                    value: 'Number',
                    description: undefined,
                    displayValue: 'Number',
                },
                {
                    id: 5,
                    value: 'Extension',
                    description: undefined,
                    displayValue: 'Extension',
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
                service: undefined,
                entity: undefined,
                types: undefined,
                enum: 'service_name',
                method: undefined,
            },
            values: [
                {
                    id: 1,
                    value: 'Auth',
                    description: undefined,
                    displayValue: 'Auth',
                },
                {
                    id: 2,
                    value: 'Core',
                    description: undefined,
                    displayValue: 'Core',
                },
                {
                    id: 3,
                    value: 'User',
                    description: undefined,
                    displayValue: 'User',
                },
            ],
        },
        {
            namespace: {
                service: 'core',
                entity: 'task_run',
                types: undefined,
                enum: 'status',
                method: undefined,
            },
            values: [
                {
                    id: 1,
                    value: 'Pending',
                    description: undefined,
                    displayValue: 'Pending',
                },
                {
                    id: 2,
                    value: 'Running',
                    description: undefined,
                    displayValue: 'Running',
                },
                {
                    id: 3,
                    value: 'Success',
                    description: undefined,
                    displayValue: 'Success',
                },
                {
                    id: 4,
                    value: 'Failed',
                    description: undefined,
                    displayValue: 'Failed',
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
        {
            namespace: {
                service: 'core',
                entity: 'task',
                types: undefined,
                enum: 'task_action',
                method: undefined,
            },
            values: [
                {
                    id: 1,
                    value: 'Run',
                    description: 'Run the task',
                    displayValue: 'Run',
                },
            ],
        },
        {
            namespace: {
                service: 'core',
                entity: 'task',
                types: undefined,
                enum: 'task_field',
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
                    value: 'Name',
                    description: undefined,
                    displayValue: 'Name',
                },
                {
                    id: 3,
                    value: 'Description',
                    description: 'string',
                    displayValue: 'Description',
                },
                {
                    id: 4,
                    value: 'Method',
                    description: undefined,
                    displayValue: 'Method',
                },
                {
                    id: 5,
                    value: 'MethodRequestData',
                    description: undefined,
                    displayValue: 'Method Request Data',
                },
                {
                    id: 6,
                    value: 'Enabled',
                    description: undefined,
                    displayValue: 'Enabled',
                },
                {
                    id: 7,
                    value: 'CreatedAt',
                    description: undefined,
                    displayValue: 'Created At',
                },
                {
                    id: 8,
                    value: 'UpdatedAt',
                    description: undefined,
                    displayValue: 'Updated At',
                },
                {
                    id: 9,
                    value: 'DeletedAt',
                    description: undefined,
                    displayValue: 'Deleted At',
                },
            ],
        },
        {
            namespace: {
                service: 'core',
                entity: 'task_run',
                types: undefined,
                enum: 'task_run_action',
                method: undefined,
            },
            values: [
                {
                    id: 1,
                    value: 'Run',
                    description: 'Run the task',
                    displayValue: 'Run',
                },
            ],
        },
        {
            namespace: {
                service: 'core',
                entity: 'task_run',
                types: undefined,
                enum: 'task_run_field',
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
                    value: 'StartedAt',
                    description: undefined,
                    displayValue: 'Started At',
                },
                {
                    id: 3,
                    value: 'CompletedAt',
                    description: undefined,
                    displayValue: 'Completed At',
                },
                {
                    id: 4,
                    value: 'Status',
                    description: undefined,
                    displayValue: 'Status',
                },
                {
                    id: 5,
                    value: 'MethodRequestData',
                    description: undefined,
                    displayValue: 'Method Request Data',
                },
                {
                    id: 6,
                    value: 'MethodResponseData',
                    description: undefined,
                    displayValue: 'Method Response Data',
                },
                {
                    id: 7,
                    value: 'TriggeredBy',
                    description: undefined,
                    displayValue: 'Triggered By',
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
                {
                    id: 11,
                    value: 'TaskID',
                    description: undefined,
                    displayValue: 'Task ID',
                },
            ],
        },
        {
            namespace: {
                service: 'user',
                entity: 'team',
                types: undefined,
                enum: 'team_field',
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
                    value: 'Name',
                    description: undefined,
                    displayValue: 'Name',
                },
                {
                    id: 3,
                    value: 'CreatedAt',
                    description: undefined,
                    displayValue: 'Created At',
                },
                {
                    id: 4,
                    value: 'UpdatedAt',
                    description: undefined,
                    displayValue: 'Updated At',
                },
                {
                    id: 5,
                    value: 'DeletedAt',
                    description: undefined,
                    displayValue: 'Deleted At',
                },
            ],
        },
        {
            namespace: {
                service: 'core',
                entity: 'task_run',
                types: undefined,
                enum: 'triggered_by',
                method: undefined,
            },
            values: [
                {
                    id: 1,
                    value: 'Cron',
                    description: undefined,
                    displayValue: 'Cron',
                },
                {
                    id: 2,
                    value: 'Manual',
                    description: undefined,
                    displayValue: 'Manual',
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
        {
            namespace: {
                service: 'user',
                entity: 'user',
                types: undefined,
                enum: 'user_field',
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
                    value: 'Name',
                    description: undefined,
                    displayValue: 'Name',
                },
                {
                    id: 3,
                    value: 'Name_ParentID',
                    description: undefined,
                    displayValue: 'Name -  Parent ID',
                },
                {
                    id: 4,
                    value: 'Name_ID',
                    description: undefined,
                    displayValue: 'Name -  ID',
                },
                {
                    id: 5,
                    value: 'Name_FirstName',
                    description: undefined,
                    displayValue: 'Name -  First Name',
                },
                {
                    id: 6,
                    value: 'Name_MiddleInitial',
                    description: undefined,
                    displayValue: 'Name -  Middle Initial',
                },
                {
                    id: 7,
                    value: 'Name_LastName',
                    description: undefined,
                    displayValue: 'Name -  Last Name',
                },
                {
                    id: 8,
                    value: 'Name_CreatedAt',
                    description: undefined,
                    displayValue: 'Name -  Created At',
                },
                {
                    id: 9,
                    value: 'Name_UpdatedAt',
                    description: undefined,
                    displayValue: 'Name -  Updated At',
                },
                {
                    id: 10,
                    value: 'Name_DeletedAt',
                    description: undefined,
                    displayValue: 'Name -  Deleted At',
                },
                {
                    id: 11,
                    value: 'Email',
                    description: undefined,
                    displayValue: 'Email',
                },
                {
                    id: 12,
                    value: 'Addresses',
                    description: undefined,
                    displayValue: 'Addresses',
                },
                {
                    id: 13,
                    value: 'Addresses_ParentID',
                    description: undefined,
                    displayValue: 'Addresses -  Parent ID',
                },
                {
                    id: 14,
                    value: 'Addresses_ID',
                    description: undefined,
                    displayValue: 'Addresses -  ID',
                },
                {
                    id: 15,
                    value: 'Addresses_Line1',
                    description: undefined,
                    displayValue: 'Addresses -  Line 1',
                },
                {
                    id: 16,
                    value: 'Addresses_Line2',
                    description: undefined,
                    displayValue: 'Addresses -  Line 2',
                },
                {
                    id: 17,
                    value: 'Addresses_City',
                    description: undefined,
                    displayValue: 'Addresses -  City',
                },
                {
                    id: 18,
                    value: 'Addresses_State',
                    description: undefined,
                    displayValue: 'Addresses -  State',
                },
                {
                    id: 19,
                    value: 'Addresses_PostalCode',
                    description: undefined,
                    displayValue: 'Addresses -  Postal Code',
                },
                {
                    id: 20,
                    value: 'Addresses_Country',
                    description: undefined,
                    displayValue: 'Addresses -  Country',
                },
                {
                    id: 21,
                    value: 'Addresses_CreatedAt',
                    description: undefined,
                    displayValue: 'Addresses -  Created At',
                },
                {
                    id: 22,
                    value: 'Addresses_UpdatedAt',
                    description: undefined,
                    displayValue: 'Addresses -  Updated At',
                },
                {
                    id: 23,
                    value: 'Addresses_DeletedAt',
                    description: undefined,
                    displayValue: 'Addresses -  Deleted At',
                },
                {
                    id: 24,
                    value: 'PastTeamIDs',
                    description: undefined,
                    displayValue: 'Past Team IDs',
                },
                {
                    id: 25,
                    value: 'CreatedAt',
                    description: undefined,
                    displayValue: 'Created At',
                },
                {
                    id: 26,
                    value: 'UpdatedAt',
                    description: undefined,
                    displayValue: 'Updated At',
                },
                {
                    id: 27,
                    value: 'DeletedAt',
                    description: undefined,
                    displayValue: 'Deleted At',
                },
                {
                    id: 28,
                    value: 'TeamsIDs',
                    description: undefined,
                    displayValue: 'Teams IDs',
                },
            ],
        },
    ],
    methods: [
        {
            namespace: {
                service: 'core',
                entity: 'task',
                types: undefined,
                enum: undefined,
                method: 'action_run',
            },
            requestTypeNamespace: {
                service: 'core',
                entity: 'task',
                types: ['standard_entity_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'core',
                entity: 'task',
                types: ['standard_entity_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'POST',
                    path: '/run',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'core',
                entity: 'task_run',
                types: undefined,
                enum: undefined,
                method: 'action_run',
            },
            requestTypeNamespace: {
                service: 'core',
                entity: 'task_run',
                types: ['standard_entity_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'core',
                entity: 'task_run',
                types: ['standard_entity_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'POST',
                    path: '/run',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'user',
                entity: 'team',
                types: undefined,
                enum: undefined,
                method: 'add',
            },
            requestTypeNamespace: {
                service: 'user',
                entity: 'team',
                types: ['add_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'user',
                entity: 'team',
                types: ['team'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'POST',
                    path: '',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'auth',
                entity: 'secret',
                types: undefined,
                enum: undefined,
                method: 'add',
            },
            requestTypeNamespace: {
                service: 'auth',
                entity: 'secret',
                types: ['add_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'auth',
                entity: 'secret',
                types: ['secret'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'POST',
                    path: '',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'user',
                entity: 'user',
                types: undefined,
                enum: undefined,
                method: 'add',
            },
            requestTypeNamespace: {
                service: 'user',
                entity: 'user',
                types: ['add_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'user',
                entity: 'user',
                types: ['user'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'POST',
                    path: '',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'core',
                entity: 'task_run',
                types: undefined,
                enum: undefined,
                method: 'add',
            },
            requestTypeNamespace: {
                service: 'core',
                entity: 'task_run',
                types: ['add_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'core',
                entity: 'task_run',
                types: ['task_run'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'POST',
                    path: '',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'core',
                entity: 'task',
                types: undefined,
                enum: undefined,
                method: 'add',
            },
            requestTypeNamespace: {
                service: 'core',
                entity: 'task',
                types: ['add_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'core',
                entity: 'task',
                types: ['task'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'POST',
                    path: '',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'core',
                entity: 'file',
                types: undefined,
                enum: undefined,
                method: 'add',
            },
            requestTypeNamespace: {
                service: 'core',
                entity: 'file',
                types: ['add_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'core',
                entity: 'file',
                types: ['file'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'POST',
                    path: '',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'auth',
                entity: undefined,
                types: undefined,
                enum: undefined,
                method: 'authenticate_token',
            },
            requestTypeNamespace: {
                service: 'auth',
                entity: undefined,
                types: ['authenticate_token_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'auth',
                entity: undefined,
                types: ['authenticate_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'POST',
                    path: '/authenticate_token',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'core',
                entity: undefined,
                types: undefined,
                enum: undefined,
                method: 'file_upload',
            },
            requestTypeNamespace: undefined,
            responseTypeNamespace: {
                service: 'core',
                entity: 'file',
                types: ['file'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'POST',
                    path: 'file/upload',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'core',
                entity: 'task',
                types: undefined,
                enum: undefined,
                method: 'get',
            },
            requestTypeNamespace: {
                service: 'core',
                entity: 'task',
                types: ['get_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'core',
                entity: 'task',
                types: ['task'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'GET',
                    path: '',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'user',
                entity: 'user',
                types: undefined,
                enum: undefined,
                method: 'get',
            },
            requestTypeNamespace: {
                service: 'user',
                entity: 'user',
                types: ['get_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'user',
                entity: 'user',
                types: ['user'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'GET',
                    path: '',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'user',
                entity: 'team',
                types: undefined,
                enum: undefined,
                method: 'get',
            },
            requestTypeNamespace: {
                service: 'user',
                entity: 'team',
                types: ['get_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'user',
                entity: 'team',
                types: ['team'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'GET',
                    path: '',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'core',
                entity: 'file',
                types: undefined,
                enum: undefined,
                method: 'get',
            },
            requestTypeNamespace: {
                service: 'core',
                entity: 'file',
                types: ['get_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'core',
                entity: 'file',
                types: ['file'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'GET',
                    path: '',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'auth',
                entity: 'secret',
                types: undefined,
                enum: undefined,
                method: 'get',
            },
            requestTypeNamespace: {
                service: 'auth',
                entity: 'secret',
                types: ['get_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'auth',
                entity: 'secret',
                types: ['secret'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'GET',
                    path: '',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'core',
                entity: 'task_run',
                types: undefined,
                enum: undefined,
                method: 'get',
            },
            requestTypeNamespace: {
                service: 'core',
                entity: 'task_run',
                types: ['get_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'core',
                entity: 'task_run',
                types: ['task_run'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'GET',
                    path: '',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'user',
                entity: 'user',
                types: undefined,
                enum: undefined,
                method: 'list',
            },
            requestTypeNamespace: {
                service: 'user',
                entity: 'user',
                types: ['list_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'user',
                entity: 'user',
                types: ['list_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'GET',
                    path: '/list',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'core',
                entity: 'task',
                types: undefined,
                enum: undefined,
                method: 'list',
            },
            requestTypeNamespace: {
                service: 'core',
                entity: 'task',
                types: ['list_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'core',
                entity: 'task',
                types: ['list_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'GET',
                    path: '/list',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'user',
                entity: 'team',
                types: undefined,
                enum: undefined,
                method: 'list',
            },
            requestTypeNamespace: {
                service: 'user',
                entity: 'team',
                types: ['list_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'user',
                entity: 'team',
                types: ['list_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'GET',
                    path: '/list',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'auth',
                entity: 'secret',
                types: undefined,
                enum: undefined,
                method: 'list',
            },
            requestTypeNamespace: {
                service: 'auth',
                entity: 'secret',
                types: ['list_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'auth',
                entity: 'secret',
                types: ['list_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'GET',
                    path: '/list',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'core',
                entity: 'task_run',
                types: undefined,
                enum: undefined,
                method: 'list',
            },
            requestTypeNamespace: {
                service: 'core',
                entity: 'task_run',
                types: ['list_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'core',
                entity: 'task_run',
                types: ['list_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'GET',
                    path: '/list',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'core',
                entity: 'file',
                types: undefined,
                enum: undefined,
                method: 'list',
            },
            requestTypeNamespace: {
                service: 'core',
                entity: 'file',
                types: ['list_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'core',
                entity: 'file',
                types: ['list_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'GET',
                    path: '/list',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'auth',
                entity: undefined,
                types: undefined,
                enum: undefined,
                method: 'login_user',
            },
            requestTypeNamespace: {
                service: 'auth',
                entity: undefined,
                types: ['login_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'auth',
                entity: undefined,
                types: ['authenticate_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'POST',
                    path: '/login',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'auth',
                entity: 'secret',
                types: undefined,
                enum: undefined,
                method: 'query_by_text',
            },
            requestTypeNamespace: {
                service: 'auth',
                entity: 'secret',
                types: ['query_by_text_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'auth',
                entity: 'secret',
                types: ['list_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'GET',
                    path: '/query_by_text',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'user',
                entity: 'user',
                types: undefined,
                enum: undefined,
                method: 'query_by_text',
            },
            requestTypeNamespace: {
                service: 'user',
                entity: 'user',
                types: ['query_by_text_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'user',
                entity: 'user',
                types: ['list_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'GET',
                    path: '/query_by_text',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'core',
                entity: 'file',
                types: undefined,
                enum: undefined,
                method: 'query_by_text',
            },
            requestTypeNamespace: {
                service: 'core',
                entity: 'file',
                types: ['query_by_text_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'core',
                entity: 'file',
                types: ['list_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'GET',
                    path: '/query_by_text',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'core',
                entity: 'task_run',
                types: undefined,
                enum: undefined,
                method: 'query_by_text',
            },
            requestTypeNamespace: {
                service: 'core',
                entity: 'task_run',
                types: ['query_by_text_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'core',
                entity: 'task_run',
                types: ['list_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'GET',
                    path: '/query_by_text',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'core',
                entity: 'task',
                types: undefined,
                enum: undefined,
                method: 'query_by_text',
            },
            requestTypeNamespace: {
                service: 'core',
                entity: 'task',
                types: ['query_by_text_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'core',
                entity: 'task',
                types: ['list_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'GET',
                    path: '/query_by_text',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'user',
                entity: 'team',
                types: undefined,
                enum: undefined,
                method: 'query_by_text',
            },
            requestTypeNamespace: {
                service: 'user',
                entity: 'team',
                types: ['query_by_text_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'user',
                entity: 'team',
                types: ['list_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'GET',
                    path: '/query_by_text',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'auth',
                entity: undefined,
                types: undefined,
                enum: undefined,
                method: 'register_user',
            },
            requestTypeNamespace: {
                service: 'auth',
                entity: undefined,
                types: ['register_user_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'auth',
                entity: undefined,
                types: ['authenticate_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'POST',
                    path: '/register',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'core',
                entity: 'task_run',
                types: undefined,
                enum: undefined,
                method: 'update',
            },
            requestTypeNamespace: {
                service: 'core',
                entity: 'task_run',
                types: ['update_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'core',
                entity: 'task_run',
                types: ['update_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'PUT',
                    path: '',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'user',
                entity: 'team',
                types: undefined,
                enum: undefined,
                method: 'update',
            },
            requestTypeNamespace: {
                service: 'user',
                entity: 'team',
                types: ['update_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'user',
                entity: 'team',
                types: ['update_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'PUT',
                    path: '',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'user',
                entity: 'user',
                types: undefined,
                enum: undefined,
                method: 'update',
            },
            requestTypeNamespace: {
                service: 'user',
                entity: 'user',
                types: ['update_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'user',
                entity: 'user',
                types: ['update_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'PUT',
                    path: '',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'core',
                entity: 'task',
                types: undefined,
                enum: undefined,
                method: 'update',
            },
            requestTypeNamespace: {
                service: 'core',
                entity: 'task',
                types: ['update_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'core',
                entity: 'task',
                types: ['update_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'PUT',
                    path: '',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'core',
                entity: 'file',
                types: undefined,
                enum: undefined,
                method: 'update',
            },
            requestTypeNamespace: {
                service: 'core',
                entity: 'file',
                types: ['update_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'core',
                entity: 'file',
                types: ['update_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'PUT',
                    path: '',
                    version: 1, // hard coded for now
                },
            ],
        },

        {
            namespace: {
                service: 'auth',
                entity: 'secret',
                types: undefined,
                enum: undefined,
                method: 'update',
            },
            requestTypeNamespace: {
                service: 'auth',
                entity: 'secret',
                types: ['update_request'],
                enum: undefined,
                method: undefined,
            },
            responseTypeNamespace: {
                service: 'auth',
                entity: 'secret',
                types: ['update_response'],
                enum: undefined,
                method: undefined,
            },
            apis: [
                {
                    method: 'PUT',
                    path: '',
                    version: 1, // hard coded for now
                },
            ],
        },
    ],
}

const typeReqs: TypeInfoReq[] = [
    {
        namespace: {
            service: undefined,
            entity: undefined,
            types: ['address'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('line_1'),
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
                name: new Name('line_2'),
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
                name: new Name('city'),
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
                name: new Name('state'),
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
                name: new Name('postal_code'),
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
                name: new Name('country'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('country'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: undefined,
                        entity: undefined,
                        types: undefined,
                        enum: 'country',
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances

            // Return the empty object
            return {
                line1: undefined,
                line2: undefined,
                city: undefined,
                state: undefined,
                postalCode: undefined,
                country: undefined,
            }
        },
    },
    {
        namespace: {
            service: undefined,
            entity: undefined,
            types: ['address_field_condition'],
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
            service: undefined,
            entity: undefined,
            types: ['address_filter'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('parent_id'),
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
                name: new Name('line_1'),
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
                name: new Name('line_2'),
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
                name: new Name('city'),
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
                name: new Name('state'),
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
                name: new Name('postal_code'),
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
                name: new Name('country'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('country_condition'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: undefined,
                        entity: undefined,
                        types: ['country_condition'],
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
                    name: new Name('address_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: undefined,
                        entity: undefined,
                        types: ['address_filter'],
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
                    name: new Name('address_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: undefined,
                        entity: undefined,
                        types: ['address_filter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const countryTypeInfo = appInfo.getTypeInfo({
                service: undefined,
                entity: undefined,
                types: ['country_condition'],
                enum: undefined,
                method: undefined,
            })
            const andTypeInfo = appInfo.getTypeInfo({
                service: undefined,
                entity: undefined,
                types: ['address_filter'],
                enum: undefined,
                method: undefined,
            })
            const orTypeInfo = appInfo.getTypeInfo({
                service: undefined,
                entity: undefined,
                types: ['address_filter'],
                enum: undefined,
                method: undefined,
            })

            // Return the empty object
            return {
                parentID: undefined,
                id: undefined,
                line1: undefined,
                line2: undefined,
                city: undefined,
                state: undefined,
                postalCode: undefined,
                country: countryTypeInfo.getEmptyObject(appInfo),
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
            service: undefined,
            entity: undefined,
            types: ['contact'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('name'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('person_name'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: undefined,
                        entity: undefined,
                        types: ['person_name'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('email'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('email'),
                    kind: fieldkind.EmailKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('address'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('address'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: undefined,
                        entity: undefined,
                        types: ['address'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const nameTypeInfo = appInfo.getTypeInfo({
                service: undefined,
                entity: undefined,
                types: ['person_name'],
                enum: undefined,
                method: undefined,
            })
            const addressTypeInfo = appInfo.getTypeInfo({
                service: undefined,
                entity: undefined,
                types: ['address'],
                enum: undefined,
                method: undefined,
            })

            // Return the empty object
            return {
                name: nameTypeInfo.getEmptyObject(appInfo),
                email: undefined,
                address: addressTypeInfo.getEmptyObject(appInfo),
            }
        },
    },
    {
        namespace: {
            service: undefined,
            entity: undefined,
            types: ['contact_field_condition'],
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
            service: undefined,
            entity: undefined,
            types: ['contact_filter'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('parent_id'),
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
                name: new Name('name'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('person_name_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: undefined,
                        entity: undefined,
                        types: ['person_name_filter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('email'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('email_condition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('address'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('address_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: undefined,
                        entity: undefined,
                        types: ['address_filter'],
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
                    name: new Name('contact_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: undefined,
                        entity: undefined,
                        types: ['contact_filter'],
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
                    name: new Name('contact_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: undefined,
                        entity: undefined,
                        types: ['contact_filter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const nameTypeInfo = appInfo.getTypeInfo({
                service: undefined,
                entity: undefined,
                types: ['person_name_filter'],
                enum: undefined,
                method: undefined,
            })
            const addressTypeInfo = appInfo.getTypeInfo({
                service: undefined,
                entity: undefined,
                types: ['address_filter'],
                enum: undefined,
                method: undefined,
            })
            const andTypeInfo = appInfo.getTypeInfo({
                service: undefined,
                entity: undefined,
                types: ['contact_filter'],
                enum: undefined,
                method: undefined,
            })
            const orTypeInfo = appInfo.getTypeInfo({
                service: undefined,
                entity: undefined,
                types: ['contact_filter'],
                enum: undefined,
                method: undefined,
            })

            // Return the empty object
            return {
                parentID: undefined,
                id: undefined,
                name: nameTypeInfo.getEmptyObject(appInfo),
                email: undefined,
                address: addressTypeInfo.getEmptyObject(appInfo),
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
            service: undefined,
            entity: undefined,
            types: ['country_condition'],
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
            service: undefined,
            entity: undefined,
            types: ['entity_name_condition'],
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
            service: undefined,
            entity: undefined,
            types: ['gender_condition'],
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
            service: undefined,
            entity: undefined,
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
            service: undefined,
            entity: undefined,
            types: ['person_name'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('first_name'),
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
                name: new Name('middle_initial'),
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
                name: new Name('last_name'),
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
                firstName: undefined,
                middleInitial: undefined,
                lastName: undefined,
            }
        },
    },
    {
        namespace: {
            service: undefined,
            entity: undefined,
            types: ['person_name_field_condition'],
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
            service: undefined,
            entity: undefined,
            types: ['person_name_filter'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('parent_id'),
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
                name: new Name('first_name'),
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
                name: new Name('middle_initial'),
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
                name: new Name('last_name'),
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
                    name: new Name('person_name_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: undefined,
                        entity: undefined,
                        types: ['person_name_filter'],
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
                    name: new Name('person_name_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: undefined,
                        entity: undefined,
                        types: ['person_name_filter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const andTypeInfo = appInfo.getTypeInfo({
                service: undefined,
                entity: undefined,
                types: ['person_name_filter'],
                enum: undefined,
                method: undefined,
            })
            const orTypeInfo = appInfo.getTypeInfo({
                service: undefined,
                entity: undefined,
                types: ['person_name_filter'],
                enum: undefined,
                method: undefined,
            })

            // Return the empty object
            return {
                parentID: undefined,
                id: undefined,
                firstName: undefined,
                middleInitial: undefined,
                lastName: undefined,
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
            service: undefined,
            entity: undefined,
            types: ['phone_number'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('country_code'),
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
            {
                name: new Name('number'),
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
                name: new Name('extension'),
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
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances

            // Return the empty object
            return {
                countryCode: undefined,
                number: undefined,
                extension: undefined,
            }
        },
    },
    {
        namespace: {
            service: undefined,
            entity: undefined,
            types: ['phone_number_field_condition'],
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
            service: undefined,
            entity: undefined,
            types: ['phone_number_filter'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('parent_id'),
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
                name: new Name('country_code'),
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
                name: new Name('number'),
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
                name: new Name('extension'),
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
                    name: new Name('phone_number_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: undefined,
                        entity: undefined,
                        types: ['phone_number_filter'],
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
                    name: new Name('phone_number_filter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: undefined,
                        entity: undefined,
                        types: ['phone_number_filter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances
            const andTypeInfo = appInfo.getTypeInfo({
                service: undefined,
                entity: undefined,
                types: ['phone_number_filter'],
                enum: undefined,
                method: undefined,
            })
            const orTypeInfo = appInfo.getTypeInfo({
                service: undefined,
                entity: undefined,
                types: ['phone_number_filter'],
                enum: undefined,
                method: undefined,
            })

            // Return the empty object
            return {
                parentID: undefined,
                id: undefined,
                countryCode: undefined,
                number: undefined,
                extension: undefined,
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
            service: undefined,
            entity: undefined,
            types: ['service_name_condition'],
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
]

const enumReqs: EnumReq[] = [
    {
        namespace: {
            service: undefined,
            entity: undefined,
            types: undefined,
            enum: 'address_field',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'ParentID',
                description: undefined,
                displayValue: 'Parent ID',
            },
            {
                id: 2,
                value: 'ID',
                description: undefined,
                displayValue: 'ID',
            },
            {
                id: 3,
                value: 'Line1',
                description: undefined,
                displayValue: 'Line 1',
            },
            {
                id: 4,
                value: 'Line2',
                description: undefined,
                displayValue: 'Line 2',
            },
            {
                id: 5,
                value: 'City',
                description: undefined,
                displayValue: 'City',
            },
            {
                id: 6,
                value: 'State',
                description: undefined,
                displayValue: 'State',
            },
            {
                id: 7,
                value: 'PostalCode',
                description: undefined,
                displayValue: 'Postal Code',
            },
            {
                id: 8,
                value: 'Country',
                description: undefined,
                displayValue: 'Country',
            },
            {
                id: 9,
                value: 'CreatedAt',
                description: undefined,
                displayValue: 'Created At',
            },
            {
                id: 10,
                value: 'UpdatedAt',
                description: undefined,
                displayValue: 'Updated At',
            },
            {
                id: 11,
                value: 'DeletedAt',
                description: undefined,
                displayValue: 'Deleted At',
            },
        ],
    },
    {
        namespace: {
            service: undefined,
            entity: undefined,
            types: undefined,
            enum: 'contact_field',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'ParentID',
                description: undefined,
                displayValue: 'Parent ID',
            },
            {
                id: 2,
                value: 'ID',
                description: undefined,
                displayValue: 'ID',
            },
            {
                id: 3,
                value: 'Name',
                description: undefined,
                displayValue: 'Name',
            },
            {
                id: 4,
                value: 'Name_ParentID',
                description: undefined,
                displayValue: 'Name -  Parent ID',
            },
            {
                id: 5,
                value: 'Name_ID',
                description: undefined,
                displayValue: 'Name -  ID',
            },
            {
                id: 6,
                value: 'Name_FirstName',
                description: undefined,
                displayValue: 'Name -  First Name',
            },
            {
                id: 7,
                value: 'Name_MiddleInitial',
                description: undefined,
                displayValue: 'Name -  Middle Initial',
            },
            {
                id: 8,
                value: 'Name_LastName',
                description: undefined,
                displayValue: 'Name -  Last Name',
            },
            {
                id: 9,
                value: 'Name_CreatedAt',
                description: undefined,
                displayValue: 'Name -  Created At',
            },
            {
                id: 10,
                value: 'Name_UpdatedAt',
                description: undefined,
                displayValue: 'Name -  Updated At',
            },
            {
                id: 11,
                value: 'Name_DeletedAt',
                description: undefined,
                displayValue: 'Name -  Deleted At',
            },
            {
                id: 12,
                value: 'Email',
                description: undefined,
                displayValue: 'Email',
            },
            {
                id: 13,
                value: 'Address',
                description: undefined,
                displayValue: 'Address',
            },
            {
                id: 14,
                value: 'Address_ParentID',
                description: undefined,
                displayValue: 'Address -  Parent ID',
            },
            {
                id: 15,
                value: 'Address_ID',
                description: undefined,
                displayValue: 'Address -  ID',
            },
            {
                id: 16,
                value: 'Address_Line1',
                description: undefined,
                displayValue: 'Address -  Line 1',
            },
            {
                id: 17,
                value: 'Address_Line2',
                description: undefined,
                displayValue: 'Address -  Line 2',
            },
            {
                id: 18,
                value: 'Address_City',
                description: undefined,
                displayValue: 'Address -  City',
            },
            {
                id: 19,
                value: 'Address_State',
                description: undefined,
                displayValue: 'Address -  State',
            },
            {
                id: 20,
                value: 'Address_PostalCode',
                description: undefined,
                displayValue: 'Address -  Postal Code',
            },
            {
                id: 21,
                value: 'Address_Country',
                description: undefined,
                displayValue: 'Address -  Country',
            },
            {
                id: 22,
                value: 'Address_CreatedAt',
                description: undefined,
                displayValue: 'Address -  Created At',
            },
            {
                id: 23,
                value: 'Address_UpdatedAt',
                description: undefined,
                displayValue: 'Address -  Updated At',
            },
            {
                id: 24,
                value: 'Address_DeletedAt',
                description: undefined,
                displayValue: 'Address -  Deleted At',
            },
            {
                id: 25,
                value: 'CreatedAt',
                description: undefined,
                displayValue: 'Created At',
            },
            {
                id: 26,
                value: 'UpdatedAt',
                description: undefined,
                displayValue: 'Updated At',
            },
            {
                id: 27,
                value: 'DeletedAt',
                description: undefined,
                displayValue: 'Deleted At',
            },
        ],
    },
    {
        namespace: {
            service: undefined,
            entity: undefined,
            types: undefined,
            enum: 'country',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'USA',
                description: undefined,
                displayValue: 'USA',
            },
            {
                id: 2,
                value: 'Canada',
                description: undefined,
                displayValue: 'Canada',
            },
            {
                id: 3,
                value: 'Mexico',
                description: undefined,
                displayValue: 'Mexico',
            },
        ],
    },
    {
        namespace: {
            service: undefined,
            entity: undefined,
            types: undefined,
            enum: 'entity_name',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'Auth_Secret',
                description: undefined,
                displayValue: 'Auth -  Secret',
            },
            {
                id: 2,
                value: 'Core_File',
                description: undefined,
                displayValue: 'Core -  File',
            },
            {
                id: 3,
                value: 'Core_Task',
                description: undefined,
                displayValue: 'Core -  Task',
            },
            {
                id: 4,
                value: 'Core_TaskRun',
                description: undefined,
                displayValue: 'Core -  Task Run',
            },
            {
                id: 5,
                value: 'User_User',
                description: undefined,
                displayValue: 'User -  User',
            },
            {
                id: 6,
                value: 'User_Team',
                description: undefined,
                displayValue: 'User -  Team',
            },
        ],
    },
    {
        namespace: {
            service: undefined,
            entity: undefined,
            types: undefined,
            enum: 'gender',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'Male',
                description: undefined,
                displayValue: 'Male',
            },
            {
                id: 2,
                value: 'Female',
                description: undefined,
                displayValue: 'Female',
            },
            {
                id: 3,
                value: 'Other',
                description: undefined,
                displayValue: 'Other',
            },
        ],
    },
    {
        namespace: {
            service: undefined,
            entity: undefined,
            types: undefined,
            enum: 'method_name',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'Auth_LoginUser',
                description: undefined,
                displayValue: 'Auth -  Login User',
            },
            {
                id: 2,
                value: 'Auth_RegisterUser',
                description: undefined,
                displayValue: 'Auth -  Register User',
            },
            {
                id: 3,
                value: 'Auth_AuthenticateToken',
                description: undefined,
                displayValue: 'Auth -  Authenticate Token',
            },
            {
                id: 4,
                value: 'Core_FileUpload',
                description: undefined,
                displayValue: 'Core -  File Upload',
            },
            {
                id: 5,
                value: 'Auth_Secret_Add',
                description: undefined,
                displayValue: 'Auth -  Secret -  Add',
            },
            {
                id: 6,
                value: 'Auth_Secret_Update',
                description: undefined,
                displayValue: 'Auth -  Secret -  Update',
            },
            {
                id: 7,
                value: 'Auth_Secret_Get',
                description: undefined,
                displayValue: 'Auth -  Secret -  Get',
            },
            {
                id: 8,
                value: 'Auth_Secret_List',
                description: undefined,
                displayValue: 'Auth -  Secret -  List',
            },
            {
                id: 9,
                value: 'Auth_Secret_QueryByText',
                description: undefined,
                displayValue: 'Auth -  Secret -  Query By Text',
            },
            {
                id: 10,
                value: 'Core_File_Add',
                description: undefined,
                displayValue: 'Core -  File -  Add',
            },
            {
                id: 11,
                value: 'Core_File_Update',
                description: undefined,
                displayValue: 'Core -  File -  Update',
            },
            {
                id: 12,
                value: 'Core_File_Get',
                description: undefined,
                displayValue: 'Core -  File -  Get',
            },
            {
                id: 13,
                value: 'Core_File_List',
                description: undefined,
                displayValue: 'Core -  File -  List',
            },
            {
                id: 14,
                value: 'Core_File_QueryByText',
                description: undefined,
                displayValue: 'Core -  File -  Query By Text',
            },
            {
                id: 15,
                value: 'Core_Task_Add',
                description: undefined,
                displayValue: 'Core -  Task -  Add',
            },
            {
                id: 16,
                value: 'Core_Task_Update',
                description: undefined,
                displayValue: 'Core -  Task -  Update',
            },
            {
                id: 17,
                value: 'Core_Task_Get',
                description: undefined,
                displayValue: 'Core -  Task -  Get',
            },
            {
                id: 18,
                value: 'Core_Task_List',
                description: undefined,
                displayValue: 'Core -  Task -  List',
            },
            {
                id: 19,
                value: 'Core_Task_QueryByText',
                description: undefined,
                displayValue: 'Core -  Task -  Query By Text',
            },
            {
                id: 20,
                value: 'Core_TaskRun_Add',
                description: undefined,
                displayValue: 'Core -  Task Run -  Add',
            },
            {
                id: 21,
                value: 'Core_TaskRun_Update',
                description: undefined,
                displayValue: 'Core -  Task Run -  Update',
            },
            {
                id: 22,
                value: 'Core_TaskRun_Get',
                description: undefined,
                displayValue: 'Core -  Task Run -  Get',
            },
            {
                id: 23,
                value: 'Core_TaskRun_List',
                description: undefined,
                displayValue: 'Core -  Task Run -  List',
            },
            {
                id: 24,
                value: 'Core_TaskRun_QueryByText',
                description: undefined,
                displayValue: 'Core -  Task Run -  Query By Text',
            },
            {
                id: 25,
                value: 'User_User_Add',
                description: undefined,
                displayValue: 'User -  User -  Add',
            },
            {
                id: 26,
                value: 'User_User_Update',
                description: undefined,
                displayValue: 'User -  User -  Update',
            },
            {
                id: 27,
                value: 'User_User_Get',
                description: undefined,
                displayValue: 'User -  User -  Get',
            },
            {
                id: 28,
                value: 'User_User_List',
                description: undefined,
                displayValue: 'User -  User -  List',
            },
            {
                id: 29,
                value: 'User_User_QueryByText',
                description: undefined,
                displayValue: 'User -  User -  Query By Text',
            },
            {
                id: 30,
                value: 'User_Team_Add',
                description: undefined,
                displayValue: 'User -  Team -  Add',
            },
            {
                id: 31,
                value: 'User_Team_Update',
                description: undefined,
                displayValue: 'User -  Team -  Update',
            },
            {
                id: 32,
                value: 'User_Team_Get',
                description: undefined,
                displayValue: 'User -  Team -  Get',
            },
            {
                id: 33,
                value: 'User_Team_List',
                description: undefined,
                displayValue: 'User -  Team -  List',
            },
            {
                id: 34,
                value: 'User_Team_QueryByText',
                description: undefined,
                displayValue: 'User -  Team -  Query By Text',
            },
            {
                id: 35,
                value: 'Core_File_HookCreatePre',
                description: undefined,
                displayValue: 'Core -  File -  Hook Create Pre',
            },
            {
                id: 36,
                value: 'Core_Task_HookSavePre',
                description: undefined,
                displayValue: 'Core -  Task -  Hook Save Pre',
            },
            {
                id: 37,
                value: 'Core_TaskRun_HookCreatePre',
                description: undefined,
                displayValue: 'Core -  Task Run -  Hook Create Pre',
            },
            {
                id: 38,
                value: 'User_User_HookInit',
                description: undefined,
                displayValue: 'User -  User -  Hook Init',
            },
            {
                id: 39,
                value: 'Core_Task_ActionRun',
                description: undefined,
                displayValue: 'Core -  Task -  Action Run',
            },
            {
                id: 40,
                value: 'Core_TaskRun_ActionRun',
                description: undefined,
                displayValue: 'Core -  Task Run -  Action Run',
            },
        ],
    },
    {
        namespace: {
            service: undefined,
            entity: undefined,
            types: undefined,
            enum: 'person_name_field',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'ParentID',
                description: undefined,
                displayValue: 'Parent ID',
            },
            {
                id: 2,
                value: 'ID',
                description: undefined,
                displayValue: 'ID',
            },
            {
                id: 3,
                value: 'FirstName',
                description: undefined,
                displayValue: 'First Name',
            },
            {
                id: 4,
                value: 'MiddleInitial',
                description: undefined,
                displayValue: 'Middle Initial',
            },
            {
                id: 5,
                value: 'LastName',
                description: undefined,
                displayValue: 'Last Name',
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
            service: undefined,
            entity: undefined,
            types: undefined,
            enum: 'phone_number_field',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'ParentID',
                description: undefined,
                displayValue: 'Parent ID',
            },
            {
                id: 2,
                value: 'ID',
                description: undefined,
                displayValue: 'ID',
            },
            {
                id: 3,
                value: 'CountryCode',
                description: undefined,
                displayValue: 'Country Code',
            },
            {
                id: 4,
                value: 'Number',
                description: undefined,
                displayValue: 'Number',
            },
            {
                id: 5,
                value: 'Extension',
                description: undefined,
                displayValue: 'Extension',
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
            service: undefined,
            entity: undefined,
            types: undefined,
            enum: 'service_name',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'Auth',
                description: undefined,
                displayValue: 'Auth',
            },
            {
                id: 2,
                value: 'Core',
                description: undefined,
                displayValue: 'Core',
            },
            {
                id: 3,
                value: 'User',
                description: undefined,
                displayValue: 'User',
            },
        ],
    },
]
