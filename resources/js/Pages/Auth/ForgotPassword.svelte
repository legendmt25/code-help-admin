<script lang="ts">
    // import '@/@types/@inertiajs__svelte'
    import '@/@types/types'

    import {useForm} from '@inertiajs/svelte';
    import AuthenticationCard from '@/Components/AuthenticationCard.svelte';
    import AuthenticationCardLogo from '@/Components/AuthenticationCardLogo.svelte';
    import InputError from '@/Components/InputError.svelte';
    import InputLabel from '@/Components/InputLabel.svelte';
    import PrimaryButton from '@/Components/PrimaryButton.svelte';
    import TextInput from '@/Components/TextInput.svelte';

    // defineProps({
    export let status: String;
    // });

    const form = useForm({
        email: '',
    });

    const submit = () => {
        $form.post(window.route('password.email'));
    };
</script>

<svelte:head>
    <title>Forgot Password</title>
</svelte:head>

<AuthenticationCard>
    <svelte:fragment slot="logo">
        <AuthenticationCardLogo/>
    </svelte:fragment>

    <div class="mb-4 text-sm text-gray-600">
        Forgot your password? No problem. Just let us know your email address and we will email you a password reset
        link that will allow you to choose a new one.
    </div>

    {#if status}
        <div class="mb-4 font-medium text-sm text-green-600">
            {{status}}
        </div>
    {/if}

    <form on:submit|preventDefault={submit}>
        <div>
            <InputLabel for="email" value="Email"/>
            <TextInput
                    id="email"
                    bind:value={$form.email}
                    type="email"
                    class="mt-1 block w-full"
                    required
                    autofocus
                    autocomplete="username"
            />
            <InputError class="mt-2" message="{$form.errors.email}"/>
        </div>

        <div class="flex items-center justify-end mt-4">
            <PrimaryButton class="{ $form.processing ? 'opacity-25': '' }" disabled={$form.processing}>
                Email Password Reset Link
            </PrimaryButton>
        </div>
    </form>
</AuthenticationCard>
