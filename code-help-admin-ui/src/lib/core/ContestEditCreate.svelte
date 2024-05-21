<script lang="ts">
  import { onMount } from "svelte";
  import { link, push } from "svelte-spa-router";
  import type { ChangeEventHandler, FormEventHandler } from "svelte/elements";
  import Button from "../../components/Button.svelte";
  import Link from "../../components/Link.svelte";
  import Spinner from "../../components/Spinner.svelte";
  import {
    ContestStatus,
    type ContestDetail,
    type ContestEditRequest,
    type ContestRequest
  } from "../../generated/admin-api";
  import { Route } from "../../routes";
  import { createContest, getContestById, updateContest } from "../../services/core/ContestService";

  export let params: { id?: string } = {};

  let loading: boolean = false;
  let value: Partial<ContestEditRequest> = {
    problems: []
  };

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
            status: contest.status,
            problems: contestEntry.problems
          };
        })
        .finally(() => (loading = false));
    }
  });

  const handleFormSubmit: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();

    const id = Number(params.id);

    let editCreatePromise: Promise<void>;

    if (!Number.isNaN(id)) {
      editCreatePromise = updateContest(id, value as ContestEditRequest);
    } else {
      editCreatePromise = createContest(value as ContestRequest);
    }

    editCreatePromise.finally(() => push(Route.contests_overview));
  };

  const handleDateChange: ChangeEventHandler<HTMLInputElement> = (event) => {
    value.startsOn = new Date(event.currentTarget.value);
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

  .p-1 {
    padding: 1rem;
  }

  .input-container {
    display: flex;
    flex-direction: column;
  }

  input,
  select {
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

  .gap-1 {
    gap: 1rem;
  }

  .column {
    display: flex;
    flex-direction: column;
  }
</style>

{#if loading}
  <Spinner />
{:else}
  <div class="column p-1 gap-1">
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
        <input
          type="date"
          id="startsOn"
          name="startsOn"
          value={value.startsOn?.toISOString().substring(0, 10)}
          on:change={handleDateChange} />
      </div>
      <div class="input-container">
        <label for="status">Status</label>
        <select id="status" name="status" bind:value={value.status}>
          {#each Object.values(ContestStatus) as status}
            <option value={status}>{status}</option>
          {/each}
        </select>
      </div>
      <Button>Submit</Button>
    </form>

    {#if contestEntry != undefined && params.id}
      <Link href={Route.contest_problems_create.replace(":contestId", params.id)}>Create problem</Link>

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
  </div>
{/if}
