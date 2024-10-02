'use client'

import { LayoutRootPublic } from '@ongoku/app-lib/src/components/admin/mantine/LayoutRootPublic'

const PublicRootLayout = ({ children }: { children: React.ReactNode }) => {
    return <LayoutRootPublic>{children}</LayoutRootPublic>
}

export default PublicRootLayout
