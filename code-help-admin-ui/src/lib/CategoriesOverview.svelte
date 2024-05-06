<script lang="ts">
  import Button from "../components/Button.svelte";
  import Spinner from "../components/Spinner.svelte";
  import { getAllCategories, updateCategory } from "../services/ProblemsService";
  import CategoryEditCreateDialog from "./CategoryEditCreateDialog.svelte";

  let categoryToEdit: string | undefined = undefined;
  let createEditCategoryDialog: HTMLDialogElement | undefined = undefined;

  const handleGetAllCategories = () =>
    getAllCategories()
      .then((data) => data.categories)
      .catch(() => []);

  let getAllCategoriesPromise = handleGetAllCategories();

  const handleDeleteCategory = (categoryName: string) => {};

  const handleSaveCategorySuccess = () => {
    getAllCategoriesPromise = handleGetAllCategories();
  };

  const handleEditCategory = (categoryName: string) => {
    categoryToEdit = categoryName;
    createEditCategoryDialog?.showModal();
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

  section {
    padding: 1rem;
    display: flex;
    flex-direction: column;
    gap: 1rem;
    position: relative;
  }

  .no-entries {
    padding: 10px;
    font-weight: bold;
  }
</style>

{#await getAllCategoriesPromise}
  <Spinner />
{:then categories}
  <section>
    <Button
      on:click={() => {
        categoryToEdit = undefined;
        createEditCategoryDialog?.showModal();
      }}>Create</Button>
    <table>
      <thead>
        <tr>
          <th>Name</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {#if categories.length === 0}
          <div class="no-entries">No entries!</div>
        {/if}
        {#each categories as categoryEntry}
          <tr>
            <td>{categoryEntry.name}</td>
            <td>
              <Button on:click={() => handleEditCategory(categoryEntry.name)}>Edit</Button>
              <Button on:click={() => handleDeleteCategory(categoryEntry.name)}>Delete</Button>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </section>
{:catch err}
  <div>An error occured! ({err})</div>
{/await}

<CategoryEditCreateDialog
  bind:oldCategoryNameValue={categoryToEdit}
  on:success={() => {
    handleSaveCategorySuccess();
    console.log("category save success");
  }}
  bind:dialog={createEditCategoryDialog} />
