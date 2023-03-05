<script lang="ts">
    import { beforeUpdate, afterUpdate, onDestroy, onMount } from 'svelte';
    import { writable } from 'svelte/store';
    import ContactList from './ContactList.svelte';
    import { contacts, get50Contacts } from './pocketbase';
    import type { ContactsResponse } from './pocketbase-types';

    let title: string;
    let content: string;
    let recipients = new Array<ContactsResponse>();

    function handleMessage(event) {
        recipients = event.detail.recipients;
    }
    function createEvent() {
        recipients.forEach((contact) => {
            console.log(contact.id);
        });
    }
    beforeUpdate(() => {
        if ($contacts.length == 0) {
            get50Contacts();
        }
    });
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

    <!-- The button to open modal -->
    <label for="my-modal-5" class="mt-12 w-1/4 btn btn-outline btn-accent"
        >{recipients.length > 0 ? 'Edit Recipients' : 'Add Recipients'}</label
    >
    <input type="checkbox" id="my-modal-5" class="modal-toggle" />
    <div class="modal">
        <div class="modal-box w-11/12 max-w-5xl">
            <ContactList on:message={handleMessage} />
            <div class="modal-action">
                <label for="my-modal-5" class="btn">Finish</label>
            </div>
        </div>
    </div>
    <div class="mt-5" />
    {#if recipients.length > 0}
        <button on:click={createEvent} class="mt-12 w-1/4 btn btn-active">Create</button>
    {/if}
</div>
