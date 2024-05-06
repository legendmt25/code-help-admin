<script lang="ts">
  import Router, { link } from "svelte-spa-router";
  import { Route, routes } from "./routes";
  import { AiFillHome, AiFillQuestionCircle, AiOutlineBarChart, AiOutlineMenu } from "svelte-icons-pack/ai";
  import { RiCommunicationChat1Fill, RiLogosCoreosFill } from "svelte-icons-pack/ri";
  import { BiCategory } from "svelte-icons-pack/bi";
  import { Icon } from "svelte-icons-pack";
  import { KEYCLOAK_KEY, getInstance, initKeycloak } from "./services/KeycloakService";
  import { onMount } from "svelte";
  import env from "./env";
  import Accordion from "./components/Accordion.svelte";
  import { FaBrandsForumbee } from "svelte-icons-pack/fa";

  let menuOpened: boolean = false;

  onMount(() => {
    if (!env.DEV || !getInstance())
      initKeycloak().then((keycloak) => localStorage.setItem(KEYCLOAK_KEY, JSON.stringify(keycloak)));
  });

  const openMenu = () => (menuOpened = true);
  const toggleMenu = () => (menuOpened = !menuOpened);
</script>

<style>
  .container {
    display: flex;
    height: 100%;

    flex: 1 0 100%;
  }

  .h-100 {
    height: 100%;
  }

  nav {
    min-width: 250px;
    display: flex;
    flex-direction: column;
    gap: 0.2rem;
    box-shadow: rgba(0, 0, 0, 0.05) 1.95px 1px 2px 0px;
    background-color: #18283b;
    color: #8392a5;
    transition: all 200ms;
  }

  main {
    width: 100%;
  }

  a {
    font-weight: 500;
    text-decoration: inherit;
    display: flex;
    align-items: center;
    gap: 0.3rem;
    color: inherit;
    transition: all 200ms;
  }

  a:hover {
    color: #000;
  }

  .link-1 {
    padding: 1rem;
  }

  .link-2 {
    padding: 1rem;
    background-color: #202e3f;
  }

  .nav-head {
    justify-content: space-between;
    user-select: none;
  }

  span {
    overflow: hidden;
    max-width: 100%;
    transition: all 200ms;
    white-space: nowrap;
    display: inline-block;
  }

  .column {
    display: flex;
    flex-direction: column;
    height: 100%;
  }

  .logged-in {
    padding: 1rem;
    overflow: hidden;

    white-space: nowrap;
    max-width: 100%;
    box-sizing: border-box;
    transition: all 200ms;
    margin-top: auto;
  }

  .hide {
    max-width: 0;
    min-width: 0;
    padding: 0;
  }

  .hide-nav {
    max-width: 60px;
    min-width: 0;
  }

  .p-1 {
    padding: 1rem;
  }
</style>

<div class="container">
  <!-- svelte-ignore a11y-no-static-element-interactions -->
  <nav class:hide-nav={!menuOpened}>
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- svelte-ignore a11y-missing-attribute -->
    <a class="nav-head link-1" on:click={toggleMenu}>
      <span class:hide={!menuOpened}>CODE HELP ADMIN</span>
      <Icon src={AiOutlineMenu} size="32" />
    </a>
    <div class="h-100 column">
      <a class="link-1" href={Route.index} use:link>
        <Icon src={AiFillHome} size="32" />
        <span class:hide={!menuOpened}>Home</span>
      </a>
      <div>
        <Accordion>
          <!-- svelte-ignore a11y-missing-attribute -->
          <a class="link-1" slot="title" on:click={openMenu} on:keydown={undefined}>
            <Icon src={RiLogosCoreosFill} size="32" />
            <span class:hide={!menuOpened}>Core</span>
          </a>
          <div>
            <a class="link-2" href={Route.categories_overview} use:link>
              <Icon src={BiCategory} size="32" />
              <span class:hide={!menuOpened}>Categories</span>
            </a>
            <a class="link-2" href={Route.problems_overview} use:link>
              <Icon src={AiFillQuestionCircle} size="32" />
              <span class:hide={!menuOpened}>Problems</span>
            </a>
            <a class="link-2" href={Route.contests_overview} use:link>
              <Icon src={AiOutlineBarChart} size="32" />
              <span class:hide={!menuOpened}>Contests</span>
            </a>
          </div>
        </Accordion>
      </div>
      <div>
        <Accordion>
          <!-- svelte-ignore a11y-missing-attribute -->
          <a class="link-1" slot="title" on:click={openMenu} on:keydown={undefined}>
            <Icon src={FaBrandsForumbee} size="32" />
            <span class:hide={!menuOpened}>Forum</span>
          </a>
          <div>
            <a class="link-2" href={Route.communities_overview} use:link>
              <Icon src={RiCommunicationChat1Fill} size="32" />
              <span class:hide={!menuOpened}>Communities</span>
            </a>
          </div>
        </Accordion>
      </div>
      <div class="logged-in" class:hide={!menuOpened}>Logged in as: Martin</div>
    </div>
  </nav>
  <main class="content">
    <Router {routes} />
  </main>
</div>
