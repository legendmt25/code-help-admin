<script lang="ts">
  import { Icon } from "svelte-icons-pack";
  import { BiEdit, BiPlus, BiTrash } from "svelte-icons-pack/bi";
  import AlertBox from "../../components/AlertBox.svelte";
  import Button from "../../components/Button.svelte";
  import Spinner from "../../components/Spinner.svelte";
  import { Route } from "../../routes";
  import { deleteContest, getAllContests } from "../../services/core/ContestService";
  import { filterContests } from "../../util";

  const handleGetAllContests = () =>
    getAllContests()
      .then((data) => data.contests)
      .catch(() => []);

  let getAllContestsPromise = handleGetAllContests();

  const handleDeleteContest = (contestId: number) =>
    deleteContest(contestId).then(() => (getAllContestsPromise = handleGetAllContests()));

  let search: string | undefined = undefined;
  $: filter = filterContests.bind(undefined, search);
</script>

<style>
  section {
    padding: 3rem 2rem;
    display: flex;
    flex-direction: column;
    gap: 4rem;
    position: relative;
  }

  .page-heading {
    display: flex;
    justify-content: space-between;
  }

  .input-container {
    width: 33%;
  }
</style>

{#await getAllContestsPromise}
  <Spinner />
{:then contests}
  <section>
    <div class="page-heading">
      <h2>Contests</h2>
      <Button maxContent type="primary-outline" href={Route.contests_create}>
        <Icon src={BiPlus} size="24" />
        <span>Create</span>
      </Button>
    </div>

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
        {#each filter(contests) as contestEntry}
          <tr>
            <td>{contestEntry.id}</td>
            <td>{contestEntry.name}</td>
            <td>{contestEntry.startsOn.toLocaleString()}</td>
            <td>{contestEntry.duration}</td>
            <td>{contestEntry.status}</td>
            <td>
              <Button
                maxContent
                href={contestEntry.id ? Route.contests_edit.replace(":id", contestEntry.id.toString()) : undefined}>
                <Icon src={BiEdit} size="24" />
                <span>Edit</span>
              </Button>
              <Button maxContent on:click={() => handleDeleteContest(contestEntry.id)}>
                <Icon src={BiTrash} size="24" />
                <span>Delete</span>
              </Button>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
    {#if contests.length === 0}
      <AlertBox type="info" message="No entries!" />
    {/if}
  </section>
{:catch err}
  <AlertBox type="error" message="An error occured! ({err})" />
{/await}
