<script lang="ts">
    export let value: string;
    export {inputId as for};

    let inputId: string;
</script>

<label for={inputId} class="block font-medium text-sm text-gray-700">
    {#if value}
        <span>{value}</span>
    {:else}
        <span><slot/></span>
    {/if}
</label>
