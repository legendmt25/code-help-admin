<script lang="ts">
  import type { FormEventHandler } from "svelte/elements";
  import Dialog from "../components/Dialog.svelte";
  import { createCategory, updateCategory } from "../services/ProblemsService";
  import Button from "../components/Button.svelte";
  import { createEventDispatcher } from "svelte";

  export let categoryId: number | undefined = undefined;
  export let categoryNameValue: string | undefined = undefined;
  export let dialog: HTMLDialogElement | undefined = undefined;

  const title = !categoryId ? "Create category" : `Update category - ${categoryId}`;

  const dispatch = createEventDispatcher();

  const handleFormSubmit: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();

    if (!categoryNameValue) {
      return;
    }

    let createEditCategoryPromise: Promise<void>;

    if (!Number.isNaN(Number(categoryId))) {
      console.log('edit category: ', categoryId);
      createEditCategoryPromise = updateCategory(categoryId!, { categoryRequest: { name: categoryNameValue } });
    } else {
      console.log('create category');
      createEditCategoryPromise = createCategory({ categoryRequest: { name: categoryNameValue } });
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
  }

  input {
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
</style>

<Dialog bind:dialog {title}>
  <form on:submit={handleFormSubmit} id="create-category-form" class="form p-1">
    <div class="input-container">
      <label for="name">Category name</label>
      <input required id="name" name="name" bind:value={categoryNameValue} />
    </div>
  </form>
  <div slot="footer">
    <Button form="create-category-form">Save</Button>
  </div>
</Dialog>
