'use client'

import { LayoutRootPrivate } from '@ongoku/app-lib/src/components/admin/mantine/LayoutRootPrivate'

const PrivateRootLayout = (props: { children: React.ReactNode }) => {
    return <LayoutRootPrivate>{props.children}</LayoutRootPrivate>
}
export default PrivateRootLayout
