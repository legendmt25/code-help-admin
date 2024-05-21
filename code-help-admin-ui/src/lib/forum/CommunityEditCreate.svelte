<script lang="ts">
  import { onMount } from "svelte";
  import { link, push } from "svelte-spa-router";
  import type { FormEventHandler } from "svelte/elements";
  import Button from "../../components/Button.svelte";
  import Spinner from "../../components/Spinner.svelte";
  import { type Community, type CommunityRequest, type ForumCategory } from "../../generated/admin-api";
  import { Route } from "../../routes";
  import {
    createCommunity,
    getAllCategories,
    getCommunityByName,
    updateCommunity
  } from "../../services/forum/ForumService";

  export let params: { name?: string } = {};

  let loading: boolean = false;
  let value: Partial<CommunityRequest> & {
    categories: { uids: string[] };
  } = {
    categories: {
      uids: []
    }
  };

  let communityEntry: Community | undefined = undefined;

  let categories: ForumCategory[] = [];
  onMount(() => {
    if (params.name) {
      loading = true;
      const promiseArray: Promise<void>[] = [
        getAllCategories().then((resp) => {
          categories = resp.categories;
        }),
        getCommunityByName(params.name).then((community) => {
          communityEntry = community;
          value = {
            ...value,
            name: community.name,
            description: community.description,
            categories: { uids: community.categories.map((x) => x.uid) },
            image: community.image
          };
        })
      ];

      Promise.all(promiseArray).finally(() => (loading = false));
    }
  });

  const handleFormSubmit: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();

    let editCreatePromise: Promise<Community>;

    if (params.name) {
      editCreatePromise = updateCommunity(params.name, value as CommunityRequest);
    } else {
      editCreatePromise = createCommunity(value as CommunityRequest);
    }

    editCreatePromise.finally(() => push(Route.communities_overview));
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

  .p-1 {
    padding: 1rem;
  }

  .input-container {
    display: flex;
    flex-direction: column;
  }

  input,
  select {
    width: 100%;
    max-width: 300px;
    padding: 5px;
    font-size: 1.05rem;
    outline: none;
  }

  form {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  .gap-1 {
    gap: 1rem;
  }

  .column {
    display: flex;
    flex-direction: column;
  }
</style>

{#if loading}
  <Spinner />
{:else}
  <div class="column p-1 gap-1">
    <form on:submit={handleFormSubmit} class="form">
      <div class="input-container">
        <label for="name">Name</label>
        <input id="name" name="name" bind:value={value.name} />
      </div>
      <div class="input-container">
        <label for="description">Description</label>
        <input id="description" name="description" bind:value={value.description} />
      </div>
      <select
        name="categories"
        multiple
        on:change={(event) =>
          (value.categories.uids = Array.from(event.currentTarget.selectedOptions).map((x) => x.value))}>
        {#each categories as category}
          <option value={category.uid} selected={!!value.categories.uids.find((uid) => category.uid === uid)}
            >{category.name}</option>
        {/each}
      </select>
      <Button>Submit</Button>
    </form>

    {#if communityEntry != undefined && params.name}
      <table>
        <thead>
          <tr>
            <th>#</th>
            <th>Title</th>
            <th>Score</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {#each communityEntry.posts as post}
            <tr>
              <td>{post.uid}</td>
              <td>{post.title}</td>
              <td>{post.user.username}</td>
              <td>
                <a href={post.uid ? Route.problems_edit.replace(":id", post.uid.toString()) : undefined} use:link
                  >Edit</a>
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    {/if}
  </div>
{/if}
