<script lang="ts">
  import { onMount } from "svelte";
  import { Icon } from "svelte-icons-pack";
  import { AiOutlinePlusCircle } from "svelte-icons-pack/ai";
  import type { FormEventHandler } from "svelte/elements";
  import Button from "../../components/Button.svelte";
  import MessageBox from "../../components/MessageBox.svelte";
  import Spinner from "../../components/Spinner.svelte";
  import Tabbar from "../../components/Tabbar.svelte";
  import { type Community, type CommunityRequest, type ForumCategory } from "../../generated/admin-api";
  import {
    addCommunityModerator,
    createCommunity,
    deletePost,
    getAllCategories,
    getCommunityByName,
    removeCommunityModerator,
    updateCommunity
  } from "../../services/forum/ForumService";
  import { FormSubmitStatus } from "../../types";
  import type { TabOption } from "../../components/types";
  import { Route } from "../../routes";

  const tabs: TabOption[] = [
    {
      key: "moderators",
      label: "Moderators"
    },
    {
      key: "posts",
      label: "Posts"
    }
  ];
  let activeTab: string = tabs[0].key;

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

  let formSubmitStatus: FormSubmitStatus;
  const handleFormSubmit: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();

    formSubmitStatus = FormSubmitStatus.SUBMITTING;
    let editCreatePromise: Promise<Community>;

    if (params.name) {
      editCreatePromise = updateCommunity(params.name, value as CommunityRequest);
    } else {
      editCreatePromise = createCommunity(value as CommunityRequest);
    }

    editCreatePromise
      .then(() => (formSubmitStatus = FormSubmitStatus.SUCCESS))
      .catch(() => (formSubmitStatus = FormSubmitStatus.ERROR));
  };

  const handleRemoveModerator = (username: string) => {
    if (!params.name) {
      return;
    }

    removeCommunityModerator(params.name, username);
  };

  const handleDeletePost = (postUid: string) => {
    if (!params.name) {
      return;
    }

    deletePost(postUid);
  };

  // const handleAddModerator = (username: string) => {
  //   if (!params.name) {
  //     return;
  //   }

  //   addCommunityModerator(params.name, username);
  // };
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

  form {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  .no-entries {
    padding: 10px;
    font-weight: bold;
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
      <div class="input-container">
        <label for="categories">Categories:</label>
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
      </div>
      <Button>Submit</Button>
      {#if formSubmitStatus === FormSubmitStatus.SUCCESS}
        <MessageBox type={FormSubmitStatus.SUCCESS}>Saved successfully!</MessageBox>
      {:else if formSubmitStatus === FormSubmitStatus.ERROR}
        <MessageBox type={FormSubmitStatus.ERROR}>Error happened while saving</MessageBox>
      {/if}
    </form>

    <div class="row gap-1">
      <Button type="primary-outline" class="icon-text"
        >Add new post <Icon src={AiOutlinePlusCircle} title="Add new post" size="24" /></Button>
      <Button type="primary-outline" class="icon-text"
        >Add new moderator <Icon src={AiOutlinePlusCircle} title="Add new moderator" size="24" /></Button>
    </div>

    <Tabbar options={tabs} defaultActive={activeTab} on:change={(x) => (activeTab = x.detail)} />
    {#if communityEntry && params.name}
      <table>
        <thead>
          <colgroup>
            <col style="width: 10%; text-align: right;" />
            <col style="width: 40%; text-align: right;" />
            <col style="width: 10%; text-align: right;" />
            <col style="width: 40%; text-align: right;" />
          </colgroup>
          <tr>
            {#if activeTab === "moderators"}
              <th>Moderator username</th>
              <th>Actions</th>
            {/if}
            {#if activeTab === "posts"}
              <th>#</th>
              <th>Title</th>
              <th>User</th>
              <th>Actions</th>
            {/if}
          </tr>
        </thead>

        <tbody>
          {#if activeTab === "moderators"}
            {#each communityEntry.moderators as moderator}
              <tr>
                <td>{moderator.username}</td>
                <td>
                  <Button on:click={() => handleRemoveModerator(moderator.username)}>Remove</Button>
                </td>
              </tr>
            {:else}
              <div class="no-entries">No entries!</div>
            {/each}
          {/if}
          {#if activeTab === "posts"}
            {#each communityEntry.posts as post}
              <tr>
                <td>{post.uid}</td>
                <td>{post.title}</td>
                <td>{post.user.username}</td>
                <td>
                  <Button maxContent href={post.uid ? Route.post_edit.replace(":uid", post.uid.toString()) : undefined}
                    >Edit</Button>
                  <Button maxContent on:click={() => handleDeletePost(post.uid)}>Delete</Button>
                </td>
              </tr>
            {:else}
              <div class="no-entries">No entries!</div>
            {/each}
          {/if}
        </tbody>
      </table>
    {/if}
  </div>
{/if}
