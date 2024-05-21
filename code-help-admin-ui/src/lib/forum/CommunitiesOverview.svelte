<script lang="ts">
  import Button from "../../components/Button.svelte";
  import Link from "../../components/Link.svelte";
  import Spinner from "../../components/Spinner.svelte";
  import { Route } from "../../routes";
  import { deleteCommunity, getAllCommunities } from "../../services/forum/ForumService";

  const getAllCommunitiesPromise = getAllCommunities().then((data) => data.communities);

  const handleDeleteCommunity = (communityName: string) => {
    deleteCommunity(communityName);
  };
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

{#await getAllCommunitiesPromise}
  <Spinner />
{:then communities}
  <section>
    <Link href={Route.communities_create}>Create</Link>
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Description</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {#if communities.length === 0}
          <div class="no-entries">No entries!</div>
        {/if}
        {#each communities as communityEntry}
          <tr>
            <td>{communityEntry.name}</td>
            <td>{JSON.stringify(communityEntry.description)}</td>
            <td>
              <Link href={Route.communities_edit.replace(':name', communityEntry.name)}>Edit</Link>
              <Button on:click={() => handleDeleteCommunity(communityEntry.name)}>Delete</Button>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </section>
{:catch err}
  <div>An error occured! ({err})</div>
{/await}
