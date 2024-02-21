// declare module '@inertiajs/svelte';


// declare module '@inertiajs/svelte' {
//     import {Writable} from "svelte/store";
//     import LinkComponent from "@inertiajs/svelte/src/Link.svelte"
//     import {Router} from '@inertiajs/core/types/router';
//     import {VisitOptions, Progress} from '@inertiajs/core/types/types';

//     type as<T> = T & {
//         isDirty: boolean,
//         hasErrors: boolean,
//         errors: unknown,
//         wasSuccessful: boolean,
//         recentlySuccessful: boolean,
//         processing: boolean,
//         progress: Progress | undefined,
//         data(): T,
//         transform(callback: (data: T) => as<T>),
//         submit(method: string, url: URL | string, options?: VisitOptions): void,
//         clearErrors(...[any]): as<T>,
//         setStore(unknown: string, unknown1: string): void,
//         get(url: URL | string, options?: VisitOptions): void,
//         post(url: URL | string, options?: VisitOptions): void,
//         put(url: URL | string, options?: VisitOptions): void,
//         patch(url: URL | string, options?: VisitOptions): void
//         delete(url: URL | string, options?: VisitOptions): void,
//         cancel(): void,
//     };

//     export const Link: LinkComponent;
//     export const router: Router

//     export function useForm<T>(historyKey: string, args: T): Writable<as<T>>;
// }

declare module '@inertiajs/svelte' {
    import {Writable} from "svelte/store";
    import LinkComponent from "@inertiajs/svelte/src/Link.svelte"
    import AppComponent from "@inertiajs/svelte/src/App.svelte"
    import {Router} from '@inertiajs/core/types/router';
    import {VisitOptions, Progress} from '@inertiajs/core/types/types';
    import { SvelteComponent } from "svelte";

    type UseFormStore<T, E> = T & {
        isDirty: boolean,
        hasErrors: boolean,
        errors: Record<keyof T, E>,
        wasSuccessful: boolean,
        recentlySuccessful: boolean,
        processing: boolean,
        progress: Progress | undefined,
        data(): T,
        transform<R>(callback: (data: T) => R): UseFormStore<R, E>,
        submit(method: string, url: URL | string, options?: VisitOptions): void,
        clearErrors(keys: (keyof T)[]): UseFormStore<T, E>,
        setError(values: Partial<Record<keyof T, E>>): UseFormStore<T, E>,
        reset(...keys: (keyof T)[]): UseFormStore<T, E>,
        defaults(values: Partial<T>): UseFormStore<T, E>,
        setStore(unknown: string, unknown1: string): void,
        get(url: URL | string, options?: VisitOptions): void,
        post(url: URL | string, options?: VisitOptions): void,
        put(url: URL | string, options?: VisitOptions): void,
        patch(url: URL | string, options?: VisitOptions): void
        delete(url: URL | string, options?: VisitOptions): void,
        cancel(): void,
    };

    export const Link: LinkComponent;
    export const router: Router;
    export function createInertiaApp<T extends {component: string}>(options: {
        id?: string,
        resolve: (name: string) => Promise<Awaited<unknown>>,
        setup:  (data: {
            el: HTMLElement,
            App: typeof SvelteComponent,
              props: {
                initialPage: T,
                resolveComponent: (name: string) =>  Promise<Awaited<SvelteComponent>>
            }
         }) => void,
        progress?: {
            delay?: number | undefined;
            color?: string | undefined;
            includeCSS?: boolean | undefined;
            showSpinner?: boolean | undefined;
        },
        page?: T
    }): unknown;

    export function useForm<T, E>(historyKey: string, args: T): Writable<UseFormStore<T, E>>;
    export function useForm<T, E>(args: T): Writable<UseFormStore<T, E>>;
}



