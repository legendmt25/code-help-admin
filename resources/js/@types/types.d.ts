// declare module '@inertiajs/svelte';

import router from "ziggy-js";

// interface Window {
//     route: typeof router;
// }

declare const window: Window & typeof globalThis & {
    route: typeof router
}
