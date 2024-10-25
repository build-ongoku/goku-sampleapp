import type { Config } from 'tailwindcss'
import path from 'path'

const config: Config = {
    content: [
        path.resolve(__dirname, './src/**/*.{js,ts,jsx,tsx,mdx}'), // Local content paths
        // Add any monorepo package which also uses tailwindcss
        path.join(path.dirname(require.resolve('@ongoku/app-lib/src')), '**/*.{ts,tsx}'),
    ],
    theme: {},
    plugins: [],
    corePlugins: {
        preflight: false,
    },
}

export default config
