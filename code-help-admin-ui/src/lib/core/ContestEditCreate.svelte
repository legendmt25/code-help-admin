<script lang="ts">
  import { onMount } from "svelte";
  import { link, push } from "svelte-spa-router";
  import type { ChangeEventHandler, FormEventHandler } from "svelte/elements";
  import Button from "../../components/Button.svelte";
  import Spinner from "../../components/Spinner.svelte";
  import {
    ContestStatus,
    type ContestDetail,
    type ContestEditRequest,
    type ContestRequest
  } from "../../generated/admin-core-api";
  import { Route } from "../../routes";
  import { createContest, getContestById, updateContest } from "../../services/core/ContestService";
  import { deleteProblem } from "../../services/core/ProblemsService";
  import { Icon } from "svelte-icons-pack";
  import { BiEdit, BiTrash } from "svelte-icons-pack/bi";

  export let params: { id?: string } = {};

  type ContestFormData = Omit<Partial<ContestEditRequest>, "startDate"> & {
    startDate: Date;
  };

  let loading: boolean = false;
  let value: ContestFormData = {
    problems: [],
    startDate: new Date()
  };

  const convertToDateTimeLocalString = (date: Date) => {
    const year = date.getFullYear();
    const month = (date.getMonth() + 1).toString().padStart(2, "0");
    const day = date.getDate().toString().padStart(2, "0");
    const hours = date.getHours().toString().padStart(2, "0");
    const minutes = date.getMinutes().toString().padStart(2, "0");

    return `${year}-${month}-${day}T${hours}:${minutes}`;
  }

  let contestEntry: ContestDetail | undefined = undefined;

  onMount(() => {
    if (!Number.isNaN(Number(params.id))) {
      loading = true;
      getContestById(Number(params.id))
        .then((contest) => {
          contestEntry = contest;
          value = {
            name: contest.name,
            description: contest.description,
            duration: contest.duration,
            startDate: contest.startDate,
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
      editCreatePromise = createContest(value as ContestRequest).then(() => {});
    }

    editCreatePromise.finally(() => push(Route.contests_overview));
  };

  const handleDateChange: ChangeEventHandler<HTMLInputElement> = (event) => {
    value.startDate = new Date(event.currentTarget.value);
  };

  const handleDeleteProblem = (id: number) => {
    deleteProblem(id);
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

  form {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    max-width: 40rem;
    width: 100%;
  }

  label {
    font-size: 0.9em;
    font-weight: bold;
  }

  .form-container {
    display: flex;
    align-items: center;
    flex-direction: column;
    padding-top: 2rem;
    gap: 4rem;
  }
</style>

{#if loading}
  <Spinner />
{:else}
  <div class="form-container p-1">
    <h2>Create new contest</h2>
    <form on:submit={handleFormSubmit} class="form">
      <div class="input-container">
        <label for="name">Name</label>
        <input id="name" name="name" bind:value={value.name} />
      </div>
      <div class="input-container">
        <label for="description">Description</label>
        <input id="description" name="description" bind:value={value.description} />
      </div>
      <div class="input-container">
        <label for="duration">Duration</label>
        <input id="duration" name="duration" bind:value={value.duration} />
      </div>
      <div class="input-container">
        <label for="startDate">Start date</label>
        <input
          type="datetime-local"
          id="startDate"
          name="startDate"
          value={convertToDateTimeLocalString(value.startDate)}
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
      <Button maxContent>Submit</Button>
    </form>

    {#if contestEntry}
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
              <td>{problem.name}</td>
              <td>{problem.score}</td>
              <td>
                <Button maxContent on:click={() => handleDeleteProblem(problem.id)}>
                  <Icon src={BiTrash} size="24" />
                  <span>Delete</span>
                </Button>
                <Button
                  maxContent
                  href={problem.id ? Route.problems_edit.replace(":id", problem.id.toString()) : undefined}>
                  <Icon src={BiEdit} size="24" />
                  <span>Edit</span>
                </Button>
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    {/if}
  </div>
{/if}
