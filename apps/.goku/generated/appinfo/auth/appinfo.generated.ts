/*eslint-disable */

import * as fieldkind from '@ongoku/app-lib/src/common/fieldkind'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { App, AppReq, ServiceReq, EntityInfoReq, TypeInfoReq, EnumReq, MethodReq } from '@ongoku/app-lib/src/common/app_v3'
import { Name, ITypeAppNamespace, IEnumAppNamespace, ITypeServiceNamespace, IEnumServiceNamespace, ITypeEntityNamespace, IEnumEntityNamespace } from '@ongoku/app-lib/src/common/namespacev2'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BoolCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

const serviceReq: ServiceReq = {
    namespace: {
        service: 'auth',
        entity: undefined,
    },
}

const typeReqs: TypeReq[] = [
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
]

const enumReqs: EnumReq[] = [
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
]
