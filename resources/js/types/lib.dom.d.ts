export {};

import router from "ziggy-js";
// declare global {
interface Window {
    route: typeof router;
}

declare var window: Window & typeof globalThis & {
    route: typeof router
}
// }
