<script lang="ts">
  import { Icon } from "svelte-icons-pack";
  import { AiFillHome, AiFillQuestionCircle, AiOutlineBarChart, AiOutlineMenu } from "svelte-icons-pack/ai";
  import { BiCategory, BiDoorOpen } from "svelte-icons-pack/bi";
  import { BsDoorOpenFill } from "svelte-icons-pack/bs";
  import { FaBrandsForumbee } from "svelte-icons-pack/fa";
  import { RiCommunicationChat1Fill, RiLogosCoreosFill } from "svelte-icons-pack/ri";
  import Router, { link } from "svelte-spa-router";
  import active from "svelte-spa-router/active";
  import Accordion from "./components/Accordion.svelte";
  import AlertBox from "./components/AlertBox.svelte";
  import Button from "./components/Button.svelte";
  import { Route, routes } from "./routes";
  import { getPrefferedUsername, isAuthenticated, login, logout } from "./services/KeycloakService";

  let menuOpened: boolean = false;
  const keycloak = window.keycloak!;

  let loggedIn: boolean = isAuthenticated(keycloak);

  const openMenu = () => (menuOpened = true);
  const toggleMenu = () => (menuOpened = !menuOpened);
</script>

<style>
  .container {
    display: flex;
    height: 100%;

    flex: 1 0 100%;
  }

  nav {
    min-width: 250px;
    display: flex;
    flex-direction: column;
    gap: 0.2rem;
    box-shadow: rgba(0, 0, 0, 0.05) 1.95px 1px 2px 0px;
    background-color: #fafafa;
    color: #757e8a;
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

  a.link.title {
    color: #000;
  }

  a.link:not(.title):hover {
    color: #000;
  }

  .link {
    padding: 1rem;
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

  .nav-head.link.title {
    cursor: pointer;
  }
</style>

<div class="container">
  <!-- svelte-ignore a11y-no-static-element-interactions -->
  <nav class:hide-nav={!menuOpened}>
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- svelte-ignore a11y-missing-attribute -->
    <a class="nav-head link title" on:click={toggleMenu}>
      <span class:hide={!menuOpened}>CODE HELP ADMIN</span>
      <Icon src={AiOutlineMenu} size="32" />
    </a>
    <div class="h-100 column">
      {#if loggedIn}
        <a class="link title" href={Route.index} use:link use:active={{ path: Route.index }}>
          <Icon src={AiFillHome} size="32" />
          <span class:hide={!menuOpened}>Home</span>
        </a>
        <div>
          <Accordion>
            <!-- svelte-ignore a11y-missing-attribute -->
            <svelte:fragment slot="title">
              <a class="link title" on:click={openMenu} on:keydown={undefined}>
                <Icon src={RiLogosCoreosFill} size="32" />
                <span class:hide={!menuOpened}>Core</span>
              </a>
            </svelte:fragment>
            <div>
              <a
                class="link"
                href={Route.categories_overview}
                use:link
                use:active={{ path: Route.categories_overview }}>
                <Icon src={BiCategory} size="32" />
                <span class:hide={!menuOpened}>Categories</span>
              </a>
              <a class="link" href={Route.problems_overview} use:link use:active={{ path: Route.problems_overview }}>
                <Icon src={AiFillQuestionCircle} size="32" />
                <span class:hide={!menuOpened}>Problems</span>
              </a>
              <a class="link" href={Route.contests_overview} use:link use:active={{ path: Route.contests_overview }}>
                <Icon src={AiOutlineBarChart} size="32" />
                <span class:hide={!menuOpened}>Contests</span>
              </a>
            </div>
          </Accordion>
        </div>
        <div>
          <Accordion>
            <!-- svelte-ignore a11y-missing-attribute -->
            <svelte:fragment slot="title">
              <a class="link title" on:click={openMenu} on:keydown={undefined}>
                <Icon src={FaBrandsForumbee} size="32" />
                <span class:hide={!menuOpened}>Forum</span>
              </a>
            </svelte:fragment>
            <div>
              <a
                class="link"
                href={Route.forum_categories_overview}
                use:link
                use:active={{ path: Route.forum_categories_overview }}>
                <Icon src={BiCategory} size="32" />
                <span class:hide={!menuOpened}>Categories</span>
              </a>
              <a
                class="link"
                href={Route.communities_overview}
                use:link
                use:active={{ path: Route.communities_overview }}>
                <Icon src={RiCommunicationChat1Fill} size="32" />
                <span class:hide={!menuOpened}>Communities</span>
              </a>
            </div>
          </Accordion>
        </div>
      {:else}
        <!-- svelte-ignore a11y-missing-attribute -->
        <a class="link" on:click={() => login(keycloak)} on:keyup={undefined}>
          <Icon src={BiDoorOpen} size="32" />
          <span class:hide={!menuOpened}>Login</span>
        </a>
      {/if}

      {#if getPrefferedUsername(keycloak)}
        <div class="logged-in column gap-1">
          <Button type="link" on:click={() => logout(keycloak)}>
            <Icon src={BsDoorOpenFill} size="32" />
            <div class="column">
              <span class:hide={!menuOpened}>Logged in as:</span>
              <span style:font-size="12px" class:hide={!menuOpened}>{getPrefferedUsername(keycloak)}</span>
            </div>
          </Button>
        </div>
      {:else}
        <div class="logged-in" class:hide={!menuOpened}>Not logged in</div>
      {/if}
    </div>
  </nav>
  <main class="w-100">
    {#if !loggedIn}
      <AlertBox type="info" message="You must be logged in!" />
    {:else}
      <Router {routes} />
    {/if}
  </main>
</div>
