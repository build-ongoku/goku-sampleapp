/*eslint-disable */

import * as fieldkind from '@ongoku/app-lib/src/common/fieldkind'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { App, AppReq, ServiceReq, EntityInfoReq, TypeInfoReq, EnumReq, MethodReq } from '@ongoku/app-lib/src/common/app_v3'
import { Name, ITypeAppNamespace, IEnumAppNamespace, ITypeServiceNamespace, IEnumServiceNamespace, ITypeEntityNamespace, IEnumEntityNamespace } from '@ongoku/app-lib/src/common/namespacev2'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BoolCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

const serviceReq: ServiceReq = {
    namespace: {
        service: 'core',
        entity: undefined,
    },
}

const typeReqs: TypeReq[] = [
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
]

const enumReqs: EnumReq[] = [
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
]
