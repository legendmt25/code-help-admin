<script lang="ts">
  import { deleteProblem, getAllProblems } from "../../services/core/ProblemsService";
  import { link } from "svelte-spa-router";
  import Spinner from "../../components/Spinner.svelte";
  import { Route } from "../../routes";
  import Link from "../../components/Link.svelte";
  import type { Problem } from "../../generated/admin-api";
  import Button from "../../components/Button.svelte";

  const getAllProblemsPromise = getAllProblems()
    .then((x) => x.problems)
    .catch(() => []);

  const handleDeleteProblem = (problemId: number) => deleteProblem(problemId);

  let search: string | undefined = undefined;
  $: filter = (problems: Problem[]) =>
    problems.filter(
      (problem) =>
        !search ||
        problem.title.toLowerCase().includes(search.toLowerCase()) ||
        problem.difficulty.toLowerCase().includes(search.toLowerCase()) ||
        problem.id.toString().toLowerCase().includes(search.toLowerCase()) ||
        problem.category?.name.toLowerCase().includes(search.toLowerCase())
    );
</script>

<style>
  table {
    border-collapse: collapse;
    width: 100%;
  }

  table,
  th,
  td {
    border: 1px solid #eee;
    padding: 10px;
  }

  section {
    padding: 1rem;
    display: flex;
    flex-direction: column;
    gap: 1rem;
    position: relative;
  }

  .no-entries {
    padding: 10px;
    font-weight: bold;
  }
</style>

{#await getAllProblemsPromise}
  <Spinner />
{:then problems}
  <section>
    <Link href={Route.problems_create}>Create</Link>
    <h2>Problems</h2>
    <hr />
    <div class="input-container">
      <input id="search" placeholder="Search" name="search" bind:value={search} />
    </div>
    <table>
      <thead>
        <tr>
          <th>#</th>
          <th>Title</th>
          <th>Category</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {#if problems.length === 0}
          <div class="no-entries">No entries!</div>
        {/if}
        {#each filter(problems) as problemEntry}
          <tr>
            <td>{problemEntry.id}</td>
            <td>{problemEntry.title}</td>
            <td>{problemEntry.category?.name}</td>
            <td>
              <Link href={problemEntry.id ? Route.problems_edit.replace(":id", problemEntry.id.toString()) : undefined}
                >Edit</Link>
              <!-- svelte-ignore a11y-click-events-have-key-events -->
              <!-- svelte-ignore a11y-no-static-element-interactions -->
              <!-- svelte-ignore a11y-missing-attribute -->
              <Button on:click={() => handleDeleteProblem(problemEntry.id)}>Delete</Button>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </section>
{:catch err}
  <div>An error occured!</div>
{/await}
