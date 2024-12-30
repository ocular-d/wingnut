// @ts-check
import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

// https://astro.build/config
export default defineConfig({
	integrations: [
		starlight({
			title: 'Wingnut',
			logo: {
				src: './src/assets/ocld-logo.png',
			},
            customCss: [
        	// Relativer Pfad zu deiner benutzerdefinierten CSS-Datei
            './src/styles/custom.css',
            ],
			social: {
				github: 'https://github.com/ocular-d/wingnut',
			},
			sidebar: [
				{
					label: 'Guides',
					items: [
						// Each item here is one entry in the navigation menu.
						{ label: 'Example Guide', slug: 'guides/example' },
					],
				},
				{
					label: 'Reference',
					autogenerate: { directory: 'reference' },
				},
                {
                    label: 'CLI',
                    autogenerate: { directory: 'cli' },
                },
			],
		}),
	],
});
