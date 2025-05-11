<script lang="ts">
  import type { FormEventHandler } from "svelte/elements";
  import Dialog from "../../components/Dialog.svelte";
  import { createCategory, updateCategory } from "../../services/core/ProblemsService";
  import Button from "../../components/Button.svelte";
  import { createEventDispatcher } from "svelte";
  import type {CategoryRequest} from "../../generated/admin-core-api";

  export let categoryId: number | undefined = undefined;
  export let formData: Partial<CategoryRequest> = {};
  export let dialog: HTMLDialogElement | undefined = undefined;

  const title = !categoryId ? "Create category" : `Update category - ${categoryId}`;

  const dispatch = createEventDispatcher<{
    error: Error,
    success: void
  }>();

  const handleFormSubmit: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();

    if (!formData || !formData.name || !formData.description) {
      return;
    }

    let createEditCategoryPromise: Promise<void>;
    const request = {
      name: formData.name,
      description: formData.description,
    };

    if (!Number.isNaN(Number(categoryId))) {
      createEditCategoryPromise = updateCategory(categoryId!, {categoryRequest: request });
    } else {
      createEditCategoryPromise = createCategory({ categoryRequest: request }).then(() => {});
    }

    createEditCategoryPromise
      .catch((err) => dispatch("error", err))
      .then(() => dispatch("success"))
      .finally(() => dialog?.close());
  };
</script>

<style>
  .p-1 {
    padding: 1rem;
  }

  .input-container {
    display: flex;
    flex-direction: column;
    gap: 0.25rem;
  }

  label {
    font-size: 0.8em;
    font-weight: bold;
  }

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
  <form on:submit={handleFormSubmit} id="create-category-form" class="form p-1">
    <div class="input-container">
      <label for="name">Category name</label>
      <input required id="name" name="name" bind:value={formData.name} />
    </div>
    <div class="input-container">
      <label for="description">Category description</label>
      <input required id="description" name="description" bind:value={formData.description} />
    </div>
  </form>
  <div slot="footer">
    <Button form="create-category-form">Save</Button>
  </div>
</Dialog>
