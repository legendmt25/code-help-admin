<script lang="ts">
  import type { FormEventHandler } from "svelte/elements";
  import Dialog from "../components/Dialog.svelte";
  import { createCategory, updateCategory } from "../services/ProblemsService";
  import Button from "../components/Button.svelte";

  export const oldCategoryNameValue: string | undefined = undefined;
  let categoryNameValue: string | undefined = undefined;
  export let dialog: HTMLDialogElement | undefined = undefined;

  const title = !oldCategoryNameValue ? "Create category" : `Update category - ${oldCategoryNameValue}`;

  const handleFormSubmit: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();

    if (!categoryNameValue) {
      return;
    }

    if (oldCategoryNameValue) {
      updateCategory(oldCategoryNameValue, { category: { name: categoryNameValue } });
    } else {
      createCategory({ category: { name: categoryNameValue } });
    }
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

<Dialog bind:dialog title={title}>
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
