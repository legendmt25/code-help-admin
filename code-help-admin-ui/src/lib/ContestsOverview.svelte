<script lang="ts">
  import Spinner from "../components/Spinner.svelte";
  import { Route } from "../routes";
  import { getAllContests } from "../services/ContestService";

  const getAllContestsPromise = getAllContests().then((data) => data.contests);
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

{#await getAllContestsPromise}
  <Spinner />
{:then contests}
  <section>
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
        {#if contests.length === 0}
          <div class="no-entries">No entries!</div>
        {/if}
        {#each contests as contestEntry}
          <tr>
            <td>{contestEntry.id}</td>
            <td>{contestEntry.name}</td>
            <td>{contestEntry.name}</td>
            <td>
              <a
                href={contestEntry.id ? Route.contests_edit.replace(":id", contestEntry.id.toString()) : undefined}
                use:link>Edit</a>
              /
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </section>
{:catch err}
  <div>An error occured!</div>
{/await}
