<script>
  import { Icon } from "svelte-icons-pack";
  import { AiFillQuestionCircle, AiOutlineBarChart } from "svelte-icons-pack/ai";
  import { RiCommunicationChat1Fill } from "svelte-icons-pack/ri";
  import { link } from "svelte-spa-router";
  import { Route } from "../routes";
  import { BiCategory } from "svelte-icons-pack/bi";
</script>

<style>
  .column {
    display: flex;
    flex-direction: column;
    height: 100%;
  }

  .row {
    display: flex;
  }

  .gap-1 {
    gap: 1rem;
  }

  .p-1 {
    padding: 1rem;
  }

  .tile {
    width: 200px;
    height: 120px;
    background-color: #18283b;
    color: #8392a5;
    padding: 1rem;
    border-radius: 0.625rem;
    display: flex;
    gap: 0.5rem;
    text-decoration: none;
  }
</style>

<div class="p-1 column gap-1">
  <h2>Core</h2>
  <hr />
  <div class="row gap-1">
    <a class="tile" href={Route.categories_overview} use:link>
      <Icon src={BiCategory} size="32" />
      <span>Categories</span>
    </a>
    <a class="tile" href={Route.problems_overview} use:link>
      <Icon src={AiFillQuestionCircle} size="32" />
      <span>Problems</span>
    </a>
    <a class="tile" href={Route.contests_overview} use:link>
      <Icon src={AiOutlineBarChart} size="32" />
      <span>Contests</span>
    </a>
  </div>
  <h2>Forum</h2>
  <hr />
  <div class="row gap-1">
    <a class="tile" href={Route.communities_overview} use:link>
      <Icon src={RiCommunicationChat1Fill} size="32" />
      <span>Communities</span>
    </a>
  </div>
</div>
