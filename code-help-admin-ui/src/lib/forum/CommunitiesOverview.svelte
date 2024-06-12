<script lang="ts">
  import { Icon } from "svelte-icons-pack";
  import { AiOutlinePlus } from "svelte-icons-pack/ai";
  import { BiEdit, BiTrash } from "svelte-icons-pack/bi";
  import AlertBox from "../../components/AlertBox.svelte";
  import Button from "../../components/Button.svelte";
  import Spinner from "../../components/Spinner.svelte";
  import { Route } from "../../routes";
  import { deleteCommunity, getAllCommunities } from "../../services/forum/ForumService";
  import { filterCommunities } from "../../util";

  const getAllCommunitiesPromise = getAllCommunities().then((data) => data.communities);

  const handleDeleteCommunity = (communityName: string) => {
    deleteCommunity(communityName);
  };

  let search: string | undefined = undefined;
  $: filter = filterCommunities.bind(undefined, search);
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

{#await getAllCommunitiesPromise}
  <Spinner />
{:then communities}
  <section>
    <div class="page-heading">
      <h2>Communities</h2>
      <Button type="primary-outline" href={Route.communities_create} maxContent circled>
        <span>Create</span>
        <Icon src={AiOutlinePlus} size="18" />
      </Button>
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
        {#each filter(communities) as communityEntry}
          <tr>
            <td>{communityEntry.name}</td>
            <td>{communityEntry.description}</td>
            <td>
              <Button maxContent href={Route.communities_edit.replace(":name", communityEntry.name)}>
                <Icon src={BiEdit} size="24" />
                <span>Edit</span>
              </Button>
              <Button maxContent on:click={() => handleDeleteCommunity(communityEntry.name)}>
                <Icon src={BiTrash} size="24" />
                <span>Delete</span>
              </Button>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
    {#if communities.length === 0}
      <AlertBox type="info" message="No entries!" />
    {/if}
  </section>
{:catch err}
  <AlertBox type="error" message="An error occured! ({err})" />
{/await}
