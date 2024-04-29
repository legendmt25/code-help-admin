<script lang="ts">
  import Button from "../components/Button.svelte";
  import Spinner from "../components/Spinner.svelte";
  import { getAllCategories } from "../services/ProblemsService";
  import CategoryEditCreateDialog from "./CategoryEditCreateDialog.svelte";

  const getAllCategoriesPromise = getAllCategories()
    .then((data) => data.categories)
    .catch(() => []);

  const handleDeleteCategory = (categoryName: string) => {};

  let createCategoryDialog: HTMLDialogElement | undefined = undefined;
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
    <Button on:click={() => createCategoryDialog?.showModal()}>Create</Button>
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

<CategoryEditCreateDialog bind:dialog={createCategoryDialog} />
