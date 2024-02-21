<script lang="ts">
    import {createEventDispatcher} from "svelte";

    const dispatch = createEventDispatcher<{
        update: boolean
    }>();

    export let checked: boolean = false;
    export let value: string | undefined = undefined;
    export let name: string = "";
</script>


<input type="checkbox"
       bind:value={value}
       bind:checked={checked}
       {name}
       class="rounded border-gray-300 text-indigo-600 shadow-sm focus:ring-indigo-500"
       on:change={(e) => dispatch("update", e.currentTarget.checked) }
>
