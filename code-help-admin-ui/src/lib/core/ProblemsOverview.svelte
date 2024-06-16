<script lang="ts">
  import { Icon } from "svelte-icons-pack";
  import { BiEdit, BiPlus, BiTrash } from "svelte-icons-pack/bi";
  import AlertBox from "../../components/AlertBox.svelte";
  import Button from "../../components/Button.svelte";
  import Spinner from "../../components/Spinner.svelte";
  import { Route } from "../../routes";
  import { deleteProblem, getAllProblems } from "../../services/core/ProblemsService";
  import { filterProblems } from "../../util";

  const getAllProblemsPromise = getAllProblems()
    .then((x) => x.problems)
    .catch(() => []);

  const handleDeleteProblem = (problemId: number) => deleteProblem(problemId);

  let search: string | undefined = undefined;
  $: filter = filterProblems.bind(undefined, search);
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

{#await getAllProblemsPromise}
  <Spinner />
{:then problems}
  <section>
    <div class="page-heading">
      <h2>Problems</h2>
      <Button maxContent type="primary-outline" href={Route.problems_create}>
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
      <colgroup>
        <col style="width: 10%; text-align: right;" />
        <col style="width: 40%; text-align: right;" />
        <col style="width: 10%; text-align: right;" />
        <col style="width: 40%; text-align: right;" />
      </colgroup>
      <thead>
        <tr>
          <th>#</th>
          <th>Title</th>
          <th>Category</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {#each filter(problems) as problemEntry}
          <tr>
            <td>{problemEntry.id}</td>
            <td>{problemEntry.title}</td>
            <td>{problemEntry.category?.name}</td>
            <td>
              <Button
              maxContent
                href={problemEntry.id ? Route.problems_edit.replace(":id", problemEntry.id.toString()) : undefined}>
                <Icon src={BiEdit} size="24" />
                <span>Edit</span>
              </Button>
              
              <Button maxContent on:click={() => handleDeleteProblem(problemEntry.id)}>
                <Icon src={BiTrash} size="24" />
                <span>Delete</span>
              </Button>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
    {#if problems.length === 0}
      <AlertBox type="info" message="No entries!" />
    {/if}
  </section>
{:catch err}
  <AlertBox type="error" message="An error occured! ({err})" />
{/await}
