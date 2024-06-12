<script lang="ts">
  import { onMount } from "svelte";
  import { Icon } from "svelte-icons-pack";
  import { AiOutlinePlusCircle } from "svelte-icons-pack/ai";
  import { BiEdit, BiTrash } from "svelte-icons-pack/bi";
  import { BsEye } from "svelte-icons-pack/bs";
  import type { FormEventHandler } from "svelte/elements";
  import Accordion from "../../components/Accordion.svelte";
  import AlertBox from "../../components/AlertBox.svelte";
  import Button from "../../components/Button.svelte";
  import MessageBox from "../../components/MessageBox.svelte";
  import Spinner from "../../components/Spinner.svelte";
  import { type Community, type CommunityRequest, type ForumCategory } from "../../generated/admin-api";
  import { Route } from "../../routes";
  import {
    createCommunity,
    deletePost,
    getAllCategories,
    getCommunityByName,
    getPosts,
    removeCommunityModerator,
    updateCommunity
  } from "../../services/forum/ForumService";
  import { FormSubmitStatus } from "../../types";
  import { filterCommunityModerators, filterPosts } from "../../util";
  import PostEditCreateDialog from "./PostEditCreateDialog.svelte";
  import AddCommunityModeratorDialog from "./AddCommunityModeratorDialog.svelte";

  export let params: { name?: string } = {};

  let error: any | undefined = undefined;
  let loading: boolean = false;
  let value: Partial<CommunityRequest> & {
    categories: { uids: string[] };
  } = {
    categories: {
      uids: []
    }
  };

  let communityEntry: Community | undefined = undefined;
  let postEditCreateDialog: HTMLDialogElement | undefined = undefined;
  let postUidToEdit: string | undefined = undefined;

  let addCommunityModeratorDialog: HTMLDialogElement | undefined = undefined;

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

      Promise.all(promiseArray)
        .catch((err) => (error = err))
        .finally(() => (loading = false));
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
    if (!params.name || !communityEntry) {
      return;
    }

    removeCommunityModerator(params.name, username).then(() => {
      communityEntry = {
        ...communityEntry,
        moderators: communityEntry?.moderators.filter((x) => x.username !== username) ?? []
      };
    });
  };

  const handleCreatePost = () => {
    postUidToEdit = undefined;

    postEditCreateDialog?.showModal();
  };

  const handleEditPost = (postUid: string) => {
    postUidToEdit = postUid;

    postEditCreateDialog?.showModal();
  };

  const handleDeletePost = (postUid: string) => {
    if (!params.name || !communityEntry) {
      return;
    }

    deletePost(postUid).then(() => {
      getPosts(params.name).then((resp) => {
        communityEntry = { ...communityEntry, posts: resp.posts ?? [] };
      });
    });
  };

  const handleAddModerator = () => addCommunityModeratorDialog?.showModal();

  let moderatorsSearch: string | undefined = undefined;
  let postsSearch: string | undefined = undefined;

  $: filteredPosts = filterPosts.bind(undefined, postsSearch);
  $: filteredModerators = filterCommunityModerators.bind(undefined, moderatorsSearch);
</script>

<style>
  form {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }
</style>

{#if loading}
  <Spinner />
{:else if error}
  <AlertBox type="error" message="An error occured while getting the community data" />
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
            <option value={category.uid} selected={!!value.categories.uids.find((uid) => category.uid === uid)}>
              {category.name}
            </option>
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
    {#if communityEntry && params.name}
      <div class="gap-1">
        <Accordion>
          <svelte:fragment slot="title">Moderators</svelte:fragment>
          <div class="column gap-1 py-1">
            <div class="row justify-space-between gap-1">
              <div class="input-container">
                <input id="search" placeholder="Search" name="search" bind:value={moderatorsSearch} />
              </div>
              <Button type="primary-outline" class="icon-text" maxContent on:click={handleAddModerator}>
                Add moderator
                <Icon src={AiOutlinePlusCircle} title="Add new moderator" size="24" />
              </Button>
            </div>
            <table>
              <thead>
                <tr>
                  <th>Moderator username</th>
                  <th>Actions</th>
                </tr>
              </thead>
              <tbody>
                {#each filteredModerators(communityEntry.moderators) as moderator}
                  <tr>
                    <td>{moderator.username}</td>
                    <td>
                      <Button maxContent on:click={() => handleRemoveModerator(moderator.username)}>
                        <Icon src={BiTrash} size="24" />
                        <span>Remove</span>
                      </Button>
                    </td>
                  </tr>
                {/each}
              </tbody>
            </table>
            {#if communityEntry.moderators.length === 0}
              <AlertBox type="info" message="No moderators!" />
            {/if}
          </div>
        </Accordion>
        <Accordion>
          <svelte:fragment slot="title">Posts</svelte:fragment>
          <div class="py-1 column gap-1">
            <div class="row justify-space-between gap-1">
              <div class="input-container">
                <input id="search" placeholder="Search" name="search" bind:value={postsSearch} />
              </div>
              <Button type="primary-outline" class="icon-text" maxContent on:click={handleCreatePost}>
                Add new post
                <Icon src={AiOutlinePlusCircle} title="Add new post" size="24" />
              </Button>
            </div>
            <table>
              <thead>
                <colgroup>
                  <col style="width: 10%; text-align: right;" />
                  <col style="width: 40%; text-align: right;" />
                  <col style="width: 10%; text-align: right;" />
                  <col style="width: 40%; text-align: right;" />
                </colgroup>
                <tr>
                  <th>#</th>
                  <th>Title</th>
                  <th>User</th>
                  <th>Actions</th>
                </tr>
              </thead>
              <tbody>
                {#each filteredPosts(communityEntry.posts) as post}
                  <tr>
                    <td>{post.uid}</td>
                    <td>{post.title}</td>
                    <td>{post.user.username}</td>
                    <td>
                      <Button
                        maxContent
                        href={post.uid ? Route.post_detail.replace(":uid", post.uid.toString()) : undefined}>
                        <Icon src={BsEye} size="24" />
                        <span>View</span>
                      </Button>
                      <Button maxContent on:click={() => handleEditPost(post.uid)}>
                        <Icon src={BiEdit} size="24" />
                        <span>Edit</span>
                      </Button>
                      <Button maxContent on:click={() => handleDeletePost(post.uid)}>
                        <Icon src={BiTrash} size="24" />
                        <span>Delete</span>
                      </Button>
                    </td>
                  </tr>
                {/each}
              </tbody>
            </table>
            {#if communityEntry.posts.length === 0}
              <AlertBox type="info" message="No posts!" />
            {/if}
          </div>
        </Accordion>
      </div>
    {/if}
  </div>
{/if}

{#if params.name}
  <PostEditCreateDialog
    bind:communityName={params.name}
    bind:dialog={postEditCreateDialog}
    bind:postUid={postUidToEdit}
    on:success={() => {
      if (params.name && communityEntry) {
        getPosts(params.name).then((allPosts) => {
          communityEntry = { ...communityEntry, posts: allPosts.posts ?? [] };
        });
      }
    }} />

  <AddCommunityModeratorDialog
    bind:communityName={params.name}
    bind:dialog={addCommunityModeratorDialog}
    on:success={({ detail: username }) => {
      communityEntry = {
        ...communityEntry,
        moderators: [...(communityEntry?.moderators ?? []), { username }]
      };
    }} />
{/if}
