/*eslint-disable */

import * as fieldkind from '@ongoku/app-lib/src/common/fieldkind'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { App, AppReq, ServiceReq, EntityInfoReq, TypeInfoReq, EnumReq, MethodReq } from '@ongoku/app-lib/src/common/app_v3'
import { Name, ITypeAppNamespace, IEnumAppNamespace, ITypeServiceNamespace, IEnumServiceNamespace, ITypeEntityNamespace, IEnumEntityNamespace } from '@ongoku/app-lib/src/common/namespacev2'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BoolCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

const entityInfoReq: EntityInfoReq = {
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
}

const typeReqs: TypeReq[] = [
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
]

const enumReqs: EnumReq[] = [
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
]
