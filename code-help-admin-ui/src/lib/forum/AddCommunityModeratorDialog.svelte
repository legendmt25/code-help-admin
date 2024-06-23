<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import type { FormEventHandler } from "svelte/elements";
  import Button from "../../components/Button.svelte";
  import Dialog from "../../components/Dialog.svelte";
  import { addCommunityModerator } from "../../services/forum/ForumService";

  export let communityName: string;
  let username: string | undefined = undefined;
  export let dialog: HTMLDialogElement | undefined = undefined;

  const title = "Add community moderator";

  const dispatch = createEventDispatcher<{
    error: Error,
    success: string
  }>();

  const handleFormSubmit: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();

    if (!username) {
      return;
    }

    addCommunityModerator(communityName, username)
      .catch((err) => dispatch("error", err))
      .then(() => dispatch("success", username!))
      .finally(() => dialog?.close());
  };
</script>

<style>
  form {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }
  
  .form {
    padding-inline: 2rem;
  }
</style>

<Dialog bind:dialog {title}>
  <form on:submit={handleFormSubmit} id="add-community-moderator-form" class="form p-1">
    <div class="input-container">
      <label for="username">Moderator username</label>
      <input required id="username" name="username" bind:value={username} />
    </div>
  </form>
  <div slot="footer">
    <Button form="add-community-moderator-form">Save</Button>
  </div>
</Dialog>
