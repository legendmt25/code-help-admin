<script lang="ts">
  import { onMount } from "svelte";
  import { getAllProblems } from "../services/ProblemsService";
  import type { ProblemEntry } from "../generated/core-api";
  import { link } from "svelte-spa-router";
  import Spinner from "../components/Spinner.svelte";
  import { Route } from "../routes";
  import Button from "../components/Button.svelte";

  let loading: boolean = false;
  let problems: ProblemEntry[] = [];

  onMount(() => {
    loading = true;
    getAllProblems()
      .then((x) => (problems = x.problems))
      .finally(() => (loading = false));
  });
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

{#if loading}
  <Spinner />
{:else}
  <section>
    <Button>Create</Button>
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
              <a>Delete</a>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </section>
{/if}
