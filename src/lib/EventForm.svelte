<script lang="ts">
    import { beforeUpdate, afterUpdate, onDestroy } from 'svelte';
    import { writable } from 'svelte/store';
    import Contacts from './Contacts.svelte';

    let title: string;
    let content: string;
    let recipients: [];
        
    onDestroy(() => {
        // save title and content so they don't get lost
    });
    
    function createEvent() {
        alert('out');
    }

</script>

<h2 class="mb-10 text-4xl font-extrabold dark:text-white">Create New Event</h2>

<div class="form-control w-full ">
    <div>
        <label class="label" for="event-title">
            <span class="text-xl font-bold text-xllabel-text">Title</span>
        </label>
        <input
            bind:value={title}
            type="text"
            placeholder="Sunday Funday St Paddy's Vball"
            class="input input-bordered w-full max-w"
            id="event-title"
        />
        <label class="label" for="event-title">
            <span class="label-text-alt">This is just for our records</span>
        </label>
    </div>
    <div class="mt-5">
        <label class="label" for="event-content">
            <span class="text-xl font-bold text-xllabel-text">SMS Content</span>
        </label>
        <textarea
            bind:value={content}
            class="w-full h-96 pb-10 textarea textarea-bordered"
            id="event-content"
        />
    </div>
    <!-- <button on:click={login} class="mt-12 w-1/4 btn btn-outline btn-accent">Add recipients</button> -->

    <!-- The button to open modal -->
    <label for="my-modal-5" class="mt-12 w-1/4 btn btn-outline btn-accent">Add recipients</label>

    <input type="checkbox" id="my-modal-5" class="modal-toggle" />
    <div class="modal">
        <div class="modal-box w-11/12 max-w-5xl">
            <h3 class="font-bold text-lg">Congratulations random Internet user!</h3>
            <Contacts/>
            <div class="modal-action">
                <label for="my-modal-5" class="btn">Finish</label>
            </div>
        </div>
    </div>

    <div class="mt-5" />
    {#if recipients}
        <button on:click={createEvent} class="mt-12 w-1/4 btn btn-active">Create</button>
    {/if}
</div>
