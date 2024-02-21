/// <reference types="@inertiajs/core" />
// declare module '@inertiajs/svelte' {
// }

// import '' from "Link.svel"

// const s = (await import('@inertiajs/svelte')).Link;


// declare module '@inertiajs/svelte' {
//     // import {Writable} from "svelte";
//     // import LinkComponent from "@inertiajs/svelte/src/Link.svelte"
//     import {Router} from '@inertiajs/core/types/router';
//     import {VisitOptions, Progress} from '@inertiajs/core/types/types';

//     export const Link: LinkComponent;
//     export const router: Router

//     export function useForm<T>(historyKey: string, args: T): Writable<T & {
//         isDirty: boolean,
//         hasErrors: boolean,
//         errors: unknown,
//         wasSuccessful: boolean,
//         recentlySuccessful: boolean,
//         processing: boolean,
//         progress: Progress | undefined,
//         data(): T,
//         transform(callback: (data: T) => this),
//         submit(method: string, url: URL | string, options?: VisitOptions): void,
//         clearErrors(...[any]): this,
//         setStore(any, any): void,
//         get(url: URL | string, options?: VisitOptions): void,
//         post(url: URL | string, options?: VisitOptions): void,
//         put(url: URL | string, options?: VisitOptions): void,
//         patch(url: URL | string, options?: VisitOptions): void
//         delete(url: URL | string, options?: VisitOptions): void,
//         cancel(): void,
//     }>;
// }
