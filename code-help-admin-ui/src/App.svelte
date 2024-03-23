<script lang="ts">
  import Router, { link } from "svelte-spa-router";
  import { Route, routes } from "./routes";
  import Button from "./components/Button.svelte";
</script>

<style>
  .container {
    height: 100vh;
    display: flex;
  }

  nav {
    min-width: 300px;
    display: flex;
    flex-direction: column;
    gap: 0.2rem;
    box-shadow: rgba(0, 0, 0, 0.05) 1.95px 1px 2px 0px;
  }

  main {
    width: 100%;
  }

  a {
    font-weight: 500;
    color: #646cff;
    text-decoration: inherit;
  }
  a:hover {
    color: #535bf2;
  }

  header {
    padding: 1rem;
    display: flex;
    justify-content: flex-end;
    box-shadow: rgba(0, 0, 0, 0.05) 300px 1px 2px 0px;
    border: 1px solid #eee;
  }
</style>

<header>
  <Button>Login</Button>
</header>
<div class="container">
  <nav>
    <a href={Route.problems_overview} use:link>
      <Button fullwidth>Problems</Button>
    </a>
  </nav>
  <main class="content">
    <Router {routes} />
  </main>
</div>
