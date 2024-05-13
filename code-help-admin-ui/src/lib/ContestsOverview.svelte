<script lang="ts">
  import { link } from "svelte-spa-router";
  import Spinner from "../components/Spinner.svelte";
  import { Route } from "../routes";
  import { deleteContest, getAllContests } from "../services/ContestService";
  import Link from "../components/Link.svelte";

  const handleGetAllContests = () =>
    getAllContests()
      .then((data) => data.contests)
      .catch(() => []);

  let getAllContestsPromise = handleGetAllContests();

  const handleDeleteContest = (contestId: number) =>
    deleteContest(contestId).then(() => (getAllContestsPromise = handleGetAllContests()));
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
    <Link href={Route.contests_create}>Create</Link>
    <table>
      <thead>
        <tr>
          <th>#</th>
          <th>Name</th>
          <th>Starts on</th>
          <th>Duration</th>
          <th>Status</th>
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
            <td>{contestEntry.startsOn}</td>
            <td>{contestEntry.duration}</td>
            <td>{contestEntry.status}</td>
            <td>
              <a
                href={contestEntry.id ? Route.contests_edit.replace(":id", contestEntry.id.toString()) : undefined}
                use:link>Edit</a>
              <a on:click={() => handleDeleteContest(contestEntry.id)}>Delete</a>
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
