<script lang="ts">
  import { onMount } from "svelte";
  import { link } from "svelte-spa-router";
  import type { FormEventHandler } from "svelte/elements";
  import Button from "../components/Button.svelte";
  import Spinner from "../components/Spinner.svelte";
  import { type ContestDetail, type ContestEditRequest, type ContestRequest } from "../generated/admin-api";
  import { Route } from "../routes";
  import { createContest, getContestById, updateContest } from "../services/ContestService";

  export let params: { id?: string } = {};

  let loading: boolean = false;
  let value: Partial<ContestEditRequest> = {};

  let contestEntry: ContestDetail | undefined = undefined;

  onMount(() => {
    if (!Number.isNaN(Number(params.id))) {
      loading = true;
      getContestById(Number(params.id))
        .then((contest) => {
          contestEntry = contest;
          value = {
            name: contest.name,
            duration: contest.duration,
            startsOn: contest.startsOn,
            status: contest.status
          };
        })
        .finally(() => (loading = false));
    }
  });

  const handleFormSubmit: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();

    const id = Number(params.id);
    if (!Number.isNaN(id)) {
      updateContest(id, value as ContestEditRequest);
    } else {
      createContest(value as ContestRequest);
    }

    console.log(value);
  };
</script>

<style>
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

{#if loading}
  <Spinner />
{:else}
  <form on:submit={handleFormSubmit} class="form">
    <div class="input-container">
      <label for="name">Name</label>
      <input id="name" name="name" bind:value={value.name} />
    </div>
    <div class="input-container">
      <label for="duration">Duration</label>
      <input id="duration" name="duration" bind:value={value.duration} />
    </div>
    <div class="input-container">
      <label for="startsOn">Starts on</label>
      <input id="startsOn" name="startsOn" bind:value={value.startsOn} />
    </div>
    <Button>Submit</Button>
  </form>

  {#if contestEntry != undefined}
    <table>
      <thead>
        <tr>
          <th>#</th>
          <th>Title</th>
          <th>Score</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {#each contestEntry.problems as problem}
          <tr>
            <td>{problem.id}</td>
            <td>{problem.title}</td>
            <td>{problem.score}</td>
            <td>
              <a href={problem.id ? Route.problems_edit.replace(":id", problem.id.toString()) : undefined} use:link
                >Edit</a>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  {/if}
{/if}
