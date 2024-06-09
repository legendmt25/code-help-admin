<script lang="ts">
  import { Icon } from "svelte-icons-pack";
  import Button from "../../components/Button.svelte";
  import Spinner from "../../components/Spinner.svelte";
  import type { ShortCommunity } from "../../generated/admin-api";
  import { Route } from "../../routes";
  import { deleteCommunity, getAllCommunities } from "../../services/forum/ForumService";
  import { AiOutlinePlus } from "svelte-icons-pack/ai";

  const getAllCommunitiesPromise = getAllCommunities().then((data) => data.communities);

  const handleDeleteCommunity = (communityName: string) => {
    deleteCommunity(communityName);
  };

  let search: string | undefined = undefined;
  $: filter = (communities: ShortCommunity[]) =>
    communities.filter(
      (community) =>
        !search ||
        community.name.toLowerCase().includes(search.toLowerCase()) ||
        community.description.toLowerCase().includes(search.toLowerCase())
    );
</script>

<style>
  section {
    padding: 3rem 2rem;
    display: flex;
    flex-direction: column;
    gap: 4rem;
    position: relative;
  }

  .no-entries {
    padding: 10px;
    font-weight: bold;
  }

  .page-heading {
    display: flex;
    justify-content: space-between;
  }

  .input-container {
    width: 33%;
  }
</style>

{#await getAllCommunitiesPromise}
  <Spinner />
{:then communities}
  <section>
    <div class="page-heading">
      <h2>Communities</h2>
      <Button type="primary-outline" href={Route.communities_create} maxContent circled>Create <Icon src={AiOutlinePlus} size="18" /></Button>
    </div>
    <div>
      <div class="input-container">
        <input id="search" placeholder="Search" name="search" bind:value={search} />
      </div>
    </div>
    <hr />
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Description</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {#if !communities || communities.length === 0}
          <div class="no-entries">No entries!</div>
        {:else}
          {#each filter(communities) as communityEntry}
            <tr>
              <td>{communityEntry.name}</td>
              <td>{communityEntry.description}</td>
              <td>
                <Button href={Route.communities_edit.replace(":name", communityEntry.name)}>Edit</Button>
                <Button on:click={() => handleDeleteCommunity(communityEntry.name)}>Delete</Button>
              </td>
            </tr>
          {/each}
        {/if}
      </tbody>
    </table>
  </section>
{:catch err}
  <div>An error occured! ({err})</div>
{/await}
