<script lang="ts">
  import Button from "../../components/Button.svelte";
  import Spinner from "../../components/Spinner.svelte";
  import { Route } from "../../routes";
  import { deleteContest, getAllContests } from "../../services/core/ContestService";
  import Link from "../../components/Link.svelte";
  import type { Contest } from "../../generated/admin-api";

  const handleGetAllContests = () =>
    getAllContests()
      .then((data) => data.contests)
      .catch(() => []);

  let getAllContestsPromise = handleGetAllContests();

  const handleDeleteContest = (contestId: number) =>
    deleteContest(contestId).then(() => (getAllContestsPromise = handleGetAllContests()));

  let search: string | undefined = undefined;
  $: filter = (contests: Contest[]) =>
    contests.filter(
      (contest) =>
        !search ||
        contest.name.toLowerCase().includes(search.toLowerCase()) ||
        contest.duration.toLowerCase().includes(search.toLowerCase()) ||
        contest.id.toString().toLowerCase().includes(search.toLowerCase())
    );
</script>

<style>

  section {
    padding: 3rem 2rem;
    display: flex;
    flex-direction: column;
    gap: 8rem;
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
    <h2>Contests</h2>
    <hr />
    <div class="input-container">
      <input id="search" placeholder="Search" name="search" bind:value={search} />
    </div>
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
        {#each filter(contests) as contestEntry}
          <tr>
            <td>{contestEntry.id}</td>
            <td>{contestEntry.name}</td>
            <td>{contestEntry.startsOn.toLocaleString()}</td>
            <td>{contestEntry.duration}</td>
            <td>{contestEntry.status}</td>
            <td>
              <Link href={contestEntry.id ? Route.contests_edit.replace(":id", contestEntry.id.toString()) : undefined}
                >Edit</Link>
              <Button on:click={() => handleDeleteContest(contestEntry.id)}>Delete</Button>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </section>
{:catch err}
  <div>An error occured!</div>
{/await}
