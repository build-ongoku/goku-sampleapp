/*eslint-disable */

import * as fieldkind from '@ongoku/app-lib/src/common/fieldkind'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { AppReq, ServiceReq, EntityInfoReq, TypeInfoReq, EnumReq, MethodReq } from '@ongoku/app-lib/src/common/app_v3'
import { Name, ITypeAppNamespace, IEnumAppNamespace, ITypeServiceNamespace, IEnumServiceNamespace, ITypeEntityNamespace, IEnumEntityNamespace } from '@ongoku/app-lib/src/common/namespacev2'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BooleanCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

const entityInfoReq: EntityInfoReq = {
    namespace: {
        service: 'Auth',
        entity: 'Secret',
    },
    associations: [],
    actions: [],
}

const typeReqs: TypeReq[] = [
    {
        namespace: {
            service: 'Auth',
            entity: 'Secret',
            types: ['Secret'],
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
                name: new Name('userID'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('User_User'),
                    kind: fieldkind.ForeignEntityKind,
                    namespace: {
                        service: 'User',
                        entity: 'User',
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
                    name: new Name('Auth_Secret_Type'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'Auth',
                        entity: 'Secret',
                        types: undefined,
                        enum: 'Type',
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
                    name: new Name('String'),
                    kind: fieldkind.StringKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('expiresAt'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Timestamp'),
                    kind: fieldkind.TimestampKind,
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
            service: 'Auth',
            entity: 'Secret',
            types: ['SecretFieldCondition'],
            enum: undefined,
            method: undefined,
        },
        fields: [],
    },
    {
        namespace: {
            service: 'Auth',
            entity: 'Secret',
            types: ['SecretFilter'],
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
                name: new Name('userID'),
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
                name: new Name('type'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('TypeCondition'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Auth',
                        entity: 'Secret',
                        types: ['TypeCondition'],
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
                    name: new Name('StringCondition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('expiresAt'),
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
                    name: new Name('Auth_Secret_SecretFilter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Auth',
                        entity: 'Secret',
                        types: ['SecretFilter'],
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
                    name: new Name('Auth_Secret_SecretFilter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Auth',
                        entity: 'Secret',
                        types: ['SecretFilter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Auth',
            entity: 'Secret',
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
                    name: new Name('Auth_Secret'),
                    kind: fieldkind.ForeignEntityKind,
                    namespace: {
                        service: 'Auth',
                        entity: 'Secret',
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Auth',
            entity: 'Secret',
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
                    name: new Name('Auth_Secret'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Auth',
                        entity: 'Secret',
                        types: ['Secret'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Auth',
            entity: 'Secret',
            types: ['TypeCondition'],
            enum: undefined,
            method: undefined,
        },
        fields: [],
    },
    {
        namespace: {
            service: 'Auth',
            entity: 'Secret',
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
                    name: new Name('Auth_Secret'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Auth',
                        entity: 'Secret',
                        types: ['Secret'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Auth',
            entity: 'Secret',
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
                    name: new Name('Auth_Secret'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Auth',
                        entity: 'Secret',
                        types: ['Secret'],
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
                    name: new Name('Auth_Secret_SecretField'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'Auth',
                        entity: 'Secret',
                        types: undefined,
                        enum: 'SecretField',
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
                    name: new Name('Auth_Secret_SecretField'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'Auth',
                        entity: 'Secret',
                        types: undefined,
                        enum: 'SecretField',
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Auth',
            entity: 'Secret',
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
                    name: new Name('Auth_Secret'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Auth',
                        entity: 'Secret',
                        types: ['Secret'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Auth',
            entity: 'Secret',
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
            service: 'Auth',
            entity: 'Secret',
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
                    name: new Name('Auth_Secret_SecretFilter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Auth',
                        entity: 'Secret',
                        types: ['SecretFilter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Auth',
            entity: 'Secret',
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
                    name: new Name('Auth_Secret'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Auth',
                        entity: 'Secret',
                        types: ['Secret'],
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
            service: 'Auth',
            entity: 'Secret',
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
            service: 'Auth',
            entity: 'Secret',
            types: undefined,
            enum: 'SecretField',
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
                value: 'UserID',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 3,
                value: 'Type',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 4,
                value: 'Secret',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 5,
                value: 'ExpiresAt',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 6,
                value: 'CreatedAt',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 7,
                value: 'UpdatedAt',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 8,
                value: 'DeletedAt',
                description: undefined,
                displayValue: undefined,
            },
        ],
    },
    {
        namespace: {
            service: 'Auth',
            entity: 'Secret',
            types: undefined,
            enum: 'Type',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'Password',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 2,
                value: 'GithubToken',
                description: undefined,
                displayValue: undefined,
            },
        ],
    },
]
