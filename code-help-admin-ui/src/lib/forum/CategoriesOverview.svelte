<script lang="ts">
  import Button from "../../components/Button.svelte";
  import Spinner from "../../components/Spinner.svelte";
  import { deleteCategory, getAllCategories } from "../../services/forum/ForumService";
  import CategoryEditCreateDialog from "./CategoryEditCreateDialog.svelte";

  let categoryValueToEdit: string | undefined = undefined;
  let categoryUIdToEdit: string | undefined = undefined;
  let createEditCategoryDialog: HTMLDialogElement | undefined = undefined;

  const handleGetAllCategories = () =>
    getAllCategories()
      .then((data) => data.categories)
      .catch(() => []);

  let getAllCategoriesPromise = handleGetAllCategories();

  const handleDeleteCategory = (categoryUid: string) =>
    deleteCategory(categoryUid).then(() => (getAllCategoriesPromise = handleGetAllCategories()));

  const handleSaveCategorySuccess = () => {
    getAllCategoriesPromise = handleGetAllCategories();
  };

  const handleEditCategory = (categoryUd: string, categoryName: string) => {
    categoryUIdToEdit = categoryUd;
    categoryValueToEdit = categoryName;
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
        categoryValueToEdit = undefined;
        categoryUIdToEdit = undefined;
        createEditCategoryDialog?.showModal();
      }}>Create</Button>
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
        {#each categories as categoryEntry}
          <tr>
            <td>{categoryEntry.uid}</td>
            <td>{categoryEntry.name}</td>
            <td>
              <Button on:click={() => handleEditCategory(categoryEntry.uid, categoryEntry.name)}>Edit</Button>
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
  bind:categoryUid={categoryUIdToEdit}
  on:success={handleSaveCategorySuccess}
  bind:dialog={createEditCategoryDialog} />
