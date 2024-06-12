<script lang="ts">
  import { Icon } from "svelte-icons-pack";
  import { BiEdit, BiPlus, BiTrash } from "svelte-icons-pack/bi";
  import AlertBox from "../../components/AlertBox.svelte";
  import Button from "../../components/Button.svelte";
  import Spinner from "../../components/Spinner.svelte";
  import { deleteCategory, getAllCategories } from "../../services/forum/ForumService";
  import { filterForumCategories } from "../../util";
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

  let search: string | undefined = undefined;
  $: filter = filterForumCategories.bind(undefined, search);
</script>

<style>
  section {
    padding: 3rem 2rem;
    display: flex;
    flex-direction: column;
    gap: 4rem;
    position: relative;
  }
</style>

{#await getAllCategoriesPromise}
  <Spinner />
{:then categories}
  <section>
    <h2>Forum Categories</h2>
    <div class="row justify-space-between">
      <div class="input-container">
        <input id="search" placeholder="Search" name="search" bind:value={search} />
      </div>
      <Button
        maxContent
        type="primary-outline"
        on:click={() => {
          categoryValueToEdit = undefined;
          categoryUIdToEdit = undefined;
          createEditCategoryDialog?.showModal();
        }}>
        <Icon src={BiPlus} size="24" />
        <span>Create</span>
      </Button>
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
            <td>{categoryEntry.uid}</td>
            <td>{categoryEntry.name}</td>
            <td>
              <Button maxContent on:click={() => handleDeleteCategory(categoryEntry.uid)}>
                <Icon src={BiTrash} size="24" />
                <span>Delete</span>
              </Button>
              <Button maxContent on:click={() => handleEditCategory(categoryEntry.uid, categoryEntry.name)}>
                <Icon src={BiEdit} size="24" />
                <span>Edit</span>
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
  bind:categoryUid={categoryUIdToEdit}
  on:success={handleSaveCategorySuccess}
  bind:dialog={createEditCategoryDialog} />
