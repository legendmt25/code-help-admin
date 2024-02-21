<script lang="ts">
    export let message: string;
    export {className as class};

    let className: string;
</script>

{#if message}
    <div class={className}>
        <p class="text-sm text-red-600">
            {{message}}
        </p>
    </div>
{/if}
