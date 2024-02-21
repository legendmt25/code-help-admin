<script lang="ts">
    import AppLayout from '@/Layouts/AppLayout.svelte';
    import Welcome from "@/Pages/Welcome.svelte";

    // import Welcome from '@/Components/Welcome.vue';

    console.log($$props)



    export let jetstream: unknown;

    console.log(jetstream);
    export let laravelVersion: string;
    export let phpVersion: string;
</script>

<!--<template>-->
<AppLayout title="Dashboard">
    <!--        <template>-->
    <svelte:fragment slot="header">
        <h2 class="font-semibold text-xl text-gray-800 leading-tight">
            Dashboard
        </h2>
    </svelte:fragment>


    <div class="py-12">
        <div class="max-w-7xl mx-auto sm:px-6 lg:px-8">
            <div class="bg-white overflow-hidden shadow-xl sm:rounded-lg">
                <Welcome {laravelVersion} {phpVersion}/>
            </div>
        </div>
    </div>
</AppLayout>
