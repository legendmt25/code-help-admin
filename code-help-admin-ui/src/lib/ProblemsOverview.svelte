<script lang="ts">
  import { deleteProblem, getAllProblems } from "../services/ProblemsService";
  import { link } from "svelte-spa-router";
  import Spinner from "../components/Spinner.svelte";
  import { Route } from "../routes";
  import Link from "../components/Link.svelte";

  const getAllProblemsPromise = getAllProblems()
    .then((x) => x.problems)
    .catch(() => []);

  const handleDeleteProblem = (problemId: number) => deleteProblem(problemId);
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
    <table>
      <thead>
        <tr>
          <th>#</th>
          <th>Name</th>
          <th>Language</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {#if problems.length === 0}
          <div class="no-entries">No entries!</div>
        {/if}
        {#each problems as problemEntry}
          <tr>
            <td>{problemEntry.id}</td>
            <td>{problemEntry.title}</td>
            <td>
              <a
                href={problemEntry.id ? Route.problems_edit.replace(":id", problemEntry.id.toString()) : undefined}
                use:link>Edit</a>
              /
              <!-- svelte-ignore a11y-click-events-have-key-events -->
              <!-- svelte-ignore a11y-no-static-element-interactions -->
              <!-- svelte-ignore a11y-missing-attribute -->
              <a on:click={() => handleDeleteProblem(problemEntry.id)}>Delete</a>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </section>
{:catch err}
  <div>An error occured!</div>
{/await}
