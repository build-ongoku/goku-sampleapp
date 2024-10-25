/*eslint-disable */

import * as fieldkind from '@ongoku/app-lib/src/common/fieldkind'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { App, AppReq, ServiceReq, EntityInfoReq, TypeInfoReq, EnumReq, MethodReq } from '@ongoku/app-lib/src/common/app_v3'
import { Name, ITypeAppNamespace, IEnumAppNamespace, ITypeServiceNamespace, IEnumServiceNamespace, ITypeEntityNamespace, IEnumEntityNamespace } from '@ongoku/app-lib/src/common/namespacev2'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BoolCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

const entityInfoReq: EntityInfoReq = {
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
}

const typeReqs: TypeReq[] = [
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
]

const enumReqs: EnumReq[] = [
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
]
