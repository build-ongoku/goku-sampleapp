/*eslint-disable */

import * as fieldkind from '@ongoku/app-lib/src/common/fieldkind'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { AppReq, ServiceReq, EntityInfoReq, TypeInfoReq, EnumReq, MethodReq } from '@ongoku/app-lib/src/common/app_v3'
import { Name, ITypeAppNamespace, IEnumAppNamespace, ITypeServiceNamespace, IEnumServiceNamespace, ITypeEntityNamespace, IEnumEntityNamespace } from '@ongoku/app-lib/src/common/namespacev2'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BooleanCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

const entityInfoReq: EntityInfoReq = {
    namespace: {
        service: 'User',
        entity: 'User',
    },
    associations: [],
    actions: [],
}

const typeReqs: TypeReq[] = [
    {
        namespace: {
            service: 'User',
            entity: 'User',
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
                    name: new Name('User_User'),
                    kind: fieldkind.ForeignEntityKind,
                    namespace: {
                        service: 'User',
                        entity: 'User',
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'User',
            entity: 'User',
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
                    name: new Name('User_User'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'User',
                        entity: 'User',
                        types: ['User'],
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
            entity: 'User',
            types: ['User'],
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
                    name: new Name('PersonNamemetaFields'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: undefined,
                        entity: undefined,
                        types: ['PersonName'],
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
                    name: new Name('Email'),
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
                    name: new Name('AddressmetaFields'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: undefined,
                        entity: undefined,
                        types: ['Address'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('teamID'),
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
            {
                name: new Name('pastTeamIDs'),
                isRepeated: true,
                isOptional: true,
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
            entity: 'User',
            types: ['UserFieldCondition'],
            enum: undefined,
            method: undefined,
        },
        fields: [],
    },
    {
        namespace: {
            service: 'User',
            entity: 'User',
            types: ['UserFilter'],
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
                    name: new Name('PersonNameFilter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: undefined,
                        entity: undefined,
                        types: ['PersonNameFilter'],
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
                    name: new Name('EmailCondition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('havingAddresses'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('AddressFilter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: undefined,
                        entity: undefined,
                        types: ['AddressFilter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('teamID'),
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
                name: new Name('havingPastTeamIDs'),
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
                    name: new Name('User_User_UserFilter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'User',
                        entity: 'User',
                        types: ['UserFilter'],
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
                    name: new Name('User_User_UserFilter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'User',
                        entity: 'User',
                        types: ['UserFilter'],
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
            entity: 'User',
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
                    name: new Name('User_User'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'User',
                        entity: 'User',
                        types: ['User'],
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
            entity: 'User',
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
                    name: new Name('User_User'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'User',
                        entity: 'User',
                        types: ['User'],
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
                    name: new Name('User_User_UserField'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'User',
                        entity: 'User',
                        types: undefined,
                        enum: 'UserField',
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
                    name: new Name('User_User_UserField'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'User',
                        entity: 'User',
                        types: undefined,
                        enum: 'UserField',
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'User',
            entity: 'User',
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
                    name: new Name('User_User'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'User',
                        entity: 'User',
                        types: ['User'],
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
            entity: 'User',
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
            entity: 'User',
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
                    name: new Name('User_User_UserFilter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'User',
                        entity: 'User',
                        types: ['UserFilter'],
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
            entity: 'User',
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
                    name: new Name('User_User'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'User',
                        entity: 'User',
                        types: ['User'],
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
            entity: 'User',
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
            entity: 'User',
            types: undefined,
            enum: 'UserField',
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
                value: 'Name_ParentID',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 4,
                value: 'Name_ID',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 5,
                value: 'Name_FirstName',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 6,
                value: 'Name_MiddleInitial',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 7,
                value: 'Name_LastName',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 8,
                value: 'Name_CreatedAt',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 9,
                value: 'Name_UpdatedAt',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 10,
                value: 'Name_DeletedAt',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 11,
                value: 'Email',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 12,
                value: 'Addresses',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 13,
                value: 'Addresses_ParentID',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 14,
                value: 'Addresses_ID',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 15,
                value: 'Addresses_Line1',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 16,
                value: 'Addresses_Line2',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 17,
                value: 'Addresses_City',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 18,
                value: 'Addresses_State',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 19,
                value: 'Addresses_PostalCode',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 20,
                value: 'Addresses_Country',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 21,
                value: 'Addresses_CreatedAt',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 22,
                value: 'Addresses_UpdatedAt',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 23,
                value: 'Addresses_DeletedAt',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 24,
                value: 'TeamID',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 25,
                value: 'PastTeamIDs',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 26,
                value: 'CreatedAt',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 27,
                value: 'UpdatedAt',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 28,
                value: 'DeletedAt',
                description: undefined,
                displayValue: undefined,
            },
        ],
    },
]
