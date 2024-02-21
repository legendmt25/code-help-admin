<script lang="ts">
    import {onMount, createEventDispatcher} from "svelte";

    const dispatch = createEventDispatcher<{
        update: string
    }>();

    export let id: string;
    export let type: string;
    export let value: string;
    export let required: boolean = false;
    export let autofocus: boolean = false;
    export let autocomplete: string | undefined = undefined;

    console.log(value)


    export {className as class};

    let className: string = "";
    let input: HTMLInputElement;


    onMount(() => {
        if (autofocus) {
            input.focus();
        }
    })

    export const focus = () => {
        input.focus()
    }

    function handleInput(newValue: string) {
        value = newValue;
        dispatch("update", newValue);
    }
</script>

<input {id} bind:this={input} type={type} {required} autocomplete={autocomplete}
       class="border-gray-300 focus:border-indigo-500 focus:ring-indigo-500 rounded-md shadow-sm {className}"
       value={value} on:input={(e) => handleInput(e.currentTarget.value) }>
