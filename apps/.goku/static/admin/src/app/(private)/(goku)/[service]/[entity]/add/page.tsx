'use client'

import { Title } from '@mantine/core'
import { AppInfoContext } from '@ongoku/app-lib/src/common/AppContext'
import { useContext } from 'react'
import { EntityAddForm } from '@ongoku/app-lib/src/components/mantine/EntityAdd'
import { EntityMinimal } from '@ongoku/app-lib/src/common/Entity'

interface Props {
    params: {
        service: string
        entity: string
    }
}

const Page = <E extends EntityMinimal = any>(props: Props) => {
    const { service: serviceName, entity: entityName } = props.params
    const { appInfo } = useContext(AppInfoContext)
    if (!appInfo) {
        throw new Error('AppInfo not loaded')
    }

    const serviceInfo = appInfo.getServiceInfo(serviceName)
    const entityInfo = serviceInfo.getEntityInfo<E>(entityName)

    return (
        <div>
            <Title order={1}>Add {entityInfo.getNameFormatted()}</Title>
            <EntityAddForm<E> entityInfo={entityInfo} initialData={entityInfo.getEmptyInstance()} />
        </div>
    )
}

export default Page
