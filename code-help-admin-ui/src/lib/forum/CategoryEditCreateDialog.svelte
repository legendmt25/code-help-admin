<script lang="ts">
  import type { FormEventHandler } from "svelte/elements";
  import Dialog from "../../components/Dialog.svelte";
  import { createCategory, updateCategory } from "../../services/forum/ForumService";
  import Button from "../../components/Button.svelte";
  import { createEventDispatcher } from "svelte";

  export let categoryUid: string | undefined = undefined;
  export let categoryNameValue: string | undefined = undefined;
  export let dialog: HTMLDialogElement | undefined = undefined;

  const title = !categoryUid ? "Create category" : `Update category - ${categoryUid}`;

  const dispatch = createEventDispatcher<{
      error: Error,
      success: void
  }>();

  const handleFormSubmit: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();

    if (!categoryNameValue) {
      return;
    }

    let createEditCategoryPromise: Promise<void>;

    if (categoryUid) {
      createEditCategoryPromise = updateCategory(categoryUid!, categoryNameValue).then(() => undefined);
    } else {
      createEditCategoryPromise = createCategory(categoryNameValue).then(()=>{});
    }

    createEditCategoryPromise
      .catch((err) => dispatch("error", err))
      .then(() => dispatch("success"))
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
  <form on:submit={handleFormSubmit} id="edit-create-category-form" class="form p-1">
    <div class="input-container">
      <label for="name">Category name</label>
      <input required id="name" name="name" bind:value={categoryNameValue} />
    </div>
  </form>
  <div slot="footer">
    <Button form="edit-create-category-form">Save</Button>
  </div>
</Dialog>
