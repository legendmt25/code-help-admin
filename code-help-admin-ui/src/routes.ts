import { type RouteDefinition } from "svelte-spa-router";
import { wrap } from "svelte-spa-router/wrap";
import { Route } from "./route-utils";
import Counter from "./lib/Counter.svelte";
import Other from "./lib/Other.svelte";

export const routes: RouteDefinition = {
  [Route.counter]: wrap({
    component: Counter,
  }),
  [Route.other]: wrap({
    component: Other
  })
};
