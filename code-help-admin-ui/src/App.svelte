<script lang="ts">
  import Router, { link } from "svelte-spa-router";
  import { routes } from "./routes";
  import { Route } from "./route-utils";
</script>

<style>
  .content {
    width: 100%;
    height: 100%;
    justify-content: center;
    align-items: center;
    display: flex;
  }

  .container {
    height: 100vh;
  }

  nav {
    padding-block: 1rem;
  }
</style>

<div class="container">
  <nav>
    <a href={Route.counter} use:link>
      <button>Counter</button>
    </a>
    <a href={Route.other} use:link>
      <button>Other</button>
    </a>
  </nav>
  <main class="content">
    <Router {routes} />
  </main>
</div>
