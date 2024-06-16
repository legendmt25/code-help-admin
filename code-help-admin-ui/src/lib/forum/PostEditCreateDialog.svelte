<script lang="ts">
  import { onMount } from "svelte";
  import { createEventDispatcher } from "svelte";
  import type { FormEventHandler } from "svelte/elements";
  import Button from "../../components/Button.svelte";
  import Dialog from "../../components/Dialog.svelte";
  import Spinner from "../../components/Spinner.svelte";
  import { createPost, getPost, updatePost } from "../../services/forum/ForumService";
  import type { PostRequest } from "../../generated/admin-api";

  export let postUid: string | undefined = undefined;
  let postRequest: Partial<PostRequest> = {};
  let loading: boolean = false;

  export let communityName: string;

  export let dialog: HTMLDialogElement | undefined = undefined;

  $: postUid,
    (() => {
      if (postUid) {
        getPost(postUid)
          .then((postEntry) => (postRequest = postEntry))
          .finally(() => (loading = false));
      } else {
        postRequest = {};
      }
    })();

  const title = !postUid ? "Create Post" : `Update Post - ${postUid}`;

  const dispatch = createEventDispatcher();

  const handleFormSubmit: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();

    let createEditPostPromise: Promise<void>;

    const value = postRequest as PostRequest;

    if (!postUid) {
      createEditPostPromise = createPost(communityName, value).then(() => undefined);
    } else {
      createEditPostPromise = updatePost(postUid, value).then(() => undefined);
    }

    createEditPostPromise
      .catch((err) => dispatch("error", err))
      .then(() => dispatch("success"))
      .finally(() => dialog?.close());
  };
</script>

<style>
  form {
    display: flex;
    flex-direction: column;
    gap: 2rem;
  }

  .form {
    padding-inline: 2rem;
  }
</style>

<Dialog bind:dialog {title}>
  {#if loading}
    <Spinner />
  {:else}
    <!-- else content here -->
  {/if}
  <form on:submit={handleFormSubmit} id="post-edit-create-form" class="form p-1">
    <div class="input-container">
      <label for="title">Title</label>
      <input required id="title" name="title" bind:value={postRequest.title} />
    </div>
    <div class="input-container">
      <label for="content">Content</label>
      <textarea style:height="100px" style:resize="none" required id="content" name="content" bind:value={postRequest.content} />
    </div>
  </form>
  <div slot="footer">
    <Button form="post-edit-create-form">Save</Button>
  </div>
</Dialog>
