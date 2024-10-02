'use client'

import { AppInfoContext } from '@ongoku/app-lib/src/common/AppContext'
import { useContext } from 'react'
import { EntityMinimal } from '@ongoku/app-lib/src/common/Entity'
import { EntityDetail } from '@ongoku/app-lib/src/components/mantine/EntityDetail'

interface Props {
    params: {
        service: string
        entity: string
        identifier: string
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

    // Get the entity from the server
    const { identifier } = props.params

    return <EntityDetail entityInfo={entityInfo} identifier={identifier} />
}

export default Page
