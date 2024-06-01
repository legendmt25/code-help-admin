<script lang="ts">
  import Button from "../../components/Button.svelte";
  import Spinner from "../../components/Spinner.svelte";
  import type { Category } from "../../generated/admin-api";
  import { getAllCategories, deleteCategory } from "../../services/core/ProblemsService";
  import CategoryEditCreateDialog from "./CategoryEditCreateDialog.svelte";

  let categoryValueToEdit: string | undefined = undefined;
  let categoryIdToEdit: number | undefined = undefined;
  let createEditCategoryDialog: HTMLDialogElement | undefined = undefined;

  const handleGetAllCategories = () =>
    getAllCategories()
      .then((data) => data.categories)
      .catch(() => []);

  let getAllCategoriesPromise = handleGetAllCategories();

  const handleDeleteCategory = (categoryId: number) =>
    deleteCategory(categoryId).then(() => (getAllCategoriesPromise = handleGetAllCategories()));

  const handleSaveCategorySuccess = () => {
    getAllCategoriesPromise = handleGetAllCategories();
  };

  const handleEditCategory = (categoryId: number, categoryName: string) => {
    categoryIdToEdit = categoryId;
    categoryValueToEdit = categoryName;
    createEditCategoryDialog?.showModal();
  };

  let search: string | undefined = undefined;
  $: filter = (categories: Category[]) =>
    categories.filter(
      (category) =>
        !search ||
        category.id.toString().toLowerCase().includes(search.toLowerCase()) ||
        category.name.toLowerCase().includes(search.toLowerCase())
    );
</script>

<style>

  section {
    padding: 3rem 2rem;
    display: flex;
    flex-direction: column;
    gap: 8rem;
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
        categoryValueToEdit = undefined;
        categoryIdToEdit = undefined;
        createEditCategoryDialog?.showModal();
      }}>Create</Button>
    <h2>Categories</h2>
    <hr />
    <div class="input-container">
      <input id="search" placeholder="Search" name="search" bind:value={search} />
    </div>
    <table>
      <thead>
        <tr>
          <th>#</th>
          <th>Name</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {#if categories.length === 0}
          <div class="no-entries">No entries!</div>
        {/if}
        {#each filter(categories) as categoryEntry}
          <tr>
            <td>{categoryEntry.id}</td>
            <td>{categoryEntry.name}</td>
            <td>
              <Button on:click={() => handleEditCategory(categoryEntry.id, categoryEntry.name)}>Edit</Button>
              <Button on:click={() => handleDeleteCategory(categoryEntry.id)}>Delete</Button>
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
  bind:categoryNameValue={categoryValueToEdit}
  bind:categoryId={categoryIdToEdit}
  on:success={handleSaveCategorySuccess}
  bind:dialog={createEditCategoryDialog} />
