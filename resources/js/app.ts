import '../css/app.css';
import '@types/@inertiajs__svelte'

import { createInertiaApp } from '@inertiajs/svelte';
import { resolvePageComponent } from 'laravel-vite-plugin/inertia-helpers';
import router from 'ziggy-js';

// declare global {
//     interface Window {
//         route: typeof router
//     }
// }


declare let window: Window & typeof globalThis & {
    route: typeof router
};
// window.route = window.route;

// window.route()
// window.route()


// window.route()

createInertiaApp({
    resolve: (name) => resolvePageComponent(`./Pages/${name}.svelte`, import.meta.glob('./Pages/**/*.svelte')),
    setup({ el, App, props }) {
        return new App({target: el, props})
    },
    progress: {
        color: '#4B5563',
    },
});
