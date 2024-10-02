/*eslint-disable */

import * as fieldkind from '@ongoku/app-lib/src/common/fieldkind'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { AppReq, ServiceReq, EntityInfoReq, TypeInfoReq, EnumReq, MethodReq } from '@ongoku/app-lib/src/common/app_v3'
import { Name, ITypeAppNamespace, IEnumAppNamespace, ITypeServiceNamespace, IEnumServiceNamespace, ITypeEntityNamespace, IEnumEntityNamespace } from '@ongoku/app-lib/src/common/namespacev2'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BooleanCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

const entityInfoReq: EntityInfoReq = {
    namespace: {
        service: 'User',
        entity: 'Team',
    },
    associations: [],
    actions: [],
}

const typeReqs: TypeReq[] = [
    {
        namespace: {
            service: 'User',
            entity: 'Team',
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
                    name: new Name('User_Team'),
                    kind: fieldkind.ForeignEntityKind,
                    namespace: {
                        service: 'User',
                        entity: 'Team',
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'User',
            entity: 'Team',
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
                    name: new Name('User_Team'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'User',
                        entity: 'Team',
                        types: ['Team'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'User',
            entity: 'Team',
            types: ['Team'],
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
                name: new Name('name'),
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
            service: 'User',
            entity: 'Team',
            types: ['TeamFieldCondition'],
            enum: undefined,
            method: undefined,
        },
        fields: [],
    },
    {
        namespace: {
            service: 'User',
            entity: 'Team',
            types: ['TeamFilter'],
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
                name: new Name('name'),
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
                    name: new Name('User_Team_TeamFilter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'User',
                        entity: 'Team',
                        types: ['TeamFilter'],
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
                    name: new Name('User_Team_TeamFilter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'User',
                        entity: 'Team',
                        types: ['TeamFilter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'User',
            entity: 'Team',
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
                    name: new Name('User_Team'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'User',
                        entity: 'Team',
                        types: ['Team'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'User',
            entity: 'Team',
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
                    name: new Name('User_Team'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'User',
                        entity: 'Team',
                        types: ['Team'],
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
                    name: new Name('User_Team_TeamField'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'User',
                        entity: 'Team',
                        types: undefined,
                        enum: 'TeamField',
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
                    name: new Name('User_Team_TeamField'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'User',
                        entity: 'Team',
                        types: undefined,
                        enum: 'TeamField',
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'User',
            entity: 'Team',
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
                    name: new Name('User_Team'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'User',
                        entity: 'Team',
                        types: ['Team'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'User',
            entity: 'Team',
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
            service: 'User',
            entity: 'Team',
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
                    name: new Name('User_Team_TeamFilter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'User',
                        entity: 'Team',
                        types: ['TeamFilter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'User',
            entity: 'Team',
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
                    name: new Name('User_Team'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'User',
                        entity: 'Team',
                        types: ['Team'],
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
            service: 'User',
            entity: 'Team',
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
            service: 'User',
            entity: 'Team',
            types: undefined,
            enum: 'TeamField',
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
                value: 'Name',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 3,
                value: 'CreatedAt',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 4,
                value: 'UpdatedAt',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 5,
                value: 'DeletedAt',
                description: undefined,
                displayValue: undefined,
            },
        ],
    },
]
