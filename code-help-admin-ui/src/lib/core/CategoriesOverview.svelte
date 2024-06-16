<script lang="ts">
  import { Icon } from "svelte-icons-pack";
  import { BiEdit, BiPlus, BiTrash } from "svelte-icons-pack/bi";
  import AlertBox from "../../components/AlertBox.svelte";
  import Button from "../../components/Button.svelte";
  import Spinner from "../../components/Spinner.svelte";
  import { deleteCategory, getAllCategories } from "../../services/core/ProblemsService";
  import { filterCategories } from "../../util";
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
  $: filter = filterCategories.bind(undefined, search);
</script>

<style>
  section {
    padding: 3rem 2rem;
    display: flex;
    flex-direction: column;
    gap: 4rem;
    position: relative;
  }

  .page-heading {
    display: flex;
    justify-content: space-between;
  }

  .input-container {
    width: 33%;
  }
</style>

{#await getAllCategoriesPromise}
  <Spinner />
{:then categories}
  <section>
    <div class="page-heading">
      <h2>Categories</h2>
      <Button
        maxContent
        type="primary-outline"
        on:click={() => {
          categoryValueToEdit = undefined;
          categoryIdToEdit = undefined;
          createEditCategoryDialog?.showModal();
        }}>
        <Icon src={BiPlus} size="24" />
        <span>Create</span>
      </Button>
    </div>
    <div>
      <div class="input-container">
        <input id="search" placeholder="Search" name="search" bind:value={search} />
      </div>
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
        {#each filter(categories) as categoryEntry}
          <tr>
            <td>{categoryEntry.id}</td>
            <td>{categoryEntry.name}</td>
            <td>
              <Button maxContent on:click={() => handleEditCategory(categoryEntry.id, categoryEntry.name)}>
                <Icon src={BiEdit} size="24" />
                <span>Edit</span>
              </Button>
              <Button maxContent on:click={() => handleDeleteCategory(categoryEntry.id)}>
                <Icon src={BiTrash} size="24" />
                <span>Delete</span>
              </Button>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
    {#if categories.length === 0}
      <AlertBox type="info" message="No entries!" />
    {/if}
  </section>
{:catch err}
  <AlertBox type="error" message="An error occured! ({err})" />
{/await}

<CategoryEditCreateDialog
  bind:categoryNameValue={categoryValueToEdit}
  bind:categoryId={categoryIdToEdit}
  on:success={handleSaveCategorySuccess}
  bind:dialog={createEditCategoryDialog} />
