<script lang="ts">

    export {className as class};
    export let type: "submit" | "button" | "reset" | null | undefined = "submit";
    export let disabled = false;

    let className: string;
</script>

<button type={type} {disabled}
        class="inline-flex items-center px-4 py-2 bg-gray-800 border border-transparent rounded-md font-semibold text-xs text-white uppercase tracking-widest hover:bg-gray-700 focus:bg-gray-700 active:bg-gray-900 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 transition ease-in-out duration-150 {className}">
    <slot/>
</button>
